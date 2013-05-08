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
	Description string
	Public      bool
	Files       map[string]map[string]string
}

func CreateFilesMap(f []string) map[string]map[string]string {
	for _, file := range f {
		filename := path.Base(file)
		content, err := ioutil.ReadFile(file)
		if err != nil {
			log.Println(err)
		}
		files[filename] = map[string]string{"content": string(content)}
	}
	return files
}

func CreateJson(d string, p bool, f map[string]map[string]string) string {
	gistJson := Gist{d, p, f}
	g, err := json.Marshal(gistJson)
	if err != nil {
		log.Println("error: ", err)
	}
	return strings.ToLower(string(g))
}

func PostGist(gist string) GistResponse {
	r := strings.NewReader(gist)
	resp, err := http.Post(url, "application/json", r)
	if err != nil {
		log.Println("error: ", err)
	}
	postResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error ", err)
	}
	resp.Body.Close()
	var g GistResponse
	err = json.Unmarshal(postResp, &g)
	if err != nil {
		log.Println("error ", err)
	}
	return g
}
