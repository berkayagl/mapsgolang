package main

import "fmt"

func main() {

	mapData := map[string]string{
		"name":   "Berkay",
		"age":    "20",
		"status": "Busy",
	}

	for key, value := range mapData {
		fmt.Println("\nKey:", key, "\nValue:", value)
	}

}

// Çıktısı

// Key: name 
// Value: Berkay

// Key: age 
// Value: 20

// Key: status 
// Value: Busy
