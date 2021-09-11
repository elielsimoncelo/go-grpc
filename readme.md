# Install packages & tools
- go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
- go install google.golang.org/protobuf/cmd/protoc-gen-go
- go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Install dependencies
- go mod tidy

# Generate protos
- protoc --proto_path=proto proto/*.proto --go_out=pb
- protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb # execute this command always that the .proto files changed

# GRPC Clients
- evans (https://github.com/ktr0731/evans) # Macos: brew tap ktr0731/evans // brew install evans
- bloomrpc https://github.com/uw-labs/bloomrpc # Macos: brew install --cask bloomrpc
- insomnia https://insomnia.rest/download # Download: https://insomnia.rest/download
- kreya (https://kreya.app/) # Download: https://kreya.app/downloads/
- wombat https://github.com/rogchap/wombat # Macos: brew install --cask wombat

# Using evans
- evans -r repl --host localhost --port 50051
