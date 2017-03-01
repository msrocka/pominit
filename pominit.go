package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	pom()
	os.MkdirAll("src/main/java", os.ModePerm)
	os.MkdirAll("src/test/java/tests", os.ModePerm)
	test()
}

func pom() {
	if fileExists("pom.xml") {
		fmt.Println("pom.xml already exists")
		return
	}
	file, err := os.Create("pom.xml")
	check(err)
	file.WriteString(strings.TrimLeft(pomText, "\n"))
}

func test() {
	path := "src/test/java/tests/ATest.java"
	if fileExists(path) {
		fmt.Println(path, "already exists")
		return
	}
	file, err := os.Create(path)
	check(err)
	file.WriteString(strings.TrimLeft(testText, "\n"))
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		return false
	}
	return true
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
