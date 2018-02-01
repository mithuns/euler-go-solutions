package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var numCases int
    fmt.Scanf("%d", &numCases)
    
    for i:=1;i<=numCases;i++{
        var N int
        fmt.Scanf("%d",&N)
        LargestPrimeFactor(N)
    }        
}
func LargestPrimeFactor(N int){
    var copyNumber = N
    var i=2
    for i<=copyNumber{
        if copyNumber%i==0{
            copyNumber = copyNumber/i
        }
        i=i+1
    }
    fmt.Println(i-1)
}
