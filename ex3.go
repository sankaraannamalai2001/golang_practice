package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
	City string
}

func main() {
	user := &User{Name: "Sankar", Age: 20, City: "Tirunelveli"}
	data, _ := json.Marshal(user)
	fmt.Println(string(data))
}