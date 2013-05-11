package main

import (
	"fmt"
	"github.com/imwally/pasteup/gist"
	"io/ioutil"
	"os"
	"os/exec"
)

var resp gist.GistResponse 
var json string 
var files map[string]map[string]string

func newPaste() (name string, err error) {
	tempf, err := ioutil.TempFile("", "paste_")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	name = tempf.Name()	
	cmd := exec.Command("vi", "+star", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	return
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		paste, err := newPaste()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		files, err := gist.CreateFilesMap([]string{paste})
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			json = gist.CreateJson("", false, files)
			resp = gist.PostGist(json)
			fmt.Println(resp.HtmlUrl)
		}
	} else {
		files, err := gist.CreateFilesMap(args)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			json = gist.CreateJson("", false, files)
			resp = gist.PostGist(json)
			fmt.Println(resp.HtmlUrl)
		}	
	}
}
