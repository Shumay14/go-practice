// if & for
package main

// fmt 터미널 관리
import "fmt"

// 내부에서만 사용될 함수이기 때문에 소문자로 사용
func printhello(x int) int {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Hello World")
		}	
	} 
	return x	
} 

func main () {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	printhello(5)
}

