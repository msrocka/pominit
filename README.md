# pominit
`pominit` is a small command line tool for quickly creating a Maven project: 

```bash
cd <your project folder>
pominit
mvn test
```

This will create the following things:

* a `pom.xml` with just a JUnit dependency
* a `src/main/java` folder
* a simple JUnit test: `src/test/java/tests/ATest.java`

You can run the `mvn test` to execute the unit tests. The created project can
be used in combination with the 
[Language support for Java ™ for Visual Studio Code](https://github.com/redhat-developer/vscode-java).

## Kotlin support
Adding `kotlin` or `kt` as argument will create a Kotlin project targeting
Java 8 (remove the `-jre8` suffix from the `stdlib` dependency if you want to
use the standard Kotlin library):

```
cd <your project folder>
pominit kt
mvn test
```

As for the Java based setup, the project will contain a simple unit test so that
you can run `mvn test`. Additionally, it will contain a simple `main`-method
and will create a jar with dependencies when you run `mvn package`. Thus, you
should be able to run the package result via:

```
mvn package
java -jar .\target\thing-0.0.1.jar
```

Remove the respective parts if you do not want to use them.
