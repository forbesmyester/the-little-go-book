package main

import (
    "fmt"
    "os"
    "strings"
    "io"
    "net/http"
)

func main() {
    for _, url := range os.Args[1:] {
        if !strings.HasPrefix(url, "http") {
            url = "http://" + url
        }
        resp, err := http.Get(url)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error retrieving URL: %s\n\n%v", url, err);
        }
        _, err2 := io.Copy(os.Stdout, resp.Body)
        fmt.Fprintf(os.Stderr, "\n\nStatus Code: %d", resp.StatusCode)
        if err2 != nil {
            fmt.Fprintf(os.Stderr, "Error Outputting URL: %s\n\n%v", url, err);
        }
    }
}
