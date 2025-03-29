.PHONY: spring golang all

spring:
	$(MAKE) -C spring-http-server all

golang:
	$(MAKE) -C go-http-server all

all: spring golang
