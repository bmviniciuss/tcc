#!/bin/bash

FILES="./cenarios/*"
for f in $FILES
do
    for i in {1..1}
    do
        echo "$f:$GRPC_ENABLED"
        k6 run $f -e GENERATE_SUMMARY=true -e GRPC_ENABLED=$GRPC_ENABLED
    done
done
