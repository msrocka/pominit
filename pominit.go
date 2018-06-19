package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {

	var dir string
	var artifact string

	if len(os.Args) > 1 {
		artifact = os.Args[1]
		dir = "./" + artifact + "/"
		os.MkdirAll(dir, os.ModePerm)
	} else {
		absPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		check(err)
		artifact = filepath.Base(absPath)
		dir = "./"
	}
	fmt.Println("Create project files", artifact, "in folder", dir)

	writeTemplateFile(dir+"pom.xml", pomTemplate,
		struct{ Artifact string }{artifact})

	writeFile(dir+".gitignore", gitignoreText)

	os.MkdirAll(dir+"src/main/java/"+artifact, os.ModePerm)
	os.MkdirAll(dir+"src/test/java/tests/"+artifact, os.ModePerm)
	writeFile(dir+"src/test/java/tests/ATest.java", jtestText)
}

func writeTemplateFile(path string, templ string, args interface{}) {
	t, err := template.New("pom").Parse(templ)
	check(err)
	var buffer strings.Builder
	t.Execute(&buffer, args)
	writeFile(path, buffer.String())
}

func writeFile(path string, content string) {
	if fileExists(path) {
		fmt.Println(path, "already exists")
		return
	}
	fmt.Println("Write file", path)
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
