run:
	go run main.go

test:
	go test -v ./src/domains/images
	go test -v ./src/infrastructures/storage
	go test -v ./src/infrastructures/db/repositories

test-no-cache:
	go clean -cache
	make test

clean:
	rm -rf ./bin ./vendor Gopkg.lock

localstack-setup:
	aws --profile localstack --endpoint-url=${LOCALSTACK_URL_S3} s3 mb s3://${S3_BUCKET}
	aws --profile localstack --endpoint-url=${LOCALSTACK_URL_S3} s3api put-bucket-acl --bucket ${S3_BUCKET} --acl public-read

localstack-uploads:
	aws --endpoint-url ${LOCALSTACK_URL_S3} s3 ls s3://${S3_BUCKET}
