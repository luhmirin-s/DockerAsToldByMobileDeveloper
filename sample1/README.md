# The basic Dockerfile

This sample contains more optimized Dockerfile that uses builder image to create executable binary that is bundled into much smaller base image resulting more optimized deliverable image.

* Build image using improved [Dockerfile](Dockerfile)

``` bash
docker build -t sample1 .
```

* Verify that image has been create

``` bash
docker image ls
```

* Run container with created image

``` bash
docker run -d -p 8080:8080 --name sample sample1
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

* Note that basic hello world application image is more than 700MB, we are going to fix in [sample2](../sample2/README.md)
