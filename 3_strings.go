package main

import (
	"fmt"
	"strings"
)

func main() {
	m1 := "my name"
	m2 := "name"
	
	fmt.Println("m1 + m2 ==", m1 + m2)
	fmt.Println("strings.Contains(m1, m2) ==", strings.Contains(m1, m2))
	fmt.Println("strings.Split(m1, ' ') ==", strings.Split(m1, " "))
	fmt.Println("strings.ReplaceAll(m1, 'a', 'A') ==", strings.ReplaceAll(m1, "a", "A"))
}