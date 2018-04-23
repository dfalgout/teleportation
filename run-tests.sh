FOLDER=./test

rm -rf $FOLDER
mkdir $FOLDER

go test -coverprofile=$FOLDER/coverage.out ./graph
go tool cover -html=$FOLDER/coverage.out