package main

import (
	"fmt"
)

type Phone interface {
	call()
	sendmail(content string)
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (nokiaPhone NokiaPhone) sendmail(c string) {
	fmt.Println("I am Nokia, sendmail !" + c)
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}
func (iPhone IPhone) sendmail(c string) {
	fmt.Println("I am iPhone, I can call sendmail!" + c)
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()
	phone.sendmail("aaa")
	phone = new(IPhone)
	phone.call()
	phone.sendmail("bbb")
}
