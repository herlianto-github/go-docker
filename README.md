[![Docker Image CI/CD](https://github.com/herlianto-github/go-docker/actions/workflows/docker-image.yml/badge.svg)](https://github.com/herlianto-github/go-docker/actions/workflows/docker-image.yml)

# go-docker

Golang Rest API and deploy with Docker.

## Clone this repository

```sh
$ git clone https://github.com/herlianto-github/go-docker.git
> Cloning into `go-docker`...
> remote: Counting objects: 10, done.
> remote: Compressing objects: 100% (8/8), done.
> remove: Total 10 (delta 1), reused 10 (delta 1)
> Unpacking objects: 100% (10/10), done.
```

## Build

```sh
docker build -t go-docker .
```

### Run

```sh
docker run -dp 8080:8080 go-docker
```
  
Open your favorite browser set URL to <http://localhost:8080/> and hit ENTER
