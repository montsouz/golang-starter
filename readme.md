# Go, MongoDB and Docker 

A simple Golang API skeleton heavly inspired by NodeJs + Express + MongoDB aplications.That way it is easier for 
Node developers to start creating new performant APIs with Go. 

This is a work in progress.

## How to run 

This project has Fresh instaled, so to run just simply type on terminal inside the root folder.

`fresh`

If you're going to run via docker

`docker-compose up`

## Organization

It is very similar to commom Node + Express project, with routes,models and handlers (controllers) folder.

  .
    ├── ...
    ├── src                   
    │   ├── db                 # package that integrates with MongoDb
    │   ├── handlers             # here is where you're going to deal with incoming requests
    |   ├── models               # The models for MongoDB
    │   └── routes
    |       |── routes.go               # Routes main file
    └── ...




