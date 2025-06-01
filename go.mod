module github.com/LizaMeytner/Form/modules/common

go 1.21

require (
	github.com/google/uuid v1.6.0
	github.com/gorilla/websocket v1.5.1
	github.com/lib/pq v1.10.9
	github.com/rs/zerolog v1.32.0
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
)

replace github.com/LizaMeytner/Form/modules/proto => ../proto
