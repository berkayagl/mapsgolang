package main

import "fmt"

func main() {

	mapData := map[string]string{
		"name":   "Berkay",
		"age":    "20",
		"status": "Busy",
	}

	l := len(mapData)

	fmt.Println(l)
}

// Çıktısı : 3
