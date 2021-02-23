#! /usr/bin/env sh

nohup ./elg.idk.sh 2>elg.log.err 1>elg.log &

nohup ./ecies.idk.sh 2>ecies.log.err 1>ecies.log &

tail -f *.log*
