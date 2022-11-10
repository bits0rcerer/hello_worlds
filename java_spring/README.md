# Hello World!

..in Java with [gin-gonic](https://spring.io/)

## Setup

- Install [JDK 17](https://www.oracle.com/java/technologies/javase/jdk17-archive-downloads.html)¹

¹IntelliJ probably asks you if it should install the JDK for you

## Build Project

```bash
# Clone repo
$ git clone https://github.com/bits0rcerer/hello_worlds
$ cd java_spring
```

- Open project in [IntelliJ IDEA](https://www.jetbrains.com/idea/)
- Open the **Gradle** tool window (`View` > `Tool Windows` > `Gradle`)
- Run `demo` > `Tasks` > `application` > `bootRun`

The web server is now running and listening on port `8080`.

Try with your browser:

- http://localhost:8080/hello
- http://localhost:8080/hello?name=Go
- http://localhost:8080/hello_json
- http://localhost:8080/hello_json?name=Go

## Run Tests

```bash
# Clone repo
$ git clone https://github.com/bits0rcerer/hello_worlds
$ cd java_spring
```

- Open project in [IntelliJ IDEA](https://www.jetbrains.com/idea/)
- Open the **Gradle** tool window (`View` > `Tool Windows` > `Gradle`)
- Run `demo` > `Tasks` > `verification` > `test`