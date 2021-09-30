package ddbconversions

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// AttributeValueFrom converts from events.DynamoDBAttributeValue to dynamodb.AttributeValue
func AttributeValueFrom(from events.DynamoDBAttributeValue) types.AttributeValue {
	switch from.DataType() {
	case events.DataTypeBinary:
		return &types.AttributeValueMemberB{Value: from.Binary()}
	case events.DataTypeBinarySet:
		return &types.AttributeValueMemberBS{Value: from.BinarySet()}
	case events.DataTypeBoolean:
		return &types.AttributeValueMemberBOOL{Value: from.Boolean()}
	case events.DataTypeList:
		vs := &types.AttributeValueMemberL{}
		for _, v := range from.List() {
			lv := AttributeValueFrom(v)
			vs.Value = append(vs.Value, lv)
		}
		return vs
	case events.DataTypeMap:
		mv := &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{}}
		for k, v := range from.Map() {
			mv.Value[k] = AttributeValueFrom(v)
		}
		return mv
	case events.DataTypeNull:
		return &types.AttributeValueMemberNULL{Value: from.IsNull()}
	case events.DataTypeNumber:
		return &types.AttributeValueMemberN{Value: from.Number()}
	case events.DataTypeNumberSet:
		ns := &types.AttributeValueMemberNS{}
		for _, v := range from.NumberSet() {
			n := v
			ns.Value = append(ns.Value, n)
		}
		return ns
	case events.DataTypeString:
		return &types.AttributeValueMemberS{Value: from.String()}
	case events.DataTypeStringSet:
		ss := &types.AttributeValueMemberSS{}
		for _, v := range from.StringSet() {
			s := v
			ss.Value = append(ss.Value, s)
		}
		return ss
	default:
		panic(errors.New("unknown type"))
	}
}

// AttributeValueMapFrom converts from events.DynamoDBAttributeValuemap to dynamodb.AttributeValue map
func AttributeValueMapFrom(from map[string]events.DynamoDBAttributeValue) map[string]types.AttributeValue {
	result := map[string]types.AttributeValue{}
	for k, v := range from {
		result[k] = AttributeValueFrom(v)
	}
	return result
}
