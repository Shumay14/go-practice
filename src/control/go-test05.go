// array
package main

import "fmt"

// 3개의 정수로 되어있는 ar이라는 이름의 배열을 구성
var ar [3]int

func Fruits() {
	// 3개의 문자열을 사과 바나나 토마토 로 구성된 fruits라는 배열 구성
	fruits := [3]string{"사과", "바나나", "토마토"}

	// 컴파일러가 배열의 개수를 알아내서 '...'를 '3'으로 생성
	fruits := [...]string{"사과", "바나나", "토마토"}

	// slice
	// 배열에 기초하여 만들어짐
	// 배열에 비해 유연한 구조를 가지고 있어 배열보다 자주 사용
	// 슬라이스는 길이와 용량을 갖고 길이가 변할 수 있는 구조
	// 배열 사용 방법

	var ar []int // 크기가 없는 배열

	// n개의 문자열 공간을 할당한 'fruits'라는 슬라이스 생성
	fruits := make([]string, n)

	// 슬라이스 -> 잘라내기
	fruits[:2]
	// 2개를 잘라서 다시 frutis에 할당
	fruits = fruits[:2]

}

func main() {
	fmt.Println("Fruits")
	Fruits()
}
