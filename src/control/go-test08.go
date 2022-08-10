// closure 클로저
// 함수 바깥에 있는 변수를 참조하는 함수값(function value)의미
// 외부의 변수를 마치 함수 안으로 끌어들인 듯이 그 변수를 읽거나 쓸 수 있게 함
package main

import "fmt"

func main() {
	//
	f := adder()
	fmt.Println(f(1))
	fmt.Println(f(20))
	fmt.Println(f(300))
}

// 외부함수 adder 함수에서 선언한 x 가 main 함수에 영향
// adder()입력인자 없고 --> func(int) int 반환 함수 등장
func adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
