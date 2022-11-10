package com.example.demo;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.test.web.servlet.MockMvc;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest
class HelloWorldTest {

    @Autowired
    private MockMvc mockMvc;

    @Test
    void hello() throws Exception {
        // Default parameter
        mockMvc.perform(get("/hello"))
                .andExpect(status().isOk())
                .andExpect(content().string("Hello World!"));

        mockMvc.perform(get("/hello")
                        .param("name", "Java"))
                .andExpect(status().isOk())
                .andExpect(content().string("Hello Java!"));
    }

    @Test
    void helloJSON() throws Exception {
        // Default parameter
        mockMvc.perform(get("/hello_json"))
                .andExpect(status().isOk())
                .andExpect(content().json("{\"name\":\"World\"}"));

        mockMvc.perform(get("/hello_json")
                        .param("name", "Java"))
                .andExpect(status().isOk())
                .andExpect(content().json("{\"name\":\"Java\"}"));
    }
}