package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file := "a.txt"
	//defer os.Remove(file)
	read, _ := os.OpenFile(file, os.O_RDONLY|os.O_CREATE, 0644)
	write, _ := os.OpenFile(file, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)

	off1, _ := write.WriteString("123456789\n")
	fmt.Println(off1)
	read.Seek(3, 0)
	sc := bufio.NewScanner(read)
	sc.Scan()
	fmt.Println(sc.Text())

	off2, _ := write.WriteString("123456789\n")
	fmt.Println(off2)
	read.Seek(4, 0)
	sc.Scan()
	fmt.Println(sc.Text())
	write.Close()
}
