# hao-tool
a coding tool hub for go
# 开始使用
## 拉取
```shell
go get github/ricardosimple/hao-tool
```
## 使用
判断元素是否在数组中:
```go
	arr := []string{"faa", "bb", "", "cc"}
	str := "ccS"
	res := InList(arr, str)
```
传入自定义方法来判断元素是否在数组中：
```go
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
```
