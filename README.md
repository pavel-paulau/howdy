Telegram Chat Bot Emulator
==========================

[![Go Report Card](https://goreportcard.com/badge/github.com/pavel-paulau/howdy)](https://goreportcard.com/report/github.com/pavel-paulau/howdy)

Howdy is an easy-to-use Telegram emulator for developing chat bots on Windows, Mac, and Linux. It doesn't require any certificates, public servers, or real Telegram clients.

Just use your web browser to test and debug chat bots while developing new cool features.

Demo
----

Click on the image below to watch a demonstration of chat bot emulator:

[![Howdy Demo](http://i.imgur.com/yXufef8.png)](http://www.youtube.com/watch?v=nVvvGp0HtbQ)

Installation
------------

Howdy binaries have no external dependencies.

To get the binary just download the latest release for your OS/Arch from the [release page](https://github.com/pavel-paulau/howdy/releases) and put the binary somewhere convenient. You can run it from any location:

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

Usage
-----

Simply follow these steps:
* Modify your bot so that it uses "http://127.0.0.1:8081/" instead of "https://api.telegram.org/"
* Use "token" as an authentication token
* Start both chat bot and bot emulator
* Navigate to "http://127.0.0.1:8081/index.html" in your favourite browser
* Enjoy!
