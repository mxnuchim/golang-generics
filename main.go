package main

import (
	"fmt"

	"github.com/mxnuchim/golang-generics/list"
)

type User struct {
	name string
	email string
	password string
}


func main() {
	// Golang generics example with linked list for any data type
	linkedList := list.New[int]()
	linkedList.PushFront(90)
	linkedList.PushFront(243)
	linkedList.PushFront(876) // add value to start (head) of linked list
	linkedList.PushBack(1034) // add value to end (tail) of linked list

	node := linkedList.Head()
	for node != nil {
		fmt.Println("value --> ", node.Value())
		node = node.Next()
	}

	// Real world use case with users //
	users := list.New[User]()
	users.PushBack(User{name: "Angela Simmons", email: "angela@example.com", password: "xxxx"})
	users.PushBack(User{name: "Manuchim Oliver", email: "manuchim@example.com", password: "xxxx"})
	users.PushBack(User{name: "Posha Alabi", email: "posha@example.com", password: "xxxx"})

	userNode := users.Head()
	for userNode != nil {
		fmt.Println("User name --> ", userNode.Value().name)
		userNode = userNode.Next()
	}
}