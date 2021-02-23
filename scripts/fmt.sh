#!/usr/bin/env sh

find . \
  -type d -name .git -prune -o \
  -type f -name "*.go" -print0 |
  xargs -0 gofumpt -l -s -w
