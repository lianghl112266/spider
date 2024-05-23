package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

type Person struct {
	Name   string
	Age    int
	Action Run
}

type Run struct {
	Speed int
}

func Test(t *testing.T) {
	var dao bytes.Buffer

	var encoder = gob.NewEncoder(&dao)
	var decoder = gob.NewDecoder(&dao)

	p := Person{Name: "chen", Age: 18, Action: Run{80}}

	err := encoder.Encode(&p)
	if err != nil {
		panic(err)
	}

	fmt.Println(dao.String())

	var d Person
	err = decoder.Decode(&d)
	if err != nil {
		panic(err)
	}

	fmt.Println(d)

}
