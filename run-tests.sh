rm -rf ./test
mkdir ./test

go test ./graph
go test -coverprofile=./test/coverage.out ./graph
go tool cover -html=./test/coverage.out