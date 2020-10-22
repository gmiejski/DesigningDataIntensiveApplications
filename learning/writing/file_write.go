package main

import (
	"fmt"
	"os"
)

func main() {

	file := "a.txt"
	//defer os.Remove(file)
	write, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)

	write.WriteString("123456789\n")
	off1, err := write.WriteString("123456789\n")
	fmt.Println(off1)
	if err != nil {
		fmt.Println(err.Error())
	}
	write.Close()
}
