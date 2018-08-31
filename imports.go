package main

import (
   "fmt"
   "math"
)
 
func main() {
    fmt.Printf("My max marks in maths is %g ", math.Sqrt(121))
    
    fmt.Println(swap("Hello" , "World"))	
} 

func swap(x,y string) (string, string)  {
	return y,x

}


func add(x int, y int) int {
	return x+y
}



