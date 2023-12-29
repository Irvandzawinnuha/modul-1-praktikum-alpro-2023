package main

import (
	"fmt"
)

func jumlahBus(penumpang, kapasitas int) int {
	var a int
	a = penumpang / kapasitas
	if penumpang%kapasitas != 0 {
		a++
	}
	return a
}
func main() {
	var n, m int
    fmt.Scan(&n, &m)
    fmt.Println(jumlahBus(n, m),"bus")
}
