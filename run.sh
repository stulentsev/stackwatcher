#! /bin/bash

go build stackwatcher.go types.go notifications.go utils.go pipeline.go cache.go && ./stackwatcher && rm stackwatcher

if [[ $? != 0 ]]; then
  say stackwatcher has exited;
fi
