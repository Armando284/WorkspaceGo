#!/bin/bash

# Run tests
go test ./...

# Build the executable
go build -o ./output/wkgo.exe ./...

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
