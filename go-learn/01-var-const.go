// 변수는 Go 키워드 var 를 사용하여 선언한다. var 키워드 뒤에 변수명을 적고, 그 뒤에 변수타입을 적는다.
// 예를 들어, 아래는 a 라는 정수(int)변수를 선언한 것이다.

// Go 언어 예약 키워드
/**
  break 		default 	 func 		interface 	select
  case 		defer		 go 		map			struct
  chan 		else		 goto		package		switch
  const		fallthrough  if			range		type
  continue	for			 import		return	 	var
  **/

// 코드 작성시 사용하지 않은 변수는 프로그램에서 삭제됨

// 1. 변수 variable var
var a int // var + 변수명 + type
var f float32 = 11

a = 10
f = 12.0

// 변수 선언시 초기값을 지정하지 않으면 Zero Value 할당된다.
// int 0, string "", bool false
var i, j, k int
var i, j, k int = 1, 2, 3

var i = 1 var s = "Hi"

// 아래와 같은 선언은 함수 내에서만 사용 가능하며, 함수 밖에서는 var 사용해야한다.
// Go에서 일반 변수와 상수는 함수 밖에서도 사용가능
i := 1

// 새로운 변수 선언 방식
// Short Assignment Statement  
// := 


// 2. 상수 constant const
// const 변수명 상수타입 = 상수값
const c int = 10
const s string = "Hello"
// Go 에서는 상수 타입을 적지 않아도 자동으로 타입을 추정한다.

// 묶음으로 상수 지정
const (
	Visa = "Visa"
	Master = "MasterCard"
	Amex = "American Express"

)

// TIP
// 상수값을 0부터 순차적으로 부여하기 위해 iota라는 identifier를 사용할 수 있다.
// 이 경우 iota가 지정된 Apple 에는 0이 할당 되고, 나머지 상수들을 순서대로 1씩 증가된 값을 부여받는다.
const (
	Apple = iota // 0
	Grape		 //	1
	Orange 		 // 2
)

