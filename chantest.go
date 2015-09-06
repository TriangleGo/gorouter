package main



import (
	"fmt"
)


func main () {
	
	specChan := make(chan interface{})
	
	go func(ch chan interface{}) {
		data := <-ch
		
		fmt.Printf("data %v\n",data)
		data = <-ch
		fmt.Printf("data %v\n",data)
		data = <-ch
		fmt.Printf("data %v\n",(data).(int))
	}( specChan)
	
	specChan <- "12314"
	specChan <- []byte("abcde")
	specChan <- 321
}