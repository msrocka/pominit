package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	kotlin := false
	if len(os.Args) > 1 {
		arg := os.Args[1]
		arg = strings.ToLower(arg)
		kotlin = arg == "kotlin" || arg == "kt"
	}
	pom(kotlin)
	writeFile(".gitignore", gitignoreText)
	if kotlin {
		os.MkdirAll("src/main/kotlin/app", os.ModePerm)
		os.MkdirAll("src/test/kotlin/tests", os.ModePerm)
	} else {
		os.MkdirAll("src/main/java", os.ModePerm)
		os.MkdirAll("src/test/java/tests", os.ModePerm)
	}
	test(kotlin)
	if kotlin {
		writeFile("src/main/kotlin/app/app.kt", kappText)
	}
}

func pom(kotlin bool) {
	if kotlin {
		writeFile("pom.xml", kpomText)
	} else {
		writeFile("pom.xml", jpomText)
	}
}

func test(kotlin bool) {
	if kotlin {
		path := "src/test/kotlin/tests/ATest.kt"
		writeFile(path, ktestText)
	} else {
		path := "src/test/java/tests/ATest.java"
		writeFile(path, jtestText)
	}
}

func writeFile(path string, content string) {
	if fileExists(path) {
		fmt.Println(path, "already exists")
		return
	}
	file, err := os.Create(path)
	check(err)
	defer file.Close()
	file.WriteString(strings.TrimLeft(content, "\n"))
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
