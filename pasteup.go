package main

import (
	"fmt"
	"github.com/imwally/pasteup/gist"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var paste string
var files map[string]map[string]string
var json string 

func newPaste() (name string, err error) {
	
	tempFile, err := ioutil.TempFile("", "paste_")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	
	cmd := exec.Command("vi", "+star", tempFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		log.Printf("Error: %v\n", err)
		log.Fatal(err)
		return
	}
	
	err = cmd.Wait()
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	return
}

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		paste, err := newPaste()
		if err != nil {
			log.Printf("%v", err)
		}
		files, err := gist.CreateFilesMap([]string{paste})
		if err != nil {
			log.Printf("%v", err)
		}
		json = gist.CreateJson("", false, files)
	} else {
		files, err := gist.CreateFilesMap(args)
		if err != nil {
			log.Printf("%v", err)
		} else {
			json = gist.CreateJson("", false, files)
		}	
	}

	fmt.Println(json)
	os.Remove(paste)

}
