# Benchmark

Projetos para fazer benchmark de web servers

> _NOTA_: O projeto está com uma estrutura maluca pois o HEROKU é uma porcaria! Em suma, tive que por o app Java na raíz e deixei o Go numa pasta.

# TODO

- [ ] Adicionar algum framework web de Go pra ser justo com o Spring
- [X] Adicionar algum ORM ao Go pra ser justo com o Spring (gorm)

# Testes

Faça deploy no Heroku:

```shell
heroku apps:create benchmark-spring-api
heroku apps:create benchmark-golang-api
heroku addons:create heroku-postgresql:hobby-dev -a benchmark-spring-api
heroku addons:create heroku-postgresql:hobby-dev -a benchmark-golang-api
heroku buildpacks:add -a benchmark-spring-api heroku-community/multi-procfile
heroku buildpacks:add -a benchmark-golang-api heroku-community/multi-procfile
heroku config:set -a benchmark-spring-api PROCFILE=Procfile
heroku config:set -a benchmark-golang-api PROCFILE=go-http-server/Procfile
git push https://git.heroku.com/benchmark-spring-api.git HEAD:master
git push https://git.heroku.com/benchmark-golang-api.git HEAD:master
```

Você pode executar os testes com:

```shell
ab -n 100000 -c 10 <host>:<port>/api/
```

## Resultados Locais (somente API)

Comando: `ab -n 10000000 -c 100 localhost:<port>/api/`

<details>
    <summary>Resultados Golang (apenas bibliotecas nativas)</summary>
    ```
    This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/

    Benchmarking localhost (be patient)


    Server Software:        
    Server Hostname:        localhost
    Server Port:            8085

    Document Path:          /api/
    Document Length:        14 bytes

    Concurrency Level:      100
    Time taken for tests:   435.847 seconds
    Complete requests:      10000000
    Failed requests:        0
    Total transferred:      1310000000 bytes
    HTML transferred:       140000000 bytes
    Requests per second:    22943.84 [#/sec] (mean)
    Time per request:       4.358 [ms] (mean)
    Time per request:       0.044 [ms] (mean, across all concurrent requests)
    Transfer rate:          2935.20 [Kbytes/sec] received

    Connection Times (ms)
                min  mean[+/-sd] median   max
    Connect:        0    2   0.5      2       6
    Processing:     0    2   0.5      2      13
    Waiting:        0    1   0.5      1      12
    Total:          1    4   0.4      4      16

    Percentage of the requests served within a certain time (ms)
    50%      4
    66%      4
    75%      5
    80%      5
    90%      5
    95%      5
    98%      5
    99%      5
    100%     16 (longest request)

    ```
</details>

<details>
    <summary>Resultados Kotlin (com Spring)</summary>
    ```
    This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/

    Benchmarking localhost (be patient)


    Server Software:        
    Server Hostname:        localhost
    Server Port:            9095

    Document Path:          /api/
    Document Length:        26 bytes

    Concurrency Level:      100
    Time taken for tests:   427.554 seconds
    Complete requests:      10000000
    Failed requests:        0
    Total transferred:      1590000000 bytes
    HTML transferred:       260000000 bytes
    Requests per second:    23388.83 [#/sec] (mean)
    Time per request:       4.276 [ms] (mean)
    Time per request:       0.043 [ms] (mean, across all concurrent requests)
    Transfer rate:          3631.66 [Kbytes/sec] received

    Connection Times (ms)
                min  mean[+/-sd] median   max
    Connect:        0    2   0.5      2       6
    Processing:     0    2   0.5      2       8
    Waiting:        0    1   0.5      1       6
    Total:          2    4   0.3      4      12

    Percentage of the requests served within a certain time (ms)
    50%      4
    66%      4
    75%      4
    80%      4
    90%      5
    95%      5
    98%      5
    99%      5
    100%     12 (longest request)

    ```
</details>

## Resultados Locais (com Banco de Dados)

<details>
    <summary>Resultados Kotlin (com Spring)</summary
    ```
    >> ab -n 100000 -c 10 localhost:9095/api/top-entities
    This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/
    
    Benchmarking localhost (be patient)
    Completed 10000 requests
    Completed 20000 requests
    Completed 30000 requests
    Completed 40000 requests
    Completed 50000 requests
    Completed 60000 requests
    Completed 70000 requests
    Completed 80000 requests
    Completed 90000 requests
    Completed 100000 requests
    Finished 100000 requests
    
    
    Server Software:
    Server Hostname:        localhost
    Server Port:            9095
    
    Document Path:          /api/top-entities
    Document Length:        29267 bytes
    
    Concurrency Level:      10
    Time taken for tests:   297.913 seconds
    Complete requests:      100000
    Failed requests:        0
    Total transferred:      2937200000 bytes
    HTML transferred:       2926700000 bytes
    Requests per second:    335.67 [#/sec] (mean)
    Time per request:       29.791 [ms] (mean)
    Time per request:       2.979 [ms] (mean, across all concurrent requests)
    Transfer rate:          9628.18 [Kbytes/sec] received
    
    Connection Times (ms)
    min  mean[+/-sd] median   max
    Connect:        0    0   0.0      0       1
    Processing:    23   30   3.7     29     204
    Waiting:       13   16   2.3     16     165
    Total:         23   30   3.7     29     204
    
    Percentage of the requests served within a certain time (ms)
    50%     29
    66%     30
    75%     30
    80%     31
    90%     33
    95%     35
    98%     39
    99%     43
    100%    204 (longest request)
    ```
</details>

<details>
    <summary>Resultados Golang (note que muitas requisições falharam)</summary>
    ```
    >> ab -n 100000 -c 10 localhost:9096/api/top-entities
    This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/
    
    Benchmarking localhost (be patient)
    Completed 10000 requests
    Completed 20000 requests
    Completed 30000 requests
    Completed 40000 requests
    Completed 50000 requests
    Completed 60000 requests
    Completed 70000 requests
    Completed 80000 requests
    Completed 90000 requests
    Completed 100000 requests
    Finished 100000 requests
    
    
    Server Software:
    Server Hostname:        localhost
    Server Port:            9096
    
    Document Path:          /api/top-entities
    Document Length:        28968 bytes
    
    Concurrency Level:      10
    Time taken for tests:   87.012 seconds
    Complete requests:      100000
    Failed requests:        5591
    (Connect: 0, Receive: 0, Length: 5591, Exceptions: 0)
    Total transferred:      2743930026 bytes
    HTML transferred:       2735022791 bytes
    Requests per second:    1149.26 [#/sec] (mean)
    Time per request:       8.701 [ms] (mean)
    Time per request:       0.870 [ms] (mean, across all concurrent requests)
    Transfer rate:          30795.90 [Kbytes/sec] received
    
    Connection Times (ms)
    min  mean[+/-sd] median   max
    Connect:        0    0   0.1      0       5
    Processing:     2    9   6.4      6      62
    Waiting:        2    9   6.4      6      62
    Total:          2    9   6.4      6      62
    
    Percentage of the requests served within a certain time (ms)
    50%      6
    66%     10
    75%     11
    80%     12
    90%     15
    95%     25
    98%     30
    99%     32
    100%     62 (longest request)
    ```
</details>
