# Tri-NIT

___

## Tech-Stack

|![react](images/react.webp?raw=true "Title") | ![go](images/go.webp?raw=true "Title")
|:---:|:---:|
|![pg](images/pg.webp?raw=true "Title") | ![docker](images/docker.webp?raw=true "Title")
|![grpc](images/grpc.webp?raw=true "Title") | ![spark](images/spark.png?raw=true "Title")
|![proto](images/protobuf.webp?raw=true "Title") | ![python](images/python.webp?raw=true "Title")
|![neo4j](images/neo4j.png?raw=true "Title") | ![kafka](images/kafka.png?raw=true "Title")

## Setup

* ### Requirements
    * [Node](https://nodejs.org/en/)
    * [Go](https://go.dev/)
    * [Docker](https://www.docker.com/)


## Architecture

* Micro-services architecture with Database per service pattern

* Admin-Service is used for registering user and then the user registers the app he wants the use the cluster-service.

* Cluster-Service authenticates the app-client by dialling a rpc call to admin-service and recieves the data via **kafka** topic `to_cluster` and creates cluster and stores in Neo4j database and emits message to `clustered` topic

![architecture](images/archi.png?raw=true "Title")


* ### Run
    * ##### Admin-Service
        ```bash
        cp .env.example .env && docker-compose up
        ```
    * ##### Kafka
        ```bash
        docker-compose up
        ```
    * ##### Cluster-Service
        ```bash
        ./build.sh && docker-compose up
        ```    
    * ##### Dashboard
        ```bash
        npm i && npm run dev
        ```