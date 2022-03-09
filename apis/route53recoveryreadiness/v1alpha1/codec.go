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
		jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResource{}).Type1()):                          ResourceSetSpecResourcesDnsTargetResourceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResource{}).Type1()):            ResourceSetSpecResourcesDnsTargetResourceTargetResourceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource{}).Type1()): ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResourceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource{}).Type1()): ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53ResourceCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResource{}).Type1()):                          ResourceSetSpecResourcesDnsTargetResourceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResource{}).Type1()):            ResourceSetSpecResourcesDnsTargetResourceTargetResourceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource{}).Type1()): ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResourceCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource{}).Type1()): ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53ResourceCodec{},
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
type ResourceSetSpecResourcesDnsTargetResourceCodec struct {
}

func (ResourceSetSpecResourcesDnsTargetResourceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ResourceSetSpecResourcesDnsTargetResource)(ptr) == nil
}

func (ResourceSetSpecResourcesDnsTargetResourceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ResourceSetSpecResourcesDnsTargetResource)(ptr)
	var objs []ResourceSetSpecResourcesDnsTargetResource
	if obj != nil {
		objs = []ResourceSetSpecResourcesDnsTargetResource{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResource{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ResourceSetSpecResourcesDnsTargetResourceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ResourceSetSpecResourcesDnsTargetResource)(ptr) = ResourceSetSpecResourcesDnsTargetResource{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ResourceSetSpecResourcesDnsTargetResource

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResource{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ResourceSetSpecResourcesDnsTargetResource)(ptr) = objs[0]
			} else {
				*(*ResourceSetSpecResourcesDnsTargetResource)(ptr) = ResourceSetSpecResourcesDnsTargetResource{}
			}
		} else {
			*(*ResourceSetSpecResourcesDnsTargetResource)(ptr) = ResourceSetSpecResourcesDnsTargetResource{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ResourceSetSpecResourcesDnsTargetResource

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResource{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ResourceSetSpecResourcesDnsTargetResource)(ptr) = obj
		} else {
			*(*ResourceSetSpecResourcesDnsTargetResource)(ptr) = ResourceSetSpecResourcesDnsTargetResource{}
		}
	default:
		iter.ReportError("decode ResourceSetSpecResourcesDnsTargetResource", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ResourceSetSpecResourcesDnsTargetResourceTargetResourceCodec struct {
}

func (ResourceSetSpecResourcesDnsTargetResourceTargetResourceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ResourceSetSpecResourcesDnsTargetResourceTargetResource)(ptr) == nil
}

func (ResourceSetSpecResourcesDnsTargetResourceTargetResourceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ResourceSetSpecResourcesDnsTargetResourceTargetResource)(ptr)
	var objs []ResourceSetSpecResourcesDnsTargetResourceTargetResource
	if obj != nil {
		objs = []ResourceSetSpecResourcesDnsTargetResourceTargetResource{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResource{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ResourceSetSpecResourcesDnsTargetResourceTargetResourceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ResourceSetSpecResourcesDnsTargetResourceTargetResource)(ptr) = ResourceSetSpecResourcesDnsTargetResourceTargetResource{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ResourceSetSpecResourcesDnsTargetResourceTargetResource

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResource{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ResourceSetSpecResourcesDnsTargetResourceTargetResource)(ptr) = objs[0]
			} else {
				*(*ResourceSetSpecResourcesDnsTargetResourceTargetResource)(ptr) = ResourceSetSpecResourcesDnsTargetResourceTargetResource{}
			}
		} else {
			*(*ResourceSetSpecResourcesDnsTargetResourceTargetResource)(ptr) = ResourceSetSpecResourcesDnsTargetResourceTargetResource{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ResourceSetSpecResourcesDnsTargetResourceTargetResource

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResource{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ResourceSetSpecResourcesDnsTargetResourceTargetResource)(ptr) = obj
		} else {
			*(*ResourceSetSpecResourcesDnsTargetResourceTargetResource)(ptr) = ResourceSetSpecResourcesDnsTargetResourceTargetResource{}
		}
	default:
		iter.ReportError("decode ResourceSetSpecResourcesDnsTargetResourceTargetResource", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResourceCodec struct {
}

func (ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResourceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource)(ptr) == nil
}

func (ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResourceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource)(ptr)
	var objs []ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource
	if obj != nil {
		objs = []ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResourceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource)(ptr) = ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource)(ptr) = objs[0]
			} else {
				*(*ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource)(ptr) = ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource{}
			}
		} else {
			*(*ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource)(ptr) = ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource)(ptr) = obj
		} else {
			*(*ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource)(ptr) = ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource{}
		}
	default:
		iter.ReportError("decode ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53ResourceCodec struct {
}

func (ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53ResourceCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource)(ptr) == nil
}

func (ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53ResourceCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource)(ptr)
	var objs []ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource
	if obj != nil {
		objs = []ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53ResourceCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource)(ptr) = ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource)(ptr) = objs[0]
			} else {
				*(*ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource)(ptr) = ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource{}
			}
		} else {
			*(*ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource)(ptr) = ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource)(ptr) = obj
		} else {
			*(*ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource)(ptr) = ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource{}
		}
	default:
		iter.ReportError("decode ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource", "unexpected JSON type")
	}
}
