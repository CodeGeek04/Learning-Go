package main

import "fmt"

func main() {
    // Slice
    fmt.Println("Slices:")
    s := make([]int, 3, 5)
    fmt.Printf("Slice: %v\n", s)
    fmt.Printf("Length: %d, Capacity: %d\n\n", len(s), cap(s))

    // Map
    fmt.Println("Maps:")
    m := make(map[string]int)
    m["key"] = 10
    fmt.Printf("Map: %v\n\n", m)

    // Channel
    fmt.Println("Channels:")
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    fmt.Printf("Channel: %v %v\n", <-ch, <-ch)
}
