#! /usr/bin/env sh

while true; do
  ./Elgamal-vs-ECIES -file $(date +%H)/elg.idk.md 2> $(date +%H)/elg.idk.total.md
  ./Elgamal-vs-ECIES -file $(date +%H)/ecies.idk.md 2> $(date +%H)/ecies.idk.total.md
  sleep 4m
done
