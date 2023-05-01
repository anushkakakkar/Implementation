/*package main
import "fmt"
func main(){
    var x float64
    x=15.2564
    y:=42
    fmt.Println(x)
    fmt.Println(y)
    fmt.Printf("y is of type %T",y)
}*/


//If-else
/*package main 
import "fmt"
func main(){
    var x int = 20
    if(x<20){
        fmt.Printf("x is less than 20 \n")
    }else{
        fmt.Printf("x is not less than 20 \n")
    }
}*/


//else-if 
/*package main
import "fmt"
func main(){
    
    score := 75
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else if score >= 70 {
    fmt.Println("C")
} else {
    fmt.Println("D")
}

}*/


//Switch Statement
/*package main
import "fmt"
func main(){
    day := "Sunday"

switch day {
case "Monday":
    fmt.Println("Today is Monday.")
case "Tuesday":
    fmt.Println("Today is Tuesday.")
case "Wednesday":
    fmt.Println("Today is Wednesday.")
case "Thursday":
    fmt.Println("Today is Thursday.")
case "Friday":
    fmt.Println("Today is Friday.")
case "Saturday":
    fmt.Println("Today is Saturday.")
case "Sunday":
    fmt.Println("Today is Sunday.")
default:
    fmt.Println("Invalid day.")
}


}*/


//For Loop
/*package main
import "fmt"
func main(){
    for i:=1;i<10;i++{
        fmt.Printf("value of i is %d\n",i)
    }
}*/


/*package main
import "fmt"
func main(){
    i:=1
    for i<10{
        fmt.Println(i)
        i++
    }
}*/


/*package main
import "fmt"
func main(){
    i:=1
    for{
        fmt.Println(i)
        i++
        if i>9{
            break
        }
    }
}*/


/*package main
import "fmt"
func main(){
    for i:=1;i<=5;i++{
        if i==3{
            break
        }
        fmt.Println(i)

    }
}*/

/*package main
import "fmt"
func main(){
    for i:=1;i<=5;i++{
        if i==3{
            continue
        }
        fmt.Println(i)
    }
}*/

/*
package main
import "fmt"
func main(){
    var a int = 100
    var b int = 200
    var ret int
    ret = max(a,b)
    fmt.Printf("max value is %d\n",ret)
}
func max(num1,num2 int)int{
    var result int
    if(num1>num2){
        result=num1
    }else{
        result=num2
    }
    return result
}*/


/*package main
import "fmt"
func printArray(arr[5]int){
    for i:=0;i<len(arr);i++{
        fmt.Println(arr[i])
    }

}
func main(){
    arr:=[5]int{1,2,3,4,5}
    printArray(arr)
}*/


/*package main
import "fmt"
func main(){
    var a int = 20
    var ip *int
    ip = &a 
    fmt.Printf("address of a variable %x\n",&a)
    fmt.Printf("address stored in ip variable %x\n",ip)
}*/

/*
package main
import "fmt"
type Books struct{
    title string
    id    int
}
func main(){
    var book1 Books
    var book2 Books

    book1.title="Hindi"
    book1.id=1

    book2.title="English"
    book2.id=2

    fmt.Printf("book1 title %s\n",book1.title)
    fmt.Printf("book2 title %d\n",book1.id)
}*/


/*package main
import "fmt"
func main(){
    arr:=[5]int{1,2,3,4,5}
    slc:=arr[0:3]
    fmt.Println(slc,len(slc))
}*/


/*package main
import "fmt"
func main(){
    var myMap map[string]int
    myMap = map[string]int{}
    myMap["a"]=12
    myMap["b"]=1200
    fmt.Println(myMap)
}*/

/*package main
import "fmt"
func main(){
    ch:=make(chan int)
    go func(){
        ch<-42

    }()
    val:=<-ch
    fmt.Println(val)
}*/




