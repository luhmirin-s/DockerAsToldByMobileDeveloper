# A bit of composition

In more realistic scenario application need access to different services e.g. databases, caches. In this sample we have go service that can receive key-value pairs in REST api and stores them in Redis.

## But first lets do it by hand

* Create virtual network for future containers

``` bash
docker network create sample_network
```

* Pull and run redis instance

``` bash
docker pull redis
docker run --name go-redis -d -p 6379:6379 --network=sample_network  redis
```

* Find IP of redis container inside of the network

``` bash
docker network inspect sample_network
```

* Build and run sample service

``` bash
docker build -t sample3 .
docker run -d -p 8080:8080 --network=sample_network -e REDIS_HOST=<redis_IP>:6379 --name sample sample3
```

* Verify that it works at  `http://localhost:8080/`

``` bash
curl -d "" -X POST "http://localhost:8080/newKey?value=newValue"
curl "http://localhost:8080/newKey"
```

* Stop everything and cleanup

``` bash
docker stop sample go-redis
docker rm sample go-redis
docker network rm sample_network
```

## And now make the machine work for you

* All of those operations can be defined in [docker-compose](docker-compose.yaml) file and started up be single line:

``` bash
docker-compose up -d
```

* Verify that it still works at `http://localhost:8080/`

``` bash
curl -d "" -X POST "http://localhost:8080/newKey?value=newValue"
curl "http://localhost:8080/newKey"
```

* Note that common network was created automatically

``` bash
docker network ls
```

* Stop everything and cleanup

``` bash
docker-compose down
```
