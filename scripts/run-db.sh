#!/bin/zsh
current_folder=$( cd "$(dirname "$0")" ; pwd -P )
cd ${current_folder}/../resources/docker
docker-compose up