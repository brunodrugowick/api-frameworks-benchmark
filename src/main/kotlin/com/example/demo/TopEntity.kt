package com.example.demo

import jakarta.persistence.*

@Entity
class TopEntity(

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    val id: Long? = null,

    @OneToMany(cascade = [CascadeType.ALL], fetch = FetchType.LAZY)
    val middleEntities: List<MiddleEntity> = mutableListOf()
)
