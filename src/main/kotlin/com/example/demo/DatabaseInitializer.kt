package com.example.demo

import org.springframework.boot.context.event.ApplicationReadyEvent
import org.springframework.boot.context.event.ApplicationStartedEvent
import org.springframework.context.event.EventListener
import org.springframework.stereotype.Component
import org.springframework.transaction.annotation.Transactional
import java.util.logging.Logger

@Component
class DatabaseInitializer(private val topEntityRepository: TopEntityRepository) {

    private val log = Logger.getLogger(DatabaseInitializer::class.java.name)

    @EventListener(ApplicationStartedEvent::class)
    @Transactional
    fun init() {
        val topEntities = (1..10).map { top: Int ->
            TopEntity(
                middleEntities = (1..10).map { middle: Int ->
                    MiddleEntity(
                        innerEntities = (1..10).map { inner: Int ->
                            InnerEntity(text = "%s-%s-%s".format(top, middle, inner))
                        }
                    )
                }
            )
        }
        log.info("Saving might take a while, please wait for the application to report the total count of entities.")
        topEntityRepository.saveAll(topEntities)
        log.info("Total inner entities: %s".format(topEntityRepository.findAll().flatMap { it.middleEntities }.flatMap { it.innerEntities }.size))
    }
}
