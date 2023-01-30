package recoverpanic

import "fmt"

func Division(a, b int) float32 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperandome del panic ", r)
		}
	}()
	resp := validarDivision(a, b)
	return resp
}

func validarDivision(a, b int) float32 {
	if b == 0 {
		panic("🤦")
	}
	resp := float32(a / b)
	return resp
}
