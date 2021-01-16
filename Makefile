.PHONY: clean $(targets)

.DEFAULT: osg

osg:
	go build -o bin/$@ cmd/$@/main.go

clean:
	find . -iname *.go -type f -exec gofmt -w -s {} \;
	go mod tidy
	rm -rf ./bin
