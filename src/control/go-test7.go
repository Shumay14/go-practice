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






