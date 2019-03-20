# что делать по команде make по умолчанию
all: app

app: x-gen

# скомпилировать app
x-gen: dep *.go docker.makefile
	@env GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o x-gen -v
	@chmod +x x-gen

# установить зависимости
dep:
	@go get -v -d ./...
