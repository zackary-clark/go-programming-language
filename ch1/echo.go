package main

import (
    "fmt"
    "os"
)

func main() {
    //fmt.Println(strings.Join(os.Args[1:], " "))

    for i, arg := range os.Args {
        fmt.Printf("%d:\t%s\n", i, arg)
    }
}
