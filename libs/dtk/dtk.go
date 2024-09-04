package dtk

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	"github.com/patrickmn/go-cache"
	"golang.org/x/exp/rand"
)

type DtkConfig struct {
	AppKey    string `yaml:"app_key"`
	AppSecret string `yaml:"app_secret"`
	AppUrl    string `yaml:"app_url"`
}

type Dtk struct {
	Config *DtkConfig
	cache  *cache.Cache
}

func NewDtkClient(config *DtkConfig) *Dtk {
	if config == nil {
		config = &DtkConfig{}
	}

	if config.AppUrl == "" {
		panic("app_url is required")
	}

	if config.AppKey == "" {
		panic("app_key is required")
	}

	if config.AppSecret == "" {
		panic("app_secret is required")
	}
	return &Dtk{
		Config: config,
		cache:  cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (d *Dtk) rand6Num() string {
	seed := "0123456789"
	b := make([]byte, 6)
	for i := range b {
		b[i] = seed[rand.Intn(len(seed))]
	}
	return string(b)
}

func (d *Dtk) publicParams() map[string]string {
	return map[string]string{
		"appKey":  d.Config.AppKey,
		"signRan": "",
		"timer":   strconv.FormatInt(time.Now().UnixMilli(), 10),
		"nonce":   d.rand6Num(),
	}
}

func (d *Dtk) sign(params map[string]string) string {
	appKey := d.Config.AppKey
	timer := params["timer"]
	nonce := params["nonce"]
	appSecret := d.Config.AppSecret
	// 1、将当前应用的appkey，appsecret，nonce参数和timer参数进行组装，拼接成字符串：appKey=xxx&timer=xxx&nonce=xxx&key=xxx （key对应appsecret）
	str := "appKey=" + appKey + "&timer=" + timer + "&nonce=" + nonce + "&key=" + appSecret
	// 2、使用MD5进行加密，得到signRan
	sum := md5.Sum([]byte(str))
	signRan := hex.EncodeToString(sum[:])
	return strings.ToUpper(signRan)
}

func mergeParams(params1 map[string]string, params2 map[string]string) map[string]string {
	for k, v := range params2 {
		params1[k] = v
	}
	return params1
}

func (d *Dtk) hashedUrlMethodAndParams(url string, method string, params map[string]string) string {
	str := fmt.Sprintf("%s%s%s%s", url, method, params["nonce"], params["timer"])
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func (d *Dtk) RequestWithCache(path string, method string, version string, params map[string]string) ([]byte, bool) {
	hash := d.hashedUrlMethodAndParams(path, method, params)

	if v, ok := d.cache.Get(hash); ok {
		resp, ok := v.([]byte)
		if !ok {
			return []byte(""), false
		}
		return resp, true
	}

	resp, err := d.Request(path, method, version, params)

	if err != nil {
		return []byte(""), false
	}

	d.cache.Set(hash, resp, 5*time.Minute)

	return resp, false
}

func (d *Dtk) Request(path string, method string, version string, params map[string]string) ([]byte, error) {
	publicParams := d.publicParams()
	publicParams["signRan"] = d.sign(publicParams)
	mergeParams := mergeParams(publicParams, params)

	url := d.Config.AppUrl + path + "?"
	method = strings.ToUpper(method)

	var req *http.Request

	log.Info("dtk request", "url", url, "method", method, "params", mergeParams)
	if method == http.MethodGet {
		for k, v := range mergeParams {
			url += k + "=" + v + "&"
		}
		req, _ = http.NewRequest(http.MethodGet, url, nil)
	} else {
		b, _ := json.Marshal(mergeParams)
		req, _ = http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte(""), err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()

	return body, nil
}
