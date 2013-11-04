#! /bin/bash

go build stackwatcher.go types.go notifications.go utils.go pipeline.go cache.go && ./stackwatcher && rm stackwatcher