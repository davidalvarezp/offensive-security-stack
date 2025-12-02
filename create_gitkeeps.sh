#!/bin/bash
# create_gitkeeps.sh
# Adds a .gitkeep file in every empty folder to allow pushing to GitHub

echo "Adding .gitkeep to empty directories..."

# Find all directories
find . -type d | while read dir; do
    # Check if directory is empty
    if [ -z "$(ls -A "$dir")" ]; then
        touch "$dir/.gitkeep"
        echo "Created .gitkeep in $dir"
    fi
done

echo "Done!"
