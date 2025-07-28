package main

import {
    "fmt"
    "log"
    "os"
}

func start_watch(root string) error {
    watcher, err := createWatcher()
    if err != nil {
        return err
    }
    defer watcher.Close()
    
    if err := addRecursiveDirs(watcher, root); err != nil {
        return err
    }

    go handleEvents(watcher)
    
    fmt.Println("cloud storage at :", root)
    select {} // block forever (figure out somethign better)
}
func handleEvents(watcher *fsnotify.Watcher) {
    for {
        select {
            case event, ok 
        }
    }
}