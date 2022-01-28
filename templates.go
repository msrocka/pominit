package main

const pomTemplate = `
<project xmlns="http://maven.apache.org/POM/4.0.0"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">

	<modelVersion>4.0.0</modelVersion>
	<properties>
		<project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
		<maven.compiler.target>17</maven.compiler.target>
  	<maven.compiler.source>17</maven.compiler.source>
	</properties>

	<groupId>{{.Artifact}}</groupId>
	<artifactId>{{.Artifact}}</artifactId>
	<version>0.0.1</version>

	<dependencies>
		<dependency>
			<groupId>junit</groupId>
			<artifactId>junit</artifactId>
			<version>4.13.2</version>
			<scope>test</scope>
		</dependency>
	</dependencies>
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

public class SimpleTest {

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

[*.{java,xml}]
indent_style = tab
indent_size = 2
tab_width = 2

[*.{md}]
indent_style = space
indent_size = 2
tab_width = 2

[*.{py}]
indent_style = space
indent_size = 4
tab_width = 4
`

const jdtPrefsTemplate = `
eclipse.preferences.version=1
org.eclipse.jdt.core.formatter.join_wrapped_lines=false
org.eclipse.jdt.core.formatter.lineSplit=80
`

const readmeTemplate = `
# {{.Artifact}}
`
