GOOS=linux GOARCH=amd64 go build -o build/lambda -v github.com/mctofu/pushgate/cmd/lambda
build-lambda-zip --output build/lambda.zip build/lambda