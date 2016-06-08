Telegram Chat Bot Emulator
==========================

[![Go Report Card](https://goreportcard.com/badge/github.com/pavel-paulau/howdy)](https://goreportcard.com/report/github.com/pavel-paulau/howdy)

Demo
----

[![Howdy Demo](http://i.imgur.com/yXufef8.png)](http://www.youtube.com/watch?v=nVvvGp0HtbQ)

Installation
------------

To get the binary just download the latest release for your OS/Arch from the [release page](https://github.com/pavel-paulau/howdy/releases) and put the binary somewhere convenient.

You can run it from any location:

```
$ ./howdy_linux_amd64 

	.:: Please go to http://127.0.0.1:8081/index.html ::.
```

Alternatively, you can pull the latest [Docker image](https://hub.docker.com/r/pavel/howdy/):

```
$ docker pull pavel/howdy
$ docker run -t -i --net="host" -p 8081:8081 pavel/howdy

	.:: Please go to http://127.0.0.1:8081/index.html ::.
```
