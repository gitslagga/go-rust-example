
package main
 
import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    
    timer1 := time.NewTimer(2 * time.Second)
    ticker1 := time.NewTicker(2 * time.Second)
 
    go func(t *time.Ticker) {
        defer wg.Done()
        for {
            <-t.C
            fmt.Println("get ticker1", time.Now().Format("2006-01-02 15:04:05"))
        }
    }(ticker1)
 
    go func(t *time.Timer) {
        defer wg.Done()
        for {
            <-t.C
            fmt.Println("get timer", time.Now().Format("2006-01-02 15:04:05"))
            t.Reset(2 * time.Second)
        }
    }(timer1)
 
    wg.Wait()
}

//output
//get ticker1 2019-08-01 23:13:10
//get timer 2019-08-01 23:13:10
//...
