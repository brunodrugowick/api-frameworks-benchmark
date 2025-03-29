BENCH_REQUESTS = 1000

.PHONY: spring golang all

spring:
	$(MAKE) -C spring-http-server all BENCH_REQUESTS=$(BENCH_REQUESTS)

golang:
	$(MAKE) -C go-http-server all BENCH_REQUESTS=$(BENCH_REQUESTS)

all: spring golang
