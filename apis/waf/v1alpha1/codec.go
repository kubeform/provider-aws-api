/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ByteMatchSetSpecByteMatchTuplesFieldToMatch{}).Type1()):                 ByteMatchSetSpecByteMatchTuplesFieldToMatchCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(RegexMatchSetSpecRegexMatchTupleFieldToMatch{}).Type1()):                RegexMatchSetSpecRegexMatchTupleFieldToMatchCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(RuleGroupSpecActivatedRuleAction{}).Type1()):                            RuleGroupSpecActivatedRuleActionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SizeConstraintSetSpecSizeConstraintsFieldToMatch{}).Type1()):            SizeConstraintSetSpecSizeConstraintsFieldToMatchCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch{}).Type1()): SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatchCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecDefaultAction{}).Type1()):                                     WebACLSpecDefaultActionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecLoggingConfiguration{}).Type1()):                              WebACLSpecLoggingConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecLoggingConfigurationRedactedFields{}).Type1()):                WebACLSpecLoggingConfigurationRedactedFieldsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecRulesAction{}).Type1()):                                       WebACLSpecRulesActionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecRulesOverrideAction{}).Type1()):                               WebACLSpecRulesOverrideActionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(XssMatchSetSpecXssMatchTuplesFieldToMatch{}).Type1()):                   XssMatchSetSpecXssMatchTuplesFieldToMatchCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ByteMatchSetSpecByteMatchTuplesFieldToMatch{}).Type1()):                 ByteMatchSetSpecByteMatchTuplesFieldToMatchCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(RegexMatchSetSpecRegexMatchTupleFieldToMatch{}).Type1()):                RegexMatchSetSpecRegexMatchTupleFieldToMatchCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(RuleGroupSpecActivatedRuleAction{}).Type1()):                            RuleGroupSpecActivatedRuleActionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SizeConstraintSetSpecSizeConstraintsFieldToMatch{}).Type1()):            SizeConstraintSetSpecSizeConstraintsFieldToMatchCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch{}).Type1()): SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatchCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecDefaultAction{}).Type1()):                                     WebACLSpecDefaultActionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecLoggingConfiguration{}).Type1()):                              WebACLSpecLoggingConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecLoggingConfigurationRedactedFields{}).Type1()):                WebACLSpecLoggingConfigurationRedactedFieldsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecRulesAction{}).Type1()):                                       WebACLSpecRulesActionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecRulesOverrideAction{}).Type1()):                               WebACLSpecRulesOverrideActionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(XssMatchSetSpecXssMatchTuplesFieldToMatch{}).Type1()):                   XssMatchSetSpecXssMatchTuplesFieldToMatchCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type ByteMatchSetSpecByteMatchTuplesFieldToMatchCodec struct {
}

func (ByteMatchSetSpecByteMatchTuplesFieldToMatchCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ByteMatchSetSpecByteMatchTuplesFieldToMatch)(ptr) == nil
}

func (ByteMatchSetSpecByteMatchTuplesFieldToMatchCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ByteMatchSetSpecByteMatchTuplesFieldToMatch)(ptr)
	var objs []ByteMatchSetSpecByteMatchTuplesFieldToMatch
	if obj != nil {
		objs = []ByteMatchSetSpecByteMatchTuplesFieldToMatch{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ByteMatchSetSpecByteMatchTuplesFieldToMatch{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ByteMatchSetSpecByteMatchTuplesFieldToMatchCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ByteMatchSetSpecByteMatchTuplesFieldToMatch)(ptr) = ByteMatchSetSpecByteMatchTuplesFieldToMatch{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ByteMatchSetSpecByteMatchTuplesFieldToMatch

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ByteMatchSetSpecByteMatchTuplesFieldToMatch{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ByteMatchSetSpecByteMatchTuplesFieldToMatch)(ptr) = objs[0]
			} else {
				*(*ByteMatchSetSpecByteMatchTuplesFieldToMatch)(ptr) = ByteMatchSetSpecByteMatchTuplesFieldToMatch{}
			}
		} else {
			*(*ByteMatchSetSpecByteMatchTuplesFieldToMatch)(ptr) = ByteMatchSetSpecByteMatchTuplesFieldToMatch{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ByteMatchSetSpecByteMatchTuplesFieldToMatch

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ByteMatchSetSpecByteMatchTuplesFieldToMatch{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ByteMatchSetSpecByteMatchTuplesFieldToMatch)(ptr) = obj
		} else {
			*(*ByteMatchSetSpecByteMatchTuplesFieldToMatch)(ptr) = ByteMatchSetSpecByteMatchTuplesFieldToMatch{}
		}
	default:
		iter.ReportError("decode ByteMatchSetSpecByteMatchTuplesFieldToMatch", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type RegexMatchSetSpecRegexMatchTupleFieldToMatchCodec struct {
}

func (RegexMatchSetSpecRegexMatchTupleFieldToMatchCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*RegexMatchSetSpecRegexMatchTupleFieldToMatch)(ptr) == nil
}

func (RegexMatchSetSpecRegexMatchTupleFieldToMatchCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*RegexMatchSetSpecRegexMatchTupleFieldToMatch)(ptr)
	var objs []RegexMatchSetSpecRegexMatchTupleFieldToMatch
	if obj != nil {
		objs = []RegexMatchSetSpecRegexMatchTupleFieldToMatch{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(RegexMatchSetSpecRegexMatchTupleFieldToMatch{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (RegexMatchSetSpecRegexMatchTupleFieldToMatchCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*RegexMatchSetSpecRegexMatchTupleFieldToMatch)(ptr) = RegexMatchSetSpecRegexMatchTupleFieldToMatch{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []RegexMatchSetSpecRegexMatchTupleFieldToMatch

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(RegexMatchSetSpecRegexMatchTupleFieldToMatch{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*RegexMatchSetSpecRegexMatchTupleFieldToMatch)(ptr) = objs[0]
			} else {
				*(*RegexMatchSetSpecRegexMatchTupleFieldToMatch)(ptr) = RegexMatchSetSpecRegexMatchTupleFieldToMatch{}
			}
		} else {
			*(*RegexMatchSetSpecRegexMatchTupleFieldToMatch)(ptr) = RegexMatchSetSpecRegexMatchTupleFieldToMatch{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj RegexMatchSetSpecRegexMatchTupleFieldToMatch

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(RegexMatchSetSpecRegexMatchTupleFieldToMatch{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*RegexMatchSetSpecRegexMatchTupleFieldToMatch)(ptr) = obj
		} else {
			*(*RegexMatchSetSpecRegexMatchTupleFieldToMatch)(ptr) = RegexMatchSetSpecRegexMatchTupleFieldToMatch{}
		}
	default:
		iter.ReportError("decode RegexMatchSetSpecRegexMatchTupleFieldToMatch", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type RuleGroupSpecActivatedRuleActionCodec struct {
}

func (RuleGroupSpecActivatedRuleActionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*RuleGroupSpecActivatedRuleAction)(ptr) == nil
}

func (RuleGroupSpecActivatedRuleActionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*RuleGroupSpecActivatedRuleAction)(ptr)
	var objs []RuleGroupSpecActivatedRuleAction
	if obj != nil {
		objs = []RuleGroupSpecActivatedRuleAction{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(RuleGroupSpecActivatedRuleAction{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (RuleGroupSpecActivatedRuleActionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*RuleGroupSpecActivatedRuleAction)(ptr) = RuleGroupSpecActivatedRuleAction{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []RuleGroupSpecActivatedRuleAction

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(RuleGroupSpecActivatedRuleAction{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*RuleGroupSpecActivatedRuleAction)(ptr) = objs[0]
			} else {
				*(*RuleGroupSpecActivatedRuleAction)(ptr) = RuleGroupSpecActivatedRuleAction{}
			}
		} else {
			*(*RuleGroupSpecActivatedRuleAction)(ptr) = RuleGroupSpecActivatedRuleAction{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj RuleGroupSpecActivatedRuleAction

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(RuleGroupSpecActivatedRuleAction{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*RuleGroupSpecActivatedRuleAction)(ptr) = obj
		} else {
			*(*RuleGroupSpecActivatedRuleAction)(ptr) = RuleGroupSpecActivatedRuleAction{}
		}
	default:
		iter.ReportError("decode RuleGroupSpecActivatedRuleAction", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type SizeConstraintSetSpecSizeConstraintsFieldToMatchCodec struct {
}

func (SizeConstraintSetSpecSizeConstraintsFieldToMatchCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SizeConstraintSetSpecSizeConstraintsFieldToMatch)(ptr) == nil
}

func (SizeConstraintSetSpecSizeConstraintsFieldToMatchCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SizeConstraintSetSpecSizeConstraintsFieldToMatch)(ptr)
	var objs []SizeConstraintSetSpecSizeConstraintsFieldToMatch
	if obj != nil {
		objs = []SizeConstraintSetSpecSizeConstraintsFieldToMatch{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SizeConstraintSetSpecSizeConstraintsFieldToMatch{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SizeConstraintSetSpecSizeConstraintsFieldToMatchCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SizeConstraintSetSpecSizeConstraintsFieldToMatch)(ptr) = SizeConstraintSetSpecSizeConstraintsFieldToMatch{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SizeConstraintSetSpecSizeConstraintsFieldToMatch

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SizeConstraintSetSpecSizeConstraintsFieldToMatch{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SizeConstraintSetSpecSizeConstraintsFieldToMatch)(ptr) = objs[0]
			} else {
				*(*SizeConstraintSetSpecSizeConstraintsFieldToMatch)(ptr) = SizeConstraintSetSpecSizeConstraintsFieldToMatch{}
			}
		} else {
			*(*SizeConstraintSetSpecSizeConstraintsFieldToMatch)(ptr) = SizeConstraintSetSpecSizeConstraintsFieldToMatch{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj SizeConstraintSetSpecSizeConstraintsFieldToMatch

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SizeConstraintSetSpecSizeConstraintsFieldToMatch{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*SizeConstraintSetSpecSizeConstraintsFieldToMatch)(ptr) = obj
		} else {
			*(*SizeConstraintSetSpecSizeConstraintsFieldToMatch)(ptr) = SizeConstraintSetSpecSizeConstraintsFieldToMatch{}
		}
	default:
		iter.ReportError("decode SizeConstraintSetSpecSizeConstraintsFieldToMatch", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatchCodec struct {
}

func (SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatchCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch)(ptr) == nil
}

func (SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatchCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch)(ptr)
	var objs []SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch
	if obj != nil {
		objs = []SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatchCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch)(ptr) = SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch)(ptr) = objs[0]
			} else {
				*(*SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch)(ptr) = SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch{}
			}
		} else {
			*(*SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch)(ptr) = SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch)(ptr) = obj
		} else {
			*(*SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch)(ptr) = SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch{}
		}
	default:
		iter.ReportError("decode SqlInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type WebACLSpecDefaultActionCodec struct {
}

func (WebACLSpecDefaultActionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*WebACLSpecDefaultAction)(ptr) == nil
}

func (WebACLSpecDefaultActionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*WebACLSpecDefaultAction)(ptr)
	var objs []WebACLSpecDefaultAction
	if obj != nil {
		objs = []WebACLSpecDefaultAction{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecDefaultAction{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (WebACLSpecDefaultActionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*WebACLSpecDefaultAction)(ptr) = WebACLSpecDefaultAction{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []WebACLSpecDefaultAction

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecDefaultAction{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*WebACLSpecDefaultAction)(ptr) = objs[0]
			} else {
				*(*WebACLSpecDefaultAction)(ptr) = WebACLSpecDefaultAction{}
			}
		} else {
			*(*WebACLSpecDefaultAction)(ptr) = WebACLSpecDefaultAction{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj WebACLSpecDefaultAction

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecDefaultAction{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*WebACLSpecDefaultAction)(ptr) = obj
		} else {
			*(*WebACLSpecDefaultAction)(ptr) = WebACLSpecDefaultAction{}
		}
	default:
		iter.ReportError("decode WebACLSpecDefaultAction", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type WebACLSpecLoggingConfigurationCodec struct {
}

func (WebACLSpecLoggingConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*WebACLSpecLoggingConfiguration)(ptr) == nil
}

func (WebACLSpecLoggingConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*WebACLSpecLoggingConfiguration)(ptr)
	var objs []WebACLSpecLoggingConfiguration
	if obj != nil {
		objs = []WebACLSpecLoggingConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecLoggingConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (WebACLSpecLoggingConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*WebACLSpecLoggingConfiguration)(ptr) = WebACLSpecLoggingConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []WebACLSpecLoggingConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecLoggingConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*WebACLSpecLoggingConfiguration)(ptr) = objs[0]
			} else {
				*(*WebACLSpecLoggingConfiguration)(ptr) = WebACLSpecLoggingConfiguration{}
			}
		} else {
			*(*WebACLSpecLoggingConfiguration)(ptr) = WebACLSpecLoggingConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj WebACLSpecLoggingConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecLoggingConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*WebACLSpecLoggingConfiguration)(ptr) = obj
		} else {
			*(*WebACLSpecLoggingConfiguration)(ptr) = WebACLSpecLoggingConfiguration{}
		}
	default:
		iter.ReportError("decode WebACLSpecLoggingConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type WebACLSpecLoggingConfigurationRedactedFieldsCodec struct {
}

func (WebACLSpecLoggingConfigurationRedactedFieldsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*WebACLSpecLoggingConfigurationRedactedFields)(ptr) == nil
}

func (WebACLSpecLoggingConfigurationRedactedFieldsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*WebACLSpecLoggingConfigurationRedactedFields)(ptr)
	var objs []WebACLSpecLoggingConfigurationRedactedFields
	if obj != nil {
		objs = []WebACLSpecLoggingConfigurationRedactedFields{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecLoggingConfigurationRedactedFields{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (WebACLSpecLoggingConfigurationRedactedFieldsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*WebACLSpecLoggingConfigurationRedactedFields)(ptr) = WebACLSpecLoggingConfigurationRedactedFields{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []WebACLSpecLoggingConfigurationRedactedFields

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecLoggingConfigurationRedactedFields{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*WebACLSpecLoggingConfigurationRedactedFields)(ptr) = objs[0]
			} else {
				*(*WebACLSpecLoggingConfigurationRedactedFields)(ptr) = WebACLSpecLoggingConfigurationRedactedFields{}
			}
		} else {
			*(*WebACLSpecLoggingConfigurationRedactedFields)(ptr) = WebACLSpecLoggingConfigurationRedactedFields{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj WebACLSpecLoggingConfigurationRedactedFields

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecLoggingConfigurationRedactedFields{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*WebACLSpecLoggingConfigurationRedactedFields)(ptr) = obj
		} else {
			*(*WebACLSpecLoggingConfigurationRedactedFields)(ptr) = WebACLSpecLoggingConfigurationRedactedFields{}
		}
	default:
		iter.ReportError("decode WebACLSpecLoggingConfigurationRedactedFields", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type WebACLSpecRulesActionCodec struct {
}

func (WebACLSpecRulesActionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*WebACLSpecRulesAction)(ptr) == nil
}

func (WebACLSpecRulesActionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*WebACLSpecRulesAction)(ptr)
	var objs []WebACLSpecRulesAction
	if obj != nil {
		objs = []WebACLSpecRulesAction{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecRulesAction{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (WebACLSpecRulesActionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*WebACLSpecRulesAction)(ptr) = WebACLSpecRulesAction{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []WebACLSpecRulesAction

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecRulesAction{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*WebACLSpecRulesAction)(ptr) = objs[0]
			} else {
				*(*WebACLSpecRulesAction)(ptr) = WebACLSpecRulesAction{}
			}
		} else {
			*(*WebACLSpecRulesAction)(ptr) = WebACLSpecRulesAction{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj WebACLSpecRulesAction

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecRulesAction{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*WebACLSpecRulesAction)(ptr) = obj
		} else {
			*(*WebACLSpecRulesAction)(ptr) = WebACLSpecRulesAction{}
		}
	default:
		iter.ReportError("decode WebACLSpecRulesAction", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type WebACLSpecRulesOverrideActionCodec struct {
}

func (WebACLSpecRulesOverrideActionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*WebACLSpecRulesOverrideAction)(ptr) == nil
}

func (WebACLSpecRulesOverrideActionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*WebACLSpecRulesOverrideAction)(ptr)
	var objs []WebACLSpecRulesOverrideAction
	if obj != nil {
		objs = []WebACLSpecRulesOverrideAction{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecRulesOverrideAction{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (WebACLSpecRulesOverrideActionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*WebACLSpecRulesOverrideAction)(ptr) = WebACLSpecRulesOverrideAction{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []WebACLSpecRulesOverrideAction

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecRulesOverrideAction{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*WebACLSpecRulesOverrideAction)(ptr) = objs[0]
			} else {
				*(*WebACLSpecRulesOverrideAction)(ptr) = WebACLSpecRulesOverrideAction{}
			}
		} else {
			*(*WebACLSpecRulesOverrideAction)(ptr) = WebACLSpecRulesOverrideAction{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj WebACLSpecRulesOverrideAction

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(WebACLSpecRulesOverrideAction{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*WebACLSpecRulesOverrideAction)(ptr) = obj
		} else {
			*(*WebACLSpecRulesOverrideAction)(ptr) = WebACLSpecRulesOverrideAction{}
		}
	default:
		iter.ReportError("decode WebACLSpecRulesOverrideAction", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type XssMatchSetSpecXssMatchTuplesFieldToMatchCodec struct {
}

func (XssMatchSetSpecXssMatchTuplesFieldToMatchCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*XssMatchSetSpecXssMatchTuplesFieldToMatch)(ptr) == nil
}

func (XssMatchSetSpecXssMatchTuplesFieldToMatchCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*XssMatchSetSpecXssMatchTuplesFieldToMatch)(ptr)
	var objs []XssMatchSetSpecXssMatchTuplesFieldToMatch
	if obj != nil {
		objs = []XssMatchSetSpecXssMatchTuplesFieldToMatch{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(XssMatchSetSpecXssMatchTuplesFieldToMatch{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (XssMatchSetSpecXssMatchTuplesFieldToMatchCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*XssMatchSetSpecXssMatchTuplesFieldToMatch)(ptr) = XssMatchSetSpecXssMatchTuplesFieldToMatch{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []XssMatchSetSpecXssMatchTuplesFieldToMatch

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(XssMatchSetSpecXssMatchTuplesFieldToMatch{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*XssMatchSetSpecXssMatchTuplesFieldToMatch)(ptr) = objs[0]
			} else {
				*(*XssMatchSetSpecXssMatchTuplesFieldToMatch)(ptr) = XssMatchSetSpecXssMatchTuplesFieldToMatch{}
			}
		} else {
			*(*XssMatchSetSpecXssMatchTuplesFieldToMatch)(ptr) = XssMatchSetSpecXssMatchTuplesFieldToMatch{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj XssMatchSetSpecXssMatchTuplesFieldToMatch

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(XssMatchSetSpecXssMatchTuplesFieldToMatch{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*XssMatchSetSpecXssMatchTuplesFieldToMatch)(ptr) = obj
		} else {
			*(*XssMatchSetSpecXssMatchTuplesFieldToMatch)(ptr) = XssMatchSetSpecXssMatchTuplesFieldToMatch{}
		}
	default:
		iter.ReportError("decode XssMatchSetSpecXssMatchTuplesFieldToMatch", "unexpected JSON type")
	}
}
