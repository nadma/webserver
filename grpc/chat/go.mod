module github.com/tutorialedge/go-grpc-beginners-tutorial

go 1.15

require (
	github.com/tutorialedge/go-grpc-tutorial v0.0.0-20200509091100-f8d1b5b15b01 // indirect
	golang.org/x/net v0.0.0-20210119194325-5f4716e94777
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
)

//replace github.com/tutorialedge/go-grpc-beginners-tutorial/chat => ./.
