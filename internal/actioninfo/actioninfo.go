package actioninfo

import "fmt"

type DataParser interface {
	// TODO: добавить методы
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
// TODO: добавить методы	

for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Println(info)
	}
}
