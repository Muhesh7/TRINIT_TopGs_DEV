# Tri-NIT

___

## Tech-Stack

|<img src=https://www.vectorlogo.zone/logos/reactjs/reactjs-ar21.png height=200 width="100%" > | <img src=https://seekvectorlogo.com/wp-content/uploads/2018/12/docker-vector-logo.png height=200 width="100%" >
|:---:|:---:|
| <img src=https://www.python.org/static/community_logos/python-logo-master-v3-TM-flattened.png height=200 width="100%" > | <img src=https://miro.medium.com/max/920/1*CdjOgfolLt_GNJYBzI-1QQ.jpeg height=200 width="100%" >
|<img src=https://res.cloudinary.com/practicaldev/image/fetch/s--TNgs2Fd7--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://dev-to-uploads.s3.amazonaws.com/uploads/articles/8susytd9w6lxe9sreqvd.jpg  height=200 width="100%" > | <img src=https://adequatesource.com/content/images/2021/08/protobuf.webp height=200 width="100%" >
|<img src=https://www.vectorlogo.zone/logos/neo4j/neo4j-ar21.png  height=200 width="100%" > |<img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQclNzIJG6C0F6DVnv7tbfVxG9lgnmfn-iw2A&usqp=CAU"  height=200 width="100%" >
|<img src=https://www.vectorlogo.zone/logos/apache_spark/apache_spark-ar21.png  height=200 width="100%" > | <img src=https://digitalis.io/wp-content/uploads/2020/12/Kafka600x340.jpg height=200 width="100%" >



## Architecture

* Micro-services architecture with Database per service pattern

* Admin-Service is used for registering user and then the user registers the app he wants the use the cluster-service.

* Cluster-Service authenticates the app-client by dialling a rpc call to admin-service and recieves the data via **kafka** topic `to_cluster` and creates cluster and stores in Neo4j database and emits message to `clustered` topic

* ![architecture](https://imgur.com/dgsgeJd.png)

## Screenshots
* #### Dashboard
    ![login](https://imgur.com/0Ta0lcD.png)

    ![login](https://imgur.com/Z0SMnx4.png)

    ![login](https://imgur.com/EZPpVKe.png)

* #### Admin-Service
    ![Imgur](https://imgur.com/eL9KQhR.png)

    ![img](https://i.imgur.com/31kUkAA.png)

* #### Kafka
   ##### Producer
    ![producer](https://imgur.com/yD3ZBZQ.png)
   ##### Consumer
    ![consumer](https://imgur.com/EVK4fuZ.png)
   ##### Sample-data
    ![sample](https://imgur.com/Wv98jfs.png)

* #### Cluster-Service
   ##### Neo4j 
   - Rules defined by the user for his app( named app3)
    ![login](https://imgur.com/EI3OFTk.png)

   
  - 12 different records into 3 Clusters based on partial match 
    in `name`, exact match in `IP` and  exact match in `email`
    ![login](https://imgur.com/vj11iNa.png)

___

## [Video](https://drive.google.com/drive/folders/1rGbbqBHBan4KE1gkbmqHMjVePnRuophJ?usp=sharing)

___

## Setup

* ### Requirements
    * [Node](https://nodejs.org/en/)
    * [Go](https://go.dev/)
    * [Docker](https://www.docker.com/)
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