package main

import "fmt"

func main() {

	mapData := map[string]string{
		"name":   "Berkay",
		"age":    "20",
		"status": "Busy",
	}
	delete(mapData, "status")

	fmt.Println(mapData)
}

// Çıktısı : map[age:20 name:Berkay]
