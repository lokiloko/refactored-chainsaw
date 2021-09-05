#!/bin/bash

# Controller
mkdir -p mocks/controller
mockgen --source=controller/base.go -package=mock_controller --destination=mocks/controller/base.go