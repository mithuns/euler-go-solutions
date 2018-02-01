package main
import (
        "fmt"
        )

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var numCases int
    fmt.Scanf("%d", &numCases)
    
    var i=1
    for i<=numCases{
        var N int
        fmt.Scanf("%d",&N)
        doComputation(N)
        i=i+1
    }    
}
func doComputation(N int){
    var sum = 3*((N-1)/3)*(((N-1)/3)+1)/2
    sum = sum + 5*((N-1)/5)*(((N-1)/5)+1)/2
    sum = sum - 15*((N-1)/15)*(((N-1)/15)+1)/2
    fmt.Println(sum)
}
