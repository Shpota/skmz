SKMZ [![Build Status](https://travis-ci.com/Shpota/skmz.svg?branch=master)](https://travis-ci.com/Shpota/skmz) [![](https://img.shields.io/codecov/c/github/Shpota/skmz?color=green&logo=test%20coverage)](https://codecov.io/gh/Shpota/skmz)
====

A web application that allows to query programmers
with their skills via a GraphQL API. The application
is implemented with Go and 
[gqlgen](https://github.com/99designs/gqlgen) on the
backend side and and React on the front end side.

## System requirements 
You need to have [Docker](https://www.docker.com) and
[Docker Compose](https://docs.docker.com/compose/)
installed in order to build and run the application.
No additional tools required.

## How to build and run
Perform 
```shell script
docker-compose up
```
Access the application via http://localhost:8080.
Access the GraphQL API using 
http://localhost:8080/playground 
