package main
  
import (
    "fmt"
    "time"
)
func Authorname() {
    arr1 := [4]string{"Sankar", "Prathish", "Paulson", "Santhosh"}
    for t1 := 0; t1 <= 3; t1++ {   
        time.Sleep(150 * time.Millisecond)
        fmt.Printf("%s\n", arr1[t1])
    }
}
func Authorid() {
   id := [4]int{81, 82, 83, 84}
      
    for t2 := 0; t2 <= 3; t2++ {
      
        time.Sleep(160 * time.Millisecond)
        fmt.Printf("%d\n", id[t2])
    }
}
  

func main() {
  
    fmt.Println("Start")
    go Authorname()
    go Authorid()
    time.Sleep(3500 * time.Millisecond)
    fmt.Println("\nEnd")
}