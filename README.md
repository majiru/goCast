# GoCast
A simple set of golang scripts to stream from one computer to another.
This program works by setting up a html server on the file server and then sending the remote paht to the player to open the file.
This was designed to allow for other devices besides the player and server to control the player through a web interface.

## Setup
To start the server instance on the file server run: `go run server.go /path/to/media playerIP`

Then, to start the client you first have to start mpv with:
`mpv --idle --input-ipc-server=/tmp/mpvsocket`

Then start the client listener: `go run client.go`

## Useage
Simple browse to the server's ip in any web browser to control the player. 
There are simple controls at the top of the page{play, pause next} that will individual commands to the player.
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
[Blang](www.github.com/blang) for his awesome mpv api.
