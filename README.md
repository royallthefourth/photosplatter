# Photosplatter
Photosplatter serves an HTML index of photos from a directory.

## Design Goals
Photosplatter is extremely simple to operate.
Small, single-board servers can easily run it at acceptable performance. 
Provided a path and a port number, it will find photos in the directory and display them.

* No config files
* No databases
* Automatic content directory rescans
* Only one file to deploy

## Basic Operation
Place `photosplatter` somewhere and run it:

```shell
./photosplatter --path ~/images --port 8080
```

See `photosplatter.service` for an optional systemd user unit.
Systems running systemd may place this file into `~/.config/systemd/user`, then `systemctl --user enable --now photosplatter` to run Photosplatter automatically as a normal user.

## Practical Use
In my home, I have a tiny arm7 server running Syncthing.
Whenever my phone connects to my home wifi, Syncthing sends all of new photos from my phone to my server.
With the photos automatically uploaded to the server, I can use Photosplatter to browse them.
It's much nicer than using a USB cable and keeps my files off of other peoples' servers.
