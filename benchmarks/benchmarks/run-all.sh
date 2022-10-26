#!/bin/bash

FILES="./cenarios/*"
for f in $FILES
do
    for i in {1..1}
    do
        echo "$f:$GRPC_ENABLED"
        GRPC_ENABLED=$GRPC_ENABLED docker compose -f ../../docker-compose.yml up -d
        sleep 5
        k6 run $f -e GENERATE_SUMMARY=true -e GRPC_ENABLED=$GRPC_ENABLED
        sleep 5
        docker compose -f ../../docker-compose.yml down -v
        sleep 3
    done

    docker compose -f ../../docker-compose.yml down -v
done
