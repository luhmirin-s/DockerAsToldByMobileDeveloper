# Docker tutorial as told by mobile developer

Presentation slides are in [here](https://docs.google.com/presentation/d/1wonO_vH1twLFoPrX37EBLPrFgsn0T73by6f-nSziCPs/edit?usp=sharing), but most of actual materials to follow along is in README files in sample folders:

## Table of contents

* [Installation](#quick-start-guide)
* [Quick start guide](#quick-start-guide)
* [Basic Dockerfile sample](sample1/README.md)
* [More realistic Dockerfile sample](sample2/README.md)
* [docker-compose sample](sample3/README.md)
* [Useful links](#useful-links)

## Instalation instructions

* For MacOS
  * [Drag-n-drop version](https://docs.docker.com/v17.12/docker-for-mac/install/)
  * [Brew](https://pilsniak.com/how-to-install-docker-on-mac-os-using-brew/)

* [For Linux](https://docs.docker.com/v17.12/install/linux/docker-ce/ubuntu/)

* Also consider:
  * [Dockerhub account](https://hub.docker.com/)
  * [Kinematic](https://kitematic.com/docs/)
  * [Lazydocker](https://github.com/jesseduffield/lazydocker)

## Quick start guide

* Go to dockerhub and find elasticsearch base image

* Pull base image (similar to git)

``` bash
docker pull elasticsearch:7.2.0
```

* Execute command from “Run Elasticsearch” section

``` bash
docker run -d --name elastic -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.2.0
```

* Check container state

``` bash
docker ps
```

* Check if it works at `http://localhost:9200/`

* Stop container

``` bash
docker stop elastic
```

* Check container state again

``` bash
docker ps
docker ps -a
```

* Restart container

``` bash
docker start elastic
docker ps
```

* Delete container

``` bash
docker stop elastic
docker rm elastic
```

* Check container state again

``` bash
docker ps
docker ps -a
```

## Useful links

* https://en.wikipedia.org/wiki/Docker_(software)
* https://www.quora.com/What-is-Docker-Please-explain-it-in-simple-terms
* https://www.callicoder.com/docker-golang-image-container-example/
* https://docs.docker.com/v17.12/install/
* https://docs.docker.com/engine/reference/builder/
* https://docs.docker.com/develop/develop-images/dockerfile_best-practices/
* https://docs.docker.com/storage/volumes/
* https://docs.docker.com/compose/
* https://hub.docker.com/_/elasticsearch
* https://hub.docker.com/_/golang
* https://hub.docker.com/r/paciolanadmin/golang-alpine
* https://hub.docker.com/_/redis 
* https://stackify.com/elasticsearch-tutorial/
* https://www.restapiexample.com/golang-tutorial/simple-example-uses-redis-golang/
