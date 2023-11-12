#!/bin/bash

# Set the absolute path to the Go executable
GO_EXECUTABLE="\"/mnt/c/Program Files/Go/bin/go.exe\""

# Run tests
eval $GO_EXECUTABLE test ./...

# Build the executable
eval $GO_EXECUTABLE build -o ./output/wkgo.exe ./...

# Check if the build was successful
if [ $? -eq 0 ]; then
  echo "Build successful"

  # Add the built executable to Git
  git add ./output/wkgo.exe

  # Commit the build changes
  git commit -m "Feat: Add latest built executable"

  # Push the changes to GitHub
  git push origin main
else
  echo "Build failed"
fi
