# SimpleApp

Architecture: Clean Architecture

### Prequisite:

1. setup go Environment https://go.dev/doc/install
2. install posgreSql [library](https://www.postgresql.org/download/windows/), [docker](https://docs.docker.com/engine/install/), [dbeaver](https://dbeaver.io/download/).
3. get code of https://github.com/gusaul/grpcox for the gRPC client
4. clone this repo

## Running the grpcox client:

Open the terminal and run command  `make start` or follow the docker instructions from the [README.md](https://github.com/gusaul/grpcox#readme) 

<img width="359" alt="image" src="https://github.com/patrichr/SimpleApp/assets/133081619/3880eccf-4f78-4057-a9d8-b136b73b9763">

## PostgreSQL 
#### Docker container

To run the docker container, open terminal from the .yaml directory and run command : `docker-compose up --build -d` 

To stop postgresql container: `docker-compose down`

<img width="733" alt="image" src="https://github.com/patrichr/SimpleApp/assets/133081619/fc4f6317-1819-4d69-aad1-d026283f7b7c">

or simply press play/stop at the interface

#### to Run the postgre database:

1. open dbeaver
2. create a new connection by clicking the New Database Connection icon 
3. fill in the config with these infos:

`dbhost = "localhost"` 

`dbport = 5656` 

`user = "server"` 

`password = "server"` 

`dbname = "grpc-server"`

<img width="503" alt="postgresql_setting" src="https://github.com/patrichr/SimpleApp/assets/133081619/6fc9f96b-1483-401e-a45f-2a9a5031906a">

4. Right click on the database, click connect
5. Open SQL editor and insert query /copy contents from [here](https://github.com/patrichr/SimpleApp/blob/main/database/01_merchant_tiers.sql) and execute
6. There should be a tasks Table generated under the tab grpc-server>Databases>postgres>Schemas>public>Table

<img width="657" alt="image" src="https://github.com/patrichr/SimpleApp/assets/133081619/60d53439-ca14-4ad0-9145-88a348909743">

## How to run SimpleApp:

1. open terminal from the **SimpleApp/server** directory
2. run command `go mod vendor`
3. run `go build && ./server`
4. from grpcox, connect to server port 50051
5. enter gRPC Server target as `localhost:50051', services as `grpc.Greeter` and Methods as `Register`
6. input the data into the JSON, and there shall be a response bellow it after clicking submit, as pictured bellow.

![image](https://github.com/patrichr/SimpleApp/assets/133081619/712626ed-ae8c-4446-af18-5eb03531cd5b)

The result To-Do task shall appear in the dbeaver, when you execute a query as shown:
![image](https://github.com/patrichr/SimpleApp/assets/133081619/bfb57532-96ea-41db-a953-7dbc4ead6ce6)


