# Test task.

## Components:
1. service. It communicates with cryptocompare.com. It contains a scheduler and keep up a database with fresh info about configured currency pairs.
2. rest server
3. websocket server. The results could be inspected in web console.

Service, rest server and websocket server communicates through grpc.

## Configs
They could be found in configs directory. Samples are in .yaml.dist files. The .yaml.dist files should be copied if .yaml files with suitable parameters filled in. Currently there are three files: restserver.yaml.dist, wsserver.yaml.dist, service.yaml.dist 

## Logs
They could be found in logs directory

## How to run:
1. make generate-proto #generate grpc protos
1. docker-compose up #start database if there is no other
2. make run-service #start service
3. make run-rest-server #start rest server
4. make run-ws-server #start websocket server
