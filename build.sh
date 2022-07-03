#!/bin/bash
docker build --build-arg ALL="$(cat .git/refs/heads/master)+$(date +%c)" --pull --rm -f "golang-api/Dockerfile" -t go-api:latest "golang-api"
docker build --build-arg ALL="$(cat .git/refs/heads/master)+$(date +%c)" --pull --rm -f "golang-db/Dockerfile" -t go-db:latest "golang-db"