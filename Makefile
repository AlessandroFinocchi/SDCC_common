.PHONY: proto_gen proto_clean cert

# If pb folder doesn't exists, it is created
proto_gen:
	[ ! -d pb ] && mkdir "pb" || echo "pb already exists"
	protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb

proto_clean:
	rm pb/*

cert:
	cd cert; chmod +x gen.sh; ./gen.sh; cd ..
