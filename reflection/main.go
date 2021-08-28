package main

import (
	"fmt"
	"log"
	"reflect"
)

type order struct {
	ordId, customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(o order) string {
	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
	return i
}

func createNewQuery(q interface{}) string {
	var (
		t = reflect.TypeOf(q)
		v = reflect.ValueOf(q)
		k = t.Kind()
	)
	fmt.Printf("Type %v\n", t)
	fmt.Printf("Value %v\n", v)
	fmt.Printf("Value %v\n", k)
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		fmt.Println("Number of Fields", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s value(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s %d", query, v.Field(i).Int())
				}

			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("unsupported type")
				return ""
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return query
	}
	fmt.Println("unsupport type")
	return ""
}

func main() {
	log.Println("Reflection in Golang")
	name := "Gwen Stacy"
	fmt.Printf("%v %T\n", name, name)

	x := reflect.ValueOf(name).String()
	fmt.Printf("Type %v, Value %v", x, x)

	a := 99

	y := reflect.ValueOf(a).Int()
	fmt.Printf("Type %v, Value %v", y, y)

	o := order{
		ordId:      777,
		customerId: 888,
	}
	fmt.Println(o)

	fmt.Println(createQuery(o))

	emp := employee{
		name:    "Gwen Stacy",
		id:      1,
		address: "Greenwich",
		salary:  999999,
		country: "UK",
	}

	fmt.Println(emp)

	fmt.Println(createNewQuery(emp))

	fmt.Println(createNewQuery(99))
}
