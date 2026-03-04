package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	var err error
	if len(dataset) == 0 {
		return
	}
	for _, v := range dataset {

		err = dp.Parse(v)

		if err != nil {
			log.Println(err)
		}
	}
	actInfo, err := dp.ActionInfo()

	fmt.Println(actInfo)
}
