package main

import "fmt"

func main() {

	var m map[string]string

	m = make(map[string]string)

	m["name"] = "Berkay"

	fmt.Println(m)
}

// Çıktısı : map[name:Berkay]
