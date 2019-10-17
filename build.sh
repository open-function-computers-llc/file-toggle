#!/bin/bash

echo -n "Killing any previously running build... "
killall file-toggle || true
echo "Done"

echo "Building"
go build -o dist/file-toggle
./dist/file-toggle &
echo "App running..."
