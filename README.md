# GoCast
A simple set of golang scripts to stream from a computer with local media to a raspberrypi. The whole interface is controllable from a phone or anything with a web browser.

This program works by setting up a html server on the file server and then sending the remote path to the player to open the file.

Currently the client only supports mpv, but there are plans to support omxplayer as well for better acceleration.

## Setup
First install gocast from the command line with `go get github.com/majiru/gocast`

### Server
To start the server instance on the file server run: `gocast server /path/to/media playerIP`

### Client
To start the client you first have to start mpv with:
`mpv --idle --input-ipc-server=/tmp/mpvsocket`

Then start the client listener: `gocast client`

## Useage
Simple browse to the server's ip in any web browser to control the player. 
There are simple controls at the top of the page{play, pause, next} that will individual commands to the player.
Clicking on a link will either change the web page to read the directory, or send the file to the player to stream.

## Todo

#### Easy Stuff
* Make a less disgusting web page for the file server
* Add external URL support (loading youtube videos through youtube-dl)
* Volume control
* Seek control

#### Hard(er) Stuff
* Have player send back the current playlist to allow users to shift the que
* Display bar with current position in file

## Credits
[Blang](http://www.github.com/blang) for his awesome mpv api.
