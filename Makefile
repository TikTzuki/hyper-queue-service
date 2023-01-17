gen_dir := .gen
hello := ${gen_dir}/helloworld
agent := ${gen_dir}/agent

install-go:
	wget https://dl.google.com/go/go1.19.5.linux-amd64.tar.gz
	rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.5.linux-amd64.tar.gz
	export PATH=$PATH:/usr/local/go/bin
	go version
setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	apt install -y protobuf-compiler || true
	protoc --version
gen:
	mkdir ${gen_dir} || true
	mkdir ${hello} || true
	mkdir ${agent} || true
	protoc --proto_path=proto \
--go_out=${hello} --go_opt=paths=source_relative  \
--go-grpc_out=${hello} --go-grpc_opt=paths=source_relative \
  proto/helloworld.proto
	protoc --proto_path=proto \
--go_out=${agent} --go_opt=paths=source_relative  \
--go-grpc_out=${agent} --go-grpc_opt=paths=source_relative \
  proto/queue/circular/agent/*.proto
clean:
	rm -rf ${gen_dir}
run-hello:
	go run helloworld/main.go
run-server:
	go run server/main.go
run-client:
	go run client/main.go
