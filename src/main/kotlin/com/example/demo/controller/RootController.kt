package com.example.demo.controller

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class ApiController {

    @GetMapping("/api/")
    fun getApiResponse(): String {
        return "Hello, Kotlin with Spring!"
    }
}
