package main

import "fmt"

var x, y, z int


var c, python, java bool



const Pi = 3.14

func main() {
	fmt.Println(x, y, z, c, python, java)

	const World = "Hi"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}


// 배열 슬라이스
// 하나의 변수가 "같은타입"의 여러개의 값을 가지는 자료형 집합

// p := []int{2, 3, 5, 7, 11}

// 구조체
// 구조체는 첫글자 대문자
// 내부 데이터도 첫글자 대문자
// struct 는 필드(데이터)들의 조합
// 여러개의 자료형을 묶어 하나의 변수로 만들어주는 자료형
type Vertex struct {
	X int
	Y int
	Name string
}

// 맵 키 값 자료형
// key 값은 string
// var m map[string]Vertex
// m["Bell Labs"] = Vertex{}


// func swap(x, y string) 

func add(x int, y int) int {
	return x + y
}


func swap(x, y string) (string, string) {
	return y, x
}
