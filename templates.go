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
            <version>4.12</version>
            <scope>test</scope>
        </dependency>

    </dependencies>

    <build>
        <plugins>
            <plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-compiler-plugin</artifactId>
                <version>3.7.0</version>
                <configuration>
                    <source>1.8</source>
                    <target>1.8</target>
                </configuration>
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

const runBatTemplate = `
mvn compile exec:java -Dexec.mainClass="{{.Package}}.Main" -q
`

const editorConfigTemplate = `
# see https://editorconfig.org/

root = true

[*]
charset = utf-8
insert_final_newline = true
trim_trailing_whitespace = true

[*.xml]
indent_style = space
indent_size = 4

[*.java]
indent_style = tab
`
