package main

import (
    "os"
    "bufio"
    "fmt"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    prev := ""
    for scanner.Scan() {
        in := scanner.Text()
        if in != prev {
            fmt.Println(in)
            prev = in
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "cannot read from pipe:", err)
    }
}
