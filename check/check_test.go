package check

import (
	"fmt"
)

func ExampleDo() {
	type User struct {
		Username string `json:"username" bind:"required" msg:"用户名不能为空"`
		Age      int    `json:"age" bind:"required" msg:"年龄不能为空"`
	}
	//Username不通过
	user := User{
		Username: "",
	}
	if err := Do(user); err != nil {
		fmt.Println(err)
	}

	//Age不通过
	user1 := User{
		Username: "abc",
		Age:      0,
	}
	if err := Do(user1); err != nil {
		fmt.Println(err)
	}

	// 校验通过
	user2 := User{
		Username: "abc",
		Age:      20,
	}
	if err := Do(user2); err == nil {
		fmt.Println("pass")
	}
	// output:
	// 用户名不能为空
	// 年龄不能为空
	// pass
}
