package gist

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
)

const url string = "https://api.github.com/gists"
var files = make(map[string]map[string]string)
var content []byte

type GistResponse struct {
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

type Gist struct {
	Description string
	Public      bool
	Files       map[string]map[string]string
}

func CreateFilesMap(f []string) (files map[string]map[string]string, err error) {
	for _, file := range f {
		content, err = ioutil.ReadFile(file)
		if err != nil {
			return 
		}
		filename := path.Base(file)
		files[filename] = map[string]string{"content": string(content)}
	}
	return
}

func CreateJson(d string, p bool, f map[string]map[string]string) string {
	gistJson := Gist{d, p, f}
	g, err := json.Marshal(gistJson)
	if err != nil {
		log.Println(err)
	}
	return strings.ToLower(string(g))
}

func PostGist(gist string) string {
	r := strings.NewReader(gist)
	resp, err := http.Post(url, "application/json", r)
	if err != nil {
		log.Println(err)
	}
	postResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	//var g GistResponse
	//err = json.Unmarshal(postResp, &g)
	if err != nil {
		log.Println(err)
	}
	return string(postResp) 
}
