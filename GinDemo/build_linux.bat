
::Build For Linux

echo "Clean for build..."
go clean

echo "Build For Linux..."

::set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64

go build -o GinDemo


echo "--------- Build For Linux Success!"


pause

