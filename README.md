# goccupied

This is a simple webserver written in [golang](https://golang.org/). It is designed to run on a Raspberry Pi and 
access its GPIO pins.
 
On request, the server checks GPIO4 for input. The input is expected to come from a lightness sensor that sends a
Low signal when light is above a certain threshold and a high signal otherwise.

Compile this code with `GOOS=freebsd` adn `GOARCH=arm`