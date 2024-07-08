#!/bin/bash

# Get the list of changed files
if git rev-parse HEAD~1 >/dev/null 2>&1; then
  git diff --name-only HEAD~1 HEAD -- '*.md' > changed_files.txt
else
  git ls-tree --name-only -r HEAD -- '*.md' > changed_files.txt
fi

# Display file names and contents using git cat-file
while IFS= read -r file; do
  echo "File: $file"
  echo "Content:"
  git show "HEAD:$file"
  echo ""
done < changed_files.txt