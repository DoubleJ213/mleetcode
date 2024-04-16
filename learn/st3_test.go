package learn

import (
	"fmt"
	"testing"
	"time"
)

type user struct {
	name string
	age  int8
}

var u = user{name: "Ankur", age: 25}
var g = &u

func modifyUser(pu *user) {
	fmt.Println("modifyUser Received Vaule", pu)
	pu.name = "Anand"
}

func printUser(u <-chan *user) {
	time.Sleep(2 * time.Second)
	fmt.Println("printUser goRoutine called", <-u)
}

func TestSt3(t *testing.T) {
	c := make(chan *user, 5)
	c <- g
	fmt.Println(g)
	// modify g
	g = &user{name: "Ankur Anand", age: 100}
	go modifyUser(g)
	//Remember all transfer of value on the go channels happens with the copy of value.
	go printUser(c)
	time.Sleep(5 * time.Second)
	fmt.Println(g)
	//c <- g
	//go printUser(c)
	//time.Sleep(5 * time.Second)
	//https://golang.design/go-questions/channel/principal/
}
