#! /bin/bash

# FROM https://itnext.io/redis-as-a-pub-sub-engine-in-go-10eb5e6699cc

# pull docker image of redis and launch it
if [[ $( docker images | grep redis) ]]; then
    echo "redis image already pulled."
else
	echo "pulling redis docker image"
	docker pull redis    
fi

if [[ $( docker ps -a | grep redis-pubsub | head -c12 ) ]]; then
    echo "redis-pubsub container already present..."
    if [[ $(docker ps | grep redis-pubsub | head -c12 ) ]]; then 
    	echo "...and running!"
    else
    	docker start redis-pubsub
    	echo "... starting container"
    fi
else
	docker run --name redis-pubsub -p 6379:6379 -d redis --requirepass "superSecret"
fi

# build executable
cd ./pub && go build -o redis-publisher && cd ../
cd ./sub && go build -o redis-subscriber && cd ../

# launch 1 publisher app
gnome-terminal -x sh -c "./pub/redis-publisher; bash"

# launch 2 subscriber apps
gnome-terminal -x sh -c "./sub/redis-subscriber; bash"
gnome-terminal -x sh -c "./sub/redis-subscriber; bash"


