# Perkbox Software Developer Test

The project runs on Docker.

## Start the environment

1. Docker/docker-machine and docker-compose needs to be installed in the host
2. On linux, change the `docker-compose.yml` `DSN_DB` ip address to be the gateway ip address of the container `perkbox-api`. What I do is to start the environment, then ssh into the `perkbox-api` container and check the `ip routes` to get the gateway ip. The needed IP might be different in Mac or Windows machines and in docker-machine/no docker-machine set-ups.
3. `docker-compose up` to start the environment.

## How the project is set up

- A MySQL docker image will be downloaded and started. When it is healthy docker-compose will start the perkbox project.
- Gomigrate will run at the begining of the execution creating a new table and example data in the db.

## How to use the API

The project is a REST API using POST, GET and PUT methods.
The easiest way to test the API is using Postman and importing the postmanConf files provided. There can be seen what the API is capable of with example set-ups.

## To-dos

- Write more tests for the handlers
- Find bugs (for example, when offset is set to 0) and fix them
- Expand filtering on get coupons method
- Play with docker networking so the mysql ip address can be known beforehand