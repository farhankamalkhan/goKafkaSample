#!/usr/bin/env bash

docker build -t kafzoo:portsopen .
docker run -td --network=host kafzoo:portsopen
