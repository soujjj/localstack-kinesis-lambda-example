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

create kinesis stream
```
docker-compose exec localstack awslocal kinesis create-stream --stream-name test-stream --shard-count 1
```

confirm kinesis stream
```
docker-compose exec localstack awslocal kinesis describe-stream-summary --stream-name test-stream
```

create lambda function
```bash
docker-compose exec localstack bash
cd /consumer
zip -r function.zip index.js node_modules package.json package-lock.json
awslocal lambda create-function --function-name="test-kinesis-trigger" --runtime=nodejs12.x --role="arn:aws:iam::123456789012:role/service-role/lambda-sample-role" --handler=index.handler --zip-file fileb://function.zip 
```

update lambda function
```bash
docker-compose exec localstack bash
cd /consumer
zip -r function.zip index.js node_modules package.json package-lock.json
awslocal lambda update-function-code --function-name="test-kinesis-trigger" --zip-file fileb://function.zip
```

set kinesis stream event trigger
```bash
docker-compose exec localstack awslocal lambda create-event-source-mapping \
--function-name test-kinesis-trigger \
--batch-size 500 \
--event-source-arn arn:aws:kinesis:us-east-1:000000000000:stream/test-stream
```

invoke lambda function
```bash
docker-compose exec localstack bash
awslocal lambda invoke --function-name test-kinesis-trigger --payload '{"Records": ["hello", "world"]}' /dev/null --log-type Tail --query 'LogResult' --output text |  base64 -d
```

tail follow logs lambda
```
aws logs --profile localstack --endpoint-url=http://localhost:4566  tail /aws/lambda/test-kinesis-trigger --follow 
```
