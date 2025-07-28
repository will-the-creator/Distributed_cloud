package main

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

func rand_colors() string {
    r := rand.Intn(256)
    g := rand.Intn(256)
    b := rand.Intn(256)
    return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// Generate a random password
func gen_pass(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/"
    var pass strings.Builder
    for i := 0; i < length; i++ {
        pass.WriteByte(charset[rand.Intn(len(charset))])
    }
    return pass.String()
}

func main() {
    rand.Seed(time.Now().UnixNano()) 

    length := 16

    for {
        colorCode := rand_colors()
        password := gen_pass(length)
        fmt.Printf("%s%s\033[0m\n", colorCode, password) // resets and stuff
        //time.Sleep(500 * time.Millisecond) // Optional: Slow it down
    }
}
