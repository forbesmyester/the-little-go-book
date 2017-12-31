package main

import (
    "fmt"
    "math/rand"
    "sort"
)

func printIt(name string, m map[string]int) string {
    power, exists := m[name]
    if exists {
        return fmt.Sprintf("%s power is %d", name, power)
    }
    return fmt.Sprintf("%s does not exists", name)
}

func main() {
    fmt.Println("Slices")
    scores := make([]int, 100)
    fmt.Println(scores)
    for i := 0; i < 100; i++ {
        scores[i] = int(rand.Int31n(1000))
    }
    sort.Ints(scores)
    worst := make([]int, 5)
    copy(worst, scores[:20])
    fmt.Println(worst)
    fmt.Println("Maps")
    lookup := make(map[string]int)
    lookup["goku"] = 9001
    for key, _ := range lookup {
        fmt.Println(printIt(key, lookup))
    }
    fmt.Println(printIt("vegeta", lookup))
}

