package main

import "fmt"

/* fmt" is a package that provides I/O functions like Println (you can import it with import "fmt" ). */

type Vertex struct {

	X int
	Y int
}

func main(){
     
     v := Vertex{1,2}
     // if you define something here  then it will be get more priority 
     // then this.
     v.X = 4
     v.Y = 6

	fmt.Println(v.X,v.Y)
	
}