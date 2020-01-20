SKMZ [![Build Status](https://travis-ci.com/Shpota/skmz.svg?branch=master)](https://travis-ci.com/Shpota/skmz)
====

A web application that allows to query programmers
with their skills via a GraphQL API. The application
is implemented with Go and 
[gqlgen](https://github.com/99designs/gqlgen) on the
backend side and and React on the front end side.

## System requirements 
You need to have [Docker](https://www.docker.com) 
installed in order to build and run the application.
No additional tools required.

## How to build and run
1. Build the application image:
    ```shell script
    docker build -t skmz .
    ```
2. Start the application container:
    ```shell script
    docker run -p 8080:8080 skmz
    ```
Access the application via http://localhost:8080.
Access the GraphQL API using 
http://localhost:8080/playground 
