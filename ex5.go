package main
import "fmt"

func fibo(number int)int{
	if(number<=1){
		return 1;
	}
	return fibo(number-1)+fibo(number-2);
}
func main(){
	var number int
	fmt.Println("Enter the fibonacci number");
	fmt.Scanf("%d",&number);
	fmt.Printf("%d th Fibonacci number is %d",number,fibo(number));
}