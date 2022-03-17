# CRUD App

CRUD is an acronym for 
1. **C**reate
1. **R**ead
1. **U**pdate
1. **D**elete

A CRUD application is expected to expose REST apis for different entities/models
in the application to create, read, update and delete them.

## The Application
Our application will consist of a simple microservice api that exposes endpoints
for working with the `Users` collection backed by a MongoDB database

## External packages
The application makes use of the following external packages -

1. Gin Framework - This is used as our REST api framework for routing, middleware etc
2. mgo - Used as the driver for interacting with the MongoDB database