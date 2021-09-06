#!/bin/bash

# Controller
mkdir -p mocks/controller/rest
mockgen --source=controller/rest/base.go -package=mock_rest_controller --destination=mocks/controller/rest/base.go
# Controller Grpc
mkdir -p mocks/controller/grpc
mockgen --source=controller/grpc/base.go -package=mock_grpc_controller --destination=mocks/controller/grpc/base.go

# Handler
mkdir -p mocks/handler
mockgen --source=handler/base.go -package=mock_handler --destination=mocks/handler/base.go

# Service
mkdir -p mocks/service/logs
mockgen --source=service/logs/base.go -package=mock_service --destination=mocks/service/logs/base.go
# Service
mkdir -p mocks/service/omdb
mockgen --source=service/omdb/base.go -package=mock_service --destination=mocks/service/omdb/base.go