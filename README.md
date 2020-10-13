# SWEEP API
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/prasetyowira/sweep/CI?style=flat-square)
![CI](https://github.com/prasetyowira/sweep/workflows/CI/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/prasetyowira/sweep)](https://goreportcard.com/report/github.com/prasetyowira/sweep)
[![GolangCI](https://golangci.com/badges/github.com/prasetyowira/sweep.svg)](https://golangci.com/r/github.com/prasetyowira/sweep)

A Sweep API for recruitment assignment

## Getting started

Go Version: 1.14

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make up
$ make start
$ make run
```

Or running inside docker
```console
$ make up
$ make start
$ docker-compose -f docker-compose.prod.yml up -d
```


### REST API

```console
$ make start
```

```http
GET/POST http://127.0.0.1:8000/message
GET http://127.0.0.1:8000/message/{id}
```

Open openapi doc on port [127.0.0.1:81](127.0.0.1:81)



### Testing

``make test``
