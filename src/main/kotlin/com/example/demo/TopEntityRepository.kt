package com.example.demo

import org.springframework.data.jpa.repository.JpaRepository

interface TopEntityRepository : JpaRepository<TopEntity, Long>
