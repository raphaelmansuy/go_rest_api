#!/usr/bin/env bash

# Install the swag tool
go get -u github.com/swaggo/swag/cmd/swag

# Run the swag tool to generate the doc.json file
swag init

