package tfjson

import "errors"

// Parts of this package have been taken from, or inspired by
// Terraform's underlying type system, which can be found at:
//   https://godoc.org/github.com/zclconf/go-cty
//
// The attribute type system here is largely slimmed down from cty
// and only serves to facilitate the correct marshaling/unmarshaling
// of the external JSON data. It is not designed to be 1-1 compatible
// with cty itself.

// SchemaAttributeType represents any schema attribute type.
type SchemaAttributeType struct {
	attributeTypeImpl
}

// attributeTypeImpl is an internal interface that all types
// implement.
type attributeTypeImpl interface {
	// isAttrbuteTypeImpl is a no-op method that exists only to express
	// that a type is an implementation of attributeTypeImpl.
	isAttrbuteTypeImpl() attributeTypeImplSigil

	// Equals returns true if the other given Type exactly equals the
	// receiver Type.
	Equals(other SchemaAttributeType) bool

	// Name returns a name for the given type.
	Name() string
}

// attributeTypeImplSigil is a no-op object that only serves to
// express that a type is an implementation of attributeTypeImpl.
type attributeTypeImplSigil struct{}

func (t attributeTypeImplSigil) isAttrbuteTypeImpl() attributeTypeImplSigil { return t }

// Equals returns true if the other given Type exactly equals the
// receiver type.
func (t SchemaAttributeType) Equals(other SchemaAttributeType) bool {
	return t.attributeTypeImpl.Equals(other)
}

// Name returns a name for the given type.
func (t SchemaAttributeType) Name() string {
	return t.attributeTypeImpl.Name()
}

// primitiveAttributeType is the hidden implementation of the various
// primitive types. When JSON attributes are parsed, the primitive
// attribute types are assigned the singleton types given in this
// package.
type primitiveAttributeType struct {
	attributeTypeImplSigil

	kind primitiveAttributeTypeKind
}

type primitiveAttributeTypeKind string

const (
	primitiveAttributeTypeBool   primitiveAttributeTypeKind = "bool"
	primitiveAttributeTypeNumber primitiveAttributeTypeKind = "number"
	primitiveAttributeTypeString primitiveAttributeTypeKind = "string"
)

func (t primitiveAttributeType) Equals(other SchemaAttributeType) bool {
	if otherP, ok := other.attributeTypeImpl.(primitiveAttributeType); ok {
		return otherP.kind == t.kind
	}

	return false
}

func (t primitiveAttributeType) Name() string {
	return string(t.kind)
}

// SchemaAttributeTypeNumber is the numeric type. Number values are
// arbitrary-precision decimal numbers, which can then be converted
// into Go's various numeric types only if they are in the
// appropriate range.
var SchemaAttributeTypeNumber = SchemaAttributeType{primitiveAttributeType{kind: primitiveAttributeTypeNumber}}

// SchemaAttributeTypeString is the string type. String values are
// sequences of unicode codepoints encoded internally as UTF-8.
var SchemaAttributeTypeString = SchemaAttributeType{primitiveAttributeType{kind: primitiveAttributeTypeString}}

// SchemaAttributeTypeBool is the boolean type. The two values of
// this type are True and False.
var SchemaAttributeTypeBool = SchemaAttributeType{primitiveAttributeType{kind: primitiveAttributeTypeBool}}

// IsPrimitiveType returns true if and only if the reciever is a primitive
// type, which means it's either number, string, or bool. Any two primitive
// types can be safely compared for equality using the standard == operator
// without panic, which is not a guarantee that holds for all types. Primitive
// types can therefore also be used in switch statements.
func (t SchemaAttributeType) IsPrimitiveType() bool {
	_, ok := t.attributeTypeImpl.(primitiveAttributeType)
	return ok
}

// SchemaAttributeTypeCollection encompasses all collection types (list,
// sets, and maps.
type SchemaAttributeTypeCollection struct {
	attributeTypeImplSigil

	kind        collectionAttributeTypeKind
	elementType SchemaAttributeType
}

type collectionAttributeTypeKind string

const (
	collectionAttributeTypeList collectionAttributeTypeKind = "list"
	collectionAttributeTypeSet  collectionAttributeTypeKind = "set"
	collectionAttributeTypeMap  collectionAttributeTypeKind = "map"
)

// Name returns the name of the type of the collection. This does not
// print the element type.
func (t SchemaAttributeTypeCollection) Name() string {
	return string(t.kind)
}

// Equals checks for equality of this collection to another
// SchemaAttributeType.
//
// Two collection types are comparable, and equal if their kind and
// element types match.
func (t SchemaAttributeTypeCollection) Equals(other SchemaAttributeType) bool {
	if otherC, ok := other.attributeTypeImpl.(SchemaAttributeTypeCollection); ok {
		return otherC == t
	}

	return false
}

// SchemaAttributeTypeList creates a list collection type with the
// element type set to elem. This collection is comparable with other
// lists with the exact same element type.
func SchemaAttributeTypeList(elem SchemaAttributeType) SchemaAttributeType {
	return SchemaAttributeType{
		SchemaAttributeTypeCollection{
			kind:        collectionAttributeTypeList,
			elementType: elem,
		},
	}
}

// SchemaAttributeTypeSet creates a set collection type with the
// element type set to elem. This collection is comparable with other
// sets with the exact same element type.
func SchemaAttributeTypeSet(elem SchemaAttributeType) SchemaAttributeType {
	return SchemaAttributeType{
		SchemaAttributeTypeCollection{
			kind:        collectionAttributeTypeSet,
			elementType: elem,
		},
	}
}

// SchemaAttributeTypeMap creates a map collection type with the
// element type set to elem. This collection is comparable with other
// maps with the exact same element type.
func SchemaAttributeTypeMap(elem SchemaAttributeType) SchemaAttributeType {
	return SchemaAttributeType{
		SchemaAttributeTypeCollection{
			kind:        collectionAttributeTypeMap,
			elementType: elem,
		},
	}
}

// IsCollectionType returns true if the given type is a collection
// (list, set, or map).
func (t SchemaAttributeType) IsCollectionType() bool {
	_, ok := t.attributeTypeImpl.(SchemaAttributeTypeCollection)
	return ok
}

// IsListType returns true if the type is a list.
func (t SchemaAttributeType) IsListType() bool {
	if ct, ok := t.attributeTypeImpl.(SchemaAttributeTypeCollection); ok {
		return ct.kind == collectionAttributeTypeList
	}

	return false
}

// IsSetType returns true if the type is a set.
func (t SchemaAttributeType) IsSetType() bool {
	if ct, ok := t.attributeTypeImpl.(SchemaAttributeTypeCollection); ok {
		return ct.kind == collectionAttributeTypeSet
	}

	return false
}

// IsMapType returns true if the type is a map.
func (t SchemaAttributeType) IsMapType() bool {
	if ct, ok := t.attributeTypeImpl.(SchemaAttributeTypeCollection); ok {
		return ct.kind == collectionAttributeTypeMap
	}

	return false
}

// ElementType returns the element type of the receiver if it is a
// collection type, or panics if it is not. Use IsCollectionType
// first to test whether this method will succeed.
func (t SchemaAttributeType) ElementType() SchemaAttributeType {
	if ct, ok := t.attributeTypeImpl.(SchemaAttributeTypeCollection); ok {
		return ct.elementType
	}

	panic(errors.New("not a collection type"))
}
