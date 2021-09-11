# Installs
go mod tidy

# Configurations
- protoc --proto_path=proto proto/*.proto --go_out=pb
- protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb

# GRPC Clients
- evans (https://github.com/ktr0731/evans) # Macos: brew tap ktr0731/evans // brew install evans
- bloomrpc https://github.com/uw-labs/bloomrpc # Macos: brew install --cask bloomrpc
- insomnia https://insomnia.rest/download # Download: https://insomnia.rest/download
- kreya (https://kreya.app/) # Download: https://kreya.app/downloads/
- wombat https://github.com/rogchap/wombat # Macos: brew install --cask wombat

# Use Evans
- evans -r repl --host localhost --port 50051