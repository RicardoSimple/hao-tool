package lists

import (
	"fmt"
	"testing"
)

func Test_InList(t *testing.T) {
	arr := []string{"faa", "bb", "", "cc"}
	str := "ccS"
	res := InList(arr, str)
	fmt.Println(res)
}

type User struct {
	Id   int
	Name string
	Age  int
}

func Test_InListWithFn(t *testing.T) {
	users := []User{
		{
			Id:   1,
			Name: "Mike",
			Age:  20,
		},
		{
			Id:   2,
			Name: "Damin",
			Age:  20,
		},
	}
	user := User{
		Id:   2,
		Name: "Da2min",
		Age:  20,
	}
	res := InListWithFn(users, user, func(a, b User) bool {
		if a.Age == b.Age && a.Id == b.Id && a.Name == b.Name {
			return true
		}
		return false
	})
	fmt.Println(res)
}
