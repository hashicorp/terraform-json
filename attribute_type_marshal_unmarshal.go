package tfjson

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Parts of this package have been taken from, or inspired by
// Terraform's underlying type system, which can be found at:
//   https://godoc.org/github.com/zclconf/go-cty
//
// The attribute type system is largely slimmed down from cty and
// only serves to facilitate the correct marshaling/unmarshaling of
// the external JSON data. It is not designed to be 1-1 compatible
// with cty itself.

// MarshalJSON is an implementation of json.Marshaler that allows
// SchemaAttributeType instances to be serialized as JSON.
func (t SchemaAttributeType) MarshalJSON() ([]byte, error) {
	switch impl := t.attributeTypeImpl.(type) {
	case primitiveAttributeType:
		return []byte(fmt.Sprintf("%q", impl.kind)), nil

	case SchemaAttributeTypeCollection:
		buf := &bytes.Buffer{}
		etyJSON, err := t.ElementType().MarshalJSON()
		if err != nil {
			return nil, err
		}
		buf.WriteRune('[')
		fmt.Fprintf(buf, "%q", impl.kind)
		buf.WriteRune(',')
		buf.Write(etyJSON)
		buf.WriteRune(']')
		return buf.Bytes(), nil

	default:
		// should never happen
		panic("unknown type implementation")
	}
}

// UnmarshalJSON is the opposite of MarshalJSON for
// SchemaAttributeType.
func (t *SchemaAttributeType) UnmarshalJSON(buf []byte) error {
	r := bytes.NewReader(buf)
	dec := json.NewDecoder(r)

	tok, err := dec.Token()
	if err != nil {
		return err
	}

	switch v := tok.(type) {
	case string:
		switch v {
		case string(primitiveAttributeTypeBool):
			*t = SchemaAttributeTypeBool

		case string(primitiveAttributeTypeNumber):
			*t = SchemaAttributeTypeNumber

		case string(primitiveAttributeTypeString):
			*t = SchemaAttributeTypeString

		default:
			return fmt.Errorf("invalid primitive type name %q", v)
		}

		if dec.More() {
			return fmt.Errorf("extraneous data after type description")
		}

		return nil

	case json.Delim:
		if rune(v) != '[' {
			return fmt.Errorf("invalid complex type description")
		}

		tok, err = dec.Token()
		if err != nil {
			return err
		}

		kind, ok := tok.(string)
		if !ok {
			return fmt.Errorf("invalid complex type kind name")
		}

		switch kind {
		case string(collectionAttributeTypeList),
			string(collectionAttributeTypeSet),
			string(collectionAttributeTypeMap):
			var ety SchemaAttributeType
			err = dec.Decode(&ety)
			if err != nil {
				return err
			}

			*t = SchemaAttributeType{
				SchemaAttributeTypeCollection{
					kind:        collectionAttributeTypeKind(kind),
					elementType: ety,
				},
			}

		default:
			return fmt.Errorf("invalid complex type kind name")
		}

		tok, err = dec.Token()
		if err != nil {
			return err
		}

		if delim, ok := tok.(json.Delim); !ok || rune(delim) != ']' || dec.More() {
			return fmt.Errorf("unexpected extra data in type description")
		}

		return nil

	default:
		return fmt.Errorf("invalid type description")
	}
}
