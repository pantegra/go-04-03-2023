package main 

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("The number of arguments passed in:", len(os.Args))
}