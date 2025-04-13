# PROTO ###

proto_book := book.proto

proto_book: $(proto_book)
	protoc --proto_path=protos/book --go_out=gen/go --go-grpc_out=gen/go $(proto_book)