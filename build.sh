#!/bin/bash
docker build --pull --rm -f "golang-api/Dockerfile" -t go-api:latest "golang-api"
docker build --pull --rm -f "golang-db/Dockerfile" -t go-db:latest "golang-db"
