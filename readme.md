
[![GitHub issues](https://img.shields.io/github/issues/montsouz/golang-starter)](https://github.com/montsouz/golang-starter/issues)
[![GitHub forks](https://img.shields.io/github/forks/montsouz/golang-starter)](https://github.com/montsouz/golang-starter/network)
[![GitHub license](https://img.shields.io/github/license/montsouz/golang-starter)](https://github.com/montsouz/golang-starter/blob/master/LICENSE)


# Go, MongoDB and Docker 

A simple Golang API skeleton heavly inspired by NodeJs + Express + MongoDB aplications.That way it is easier for 
Node developers to start creating new performant APIs with Go. 

## How to run 

The easiest way to run this project is by installing Fresh. 

`go get github.com/pilu/fresh`

Then, just type: 

`fresh`

If you're going to run via docker change the `.env` 

Instead of... 

`DATABASE_URL=localhost` 

Use.

`DATABASE_URL=mongo` 

Then run the magic command:

`docker-compose up`

## Organization

It is very similar to a commom Node + Express project, with routes,models and handlers (controllers) folder.

        .
        ├── ...
        ├── src
        |   ├── context              # To keep project constants, db sessions ... etc 
        │   ├── db                   # package that integrates with MongoDb
        │   ├── handlers             # here is where you're going to deal with incoming requests
        |   ├── models               # The models for MongoDB
        │   ├── routes
        |   |    |── routes.go       # Routes main file 
        |   └──  utils               # utility methods   
        └── ...






