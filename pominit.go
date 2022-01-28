package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

func main() {

	var dir string
	var artifactName string

	if len(os.Args) > 1 && os.Args[1] != "." {
		artifactName = os.Args[1]
		dir = "./" + artifactName + "/"
		os.MkdirAll(dir, os.ModePerm)
	} else {
		absPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		check(err)
		artifactName = filepath.Base(absPath)
		dir = "./"
	}
	fmt.Println("Create project files", artifactName, "in folder", dir)

	// create the template parameters
	artifactParam := struct{ Artifact string }{artifactName}
	packageName := bytes.Buffer{}
	packageDir := bytes.Buffer{}
	for _, char := range artifactName {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			packageName.WriteRune(char)
			packageDir.WriteRune(char)
		} else {
			packageName.WriteRune('.')
			packageDir.WriteRune('/')
		}
	}
	packageParam := struct{ Package string }{packageName.String()}

	// resources in main folder: pom.xml; gitignore
	writeTemplateFile(dir+"README.md", readmeTemplate, artifactParam)
	writeTemplateFile(dir+"pom.xml", pomTemplate, artifactParam)
	writeFile(dir+".gitignore", gitignoreTemplate)
	writeFile(dir+".editorconfig", editorConfigTemplate)
	os.MkdirAll(dir+".settings", os.ModePerm)
	writeFile(dir+".settings/org.eclipse.jdt.core.prefs", jdtPrefsTemplate)

	// src/jave folder with sample main
	srcDir := dir + "src/main/java/" + packageDir.String()
	os.MkdirAll(srcDir, os.ModePerm)
	writeTemplateFile(srcDir+"/Main.java", mainTemplate, packageParam)

	// test folder with sample test
	testDir := dir + "src/test/java/" + packageDir.String()
	os.MkdirAll(testDir, os.ModePerm)
	writeTemplateFile(testDir+"/SimpleTest.java", testTemplate, packageParam)
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
	if strings.HasSuffix(path, ".java") {
		content = strings.Replace(content, "    ", "\t", -1)
	}
	file, err := os.Create(path)
	check(err)
	defer file.Close()
	_, err = file.WriteString(strings.TrimLeft(content, "\n"))
	check(err)
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
