package assertiontype

import (
	"fmt"
	"strings"
)

func Exec(i interface{}) {
	// v, ok := i.(string)
	// if ok {
	// 	fmt.Println("Si es String")
	// }
	switch v := i.(type) {
	case string:
		fmt.Printf("tipo %T, valor %s \n", v, strings.ToUpper(v))
	case int:
		fmt.Printf("tipo %T, valor %d \n", v, (v * 3))
	default:
		fmt.Println("Tipo de dato no mapeado")
	}
}
