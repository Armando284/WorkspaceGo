#!/bin/bash

# Set the absolute path to the Go executable
GO_EXECUTABLE="\"/c/Program Files/Go/bin/go.exe\""

# Run tests
eval $GO_EXECUTABLE test ./...

# Build the executable
eval $GO_EXECUTABLE build -o ./output/wkgo.exe ./...

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Build successful"
else
  echo "Build failed"
fi
