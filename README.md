<H1>Adventure With Strangers</H1>
This project consists of tow parts, The front-end and the back-end.
The only purpose of back-end in life for this project is to serve the front-end static pages and serve all the necessary API's for the front-end to work properly.
It is also in charge of reading from the database and writing to it.   
  
Back-end is written in Go and uses Fiber framework to for serving the static pages, API's, ....  
It communicates with a sqlite3 database declared in the root of the project

Now lets walk through each and every aspect this project

## Table of content
- [Table of content](#table-of-content)
- [configuration](#configuration)
- [cmd](#cmd)
- [database](#database)
- [internal](#internal)

## configuration
`config.json` is a json file that is located in (cmd/) and it tells the program the configuration of the server and the database. Right now it simply consists of:   
- #### host
    The ip address of the server.
- #### port   
    The port of the server.
- #### client-port
    The port of the client side.
- #### sqlite
    The source sqlite3 database file location.       
## build
`build` directory is where all the built static files of the react front-end lives. Including all the complied html, css, javascript, ... pages with all the other static files like images, etc.

## cmd
`cmd` directory is for the main services of the project which currently there is only one, the back-end which contains `main.go`, the main file with the main function in it, more on that later. 

- ### back-end
  The `back-end` directory is the place where the `main.go` lives. It is responsible to read the config and prepare the server and listen for incoming requests. it uses a Standard Go Fiber service to listen for requests and handle them.

## database
Package `database` is an implementation of sqlite3 database and consists of function
to query on the specific database designed for this project.

## internal
`internal` directory is where all the internal packages for this project live including the following packages:
- ### config
    Package `config` is for reading the config file necessary for the server interface to run
    And exports the config file to the config model declared in this package.
- ### handler
    `handler` directory is where the `api` package and `pages` package live.
    - ### api
        Package `api` handles all the incoming requests to api. It uses the `validation` package to validate the incoming requests. 
    - ### pages
        Package `pages` handles all the static pages built and imported from the react front-end. and it is set to serve the files in the `build` directory at the root of the project
- ### messages
    Package `messages` includes all the success and failure messages used across the project
- ### models
    Package `models` is for all the database and front-end form models that are used across other internal packages.
- ### server
    Package `fibServer` is the Standard Go Fiber server implementation and is called from the `main.go` file to prepare and set up the routs and apis declared in the `api` and `pages` packages and crating the cors configuration, cookie encryption, and serving and listening.
    It uses the config model provided in the `config` package.
- ### validation
    Package `validation` is for validating the incoming request to the `api` package.


      
    



