package

import "fmt"


// map

var m map[keyType]valueType
var m map[string]string

// map 에 자료 넣기
// make 함수 사용
// literal 사용 -> 소스코드의 고정된 값으 대표하는 용어

idMap = make(map[int]string)
tickers := map[string]string{
	"GOOG": "Google Inc",
	"MSFT": "Microsoft",
	"FB":   "FaceBook",
}


// 함수
// 함수 선언 방법
// 소문자는 파일내에서만 대문자는 내외부로 사용할 경우

func square(f float64) float64 { return f*f }

func MySqrt(f float64) (float64, bool) {
	if f >= 0 {
		return math.Sqrt(f), true
	}
	return 0, false
}



// 반환값 없는 함수

// 가변인자함수

func say(msg ...string) {

}

func main() {
	say("This", "is", "a", "book")
	say("Hi")
}

/// range: for반복문에서 슬라이스나 냅을 순회(iterates)할 때 사용

func say(msg ...string) {
	for _, s := range msg {
		fmt.Println(s)
	}
}

// 반환값이 있는 함수
func sum(nums ...int) int { return s}

// 복수 반환값이 있는 함수
// 가변 인자를 나타내는 ... 사용
// 가변 인자를 갖는 함수 호출 시 n개의 동일 타입 인자 전달
func main() {
	count, total := sum(1, 4, 7, 9)
	fmt.Println(count, total)
}

func sum(nums ...int) (int, int) {
	s := 0 // 합계
	count := 0 // 요소 갯수
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}



// 익명함수

func main() {
	sum := func(n ...int) int { 
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}
	result := sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}





// 일급함수
// 다른 함수의 인자로 전달
// 다른 함수의 반환 값으로 이용

func main() {
	// 변수 add 에 익명함수 할당
	add := func(i int, j int) int {
		return i + j
	}
	r1 := calc(add, 10, 20) // add 함수 전달
	fmt.Println(r1)

	// 직접 첫번째 파라미터에 익명함수 정의
	r2 := calc(func(x int, y int) int {
		return x - y
	}, 10, 20)
	fmt.Println(r2)
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

// 함수 원형 정의
// type 문

// 구조체(struct), 인터페이스 등 Custom Type
// (혹은 User Defined Type) 정의
// 또한 함수 원형을 정의하는데 사용 가능
// 변수명이 함수명과 같이 취급
// "변수명(인자들)" 형식으로 함수 호출

// 원형 정의 선언
type calculator func(int, int) int

// 사전에 정의해놔서 함수 내 여러가지 형식 설정할 필요 없음
// calculator 원형 사용 그대로 가져와서 씀
func calc(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}




