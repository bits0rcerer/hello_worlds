package com.example.demo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
@SpringBootApplication
public class HelloWorld {
    public static void main(String[] args) {
        SpringApplication.run(HelloWorld.class, args);
    }

    @GetMapping(value = "/hello")
    public @ResponseBody String hello(@RequestParam(defaultValue = "World", name = "name") String name) {
        return String.format("Hello %s!", name);
    }

    @GetMapping(value = "/hello_json")
    public @ResponseBody HelloModel helloJSON(@RequestParam(defaultValue = "World", name = "name") String name) {
        return new HelloModel(name);
    }

    private record HelloModel(String name) {
    }
}
