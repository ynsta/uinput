all:
	c-for-go -out .. uinput.yml

clean:
	rm -f cgo_helpers.h cgo_helpers.go const.go doc.go types.go

test:
	go build
