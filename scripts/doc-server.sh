#!/bin/bash
PORT=8080
echo start doc server http://localhost:$PORT
godoc -http=:$PORT
