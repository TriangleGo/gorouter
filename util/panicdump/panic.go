package panic

import (
    "fmt"
    "time"
    "os"
    "log"
)

func CreateCrashDump() {
    fd,err := os.Create("crash."+ string(randomCreateBytes(8)))
    if err != nil {
        log.Fatal(err)
    }
    redirectStderr(fd)
}


func saferoutine(c chan bool) {
    for i := 0; i < 10; i++ {
        fmt.Println("Count:", i)
        time.Sleep(1 * time.Second)
    }
    c <- true
}

func panicgoroutine(c chan bool) {
    time.Sleep(5 * time.Second)
    panic("Panic, omg ...")
    c <- true
}


func test() {
    
    CreateCrashDump()
    
    c := make(chan bool, 2)
        go saferoutine(c)
        go panicgoroutine(c)
        for i := 0; i < 2; i++ {
        <-c
    }

    defer fmt.Print("defer")

}

