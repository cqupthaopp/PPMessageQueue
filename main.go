package main

import (
	"PPMessageQueue/PPMessageQueue"
	"fmt"
)

type User struct {
	Username string
	Password string
}

func main() {

	ppQueue := PPMessageQueue.NewPPQueue(10, 10000)

	u1 := User{
		Username: "Hao_pp",
		Password: "123456",
	}

	u2 := User{
		Username: "test",
		Password: "123456",
	}

	ppQueue.PushData(u1)

	ppQueue.PushData(u2)

	fmt.Println(ppQueue.PopAllData())
}
