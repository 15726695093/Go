package main

import (
	"fmt"
	"geektime_homework_error/dao"
)

func main() {
	user, err := dao.QueryUserById("123456")
	if err != nil {
		fmt.Printf("query user err : %+v", err)
		return
	}
	fmt.Println("query user : ", user)
}
