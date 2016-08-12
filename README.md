# goccupied

This is a simple webserver written in [golang](https://golang.org/). It is designed to run on a Raspberry Pi and 
access its GPIO pins.
 
The server periodically checks GPIO14 for input. The input is expected to come from a brightness sensor that sends a
Low signal when light is above a certain threshold and a high signal otherwise.

The state can be queried by a http request to `http://<hostname>:8000/api/v1/occupied` which also returns the duration
of the current state in seconds. The sensor can be queried directly by calling `http://<hostname>:8000/api/v1/light`

Compile this code with `GOOS=linux` and `GOARCH=arm`