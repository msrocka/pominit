package main

const pomTemplate = `
<project xmlns="http://maven.apache.org/POM/4.0.0"
  xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">

  <modelVersion>4.0.0</modelVersion>
  <properties>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
  </properties>

  <groupId>{{.Artifact}}</groupId>
  <artifactId>{{.Artifact}}</artifactId>
  <version>0.0.1</version>

  <dependencies>

    <dependency>
      <groupId>junit</groupId>
      <artifactId>junit</artifactId>
      <version>4.13.1</version>
      <scope>test</scope>
    </dependency>

  </dependencies>

  <build>
    <plugins>
      <plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-compiler-plugin</artifactId>
        <version>3.8.1</version>
        <configuration>
          <source>13</source>
          <target>13</target>
        </configuration>
      </plugin>
      <plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-dependency-plugin</artifactId>
        <version>3.1.2</version>
        <executions>
          <execution>
            <phase>package</phase>
            <goals>
              <goal>copy-dependencies</goal>
            </goals>
            <configuration>
              <outputDirectory>${project.build.directory}/lib</outputDirectory>
              <includeScope>runtime</includeScope>
            </configuration>
          </execution>
        </executions>
      </plugin>
    </plugins>
  </build>
</project>
`

const gitignoreTemplate = `# Maven output
target/

# IDE files
.idea/
*.iml
.classpath
.project
.settings/
.vscode
`

const mainTemplate = `
package {{.Package}};

public class Main {

  public static void main(String[] args) {
    System.out.println("Works!");
  }
}
`

const testTemplate = `
package {{.Package}};

import org.junit.Assert;
import org.junit.Test;

public class ATest {

  @Test
  public void test() {
    System.out.println("running a test ...");
    Assert.assertTrue(true);
  }
}
`

const editorConfigTemplate = `
# see https://editorconfig.org/

root = true

[*]
charset = utf-8
insert_final_newline = true
trim_trailing_whitespace = true
indent_style = space
indent_size = 2
`

const jdtPrefsTemplate = `
eclipse.preferences.version=1
org.eclipse.jdt.core.formatter.join_wrapped_lines=false
org.eclipse.jdt.core.formatter.lineSplit=80
`
