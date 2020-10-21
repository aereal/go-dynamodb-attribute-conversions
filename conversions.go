package ddbconversions

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func AttributeValueFrom(from events.DynamoDBAttributeValue) *dynamodb.AttributeValue {
	av := &dynamodb.AttributeValue{}
	switch from.DataType() {
	case events.DataTypeBinary:
		av.SetB(from.Binary())
	case events.DataTypeBinarySet:
		av.SetBS(from.BinarySet())
	case events.DataTypeBoolean:
		av.SetBOOL(from.Boolean())
	case events.DataTypeList:
		vs := []*dynamodb.AttributeValue{}
		for _, v := range from.List() {
			lv := AttributeValueFrom(v)
			vs = append(vs, lv)
		}
		av.SetL(vs)
	case events.DataTypeMap:
		mv := map[string]*dynamodb.AttributeValue{}
		for k, v := range from.Map() {
			mv[k] = AttributeValueFrom(v)
		}
		av.SetM(mv)
	case events.DataTypeNull:
		av.SetNULL(from.IsNull())
	case events.DataTypeNumber:
		av.SetN(from.Number())
	case events.DataTypeNumberSet:
		var ns []*string
		for _, v := range from.NumberSet() {
			n := v
			ns = append(ns, &n)
		}
		av.SetNS(ns)
	case events.DataTypeString:
		av.SetS(from.String())
	case events.DataTypeStringSet:
		var ss []*string
		for _, v := range from.StringSet() {
			s := v
			ss = append(ss, &s)
		}
		av.SetSS(ss)
	}
	return av
}

func AttributeValueMapFrom(from map[string]events.DynamoDBAttributeValue) map[string]*dynamodb.AttributeValue {
	result := map[string]*dynamodb.AttributeValue{}
	for k, v := range from {
		result[k] = AttributeValueFrom(v)
	}
	return result
}
