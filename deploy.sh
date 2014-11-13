#!/bin/bash

git reset --hard
git pull

cd public/style
for file in *.css; do
  minify --output temp $file
  mv temp $file
done

cd ../js
for file in *.js; do
  minify --output temp $file
  mv temp $file
done

cd ../..
go build && ./gogobbles >>logs 2>>errors
