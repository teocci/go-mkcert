#!/bin/bash
## Build the main
go build main.go

## Rename main as a mkcert
mv -v main mkcert

#cp -v proctel /home/rtt/apps/proctel