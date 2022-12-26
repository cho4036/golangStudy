package singleton_pattern

import v1setting "golangStudy/Design_Patterns/singleton_pattern/setting"

func main() {
	setting := v1setting.GetInstance()
	setting.SetColor("red")

	setting2 := v1setting.GetInstance()
	setting2.Color()
}
