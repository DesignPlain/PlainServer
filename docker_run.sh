#!/bin/sh

## Script to start and stop design-sphere docker container
## Usage: docker_run start|stop

echo "Running command [ $0 $@ ] "
if [[ $1 == "start" ]]; then
    # Remove existing container
    echo "\n# [ Removing existing stopped containers of name ds-container ]"
    docker container rm ds-container

    # Remove all the design-sphere related images
    # echo "\n# [ Removing docker image with label 'ds' ]"
    # docker image prune --all --force --filter label=image="ds"

    # Copy, build the app and create a docker image tagger `ds-server`
    echo "\n# [ Building docker image ds-server]"
    docker build -t ds-server .

    # Run the docker container with name `ds-container`
    echo "\n# [ Starting docker image ds-server in container ds-container ]"
    docker run -dp 127.0.0.1:8080:8080 --name ds-container ds-server
elif [[ $1 == "stop" ]]; then
    echo "\n# [ Stopping docker container ]"
    docker container stop ds-container
else
    echo "\nUsage: ./docker_run start|stop\n"
fi

echo "\n# [ Done ]"
