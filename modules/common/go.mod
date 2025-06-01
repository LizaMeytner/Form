module github.com/LizaMeytner/Form/modules/common

go 1.21

require (
	github.com/gorilla/websocket v1.5.1
	github.com/lib/pq v1.10.9
	github.com/rs/zerolog v1.32.0
	github.com/LizaMeytner/Form/modules/proto v0.0.0
	google.golang.org/grpc v1.62.1
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace github.com/LizaMeytner/Form/modules/proto => ../proto 