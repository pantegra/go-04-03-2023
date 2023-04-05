
package main 

import (
    "fmt"
    "errors"
)

func testing() (string, int, error){

    return "tt", 99, errors.New("ffffff")

}



func main() {


    myname , _, _ := testing()
    fmt.Println(myname)
    var name = "Brian Erickson"
    var city, state = "Hagerman", "New Mexico"

    fmt.Println("My Name is:", name)
    fmt.Println("I live in ", city , ",", state)
}

