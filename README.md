# Photosplatter
Photosplatter serves an HTML index of photos from a directory.

## Design Goals
Photosplatter is extremely simple to operate.
Small, single-board servers can easily run it at acceptable performance. 
Provided a path and a port number, it will find photos in the directory and display them.
It's a single binary that needs no other files or servers in order to work.

## Operation
Place `photosplatter` somewhere and run it:

```shell
./photosplatter --path ~/images --port 8080
```

See `photosplatter.service` for an optional systemd user unit.
Systems running systemd may place this file into `~/.config/systemd/user`, then `systemctl --user enable --now photosplatter` to run Photosplatter automatically as a normal user.
