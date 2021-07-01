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
		jsoniter.MustGetKind(reflect2.TypeOf(SetSpecAutoDeployment{}).Type1()): SetSpecAutoDeploymentCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(SetSpecAutoDeployment{}).Type1()): SetSpecAutoDeploymentCodec{},
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
type SetSpecAutoDeploymentCodec struct {
}

func (SetSpecAutoDeploymentCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*SetSpecAutoDeployment)(ptr) == nil
}

func (SetSpecAutoDeploymentCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*SetSpecAutoDeployment)(ptr)
	var objs []SetSpecAutoDeployment
	if obj != nil {
		objs = []SetSpecAutoDeployment{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SetSpecAutoDeployment{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (SetSpecAutoDeploymentCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*SetSpecAutoDeployment)(ptr) = SetSpecAutoDeployment{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []SetSpecAutoDeployment

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SetSpecAutoDeployment{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*SetSpecAutoDeployment)(ptr) = objs[0]
			} else {
				*(*SetSpecAutoDeployment)(ptr) = SetSpecAutoDeployment{}
			}
		} else {
			*(*SetSpecAutoDeployment)(ptr) = SetSpecAutoDeployment{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj SetSpecAutoDeployment

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(SetSpecAutoDeployment{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*SetSpecAutoDeployment)(ptr) = obj
		} else {
			*(*SetSpecAutoDeployment)(ptr) = SetSpecAutoDeployment{}
		}
	default:
		iter.ReportError("decode SetSpecAutoDeployment", "unexpected JSON type")
	}
}
