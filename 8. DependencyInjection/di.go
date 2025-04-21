package main

import (
	"bytes"
	"fmt"
)

type A struct {

}

type Swimmer interface {

}

func Greet(writer *bytes.Buffer, str string) {
	fmt.Fprintf(writer, "Hello, %s", str)
}

