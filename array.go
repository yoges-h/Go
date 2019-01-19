
// array in go 

package main 

import "fmt"

func  main(){

	 var a [2]string // a [2]string is an array of Strings 

	  a[0] = "Hello" // defining the data of array as "Hello"
	  a[1] = "world" // Defining the data of array as "world"

	  fmt.Println(a[0],a[1]) // Simply Printinig the Hello Wolrd
	  fmt.Println(a) // Printing the Hello World in  Bracket

	  primes := [6]int{2,3,5,7,11,13} // Defining the Array of Prime Numbers 
	  fmt.Println(primes) // Printing the data of prime numbers 

}