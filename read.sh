#!/bin/bash

path="/home/parsa/projects/templates/named_pipes/"
FILE="${path}pipes/main.pipe"
if [ -p "$FILE" ]; then
    echo "$FILE exists."
else 
    mkfifo $FILE
fi
#start reading from pipe
while true
do
    if read line <$FILE; then
        #do something with $line
        #can pass this to the model so it knows what to do
        echo $line
    fi
done

# container > pipe > reader > model > pipe2 > container