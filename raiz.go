package main

import (
    "fmt"
    "math"
    "log"
    "net/http"
)

func greetings(message string) string {
    return "<b>" + message + "</b>"
}

func handler(w http.ResponseWriter, r *http.Request) {
    x := 0.0001

    for i := 0; i < 1000000; i++ {
        x += math.Sqrt(x)
        fmt.Fprintf(w, greetings("Code.education Rocks!"))
    }
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}