package check

import (
	"fmt"
)

func ExampleDo() {
	type User struct {
		Username string   `json:"username" binding:"required" msg:"用户名不能为空"`
		Age      int      `json:"age" binding:"required" msg:"年龄不能为空"`
		Tels     []string `binding:"required" msg:"手机不能为空"`
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

	//Tels不通过
	user2 := User{
		Username: "abc",
		Age:      20,
		Tels:     []string{""},
	}
	if err := Do(user2); err != nil {
		fmt.Println(err)
	}

	// 校验通过
	user4 := User{
		Username: "abc",
		Age:      20,
		Tels:     []string{"15888888888"},
	}
	if err := Do(user4); err == nil {
		fmt.Println("pass")
	}

	// output:
	// 用户名不能为空
	// 年龄不能为空
	// 手机不能为空
	// pass
}
