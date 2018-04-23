FOLDER=./output

rm -rf $FOLDER
mkdir $FOLDER
go build -o $FOLDER/teleport ./main.go

# Run Program
$FOLDER/teleport