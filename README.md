# Benchmark

Projetos para fazer benchmark de web servers

# TODO

- [X] Adicionar algum framework web de Go pra ser justo com o Spring
- [X] Adicionar algum ORM ao Go pra ser justo com o Spring (gorm)
- [ ] Adicionar coisas de Python

# Testes

Você pode executar os testes com:

```shell
ab -n 100000 -c 10 <host>:<port>/api/
```

## Resultados Locais (somente API)

Comando: `ab -n 10000000 -c 100 localhost:<port>/api/`

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

## Resultados Locais (com Banco de Dados)

<details>
    <summary>Resultados Kotlin (com Spring)</summary>
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

## Resultados Locais (com ORM e framework web em ambos)

<details>
    <summary>Resultados Kotlin (com Spring)</summary>

    ```shell
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
    Time taken for tests:   286.066 seconds
    Complete requests:      100000
    Failed requests:        0
    Total transferred:      2937200000 bytes
    HTML transferred:       2926700000 bytes
    Requests per second:    349.57 [#/sec] (mean)
    Time per request:       28.607 [ms] (mean)
    Time per request:       2.861 [ms] (mean, across all concurrent requests)
    Transfer rate:          10026.93 [Kbytes/sec] received
    
    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0    0   0.0      0       3
    Processing:    22   29   2.8     28     220
    Waiting:       12   16   1.8     16     179
    Total:         22   29   2.8     28     220
    
    Percentage of the requests served within a certain time (ms)
      50%     28
      66%     29
      75%     29
      80%     30
      90%     31
      95%     33
      98%     36
      99%     38
     100%    220 (longest request)
    ```

</details>

<details>
    <summary>Resultados Golang</summary>

    ```shell
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
    Document Length:        50967 bytes
    
    Concurrency Level:      10
    Time taken for tests:   93.394 seconds
    Complete requests:      100000
    Failed requests:        0
    Total transferred:      5107000000 bytes
    HTML transferred:       5096700000 bytes
    Requests per second:    1070.73 [#/sec] (mean)
    Time per request:       9.339 [ms] (mean)
    Time per request:       0.934 [ms] (mean, across all concurrent requests)
    Transfer rate:          53400.60 [Kbytes/sec] received
    
    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0    0   0.0      0       2
    Processing:     6    9   1.0      9      20
    Waiting:        5    9   1.0      9      20
    Total:          6    9   1.0      9      20
    
    Percentage of the requests served within a certain time (ms)
      50%      9
      66%     10
      75%     10
      80%     10
      90%     10
      95%     11
      98%     12
      99%     13
     100%     20 (longest request)
    ```

</details>

## Resultados Docker (com limite de `cpus=1` e `memory=512m`)

<details>

<summary>Resultados Kotlin (`make spring-bench`) </summary>

```shell
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            9095

Document Path:          /api/top-entities
Document Length:        29267 bytes

Concurrency Level:      100
Time taken for tests:   71.107 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      293720000 bytes
HTML transferred:       292670000 bytes
Requests per second:    140.63 [#/sec] (mean)
Time per request:       711.068 [ms] (mean)
Time per request:       7.111 [ms] (mean, across all concurrent requests)
Transfer rate:          4033.88 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       2
Processing:    26  706 369.5    594    4310
Waiting:       13  674 357.0    575    4294
Total:         26  706 369.5    594    4310

Percentage of the requests served within a certain time (ms)
  50%    594
  66%    684
  75%    735
  80%    793
  90%    996
  95%   1490
  98%   1894
  99%   2387
 100%   4310 (longest request)
```

</details>

<details>

<summary>Resultados Golang (`make golang-bench`)</summary>

```shell
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 1000 requests
Completed 2000 requests
Completed 3000 requests
Completed 4000 requests
Completed 5000 requests
Completed 6000 requests
Completed 7000 requests
Completed 8000 requests
Completed 9000 requests
Completed 10000 requests
Finished 10000 requests


Server Software:
Server Hostname:        localhost
Server Port:            9096

Document Path:          /api/top-entities
Document Length:        50967 bytes

Concurrency Level:      100
Time taken for tests:   86.540 seconds
Complete requests:      10000
Failed requests:        0
Total transferred:      510700000 bytes
HTML transferred:       509670000 bytes
Requests per second:    115.55 [#/sec] (mean)
Time per request:       865.401 [ms] (mean)
Time per request:       8.654 [ms] (mean, across all concurrent requests)
Transfer rate:          5763.00 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       3
Processing:     7  863 389.0    802    2995
Waiting:        7  856 388.6    801    2913
Total:          8  863 389.0    802    2995

Percentage of the requests served within a certain time (ms)
  50%    802
  66%    997
  75%   1098
  80%   1195
  90%   1398
  95%   1597
  98%   1800
  99%   1996
 100%   2995 (longest request)
```

</details>
