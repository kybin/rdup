package main

import (
    "os"
    "flag"
    "bufio"
    "fmt"
)

var nprev int

func init() {
    flag.IntVar(&nprev, "n", 1, "how many lines to remember")
}

func main() {
    flag.Parse()

    scanner := bufio.NewScanner(os.Stdin)
    prevs := make([]string, 0)
    for scanner.Scan() {
        in := scanner.Text()
        match := false
        for _, p := range prevs {
            if in == p {
                match = true
                break
            }
        }
        if !match {
            fmt.Println(in)
            prevs = append(prevs, in)
            if len(prevs) > nprev {
                prevs = prevs[1:]
            }
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "cannot read from pipe:", err)
    }
}
