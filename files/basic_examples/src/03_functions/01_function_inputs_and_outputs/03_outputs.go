package main

import (
  "fmt"
)


// here we specify the return output to be a int datatype 
func plus(a int, b int) int {

    c := a + b
    
    return c
}


func main() {

    d := plus(1, 2)
    fmt.Println("The total is:", d)

    // the above 2 lines would be combined into a single line
    fmt.Println("The total is:", plus(1, 2))    

}
