package main

import (
	"fmt"
)

func main() {
	names := []string{"Tanmay Bakshi","Baheer Kamal","Kathy"}
	for i, v := range names {
		fmt.Println(i)
		fmt.Println(v)
	}
}
