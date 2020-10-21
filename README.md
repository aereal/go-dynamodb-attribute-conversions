![CI][ci-status]
[![PkgGoDev][pkg-go-dev-badge]][pkg-go-dev]

# go-dynamodb-attribute-conversions

```sh
go get github.com/aereal/go-dynamodb-attribute-conversions
```

## Usage

```go
import (
  "context"

  "github.com/aereal/go-dynamodb-attribute-conversions"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func handler(ctx context.Context, event events.DynamoDBEvent) error {
  for _, record := range event.Records {
    m := ddbconversions.AttributeValueMapFrom(record.Change.NewImage)
    var item struct{
      Bool bool
      Str string
    }
    dynamodbattribute.UnmarshalMap(m, &item)
  }
  return nil
}
```

## License

See LICENSE file.

[pkg-go-dev]: https://pkg.go.dev/github.com/aereal/go-dynamodb-attribute-conversions
[pkg-go-dev-badge]: https://pkg.go.dev/badge/aereal/go-dynamodb-attribute-conversions
[ci-status]: https://github.com/aereal/go-dynamodb-attribute-conversions/workflows/CI/badge.svg?branch=main
