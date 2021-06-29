#! /bin/bash

# FROM https://towardsdev.com/scalable-event-streaming-with-redis-streams-and-go-dee5fbe8982c
# GITHUB: https://github.com/gmrdn/redis-streams-go

# pull docker image of redis and launch it
if [[ $( docker images | grep redis) ]]; then
    echo "redis image already pulled."
else
	echo "pulling redis docker image"
	docker pull redis    
fi

if [[ $( docker ps -a | grep redis-streams | head -c12 ) ]]; then
    echo "redis-streams container already present..."
    if [[ $(docker ps | grep redis-streams | head -c12 ) ]]; then 
    	echo "...and running!"
    else
    	docker start redis-streams
    	echo "... starting container"
    fi
else
	docker run --name redis-streams -p 6379:6379 -d redis --requirepass "superSecret"
fi

# build executable
cd ./publisher && go build -o redis-publisher && cd ../
cd ./consumer && go build -o redis-consumer && cd ../

# launch 1 publisher app
gnome-terminal -x sh -c "./publisher/redis-publisher; bash"

echo "wait for the publisher to be ready..."
sleep 5s

# launch 2 consumer apps
gnome-terminal -x sh -c "./consumer/redis-consumer; bash"
gnome-terminal -x sh -c "./consumer/redis-consumer; bash"

