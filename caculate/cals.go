package main

import (
	"os"
	"test/mathpkg"
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numDay
)

func main() {
	os.Open("test.txt")
	mathpkg.Add(1, 2)
}
