# pominit
`pominit` is a small command line tool for quickly creating a Maven project: 

```bash
pominit <artifactId>
```

This will create a Maven project for the given `artifactId`. If no `artifactId`
is given it will create the maven project directly in the respective folder
where `pominit` is called and will take the name of the surrounding folder as
`artifactId`.

It will then create the following things:

* A `Main` class and a `run` script which should print `"Works!"` when you
  execute it
* A sample test `ATest` which you can execute via `mvn test`
* Additional resource files like a `.gitignore`, `.editorconfig` etc.

The generated project should work fine with the 
[VSCode Java Extension](https://marketplace.visualstudio.com/items?itemName=vscjava.vscode-java-pack).
