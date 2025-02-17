# Allow pasting in Proxmox console

After experiencing the hassle of having typed SHA256 hashes in the Proxmox console, I decided to write a small helper script to make it easier to paste in the console. This script will take advantage of ydotool program to simulate keyboard input on wayland session.

## Requirements
- [ydotool](https://github.com/ReimuNotMoe/ydotool)
- [go](https://golang.org)

## How to use

1. Clone the repository
2. Run `go build` to build the binary
3. Run the binary
```
./proxmox-helper
```
4. Paste the text you want to type in the console
5. You have 10 seconds to focus on the Proxmox console or any other window
