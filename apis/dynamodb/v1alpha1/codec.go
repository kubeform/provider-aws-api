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
		jsoniter.MustGetKind(reflect2.TypeOf(TableSpecPointInTimeRecovery{}).Type1()):  TableSpecPointInTimeRecoveryCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TableSpecServerSideEncryption{}).Type1()): TableSpecServerSideEncryptionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TableSpecTtl{}).Type1()):                  TableSpecTtlCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(TableSpecPointInTimeRecovery{}).Type1()):  TableSpecPointInTimeRecoveryCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TableSpecServerSideEncryption{}).Type1()): TableSpecServerSideEncryptionCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(TableSpecTtl{}).Type1()):                  TableSpecTtlCodec{},
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
type TableSpecPointInTimeRecoveryCodec struct {
}

func (TableSpecPointInTimeRecoveryCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TableSpecPointInTimeRecovery)(ptr) == nil
}

func (TableSpecPointInTimeRecoveryCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TableSpecPointInTimeRecovery)(ptr)
	var objs []TableSpecPointInTimeRecovery
	if obj != nil {
		objs = []TableSpecPointInTimeRecovery{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TableSpecPointInTimeRecovery{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TableSpecPointInTimeRecoveryCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TableSpecPointInTimeRecovery)(ptr) = TableSpecPointInTimeRecovery{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TableSpecPointInTimeRecovery

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TableSpecPointInTimeRecovery{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TableSpecPointInTimeRecovery)(ptr) = objs[0]
			} else {
				*(*TableSpecPointInTimeRecovery)(ptr) = TableSpecPointInTimeRecovery{}
			}
		} else {
			*(*TableSpecPointInTimeRecovery)(ptr) = TableSpecPointInTimeRecovery{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj TableSpecPointInTimeRecovery

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TableSpecPointInTimeRecovery{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*TableSpecPointInTimeRecovery)(ptr) = obj
		} else {
			*(*TableSpecPointInTimeRecovery)(ptr) = TableSpecPointInTimeRecovery{}
		}
	default:
		iter.ReportError("decode TableSpecPointInTimeRecovery", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type TableSpecServerSideEncryptionCodec struct {
}

func (TableSpecServerSideEncryptionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TableSpecServerSideEncryption)(ptr) == nil
}

func (TableSpecServerSideEncryptionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TableSpecServerSideEncryption)(ptr)
	var objs []TableSpecServerSideEncryption
	if obj != nil {
		objs = []TableSpecServerSideEncryption{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TableSpecServerSideEncryption{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TableSpecServerSideEncryptionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TableSpecServerSideEncryption)(ptr) = TableSpecServerSideEncryption{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TableSpecServerSideEncryption

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TableSpecServerSideEncryption{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TableSpecServerSideEncryption)(ptr) = objs[0]
			} else {
				*(*TableSpecServerSideEncryption)(ptr) = TableSpecServerSideEncryption{}
			}
		} else {
			*(*TableSpecServerSideEncryption)(ptr) = TableSpecServerSideEncryption{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj TableSpecServerSideEncryption

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TableSpecServerSideEncryption{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*TableSpecServerSideEncryption)(ptr) = obj
		} else {
			*(*TableSpecServerSideEncryption)(ptr) = TableSpecServerSideEncryption{}
		}
	default:
		iter.ReportError("decode TableSpecServerSideEncryption", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type TableSpecTtlCodec struct {
}

func (TableSpecTtlCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*TableSpecTtl)(ptr) == nil
}

func (TableSpecTtlCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*TableSpecTtl)(ptr)
	var objs []TableSpecTtl
	if obj != nil {
		objs = []TableSpecTtl{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TableSpecTtl{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (TableSpecTtlCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*TableSpecTtl)(ptr) = TableSpecTtl{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []TableSpecTtl

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TableSpecTtl{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*TableSpecTtl)(ptr) = objs[0]
			} else {
				*(*TableSpecTtl)(ptr) = TableSpecTtl{}
			}
		} else {
			*(*TableSpecTtl)(ptr) = TableSpecTtl{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj TableSpecTtl

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(TableSpecTtl{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*TableSpecTtl)(ptr) = obj
		} else {
			*(*TableSpecTtl)(ptr) = TableSpecTtl{}
		}
	default:
		iter.ReportError("decode TableSpecTtl", "unexpected JSON type")
	}
}
