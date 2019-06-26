package utils

import (
	"encoding/json"
	"fmt"
)

// JSONPrettyPrint 美化 json 输出
func JSONPrettyPrint(i interface{}) {
	str, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		fmt.Println("JSONPrettyPrint error: " + err.Error())
	}

	fmt.Println(str)
}
