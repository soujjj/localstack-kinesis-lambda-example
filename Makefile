kinesis-put-records:
	docker-compose exec localstack awslocal kinesis put-records --cli-input-json file:///producer/put_records.json

lambda-tail-log:
	aws logs --profile localstack --endpoint-url=http://localhost:4566  tail /aws/lambda/local-lambda --follow

confirm-sqs-message:
	aws sqs \
		--profile localstack \
		--endpoint-url=http://localhost:4566 \
		get-queue-attributes \
		--queue-url http://localhost:4566/000000000000/local-lambda-mapping-failre-sqs.fifo \
		--attribute-names ApproximateNumberOfMessages \
		--query 'Attributes.ApproximateNumberOfMessages' \
		--output text
