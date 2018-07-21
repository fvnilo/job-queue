build: cmd/listener/main.go cmd/publisher/main.go
		go build -o bin/listener cmd/listener/main.go && \
		go build -o bin/publisher cmd/publisher/main.go 

clean: bin
		rm -rf ./bin