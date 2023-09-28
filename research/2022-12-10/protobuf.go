package protobuf

import (
	"errors"

	"github.com/mxrch/rosso/research/quotedprintable"
	"google.golang.org/protobuf/encoding/protowire"
)

type bytes []byte

func (b bytes) MarshalText() ([]byte, error) {
	return quotedprintable.Encode_Bytes(b), nil
}

type field struct {
	protowire.Number
	Value any
}

type raw struct {
	Bytes  bytes
	Fields []field
}

func unmarshal(buf []byte) ([]field, error) {
	var fields []field
	for len(buf) >= 1 {
		num, typ, length := protowire.ConsumeTag(buf)
		err := protowire.ParseError(length)
		if err != nil {
			return nil, err
		}
		buf = buf[length:]
		switch typ {
		case protowire.Fixed32Type:
			var value uint32
			value, length = protowire.ConsumeFixed32(buf)
			err := protowire.ParseError(length)
			if err != nil {
				return nil, err
			}
			fields = append(fields, field{num, value})
		case protowire.Fixed64Type:
			var value uint64
			value, length = protowire.ConsumeFixed64(buf)
			err := protowire.ParseError(length)
			if err != nil {
				return nil, err
			}
			fields = append(fields, field{num, value})
		case protowire.VarintType:
			var value uint64
			value, length = protowire.ConsumeVarint(buf)
			err := protowire.ParseError(length)
			if err != nil {
				return nil, err
			}
			fields = append(fields, field{num, value})
		case protowire.BytesType:
			var value raw
			value.Bytes, length = protowire.ConsumeBytes(buf)
			err := protowire.ParseError(length)
			if err != nil {
				return nil, err
			}
			value.Fields, _ = unmarshal(value.Bytes)
			fields = append(fields, field{num, value})
		case protowire.StartGroupType:
			var value raw
			value.Bytes, length = protowire.ConsumeGroup(num, buf)
			err := protowire.ParseError(length)
			if err != nil {
				return nil, err
			}
			value.Fields, err = unmarshal(value.Bytes)
			if err != nil {
				return nil, err
			}
			fields = append(fields, field{num, value})
		default:
			return nil, errors.New("cannot parse reserved wire type")
		}
		buf = buf[length:]
	}
	return fields, nil
}
