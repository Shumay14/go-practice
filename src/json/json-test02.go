package main

import (
“encoding/json”
“fmt”
)

type Person struct {
Name string json:"name"
Age int json:"age"
Height int json:"height"
}

func main() {
a := new(Person)
a.Name = “GCLOUD”
a.Age = 24
a.Height = 175

fmt.Println(a)
marshal, _ := json.Marshal(a)

unmarshal := Person{}
_ = json.Unmarshal(marshal, &unmarshal)

fmt.Println(string(marshal))
fmt.Println(unmarshal)

}

fmt.Println(a)
marshal, _ := json.Marshal(a)

unmarshal := Person{}
_ = json.Unmarshal(marshal, &unmarshal)

fmt.Println(string(marshal))
fmt.Println(unmarshal)