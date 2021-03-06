cd $1


go mod vendor # only useful if generated code accessible by Build Tools imports

protoc \
-I . \
-I ./vendor \
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \
$1.proto

rm -rf vendor