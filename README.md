SKMZ [![Build Status](https://travis-ci.com/Shpota/skmz.svg?branch=master)](https://travis-ci.com/Shpota/skmz) [![](https://img.shields.io/codecov/c/github/Shpota/skmz?color=green&logo=test%20coverage)](https://codecov.io/gh/Shpota/skmz)
====

A web application that allows to query programmers
with their skills via a **GraphQL** API. The
application is implemented with **Go** and 
**[gqlgen](https://github.com/99designs/gqlgen)**
on the backend side and **React** on the front end
side. **MongoDB** is used as a database.

![Showcase](showcase.gif)


## System requirements
You need to have [Docker](https://www.docker.com) and
[Docker Compose](https://docs.docker.com/compose/)
installed in oder to build and run the project. No
additional tools required.

## How to build and run in production mode
Perform 
```sh
docker-compose up
```
Access the application via http://localhost:8080.
Access the GraphQL Playground using 
http://localhost:8080/playground.

## How to develop locally

**Tools**

In order to develop the app locally the following
tools are required: [Docker](https://docs.docker.com/),
[Docker Compose](https://docs.docker.com/compose/) (if you
are on Mac or Windows it comes installed with Docker), 
[Node.js](https://nodejs.org/en/) and
[Go](https://golang.org/dl/).

Verify if your environment is ready by running the
following 4 commands:

```sh
docker --version
docker-compose --version
npm --version
go version
```

**Start the dev DB**
```sh
docker-compose -f docker-compose-dev.yml up
```
This will start a local MongoDB which will be
accessible on port `27017`. The DB will
be populated with test records from 
[mongo.init](server/db/mongo.init).

**Start the server**

Navigate to the `/server` folder and execute:

```sh
go run server.go
```
This will compile and run the back end part.
As a result, the API and [the GraphQL
playground](http://localhost:8080/playground)
will be available.

**Start the Front End dev server**

Navigate to the `/webapp` folder and execute
the following commands:

```sh
npm install
npm start
```
As a result, the web site will be accessible
on http://localhost:3000.

The changes on the front end side will be automatically
applied once a file is saved. The changes in the back
end code require restarting the back end.

## Customizations

The database starts with a preloaded set of data which
can be customized in 
[the mongo.init file](server/db/mongo.init).

Here is an example of a GraphQL query which can be
run in [the Playground](http://localhost:8080/playground):
```graphql
query {
  programmers(skill: "go") { 
    name,
    picture,
    title,
    company,
    skills {
      name,
      icon,
      importance
    }
  }
}
```
