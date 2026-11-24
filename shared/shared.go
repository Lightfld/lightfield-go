// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/Lightfld/lightfield-go/internal/apijson"
	"github.com/Lightfld/lightfield-go/packages/param"
	"github.com/Lightfld/lightfield-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// JsonValueUnion contains all possible properties and values from [string],
// [float64], [bool], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfAnyArray]
type JsonValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString   respjson.Field
		OfFloat    respjson.Field
		OfBool     respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u JsonValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JsonValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JsonValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u JsonValueUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u JsonValueUnion) RawJSON() string { return u.JSON.raw }

func (r *JsonValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JsonValueUnion to a JsonValueUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JsonValueUnionParam.Overrides()
func (r JsonValueUnion) ToParam() JsonValueUnionParam {
	return param.Override[JsonValueUnionParam](json.RawMessage(r.RawJSON()))
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type JsonValueUnionParam struct {
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u JsonValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfAnyArray)
}
func (u *JsonValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}
