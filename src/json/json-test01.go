package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Task struct {
	Title  string
	Status status
}

type status int

const (
	UNKWON status = iota
	TODO
	DONE
)

func main() {
	ExampleTask_marshalJSON()
	ExampleTask_unmarshalJSON()
}

func ExampleTask_marshalJSON() {
	t := Task{
		"Coding",
		DONE,
	}
	b, err := json.Marshal(t)

}
