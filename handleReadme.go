package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
)

type Readme struct {
    path    string
    content string
    mu      sync.RWMutex
}

func (file *Readme) watchReadme() {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
	fmt.Printf("Failed creating the watcher: %s", err.Error())
	return
    }

    defer watcher.Close()

    go func(){
	for {
	    select{
	    case event, ok := <-watcher.Events:
		if !ok{
		    return
		}
		log.Println("Event: ", event)
		if event.Has(fsnotify.Rename){
		    log.Println("README modified, updating content")
		    file.updateReadmeContent()
		}

	    case err, ok := <-watcher.Errors:
		if !ok{
		    return
		}
		log.Printf("Watcher error: %s", err.Error())
	    }
	}
    }()

    err = watcher.Add(file.path)
    if err != nil {
	log.Printf("Failed adding filepath: %s", err.Error())
	return
    }
    <-make(chan struct{})
}

func (file *Readme) updateReadmeContent() {
    data, err := os.ReadFile(file.path)
    if err != nil {
	log.Printf("Error reading file: %s", err.Error())
	return
    }

    file.mu.Lock()
    file.content = string(data)
    file.mu.Unlock()
}

func (file *Readme) serveReadme(w http.ResponseWriter, _ *http.Request){
    file.mu.Lock()
    defer file.mu.Unlock()

    if file.content == "" {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }
    //TODO:CREATE AN HTML WRITE WITH THE MARKDOWN
    w.Write([]byte(file.content))
}
