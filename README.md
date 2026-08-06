brew install golangci-lint

go mod init example/hello
go run hello.go
go test
go build
go doc fmt


##Writing a test is just like writing a function, with a few rules

    It needs to be in a file with a name like xxx_test.go

    The test function must start with the word Test

    The test function takes one argument only t *testing.T

    To use the *testing.T type, you need to import "testing", like we did with fmt in the other file



Go's second tool for viewing documentation is the pkgsite command, which powers Go's official package viewing website. You can install pkgsite with go install golang.org/x/pkgsite/cmd/pkgsite@latest, then run it with pkgsite -open .. Go's install command will download the source files from that repository and build them into an executable binary.