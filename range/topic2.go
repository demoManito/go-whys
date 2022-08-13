package _range

import "fmt"

type topic2 struct {
}

// A. &{Name:张三 Age:20} &{Name:李四 Age:21}
// B. &{Name:张三 Age:20} &{Name:张三 Age:20}
// C. &{Name:李四 Age:21} &{Name:李四 Age:21}
// D. 编译错误
func (*topic2) t2() {
	type user struct {
		Name string
		Age  int
	}
	users := []user{{Name: "张三", Age: 20}, {Name: "李四", Age: 21}}
	usersCopy := make([]*user, 0, len(users))
	for _, u := range users {
		usersCopy = append(usersCopy, &u)
	}

	for _, u := range usersCopy {
		fmt.Printf("%+v ", u)
	}
}
