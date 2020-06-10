# Excercise 1: Get the scaffold running

The first excercise is pretty simple. A scaffold has been provided to you and we want to make sure you are able to run it. This will validate your Go installation and make sure all dependencies are in order.

Change directory inside the Scaffold directory. We'll promise it won't remain a scaffold for long.

```bash
cd Excercises/RESTService/Scaffold
go run main.go
```

In a browser or with curl, you should now be able to get the fake homepage we provided to test the service.

```bash
bas@DESKTOP-RFVONSL: /mnt/c/data/golang-bits-n-bytes $ curl localhost:8080
Hello SynTouch%
bas@DESKTOP-RFVONSL: /mnt/c/data/golang-bits-n-bytes $ curl localhost:8080/health
{"Status":"OK"}
```

There is a chance, when running on Windows and not using the Docker container provided, an error will pop up when you try and run the Go application.

```bash
# github.com/mattn/go-sqlite3
exec: "gcc": executable file not found in %PATH%
```

This is due to the sqlite driver we will use later in the workshop. In order to resolve this, you need to install a C compiler. One that has been tested and works can be downloaded [HERE](https://jmeubank.github.io/tdm-gcc/).

Installation is pretty straight forward but [this blog](https://medium.com/@yaravind/go-sqlite-on-windows-f91ef2dacfe) explains how to do it when in doubt. Most important part is to add the C compiler to your PATH variable. After you installed the compiler, make sure to restart any terminals and editors so it is picked up correctly.