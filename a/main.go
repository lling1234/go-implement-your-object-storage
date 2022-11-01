package main

import "fmt"

func main() {
    a := []int{5}
    for range a {
        a = append(a, 1)
    }
    fmt.Println(len(a))
}