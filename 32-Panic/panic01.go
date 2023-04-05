/* Alta3 Research | RZFeeser
   Calling the panic() function directly */
   
   package main

   import (
	   "fmt"
   )
   
   func jim() bool {
	panic("Jim, we have a problem.")
   }

   func main() {
		jim()
	   // panic produces a quick exit
	   panic("Jim, we have a problem.") // this line will be displayed
   
	   fmt.Println("You will not even see this line. The panic creates a fast fail.")
   }
   