// JSON
// 직렬화란?
// 객체의 상태를 보관이나 전송 가능한 상태로 변환하는 것
// 바이트코드로 변환 ? bytes
// 구조체 -> JSON
 

/***** 직렬화 가능한 요소
구조체, 
숫자, 문자열과 같은 다양한 데이터 타입의 변수 또는 상수
배열, 맵 
*****/


/***** 사용 분야
저장 및 불러오기
네트워크를 통한 메시지 전송
(ex. RPC(Remote Procedure Call)) 
*****/

// JavaScript Object Notation (JSON)
// 자료 교환 형식 중 하나
// 자바스크립트에서 객체를 표현하는 방식과 비슷함
// XML에 비하여 사람이 읽기 쉽고 간단하기 때문에 현재 많이 사용됨


{
	"Title": "Laundry",
	"Status": 2,
}


// Fabcar 예제에서 JSON
// 인코딩에 제이슨이라는 라이브러리 사용
//*** import("encoding/json") ***//
// JSON을 읽고 쓰기 때문에 import 시킴

//*** json.Marshal() ***//
// carAsBytes, _ = json.Marshal(car)
// car 구조체의 데이터를 JSON 형식으로 저장
// '_'으로 에러는 처리하지 않음

//*** json.Unmarshal() ***//
// json.Unmarshal(carAsBytes, &car)
// 전달받은 'carAsBytes'라는 JSON형식의 데이터를 car 구조체 안의 값으로 채움



// 기존 구조체 필드 이름을 JSON에서 다른 이름으로 사용하고 싶을때 사용
// JSON에서 필드를 나열하고 싶지 않을 때
// ex. Internal string `json:"-"`

// Define the car structure, with 4 properties.
// Structure tags are used by encoding/json library
type Car struct {
	Make string `json:"make"`
	Model string `json:"model"`
	Colour string `json:"colour"`
	Owner string `json:"owner"`
}