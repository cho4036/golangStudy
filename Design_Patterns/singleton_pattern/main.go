package main

import (
	"fmt"
	v1setting "golangStudy/Design_Patterns/singleton_pattern/setting"
)

func main() {
	setting := v1setting.GetInstance()
	setting.SetColor("red")
	fmt.Println(setting.Color())

	setting2 := v1setting.GetInstance()
	fmt.Println(setting2.Color())
}
