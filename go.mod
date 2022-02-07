module github.com/cdecoux/golang-backpack

go 1.17

replace github.com/cdecoux/golang-backpack/protobuf => ./protobuf

require github.com/sirupsen/logrus v1.8.1

require github.com/cdecoux/golang-backpack/protobuf v0.0.0

require (
	github.com/golang/protobuf v1.5.0 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.44.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
