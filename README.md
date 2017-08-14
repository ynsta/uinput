# uinput go binding

This uinput [Go](https://golang.org/) binding permits to generate gamepad, mouse and keyboard events under Linux

Most of the constants and structures are extracted with [c-for-go](https://github.com/xlab/c-for-go).

This is a Work in Progress, the code is not yet documented nor tested.

## Setup

```
go get github.com/ynsta/uinput
```

## Example

```
cd ${GOPATH}/src/github.com/ynsta/uinput/examples
go build gamepad.go
```

The user must have the permission to write to `/dev/uinput`, the configuration is distribution specific but should be similar to this :

1. Add your user to the `input` group:
```
sudo gpasswd --add $USER input
```
2. Create `/etc/udev/rules.d/99-uinput.rules` with:

```
KERNEL=="uinput", MODE="0660", GROUP="input", OPTIONS+="static_node=uinput"
```

3. Reload udev `sudo udevadm control --reload`
