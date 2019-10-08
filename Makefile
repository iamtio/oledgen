sprites: bindata.go
bindata.go: sprites.png
	go-bindata -nocompress sprites.png
