
/* A struct literal denotes newly allocated struct value by listing the values of its filed 	 */

package main 

import "fmt"

type Vertex struct{

	X,Y int

}
var (

	v1 = Vertex{1,2} // has  type vertex 
	v2 = Vertex{X:1} // y:0 is implicit 
	v3 = Vertex{}  // x:0 and y:0
    p = &Vertex{1,2} // has type *Vertex

)
func main(){

	fmt.Println(v1,p,v2,v3);

}