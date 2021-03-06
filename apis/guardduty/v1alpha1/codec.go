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
		jsoniter.MustGetKind(reflect2.TypeOf(DetectorSpecDatasources{}).Type1()):                        DetectorSpecDatasourcesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(DetectorSpecDatasourcesS3Logs{}).Type1()):                  DetectorSpecDatasourcesS3LogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FilterSpecFindingCriteria{}).Type1()):                      FilterSpecFindingCriteriaCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(OrganizationConfigurationSpecDatasources{}).Type1()):       OrganizationConfigurationSpecDatasourcesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(OrganizationConfigurationSpecDatasourcesS3Logs{}).Type1()): OrganizationConfigurationSpecDatasourcesS3LogsCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(DetectorSpecDatasources{}).Type1()):                        DetectorSpecDatasourcesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(DetectorSpecDatasourcesS3Logs{}).Type1()):                  DetectorSpecDatasourcesS3LogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(FilterSpecFindingCriteria{}).Type1()):                      FilterSpecFindingCriteriaCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(OrganizationConfigurationSpecDatasources{}).Type1()):       OrganizationConfigurationSpecDatasourcesCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(OrganizationConfigurationSpecDatasourcesS3Logs{}).Type1()): OrganizationConfigurationSpecDatasourcesS3LogsCodec{},
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
type DetectorSpecDatasourcesCodec struct {
}

func (DetectorSpecDatasourcesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*DetectorSpecDatasources)(ptr) == nil
}

func (DetectorSpecDatasourcesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DetectorSpecDatasources)(ptr)
	var objs []DetectorSpecDatasources
	if obj != nil {
		objs = []DetectorSpecDatasources{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DetectorSpecDatasources{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (DetectorSpecDatasourcesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*DetectorSpecDatasources)(ptr) = DetectorSpecDatasources{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []DetectorSpecDatasources

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DetectorSpecDatasources{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*DetectorSpecDatasources)(ptr) = objs[0]
			} else {
				*(*DetectorSpecDatasources)(ptr) = DetectorSpecDatasources{}
			}
		} else {
			*(*DetectorSpecDatasources)(ptr) = DetectorSpecDatasources{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj DetectorSpecDatasources

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DetectorSpecDatasources{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*DetectorSpecDatasources)(ptr) = obj
		} else {
			*(*DetectorSpecDatasources)(ptr) = DetectorSpecDatasources{}
		}
	default:
		iter.ReportError("decode DetectorSpecDatasources", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type DetectorSpecDatasourcesS3LogsCodec struct {
}

func (DetectorSpecDatasourcesS3LogsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*DetectorSpecDatasourcesS3Logs)(ptr) == nil
}

func (DetectorSpecDatasourcesS3LogsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DetectorSpecDatasourcesS3Logs)(ptr)
	var objs []DetectorSpecDatasourcesS3Logs
	if obj != nil {
		objs = []DetectorSpecDatasourcesS3Logs{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DetectorSpecDatasourcesS3Logs{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (DetectorSpecDatasourcesS3LogsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*DetectorSpecDatasourcesS3Logs)(ptr) = DetectorSpecDatasourcesS3Logs{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []DetectorSpecDatasourcesS3Logs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DetectorSpecDatasourcesS3Logs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*DetectorSpecDatasourcesS3Logs)(ptr) = objs[0]
			} else {
				*(*DetectorSpecDatasourcesS3Logs)(ptr) = DetectorSpecDatasourcesS3Logs{}
			}
		} else {
			*(*DetectorSpecDatasourcesS3Logs)(ptr) = DetectorSpecDatasourcesS3Logs{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj DetectorSpecDatasourcesS3Logs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(DetectorSpecDatasourcesS3Logs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*DetectorSpecDatasourcesS3Logs)(ptr) = obj
		} else {
			*(*DetectorSpecDatasourcesS3Logs)(ptr) = DetectorSpecDatasourcesS3Logs{}
		}
	default:
		iter.ReportError("decode DetectorSpecDatasourcesS3Logs", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type FilterSpecFindingCriteriaCodec struct {
}

func (FilterSpecFindingCriteriaCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*FilterSpecFindingCriteria)(ptr) == nil
}

func (FilterSpecFindingCriteriaCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*FilterSpecFindingCriteria)(ptr)
	var objs []FilterSpecFindingCriteria
	if obj != nil {
		objs = []FilterSpecFindingCriteria{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FilterSpecFindingCriteria{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (FilterSpecFindingCriteriaCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*FilterSpecFindingCriteria)(ptr) = FilterSpecFindingCriteria{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []FilterSpecFindingCriteria

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FilterSpecFindingCriteria{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*FilterSpecFindingCriteria)(ptr) = objs[0]
			} else {
				*(*FilterSpecFindingCriteria)(ptr) = FilterSpecFindingCriteria{}
			}
		} else {
			*(*FilterSpecFindingCriteria)(ptr) = FilterSpecFindingCriteria{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj FilterSpecFindingCriteria

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(FilterSpecFindingCriteria{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*FilterSpecFindingCriteria)(ptr) = obj
		} else {
			*(*FilterSpecFindingCriteria)(ptr) = FilterSpecFindingCriteria{}
		}
	default:
		iter.ReportError("decode FilterSpecFindingCriteria", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type OrganizationConfigurationSpecDatasourcesCodec struct {
}

func (OrganizationConfigurationSpecDatasourcesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*OrganizationConfigurationSpecDatasources)(ptr) == nil
}

func (OrganizationConfigurationSpecDatasourcesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*OrganizationConfigurationSpecDatasources)(ptr)
	var objs []OrganizationConfigurationSpecDatasources
	if obj != nil {
		objs = []OrganizationConfigurationSpecDatasources{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(OrganizationConfigurationSpecDatasources{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (OrganizationConfigurationSpecDatasourcesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*OrganizationConfigurationSpecDatasources)(ptr) = OrganizationConfigurationSpecDatasources{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []OrganizationConfigurationSpecDatasources

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(OrganizationConfigurationSpecDatasources{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*OrganizationConfigurationSpecDatasources)(ptr) = objs[0]
			} else {
				*(*OrganizationConfigurationSpecDatasources)(ptr) = OrganizationConfigurationSpecDatasources{}
			}
		} else {
			*(*OrganizationConfigurationSpecDatasources)(ptr) = OrganizationConfigurationSpecDatasources{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj OrganizationConfigurationSpecDatasources

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(OrganizationConfigurationSpecDatasources{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*OrganizationConfigurationSpecDatasources)(ptr) = obj
		} else {
			*(*OrganizationConfigurationSpecDatasources)(ptr) = OrganizationConfigurationSpecDatasources{}
		}
	default:
		iter.ReportError("decode OrganizationConfigurationSpecDatasources", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type OrganizationConfigurationSpecDatasourcesS3LogsCodec struct {
}

func (OrganizationConfigurationSpecDatasourcesS3LogsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*OrganizationConfigurationSpecDatasourcesS3Logs)(ptr) == nil
}

func (OrganizationConfigurationSpecDatasourcesS3LogsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*OrganizationConfigurationSpecDatasourcesS3Logs)(ptr)
	var objs []OrganizationConfigurationSpecDatasourcesS3Logs
	if obj != nil {
		objs = []OrganizationConfigurationSpecDatasourcesS3Logs{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(OrganizationConfigurationSpecDatasourcesS3Logs{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (OrganizationConfigurationSpecDatasourcesS3LogsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*OrganizationConfigurationSpecDatasourcesS3Logs)(ptr) = OrganizationConfigurationSpecDatasourcesS3Logs{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []OrganizationConfigurationSpecDatasourcesS3Logs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(OrganizationConfigurationSpecDatasourcesS3Logs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*OrganizationConfigurationSpecDatasourcesS3Logs)(ptr) = objs[0]
			} else {
				*(*OrganizationConfigurationSpecDatasourcesS3Logs)(ptr) = OrganizationConfigurationSpecDatasourcesS3Logs{}
			}
		} else {
			*(*OrganizationConfigurationSpecDatasourcesS3Logs)(ptr) = OrganizationConfigurationSpecDatasourcesS3Logs{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj OrganizationConfigurationSpecDatasourcesS3Logs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(OrganizationConfigurationSpecDatasourcesS3Logs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*OrganizationConfigurationSpecDatasourcesS3Logs)(ptr) = obj
		} else {
			*(*OrganizationConfigurationSpecDatasourcesS3Logs)(ptr) = OrganizationConfigurationSpecDatasourcesS3Logs{}
		}
	default:
		iter.ReportError("decode OrganizationConfigurationSpecDatasourcesS3Logs", "unexpected JSON type")
	}
}
