# Benchmark

Projetos para fazer benchmark de web servers

> _NOTA_: O projeto está com uma estrutura maluca pois o HEROKU é uma porcaria! Em suma, tive que por o app Java na raíz e deixei o Go numa pasta.

# Testes

Faça deploy no Heroku:

```shell
heroku apps:create benchmark-spring-api
heroku apps:create benchmark-golang-api
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

## Teste Locais

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

# TODO

- [ ] Adicionar algum framework web de Go pra ser justo com o Spring
