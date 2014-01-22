// Go 1.2
// go run helloworld_go.go

package main

import (
    . "fmt" // Using '.' to avoid prefixing functions with their package names
    . "runtime" // This is probably not a good idea for large projects...
    . "time"
)

var i = 0

func incr(sem chan int) {
        for x := 0; x < 1000004; x++ {
			sem <- 1
                i++
			<- sem
            }
}

func decr(sem chan int){
        for x := 0; x < 1000000; x++ {
			sem <- 2
                i--
			<- sem
            }
}

func main() {
    GOMAXPROCS(NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
		var sem = make(chan int, 1)
        go incr(sem)
        go decr(sem)
            /*for x := 0; x < 50; x++ {
        Println(i)
}*/
    // No way to wait for the completion of a goroutine (without additional syncronization)
    // We'll come back to using channels in Exercise 2. For now: Sleep
    Sleep(1000*Millisecond)
    Println("Done:", i);
}
