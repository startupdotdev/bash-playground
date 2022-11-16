#! /bin/bash
if [ -z $1 ]; then
  read -p "Enter a file to look for: " file
else
  file="$1"
fi

if test -f $file; then
  echo "File exists"
else
  echo "File doesn't exist"
fi