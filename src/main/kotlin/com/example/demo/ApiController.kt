package com.example.demo

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class ApiController(
    private val topEntityRepository: TopEntityRepository
) {

    @GetMapping("/api/")
    fun getApiResponse(): String {
        return "Hello, Kotlin with Spring!"
    }

    @GetMapping("/api/top-entities")
    fun getTopEntities(): List<TopEntity> {
        return topEntityRepository.findAll()
    }

}
