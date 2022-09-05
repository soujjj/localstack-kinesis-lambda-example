
build:
	GOOS=linux GOARCH=amd64 go build -o handler handler.go
	chmod +x handler

package: build
	zip -r ../stacks/handler.zip ./handler

deploy:
	aws lambda --endpoint http://localhost:4566 --profile localstack update-function-code --function-name="local-lambda" --zip-file fileb://../stacks/handler.zip

python_package:
	zip -r ../stacks/handler.zip ./main.py
