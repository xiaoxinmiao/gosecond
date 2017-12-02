package main

import (
	"encoding/json"
	"fmt"
	"gosecond/p1"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	//first()
	//first2()
	first3()
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.GET("/pay", p1.Pay)

	e.Start(":5000")
}

func first3() {

	a := `{
		"name":"apple",
		"price":18
	}`

	fruit := p1.Fruit{}
	err := json.Unmarshal([]byte(a), &fruit)
	if err != nil {

	}
	fmt.Printf("%+v", fruit)

}

func first2() {
	fruit := p1.Fruit{}
	fruit.Name = "li"
	fruit.Prict = 20
	fmt.Println(fruit)

	fmt.Printf("%+v", fruit)
	fmt.Printf("%#v", fruit)

	fruit2 := new(p1.Fruit)
	fruit2.Name = "xiao"
	fruit2.Prict = 3
	fmt.Printf("%+v", fruit2)

	fmt.Println(Buy(&fruit))
	fmt.Println(Buy2(*fruit2))
	fmt.Println((*fruit2).Sale())

	fmt.Println(fruit2.Sale())
	fmt.Println(fruit2.Sale2())

}

func Buy(fruit *p1.Fruit) (money float64) {
	return fruit.Prict * 10
}

func Buy2(fruit p1.Fruit) (money float64) {
	return fruit.Prict * 10
}

func first() {
	fmt.Println("hello world")

	var name string = "li"
	fmt.Println(name)

	age := 18
	fmt.Println(age)

	slice1 := []int{1, 2, 3}
	for index := 0; index < len(slice1); index++ {
		fmt.Println(slice1[index])
	}

	for k, v := range slice1 {
		fmt.Println(k, v)
	}

	fmt.Println(slice1[0:1])

	fmt.Printf("%T,%v", slice1[0:1], slice1[0:1])

	var speed int64 = 10
	var speed1 int = 11
	speed = int64(speed1)
	fmt.Println(speed)

	fmt.Println("xiao" + strconv.FormatInt(speed, 10))

	flower := "11"
	i, err := strconv.ParseInt(flower, 10, 64)
	if err != nil {
		return
	}
	fmt.Println(i, err)

	switch flower {
	case "11":
		fmt.Println("get it1")
		//fallthrough
	case "12":
		fmt.Println("get it2")
	default:
		fmt.Println(11)
	}

	names := []string{"xiao", "li"}
	for k, v := range names {
		fmt.Println(k, v)
	}

	person := make(map[string]interface{}, 0)
	person["name"] = "li.kun"
	person["age"] = 18

	person2 := map[string]interface{}{
		"name": "xiao.xinmiao",
		"age":  18,
	}

	for k, v := range person {
		fmt.Println(k, v)

	}

	for k, v := range person2 {
		fmt.Println(k, v)

	}

	names = append(names, "xia")
	names = append(names[:1], names[2:]...)
	names[0] = "123"

	person["price"] = 1.5
	delete(person, "name")

	fmt.Println(p1.Run("xoap"))
}
