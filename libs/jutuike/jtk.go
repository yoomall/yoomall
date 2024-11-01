package jutuike

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
)

type JtkConfig struct {
	PubId  string `mapstructure:"pub_id" json:"pub_id" yaml:"pub_id"`
	ApiKey string `mapstructure:"api_key" json:"api_key" yaml:"api_key"`
	JtkUrl string `mapstructure:"jtk_url" json:"jtk_url" yaml:"jtk_url"`
}

type Jtk struct {
	Config *JtkConfig
	cache  *cache.Cache
}

func NewJtkFromViper(config *viper.Viper) (*Jtk, error) {
	return NewJtkClient(&JtkConfig{
		PubId:  config.GetString("jutuike.pub_id"),
		ApiKey: config.GetString("jutuike.api_key"),
		JtkUrl: config.GetString("jutuike.jtk_url"),
	})
}

func NewJtkClient(config *JtkConfig) (*Jtk, error) {
	if config == nil {
		config = &JtkConfig{}
	}

	if config.JtkUrl == "" {
		return nil, fmt.Errorf("jtk_url is required")
	}

	if config.PubId == "" {
		return nil, fmt.Errorf("pub_id is required")
	}

	if config.ApiKey == "" {
		return nil, fmt.Errorf("api_key is required")
	}

	return &Jtk{
		Config: config,
		cache:  cache.New(5*time.Minute, time.Hour*24*7),
	}, nil
}

func (j *Jtk) hashedUrlMethodAndParams(url string, method string, params map[string]string) string {
	paramsJsonStr := ""

	if params != nil {
		paramsJson, _ := json.Marshal(params)
		paramsJsonStr = string(paramsJson)
	}

	str := url + method + paramsJsonStr
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func (j *Jtk) RequestWithCache(path string, method string, sid string, params map[string]string) (*http.Response, any, bool, error) {
	hash := j.hashedUrlMethodAndParams(path, method, params)

	if v, ok := j.cache.Get(hash); ok {
		data, ok := v.(map[string]any)
		if !ok {
			return nil, nil, false, fmt.Errorf("cache value is not []byte")
		}
		return nil, data, true, nil
	}

	resp, data, err := j.Request(path, method, sid, params)

	if err != nil {
		return resp, data, false, err
	}

	j.cache.Set(hash, data, cache.DefaultExpiration)

	// cache
	return resp, data, false, nil
}

// request

func (j *Jtk) Request(path string, method string, sid string, params map[string]string) (*http.Response, any, error) {
	params["apikey"] = j.Config.ApiKey
	params["sid"] = sid

	url := j.Config.JtkUrl + path
	method = strings.ToUpper(method)
	var req *http.Request

	if method == http.MethodGet {
		url += "?"
		for k, v := range params {
			url += k + "=" + v + "&"
		}
		req, _ = http.NewRequest(http.MethodGet, url, nil)
	} else {
		b, _ := json.Marshal(params)
		req, _ = http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return resp, nil, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, nil, err
	}

	var data map[string]any
	err = json.Unmarshal(b, &data)
	if err != nil {
		return resp, nil, err
	}

	return resp, data, nil
}
