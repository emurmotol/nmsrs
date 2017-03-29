package extract

import (
	"encoding/json"
	"fmt"
)

func Struct(i interface{}) {
	out, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
