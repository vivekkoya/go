package main

import (
	"fmt"
	"os"

)

func main() {
    fmt.Println("Hello, World!")
    s, err := os.Hostname();
    if(err != nil){
    	fmt.Println(err)
}
    fmt.Println(s)
}
