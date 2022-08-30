# localstack-kinesis-lambda-example

```
$ aws configure --profile localstack
AWS Access Key ID [None]: dummy
AWS Secret Access Key [None]: dummy
Default region name [None]: us-east-1
Default output format [None]: text

$ cat ~/.aws/credentials
[localstack]
aws_access_key_id = dummy
aws_secret_access_key = dummy

$ cat ~/.aws/config
[profile localstack]
region = us-east-1
output = text
```

init
```bash
$ zip -r ./stacks/function.zip ./consumer/index.js
$ docker-compose up -d
$ docker-compose run --rm terraform init
$ docker-compose run --rm terraform plan
$ docker-compose run --rm terraform apply --auto-approve
```

confirm kinesis stream
```
aws kinesis --profile localstack --endpoint http://localhost:4566 describe-stream-summary --stream-name local-stream
```

put record kinesis stream
```
docker-compose exec localstack awslocal kinesis put-records --cli-input-json file:///producer/put_records.json
```

update lambda function
```bash
docker-compose exec localstack bash -c '
cd /consumer
zip -r function.zip index.js node_modules package.json package-lock.json
awslocal lambda update-function-code --function-name="=local-lambda" --zip-file fileb:///consumer/function.zip
'
```

invoke lambda function
```bash
docker-compose exec localstack bash -c "
awslocal lambda invoke --function-name local-lambda --payload file:///consumer/lambda_invoke_payload.json /dev/null --log-type Tail --query 'LogResult' --output text |  base64 -d
"
```

tail follow logs lambda
```
aws logs --profile localstack --endpoint-url=http://localhost:4566  tail /aws/lambda/local-lambda --follow 
```
