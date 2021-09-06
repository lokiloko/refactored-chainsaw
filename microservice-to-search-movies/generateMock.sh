#!/bin/bash

# Handler
mkdir -p mocks/handler
mockgen --source=handler/base.go -package=mock_handler --destination=mocks/handler/base.go

# Service
mkdir -p mocks/service/logs
mockgen --source=service/logs/base.go -package=mock_service --destination=mocks/service/logs/base.go
# Service
mkdir -p mocks/service/omdb
mockgen --source=service/omdb/base.go -package=mock_service --destination=mocks/service/omdb/base.go