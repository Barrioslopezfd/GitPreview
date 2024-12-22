package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
)

type File struct {
	path    string
	content string
	html    string
	mu      sync.RWMutex
}

func (f *File) watchReadme() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("Failed creating the watcher: %s", err.Error())
		return
	}

	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("Event: ", event)
				if event.Has(fsnotify.Rename) {
					log.Println("README modified, updating content")
					f.updateReadmeContent()
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("Watcher error: %s", err.Error())
			}
		}
	}()

	err = watcher.Add(f.path)
	if err != nil {
		log.Printf("Failed adding filepath: %s", err.Error())
		return
	}
	<-make(chan struct{})
}

func (f *File) updateReadmeContent() {
	data, err := os.ReadFile(f.path)
	if err != nil {
		log.Printf("Error reading f: %s", err.Error())
		return
	}

	f.mu.Lock()
	f.content = string(data)
	f.mu.Unlock()
}

func (f *File) serveReadme(w http.ResponseWriter, _ *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.content == "" {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	f.ToHTML()
	//TODO:CREATE HTML CONTENT WITH THE MARKDOWN
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(f.html))
}
