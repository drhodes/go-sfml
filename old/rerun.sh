#!/bin/bash

WATCH_PATH=`pwd`

while inotifywait -e modify -r ./ 2> /dev/null; do	
	clear
	echo "------------------------------------------------------------------"
	./build.sh

	if [ "$?" -eq "0" ]
	then 
		echo "build succeeds"
		go install
	else 			
		cat /tmp/makelog
		echo "build failure"
	fi
done
