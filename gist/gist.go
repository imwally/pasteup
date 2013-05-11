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

type GistResponse struct {
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

type Gist struct {
	Description string                       `json:"description"`
	Public      bool                         `json:"public"`
	Files       map[string]map[string]string `json:"files"`
}

func CreateFilesMap(f []string) (fm map[string]map[string]string, err error) {
	var c []byte
	for _, file := range f {
		fn := path.Base(file)
		c, err = ioutil.ReadFile(file)
		files[fn] = map[string]string{"content": string(c)}
	}
	return files, err
}

func CreateJson(d string, p bool, fm map[string]map[string]string) string {
	gistJson := Gist{d, p, fm}
	g, err := json.Marshal(gistJson)
	if err != nil {
		log.Println(err)
	}
	return string(g)
}

func PostGist(gist string) GistResponse {
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
	var g GistResponse
	err = json.Unmarshal(postResp, &g)
	if err != nil {
		log.Println(err)
	}
	return g
}
