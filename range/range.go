package _range

import "fmt"

// Topic1 .
// A. &{Name:张三 Age:20} &{Name:李四 Age:21}
// B. {Name:张三 Age:20} {Name:李四 Age:21}
// C. &{Name:李四 Age:21} &{Name:李四 Age:21}
// D. &{Name:张三 Age:20} &{Name:张三 Age:20}
func Topic1() {
	type user struct {
		Name string
		Age  int
	}
	users := []*user{{Name: "张三", Age: 20}, {Name: "李四", Age: 21}}
	usersCopy := make([]*user, 0, len(users))
	for _, u := range users {
		usersCopy = append(usersCopy, u)
	}

	for _, name := range usersCopy {
		fmt.Printf("%+v ", name)
	}
}

// Topic2 .
// A. &{Name:张三 Age:20} &{Name:李四 Age:21}
// B. {Name:张三 Age:20} {Name:李四 Age:21}
// C. &{Name:李四 Age:21} &{Name:李四 Age:21}
// D. &{Name:张三 Age:20} &{Name:张三 Age:20}
func Topic2() {
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
