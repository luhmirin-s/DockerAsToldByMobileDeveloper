# Bit more realistic Dockerfile

This sample contains the most basic Dockerfile example and instructions how to start service using one.

* Build image using this [Dockerfile](Dockerfile)

``` bash
docker build -t sample2 .
```

* Run container with created image

``` bash
docker run -d -p 8080:8080 --name sample sample2
```

* Verify that it works at `http://localhost:8080/`

* Stop and delete container

``` bash
docker stop sample
docker rm sample
```

* Open image list again

``` bash
docker image ls
```

* Note that sample2 image is waaaay smaller than sample1

That was quite a lot of typing, in [sample3](../sample3/README.md) we are going to fix that.
