# goKafkaSample
A sample kafka application written in go

## Run Kafka server

```
Note: Docker must be installed on the system since the Kafka server is run inside a docker container
```
To run the Kafka server please run the runKafka script using the command `./runKafka.sh`

The above script will run the Kafka server in a container on `localhost:9092`, removing the container will stop the server and clean the database.

## Run Messaging Application

To run the application, the consumer must be started before the publisher can send messages to it.

### Run the consumer

```
go run main.go --start consumer
```

### Run the publisher

```
go run main.go --start publisher
```

The above command will run the default publisher sending the string `Hello World !`  to the default kafka port on localhost, i.e. `localhost:9092`, this can be configured using flags for the publisher, more can be explored using `go run main.go --help`