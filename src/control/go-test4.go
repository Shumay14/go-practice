// string

package main

import (
	"fmt"
	s "strings"
)

// 기존 함수 변수에 재설정
var p = fmt.Println

func main() {
	p("Contains: ", s.Contains("test", "es")) // true
	p("Count: ", s.Count("test", "t"))        // 2
	// 앞글자들 뒷글자들 확인
	p("HasPrefix: ", s.HasPrefix("test", "te")) // true
	p("HasSuffix: ", s.HasSuffix("test", "st")) // true

	//
	p("Index: ", s.Index("test", "e"))
	p("Join: ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat: ", s.Repeat("a", 5))
	p("Replace: ", s.Replace("foo", "o", "0", -1))
	p("Replace: ", s.Replace("foo", "o", "0", 1))
	p("Split: ", s.Split("a-b-c-d-e", "-"))
	// 대소문자 변환
	p("ToLower: ", s.ToLower("TEST")) // test
	p("ToUpper: ", s.ToUpper("test")) // TEST
	p()
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])

}
