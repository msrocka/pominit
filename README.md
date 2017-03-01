# pominit
`pominit` is a small command line tool for quickly creating a Maven project. You
just type `pominit` and it will create the following things:

* a `pom.xml` with just a JUnit dependency
* a `src/main/java` folder
* a simple JUnit test: `src/test/java/tests/ATest.java`

You can run the `mvn test` to execute the unit tests. The created project can
be used in combination with the 
[Language support for Java ™ for Visual Studio Code](https://github.com/redhat-developer/vscode-java).
