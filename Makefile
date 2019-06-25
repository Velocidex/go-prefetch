all:
	go build  -o prefetch cmd/*.go

gen:
	binparsegen conversion.spec.yaml  > prefetch_gen.go

test:
	go test ./...

windows:
	GOOS=windows GOARCH=amd64 go build -o prefetch.exe cmd/*.go
