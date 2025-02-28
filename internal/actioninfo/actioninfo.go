package actioninfo

import "fmt"

// создайте интерфейс DataParser
type DataParser interface {
	Parse(string) error
	ActionInfo() string
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for _, val := range dataset {
		err := dp.Parse(val)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		dp.ActionInfo()
	}
}
