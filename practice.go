
/*package main
import "fmt"
type Person struct{
        name string
        job string
}
func main(){
        var aperson Person

        aperson.name="ANU"
        aperson.job="Doctor"

        fmt.Printf("aperson.name = %s\n",aperson.name)
        fmt.Printf("aperson.job=%s\n",aperson.job)
}*/

//Create a program that calculates the average
/*package main
import "fmt"
func main(){
        p := 10
        q := 20
        r := 30

        avg := (p+q+r)/3
        fmt.Println(avg)
}*/

//Create an array with the number 0 to 10
/*package main
import "fmt"
func main(){
        var arr = []int64{0,1,2,3,4,5,6,7,8,9,10}
        fmt.Println(arr)
}*/

//Take the string ‘hello world’ and slice it in two.

/*package main

import "fmt"

func main() {
    mystring := "hello world"
    hello := mystring[0:5]
    world := mystring[6:11]
    fmt.Println(hello)
    fmt.Println(world)
}*/

//Create a program with multiple string variables

/*package main

import "fmt"

func main() {
    s1 := "hello "
    s2 := "world"
    fmt.Println(s1,s2)
}*/

//Create an array of strings with names

/*package main

import "fmt"

func main() {
    var a = []string{ "A","B","C" }
    fmt.Println(a)
}*/

//Print "Hello World" if x is greater than y.

/*package main
import ("fmt")

func main() {
  var x = 50
  var y = 10
   if x>y{
    fmt.Print("Hello World")
   }
}*/

//Print i as long as i is less than 6.

/*package main
import "fmt"
func main(){
        for i:=0;i<6;i++{
                fmt.Println(i)
        }
}*/

//Use a for loop to print "Yes" 5 times

/*package main
import "fmt"
func main(){
        for i:=0;i<5;i++{
                fmt.Println("Yes")
        }
}*/

package main
import ("fmt")

func myFunction(x int) int {
  
     return 5 + x
}

func main() {
  fmt.Println(myFunction(3))
}