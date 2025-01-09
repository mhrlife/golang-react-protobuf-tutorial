PROTOC_GEN_TS_PROTO_PATH="./node_modules/.bin/protoc-gen-ts_proto"


dependencies:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go get google.golang.org/protobuf
	npm install ts-proto
	go install github.com/srikrsna/protoc-gen-gotag@latest


proto-ts:
	protoc \
		--plugin="protoc-gen-ts_proto=${PROTOC_GEN_TS_PROTO_PATH}" \
		--ts_proto_opt="esModuleInterop=true,forceLong=long" \
		--ts_proto_out="./ui/src" \
		./schema/*.proto

proto-go:
	protoc --gotag_out=. ./schema/*.proto