package ddbconversions

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestAttributeValueFrom(t *testing.T) {
	tests := []struct {
		name string
		from events.DynamoDBAttributeValue
		want types.AttributeValue
	}{
		{"binary", events.NewBinaryAttribute([]byte("abc")), &types.AttributeValueMemberB{Value: []byte("abc")}},
		{"binary set", events.NewBinarySetAttribute([][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}}), &types.AttributeValueMemberBS{Value: [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}}}},
		{"boolean", events.NewBooleanAttribute(false), &types.AttributeValueMemberBOOL{Value: false}},
		{"list", events.NewListAttribute([]events.DynamoDBAttributeValue{events.NewBooleanAttribute(true), events.NewStringAttribute("a")}), &types.AttributeValueMemberL{Value: []types.AttributeValue{&types.AttributeValueMemberBOOL{Value: true}, &types.AttributeValueMemberS{Value: "a"}}}},
		{"empty list", events.NewListAttribute([]events.DynamoDBAttributeValue{}), &types.AttributeValueMemberL{Value: nil}},
		{"map", events.NewMapAttribute(map[string]events.DynamoDBAttributeValue{"bool": events.NewBooleanAttribute(true), "string": events.NewStringAttribute("a")}), &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{"bool": &types.AttributeValueMemberBOOL{Value: true}, "string": &types.AttributeValueMemberS{Value: "a"}}}},
		{"empty map", events.NewMapAttribute(nil), &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{}}},
		{"null", events.NewNullAttribute(), &types.AttributeValueMemberNULL{Value: true}},
		{"number", events.NewNumberAttribute("1"), &types.AttributeValueMemberN{Value: "1"}},
		{"number set", events.NewNumberSetAttribute([]string{"1", "2"}), &types.AttributeValueMemberNS{Value: []string{"1", "2"}}},
		{"string", events.NewStringAttribute("a"), &types.AttributeValueMemberS{Value: "a"}},
		{"string set", events.NewStringSetAttribute([]string{"a", "b"}), &types.AttributeValueMemberSS{Value: []string{"a", "b"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AttributeValueFrom(tt.from)
			if diff := cmp.Diff(tt.want, got, ignoreOpts); diff != "" {
				t.Errorf("-want, +got:\n%s", diff)
			}
		})
	}
}

func TestAttributeValueMapFrom(t *testing.T) {
	tests := []struct {
		name string
		from map[string]events.DynamoDBAttributeValue
		want map[string]types.AttributeValue
	}{
		{"ok", map[string]events.DynamoDBAttributeValue{"bool": events.NewBooleanAttribute(true), "string": events.NewStringAttribute("a")}, map[string]types.AttributeValue{"bool": &types.AttributeValueMemberBOOL{Value: true}, "string": &types.AttributeValueMemberS{Value: "a"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AttributeValueMapFrom(tt.from)
			if diff := cmp.Diff(tt.want, got, ignoreOpts); diff != "" {
				t.Errorf("-want, +got:\n%s", diff)
			}
		})
	}
}

var ignoreOpts = cmpopts.IgnoreUnexported(
	types.AttributeValueMemberS{},
	types.AttributeValueMemberN{},
	types.AttributeValueMemberB{},
	types.AttributeValueMemberSS{},
	types.AttributeValueMemberNS{},
	types.AttributeValueMemberBS{},
	types.AttributeValueMemberM{},
	types.AttributeValueMemberL{},
	types.AttributeValueMemberNULL{},
	types.AttributeValueMemberBOOL{},
)
