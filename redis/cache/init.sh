#! /bin/bash

# FROM https://tutorialedge.net/golang/go-redis-tutorial/

# pull docker image of redis and launch it
if [[ $( docker images | grep redis) ]]; then
    echo "redis image already pulled."
else
	echo "pulling redis docker image"
	docker pull redis    
fi

if [[ $( docker ps -a | grep redis-test-instance | head -c12) ]]; then
    echo "redis-test-instance container already present...."
    if [[ $(docker ps | grep redis-test-instance | head -c12 ) ]]; then 
    	echo "...and running!"
    else
    	docker start redis-test-instance
    	echo "... starting container"
    fi
else
	docker run --name redis-test-instance -p 6379:6379 -d redis 
fi

# run program
go build -o redis-tutorial
./redis-tutorial

