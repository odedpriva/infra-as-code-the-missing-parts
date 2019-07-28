# localstack

## prerequisites

- point .localhost TLD to localhost

## setup the environment

```shell
>> TMPDIR=/private$TMPDIR docker-compose up
>> terraform init
>> terraform apply
```

## check queue was created successfully

```shell
>> aws --endpoint-url=http://localhost:4576 sqs list-queues
```
