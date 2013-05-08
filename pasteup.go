package main

import (
	"fmt"
	"github.com/imwally/pasteup/gist"
	"log"
	"os"
	"os/exec"
	"time"
)

var dateStamp string = time.Now().Format("20060102150530")
var fileName string = "paste_" + dateStamp + ".txt"
var tempDir string = os.TempDir()
var files = make(map[string]map[string]string)

func newPaste(fileName string) {
	cmd := exec.Command("vi", "+star", fileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Printf("2")
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("Error creating new paste. Error: %v\n", err)
	}
}

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		pasteFile := tempDir + fileName
		newPaste(pasteFile)
		files = gist.CreateFilesMap([]string{pasteFile})
	} else {
		files = gist.CreateFilesMap(args)
	}

	json := gist.CreateJson("", false, files)
	r := gist.PostGist(json)
	fmt.Println(r.HtmlUrl)

}
