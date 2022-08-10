cd go/
mkdir pkg bin
cd src/
mkdir study
cd study/

--------------first.go-------------------

package main

import "fmt"

func main() {
        fmt.Println("Hello World")
}

----------------------------------------

mkdir data
cd data

-----------datatype.go-------------------

package main

import "fmt"

var x, y, z int
var c, python, java bool

func main() {
        fmt.Println(x, y, z, c, python, java)
}


----------------------------------------

mkdir const
cd const

-----------const.go--------------------

package main

import "fmt"

const Pi = 3.14

func main() {
        const World = "ì•ˆë…•"
        fmt.Println("Hello", World)
        fmt.Println("Happy", Pi, "Day")
        const Truth = true
        fmt.Println("Go rules?", Truth)
}

----------------------------------------

mkdir pointer
cd pointer

-------------pointer.go-----------------

package main

import "fmt"

func main() {
        var num int = 1
        var numPtr *int = &num
        fmt.Println(numPtr)
        fmt.Println(*numPtr)
        fmt.Println(&num)

}

------------------------------------------


------------data.go----------------------

package main
import (
        "fmt"
        "math/cmplx"
)

var (
        ToBe bool = false
        MaxInt uint64 = 1<<64 - 1
        z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
        const f = "%T(%v)\n"
        fmt.Printf(f, ToBe, ToBe)
        fmt.Printf(f, MaxInt, MaxInt)
        fmt.Printf(f, z, z)
}

----------------------------------------------

mkdir func
cd func

mkdir func_add

-----------------func_add.go------------------

package main

import "fmt"

func add(x int, y int) int {
        return x + y
}

func main() {
        fmt.Println(add(42, 13))
}

-------------------------------------------------

mkdir control
cd control

mkdir if 

--------------if.go--------------------------------

package main

import (
        "fmt"
        "math"
)

func pow(x, n, lim float64) float64 {
        if v := math.Pow(x, n); v < lim {
        return v


        } else {
                fmt.Printf("%g >= %g\n", v, lim)
        }
        return lim
}

func main() {
        fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
        )
}

-------------------------------------------------


mkdir for

---------------for.go----------------------------

package main

import "fmt"

func main() {
        sum := 0

        for i := 0  ;   i < 10  ; i++   {

                sum += i
        }

        fmt.Println(sum)
}

--------------------------------------------------



---------------if_for.go-----------------------

package main

import "fmt"

func printhello(x int) (int) {
        for i := 0 ; i < 10 ; i++ {
                if i%2 == 0 {
                        fmt.Println("Hello World")
                }
        }
        return x
}

func main() {
        printhello(5)
}

---------------------------------------------------

mkdir str
cd str

mkdir add_str

-------------add_str.go------------------------

package main

import "fmt"

func main() {
        s := "abc"

        ps := &s

        s += "def"

        fmt.Println(s)
        fmt.Println(*ps)
}

--------------------------------------------------



-------------------str_func.go--------------------

package main

import (
        "fmt"
        s  "strings"
)

var p = fmt.Println

func main() {
        p("Contains: ", s.Contains("test", "es"))
        p("Count: ", s.Count("test", "t"))
        p("HasPrefix: ", s.HasPrefix("test", "te"))
        p("HasSuffix: ", s.HasSuffix("test", "st"))
        p("Index: ", s.Index("test", "e"))
        p("Join: ", s.Join([]string{"a", "b"}, "-"))
        p("Repeat: ", s.Repeat("a", 5))
        p("Replace: ", s.Replace("foo", "o", "0", -1))
        p("Replace: ", s.Replace("foo", "o", "0", 1))
        p("Split: ", s.Split("a-b-c-d-e", "-"))
        p("ToLower: ", s.ToLower("TEST"))
        p("ToUpper: ", s.ToUpper("test"))
        p()
        p("Len: ", len("hello"))
        p("Char:", "hello"[1])
}

-----------------------------------------------------------


mkdir slice
cd slice
mkdir slice
cd slice

-----------------slice.go----------------------------

package main

import "fmt"

func main() {
        a := make([]int, 5)
        printSlice("a", a)
        b := make([]int, 0, 5)
        printSlice("b", b)
        c := b[:2]
        printSlice("c", c)
        d := c[2:5]
        printSlice("d", d)
}

func printSlice(s string, x []int) {
        fmt.Printf("%s len=%d cap=%d %v\n",
        s, len(x), cap(x), x)
}

------------------------------------------------------------


-----------------------slice_append.go--------------------

package main

import "fmt"

func main() {
        f1 := []string{"ì‚¬ê³¼", "ë°”ë‚˜ë‚˜", "í† ë§ˆí† "}
        f2 := []string{"í¬ë„", "ë”¸ê¸°"}
        f3 := append(f1, f2...) // ì´ì–´ë¶™ì´ê¸°

        // í† ë§ˆí† ë¥¼ ì œì™¸í•˜ê³  ì´ì–´ë¶™ì´ê¸°
        f4 := append(f1[:2], f2...)
        fmt.Println(f1)
        fmt.Println(f2)
        fmt.Println(f3)
        fmt.Println(f4)
}

-----------------------------------------------------------------


mkdir map
cd map
mkdir map
cd map


---------------------run map_exam.go-----------------------

package main

import "fmt"

func main() {
        var m map[int]string
        m = make(map[int]string)
        //ì¶”ê°€ í˜¹ì€ ê°±ì‹ 
        m[901] = "Apple"
        m[134] = "Grape"
        m[777] = "Tomato"

        // í‚¤ì— ëŒ€í•œ ê°’ ì½ê¸°
        str := m[134]
        fmt.Println(str)

        noData := m[999] // ê°’ì´ ì—†ìœ¼ë©´ nil í˜¹ì€ zero ë¦¬í„´
        fmt.Println(noData)

        // ì‚­ì œ
        map777 := m[777]
        fmt.Println(map777)
        delete(m, 777)
        del := m[777]
        fmt.Println(del)
}

------------------------------------------------------------------


------------------------map.go------------------------------

package main

import "fmt"

func main() {
        myMap := map[string]string{
                "A": "Apple",
                "B": "Banana",
                "C": "Charlie",
        }
        // for range ë¬¸ì„ ì‚¬ìš©í•˜ì—¬ ëª¨ë“  ë§µ ìš”ì†Œ ì¶œë ¥
        for key, val := range myMap {
                fmt.Println(key, val)
        }
}

---------------------------------------------------------


cd func
mkdir va_arg
cd va_arg

--------------------------va_arg.go-------------------------

package main

import "fmt"

func main() {
        say("This", "is", "a", "book")
        say("Hi")
}

func say(msg ...string) {
        for _, s := range msg {
                fmt.Println(s)
        }
}

----------------------------------------------------------------


mkdir re_mul_value
cd re_mul_value

-------------re_mul_value.go--------------------------------

package main

import "fmt"

func main() {
        count, total := sum(1, 7, 3, 5, 9)

        fmt.Println(count, total)
}

func sum(nums ...int) (int, int) {
        s := 0 // í•©ê³„
        count := 0 // ìš”ì†Œ ê°¯ìˆ˜
        for _, n := range nums {
                s += n
                count++
        }
        return count, s
}

---------------------------------------------------------------


mkdir anonymous_func
cd anonymous_func


--------------anonymous_func.go--------------------------

package main

import "fmt"

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

-----------------------------------------------------------------

mkdir first-class_func
cd first-class_func

------------------first-class_func.go--------------------------

package main

import "fmt"

func main() {// ë³€ìˆ˜ add ì— ìµëª…í•¨ìˆ˜ í• ë‹¹
        add := func(i int, j int) int {
                return i + j
        }

        r1 := calc(add, 10, 20) // add í•¨ìˆ˜ ì „ë‹¬
        fmt.Println(r1)

        // ì§ì ‘ ì²«ë²ˆì§¸ íŒŒë¼ë¯¸í„°ì— ìµëª…í•¨ìˆ˜ ì •ì˜
        r2 := calc(func(x int, y int) int {return x - y }, 10, 20)
                fmt.Println(r2)
}

func calc(f func(int, int) int, a int, b int) int {
        result := f(a, b)
        return result
}

------------------------------------------------------------------


mkdir closure
cd closure


--------------------closure.go------------------------------------

package main

import "fmt"

func main() {
        f := adder()
        fmt.Print(f(1))
        fmt.Print(f(20))
        fmt.Print(f(300))
}

func adder() func(int) int {
        var x int
        return func(delta int) int {
                x += delta
                return x
        }
}


------------------------------------------------------------------


------------func.go-------------------------------------

package main

import "fmt"

// returnì´ ì—†ëŠ” ê²½ìš°
func printAdd(x int, y int) {
        fmt.Println(x + y)
}

// return ê°’ì´ í•˜ë‚˜ì¸ ê²½ìš°
func add(x int, y int) int {
        return x + y
}

// return ê°’ì´ ë‘ ê°œ ì´ìƒì¸ ê²½ìš°
func addAndMultiply(x int, y int) (int, int) {
        return x + y, x * y
}

func main() {
        fmt.Println(add(42, 13))
        printAdd(42, 13)
        fmt.Println(addAndMultiply(42, 13))
}

----------------------------------------------------------------

mkdir struct
cd struct

--------------------struct.go-----------------------------

package main

import "fmt"

// struct ì •ì˜
type person struct {
        name string
        age int
}

func main() {
        // ë¹ˆ ê°ì²´ ìƒì„± í›„ í•„ë“œê°’ ì±„ìš°ê¸°
        p := person{}
        p.name = "Lee"
        p.age = 10

        // ê°ì²´ ìƒì„± ì‹œ í•„ë“œê°’ í• ë‹¹
        p2 := person{name:"Sean", age: 50}

        // new() ì‚¬ìš©í•˜ì—¬ ê°ì²´ ë§Œë“¤ê¸°
        p3 := new(person)
        p3.name = "Lee" // p3ê°€ í¬ì¸í„°ë¼ë„ . ì„ ì‚¬ìš©í•œë‹¤

        fmt.Println(p)
        fmt.Println(p2)
        fmt.Println(p3)
}

-------------------------------------------------------------


mkdir json
cd json

------------------------json.go----------------------------

package main

import (
        "encoding/json"
        "fmt"
        "log"
)

type Task struct {
        Title string
        Status status
}

type status int

const (
        UNKNOWN status = iota
        TODO
        DONE
)

func main() {
        ExampleTask_marshalJSON()
        ExampleTask_unmarshalJSON()
}

func ExampleTask_marshalJSON() {
        t := Task{
                "Laundry",
                DONE,
        }

        b, err := json.Marshal(t)

        if err != nil {
                log.Println(err)
                return
        }

        fmt.Println(string(b))
}

func ExampleTask_unmarshalJSON() {
        b := []byte(`{"Title":"Buy Milk","Status":2}`)
        t := Task{}
        err := json.Unmarshal(b, &t)

        if err != nil {
                log.Println(err)
                return
        }

        fmt.Println(t.Title)
        fmt.Println(t.Status)
}

--------------------------------------------------------------

mkdir method
cd method

-------------------method.go--------------------------------

package main

import "fmt"

type Rect struct { //Rect - struct ì •ì˜
        width, height int
}

func (r Rect) area() int { //Rectì˜ area() ë©”ì†Œë“œ
        return r.width * r.height
}

func (r *Rect) area2() int { // í¬ì¸í„° Receiver
        r.width++
        return r.width * r.height
}

func main() {
        rect := Rect{10, 20}
        area1 := rect.area() //ë©”ì„œë“œ í˜¸ì¶œ
        area2 := rect.area2() //ë©”ì„œë“œ í˜¸ì¶œ
        fmt.Println(area1)
        fmt.Println(rect.width, area2)
}

----------------------------------------------------------


mkdir interface
cd interface

------------------interface.go-----------------------------

package main

import (
        "fmt"
        "math"
)

type Shape interface {
        area() float64
        perimeter() float64
}

type Rect struct { //Rect ì •ì˜
        width, height float64
}
type Circle struct { //Circle ì •ì˜
        radius float64
}
//Rect íƒ€ìž…ì— ëŒ€í•œ Shape ì¸í„°íŽ˜ì´ìŠ¤ êµ¬í˜„
func (r Rect) area() float64 {
        return r.width * r.height
}
func (r Rect) perimeter() float64 {
        return 2 * (r.width + r.height)
}
//Circle íƒ€ìž…ì— ëŒ€í•œ Shape ì¸í„°íŽ˜ì´ìŠ¤ êµ¬í˜„
func (c Circle) area() float64 {
        return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
        return 2 * math.Pi * c.radius
}

func main() {
        r := Rect{10., 20.}
        c := Circle{10}
        showArea(r, c)
}

func showArea(shapes ...Shape) {
        for _, s := range shapes {
                a := s.area() //ì¸í„°íŽ˜ì´ìŠ¤ ë©”ì„œë“œ í˜¸ì¶œ
                fmt.Println(a)
        }
}

---------------------------------------------------------