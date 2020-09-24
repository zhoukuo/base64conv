version=1.0

rm *.zip -f

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build base64conv.go
zip base64conv.linux64-${version}.zip ./* -x *.zip *.go *.sh
rm base64conv

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build base64conv.go
zip base64conv.darwin-${version}.zip ./* -x *.zip *.go *.sh
rm base64conv

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build base64conv.go
zip base64conv.windows64-${version}.zip ./* -x *.zip *.go *.sh
rm base64conv.exe