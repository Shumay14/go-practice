// struct 구조체 ** 매우 중요
// Custom Data Type을 표현하는데 사용
// 필드들의 모음 또는 묶음

// 구조체 이름 대문자 사용
// 구조체 내 변수들 대문자 사용
// 왜냐 외부에서도 사용할 것이기 때문

// 필드들의 모음 또는 묶음
// 배열과의 비교
// 배열: 서로 같은 자료형의 자료를 묶어놓은 것
// 구조체: 서로 다른 자료형의 자료들도 묶을 수 있는 것,
// 배열과 슬라이스 같은 동일한 자료형도 묶을 수 있다.

type Person struct {
	Name    string
	Age     int
	Id      string
	Address string
}

// 빈 객체 생성 후 필드 값 채우기
p := Person{}
p.Name = "Lee"
p.Age = 10
p.Id = "Sejong"
p.Address = "Sejong-si"

// 객체 생성 시 필드 값 할당
p2 := Person{Name: "Sean", Age: "25"}

// new() 사용
// new() 를 사용하면 모든 필드를 Zero value 로 초기화
// Person 객체 포인터(*Person) 리턴
p2 := new(Person)
p2.Name = "Shu"
p2.Age = 30

// Define the car structure, with 4 properties.
// Structure tags are used by encoding/json library
type Car struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Colour string `json:"colour"`
	Owner  string `json:"owner"`
}