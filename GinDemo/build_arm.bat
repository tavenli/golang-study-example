
echo "Clean for build..."
go clean

echo "Build For linux-arm..."

set GOOS=linux
set GOARCH=arm

go build -o GinDemo

echo "--------- Build For linux-arm Success!"


pause

