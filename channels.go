package main

import "fmt"

func numberGen(start, count int, out chan<- int) {
    for i := 0; i < count; i++ {
        out <- start + i
    }
    close(out)
}

func printNumbers(in <-chan int, done chan<- bool) {
    for num := range in {
        fmt.Printf("%d\n", num)
    }
    done <- true
}

func main() {
    numberChan := make(chan int)
    done := make(chan bool)
    go numberGen(1, 10, numberChan)
    go printNumbers(numberChan, done)

    <-done
}
