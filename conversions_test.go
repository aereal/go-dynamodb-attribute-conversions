package ddbconversions

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func TestAttributeValueFrom(t *testing.T) {
	tests := []struct {
		name string
		from events.DynamoDBAttributeValue
		want *dynamodb.AttributeValue
	}{
		{"binary", events.NewBinaryAttribute([]byte("abc")), &dynamodb.AttributeValue{B: []byte("abc")}},
		{"binary set", events.NewBinarySetAttribute([][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}}), &dynamodb.AttributeValue{BS: [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}}}},
		{"boolean", events.NewBooleanAttribute(false), &dynamodb.AttributeValue{BOOL: aws.Bool(false)}},
		{"list", events.NewListAttribute([]events.DynamoDBAttributeValue{events.NewBooleanAttribute(true), events.NewStringAttribute("a")}), &dynamodb.AttributeValue{L: []*dynamodb.AttributeValue{{BOOL: aws.Bool(true)}, {S: aws.String("a")}}}},
		{"empty list", events.NewListAttribute([]events.DynamoDBAttributeValue{}), &dynamodb.AttributeValue{L: []*dynamodb.AttributeValue{}}},
		{"map", events.NewMapAttribute(map[string]events.DynamoDBAttributeValue{"bool": events.NewBooleanAttribute(true), "string": events.NewStringAttribute("a")}), &dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{"bool": {BOOL: aws.Bool(true)}, "string": {S: aws.String("a")}}}},
		{"empty map", events.NewMapAttribute(nil), &dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{}}},
		{"null", events.NewNullAttribute(), &dynamodb.AttributeValue{NULL: aws.Bool(true)}},
		{"number", events.NewNumberAttribute("1"), &dynamodb.AttributeValue{N: aws.String("1")}},
		{"number set", events.NewNumberSetAttribute([]string{"1", "2"}), &dynamodb.AttributeValue{NS: []*string{aws.String("1"), aws.String("2")}}},
		{"string", events.NewStringAttribute("a"), &dynamodb.AttributeValue{S: aws.String("a")}},
		{"string set", events.NewStringSetAttribute([]string{"a", "b"}), &dynamodb.AttributeValue{SS: []*string{aws.String("a"), aws.String("b")}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AttributeValueFrom(tt.from); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AttributeValueFrom() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestAttributeValueMapFrom(t *testing.T) {
	tests := []struct {
		name string
		from map[string]events.DynamoDBAttributeValue
		want map[string]*dynamodb.AttributeValue
	}{
		{"ok", map[string]events.DynamoDBAttributeValue{"bool": events.NewBooleanAttribute(true), "string": events.NewStringAttribute("a")}, map[string]*dynamodb.AttributeValue{"bool": {BOOL: aws.Bool(true)}, "string": {S: aws.String("a")}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AttributeValueMapFrom(tt.from); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AttributeValueMapFrom() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
