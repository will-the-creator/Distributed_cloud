package main

import (
    "fmt"
    "github.com/fsnotify/fsnotify"
    "log"
    "time"
)

func n_watcher() {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal("error:", err)
    }
    defer watcher.Close()

    done := make(chan bool)

    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                fmt.Println("event:", event)
                if event.Op&fsnotify.Write == fsnotify.Write {
                    fmt.Println("modified file:", event.Name)
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                fmt.Println("error:", err)
            }
        }
    }()

    err = watcher.Add(".") // watch current directory
    if err != nil {
        log.Fatal("error:", err)
    }
    <-done // block forever
}

func anonymous() {
    done := make(chan bool)

    go func() {
        fmt.Println("running in anonymous goroutine")
        done <- true
    }()

    <-done // wait for goroutine to finish
}

func main() {
    fmt.Println("Starting anonymous func:")
    anonymous()

    fmt.Println("\nStarting watcher (press Ctrl+C to stop):")
    n_watcher()

    // We'll never get here because n_watcher blocks forever
    time.Sleep(1 * time.Second)
}
