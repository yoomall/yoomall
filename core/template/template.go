package template

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"strings"
	"yoomall/config"

	"github.com/charmbracelet/log"
)

/**
@ usege
```golang
html := template.Must(_template.ParseGlob(template.New("main"), "templates", "*.html")).Funcs(_template.TemplateFuncsMap)
engine.SetHTMLTemplate(html)
```
	**/

// TemplateRenderer is a custom html/template renderer for Echo framework
type (
	TemplateRenderer struct {
		Templates *template.Template
	}

	tplFile struct {
		Name string
		Path string
	}
)

// ParseGlob 自定义模版解析，扫描子目录
func ParseGlob(tpl *template.Template, dir string, pattern string) (t *template.Template, err error) {
	t = tpl
	fmt.Println("扫描模版目录：" + dir)
	files := allFiles(dir, pattern)
	for _, file := range files {
		fmt.Printf("挂载模板：%s\n", file.Path)
		b, err := os.ReadFile(file.Path)
		if err != nil {
			return t, err
		}
		s := string(b)
		name := file.Name
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

// 目录下的所有文件
func allFiles(dir string, suffix string) (arr []*tplFile) {

	files, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("扫描子目录：" + file.Name())
			arr = append(arr, allFiles(path.Join(dir, file.Name()), suffix)...)
		} else {
			ok, _ := filepath.Match(suffix, file.Name())
			if ok {
				pathName := path.Join(dir, file.Name())
				list := strings.Split(filepath.ToSlash(pathName), "/")
				if len(list) > 1 {
					list = list[1:]
				}
				// fmt.Println(pathName, list)
				name := strings.Join(list, "/")
				arr = append(arr, &tplFile{Name: name, Path: pathName})
			}
		}
	}

	return
}

var Funcs = template.FuncMap{
	"hello": func() string {
		return "hello world by template funcs!"
	},
	"vite": func(_path string) string {
		if config.VITE_DEBUG {
			return config.VITE_URL + _path
		}
		manifestPath := path.Join(config.VITE_BUILD_DIR, "manifest.json")
		manifestFile, err := os.Open(manifestPath)
		if err != nil {
			log.Error("open manifest error: " + err.Error())
			return "open manifest error: " + err.Error()
		}
		defer manifestFile.Close()
		manifestData := make(map[string]map[string]any)
		if err := json.NewDecoder(manifestFile).Decode(&manifestData); err != nil {
			log.Error("parse manifest error: " + err.Error())
			return "parse manifest error: " + err.Error()
		}

		if _, ok := manifestData[_path]; !ok {
			log.Error("path not found in manifest: " + _path)
			return "path not found in manifest: " + _path
		}

		if _, ok := manifestData[_path]["file"]; !ok {
			log.Error("file not found in manifest: " + _path)
			return "file not found in manifest: " + _path
		}

		return manifestData[_path]["file"].(string)
	},
}
