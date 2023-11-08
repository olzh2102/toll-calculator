# Toll Calculator

```
docker run -d --name kafka -p 9092:9092 \
    -e KAFKA_CFG_NODE_ID=0 \
    -e KAFKA_CFG_PROCESS_ROLES=controller,broker \
    -e KAFKA_CFG_LISTENERS=PLAINTEXT://:9092,CONTROLLER://:9093 \
    -e KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=CONTROLLER:PLAINTEXT,PLAINTEXT:PLAINTEXT \
    -e KAFKA_CFG_CONTROLLER_QUORUM_VOTERS=0@kafka-server:9093 \
    -e KAFKA_CFG_CONTROLLER_LISTENER_NAMES=CONTROLLER \
    bitnami/kafka:latest
```

## Installing protobuf compiler
for linux user or (wsl2)
```
sudo apt install -y protobuf-compiler 
```

for mac user you can use brew for this
```
brew install protobuf
```

## Installing GRPC and Protobuf plugins for Golang
1. Protobuffers
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```

2. GRPC
```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

3. install the package dependencies

```
go get google.golang.org/protobuf
```

```
go get google.golang.org/grpc
```