package main

import "fmt"

type user struct {
	name string
	email string
}

/*
	如果是要创建一个新值，该类型的方法就使用值接收者。如果是要修改当前值，就使用指针接收者
 */

//使用值接收者实现方法[调用使用值的副本执行]
func (u user) notify() {
	fmt.Printf("Call notify of user, name: %s, email: %s. \n", u.name, u.email)
}

//使用指针接收者实现方法
func (us *user) changeEmail(newEmail string)  {
	us.email = newEmail
}

type changer interface {
	changeEmail(s string)
}

func changeSomething(n changer, newEmail string)  {
	n.changeEmail(newEmail)
}

func main() {

	mike := user{"Mike", "mike@email.com"}
	//使用值接收者实现方法[调用使用值的副本执行]
	mike.notify()

	pika := &user{"pika", "pika@email.com"}
	//使用指针类型调用方法, go语言内部机制执行了转换
	// 实质上是(*pika).notify()
	pika.notify()

	//mike.changeEmail("abc@email.com")
	//mike.notify()

	//通过编译?
	changeSomething(&mike, "new@email.com")
	changeSomething(pika, "new@email.com")

	fmt.Println("After change...")
	mike.notify()
	pika.notify()
}