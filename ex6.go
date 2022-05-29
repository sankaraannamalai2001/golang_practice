
package main 
    
import ( 
    "fmt"
) 
func myfunc(ch chan int) {
  
    fmt.Println(234 + <-ch)
}
func main() { 
    
    ch1 := make(chan int, 1) 
    ch2 := make(chan string, 5)
    ch3 := make(chan float64, 8)
	ch1<-20
	ch2<-"Sankar"
	ch2<-"somu"
	ch2<-"raj"
	
	myfunc(ch1)
    fmt.Println("Capacity of ch1", cap(ch1)) 
    fmt.Println("Capacity of ch2: ", cap(ch2)) 
    fmt.Println("Capacity of ch3: ", cap(ch3)) 
    
} 