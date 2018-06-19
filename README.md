# pominit
`pominit` is a small command line tool for quickly creating a Maven project: 

```bash
pominit <artifactId>
```

This will create a Maven project for the given `artifactId`. If no `artifactId`
is given it will create the maven project directly in the respective folder
where `pominit` is called and will take the name of the surrounding folder as
`artifactId`.


* a `pom.xml` with just a JUnit dependency
* a `src/main/java` folder
* a simple JUnit test: `src/test/java/tests/ATest.java`

You can run the `mvn test` to execute the unit tests. The created project can
be used in combination with the 
[Language support for Java ™ for Visual Studio Code](https://github.com/redhat-developer/vscode-java).
