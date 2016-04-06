#!/bin/bash

ls ./presenter 2> /dev/null
if [ $? -gt 0 ]
then
  echo "Building server..."
  go build
fi
./presenter
