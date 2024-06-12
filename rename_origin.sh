#!/bin/bash
for file in *_origin*; do
  mv "$file" "${file//_origin/}"
done
