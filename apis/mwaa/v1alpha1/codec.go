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
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfiguration{}).Type1()):                  EnvironmentSpecLoggingConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationDagProcessingLogs{}).Type1()): EnvironmentSpecLoggingConfigurationDagProcessingLogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationSchedulerLogs{}).Type1()):     EnvironmentSpecLoggingConfigurationSchedulerLogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationTaskLogs{}).Type1()):          EnvironmentSpecLoggingConfigurationTaskLogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationWebserverLogs{}).Type1()):     EnvironmentSpecLoggingConfigurationWebserverLogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationWorkerLogs{}).Type1()):        EnvironmentSpecLoggingConfigurationWorkerLogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecNetworkConfiguration{}).Type1()):                  EnvironmentSpecNetworkConfigurationCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfiguration{}).Type1()):                  EnvironmentSpecLoggingConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationDagProcessingLogs{}).Type1()): EnvironmentSpecLoggingConfigurationDagProcessingLogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationSchedulerLogs{}).Type1()):     EnvironmentSpecLoggingConfigurationSchedulerLogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationTaskLogs{}).Type1()):          EnvironmentSpecLoggingConfigurationTaskLogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationWebserverLogs{}).Type1()):     EnvironmentSpecLoggingConfigurationWebserverLogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationWorkerLogs{}).Type1()):        EnvironmentSpecLoggingConfigurationWorkerLogsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecNetworkConfiguration{}).Type1()):                  EnvironmentSpecNetworkConfigurationCodec{},
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
type EnvironmentSpecLoggingConfigurationCodec struct {
}

func (EnvironmentSpecLoggingConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecLoggingConfiguration)(ptr) == nil
}

func (EnvironmentSpecLoggingConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecLoggingConfiguration)(ptr)
	var objs []EnvironmentSpecLoggingConfiguration
	if obj != nil {
		objs = []EnvironmentSpecLoggingConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecLoggingConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecLoggingConfiguration)(ptr) = EnvironmentSpecLoggingConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecLoggingConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecLoggingConfiguration)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecLoggingConfiguration)(ptr) = EnvironmentSpecLoggingConfiguration{}
			}
		} else {
			*(*EnvironmentSpecLoggingConfiguration)(ptr) = EnvironmentSpecLoggingConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecLoggingConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecLoggingConfiguration)(ptr) = obj
		} else {
			*(*EnvironmentSpecLoggingConfiguration)(ptr) = EnvironmentSpecLoggingConfiguration{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecLoggingConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecLoggingConfigurationDagProcessingLogsCodec struct {
}

func (EnvironmentSpecLoggingConfigurationDagProcessingLogsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecLoggingConfigurationDagProcessingLogs)(ptr) == nil
}

func (EnvironmentSpecLoggingConfigurationDagProcessingLogsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecLoggingConfigurationDagProcessingLogs)(ptr)
	var objs []EnvironmentSpecLoggingConfigurationDagProcessingLogs
	if obj != nil {
		objs = []EnvironmentSpecLoggingConfigurationDagProcessingLogs{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationDagProcessingLogs{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecLoggingConfigurationDagProcessingLogsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecLoggingConfigurationDagProcessingLogs)(ptr) = EnvironmentSpecLoggingConfigurationDagProcessingLogs{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecLoggingConfigurationDagProcessingLogs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationDagProcessingLogs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecLoggingConfigurationDagProcessingLogs)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecLoggingConfigurationDagProcessingLogs)(ptr) = EnvironmentSpecLoggingConfigurationDagProcessingLogs{}
			}
		} else {
			*(*EnvironmentSpecLoggingConfigurationDagProcessingLogs)(ptr) = EnvironmentSpecLoggingConfigurationDagProcessingLogs{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecLoggingConfigurationDagProcessingLogs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationDagProcessingLogs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecLoggingConfigurationDagProcessingLogs)(ptr) = obj
		} else {
			*(*EnvironmentSpecLoggingConfigurationDagProcessingLogs)(ptr) = EnvironmentSpecLoggingConfigurationDagProcessingLogs{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecLoggingConfigurationDagProcessingLogs", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecLoggingConfigurationSchedulerLogsCodec struct {
}

func (EnvironmentSpecLoggingConfigurationSchedulerLogsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecLoggingConfigurationSchedulerLogs)(ptr) == nil
}

func (EnvironmentSpecLoggingConfigurationSchedulerLogsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecLoggingConfigurationSchedulerLogs)(ptr)
	var objs []EnvironmentSpecLoggingConfigurationSchedulerLogs
	if obj != nil {
		objs = []EnvironmentSpecLoggingConfigurationSchedulerLogs{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationSchedulerLogs{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecLoggingConfigurationSchedulerLogsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecLoggingConfigurationSchedulerLogs)(ptr) = EnvironmentSpecLoggingConfigurationSchedulerLogs{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecLoggingConfigurationSchedulerLogs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationSchedulerLogs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecLoggingConfigurationSchedulerLogs)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecLoggingConfigurationSchedulerLogs)(ptr) = EnvironmentSpecLoggingConfigurationSchedulerLogs{}
			}
		} else {
			*(*EnvironmentSpecLoggingConfigurationSchedulerLogs)(ptr) = EnvironmentSpecLoggingConfigurationSchedulerLogs{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecLoggingConfigurationSchedulerLogs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationSchedulerLogs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecLoggingConfigurationSchedulerLogs)(ptr) = obj
		} else {
			*(*EnvironmentSpecLoggingConfigurationSchedulerLogs)(ptr) = EnvironmentSpecLoggingConfigurationSchedulerLogs{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecLoggingConfigurationSchedulerLogs", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecLoggingConfigurationTaskLogsCodec struct {
}

func (EnvironmentSpecLoggingConfigurationTaskLogsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecLoggingConfigurationTaskLogs)(ptr) == nil
}

func (EnvironmentSpecLoggingConfigurationTaskLogsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecLoggingConfigurationTaskLogs)(ptr)
	var objs []EnvironmentSpecLoggingConfigurationTaskLogs
	if obj != nil {
		objs = []EnvironmentSpecLoggingConfigurationTaskLogs{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationTaskLogs{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecLoggingConfigurationTaskLogsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecLoggingConfigurationTaskLogs)(ptr) = EnvironmentSpecLoggingConfigurationTaskLogs{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecLoggingConfigurationTaskLogs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationTaskLogs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecLoggingConfigurationTaskLogs)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecLoggingConfigurationTaskLogs)(ptr) = EnvironmentSpecLoggingConfigurationTaskLogs{}
			}
		} else {
			*(*EnvironmentSpecLoggingConfigurationTaskLogs)(ptr) = EnvironmentSpecLoggingConfigurationTaskLogs{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecLoggingConfigurationTaskLogs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationTaskLogs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecLoggingConfigurationTaskLogs)(ptr) = obj
		} else {
			*(*EnvironmentSpecLoggingConfigurationTaskLogs)(ptr) = EnvironmentSpecLoggingConfigurationTaskLogs{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecLoggingConfigurationTaskLogs", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecLoggingConfigurationWebserverLogsCodec struct {
}

func (EnvironmentSpecLoggingConfigurationWebserverLogsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecLoggingConfigurationWebserverLogs)(ptr) == nil
}

func (EnvironmentSpecLoggingConfigurationWebserverLogsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecLoggingConfigurationWebserverLogs)(ptr)
	var objs []EnvironmentSpecLoggingConfigurationWebserverLogs
	if obj != nil {
		objs = []EnvironmentSpecLoggingConfigurationWebserverLogs{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationWebserverLogs{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecLoggingConfigurationWebserverLogsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecLoggingConfigurationWebserverLogs)(ptr) = EnvironmentSpecLoggingConfigurationWebserverLogs{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecLoggingConfigurationWebserverLogs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationWebserverLogs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecLoggingConfigurationWebserverLogs)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecLoggingConfigurationWebserverLogs)(ptr) = EnvironmentSpecLoggingConfigurationWebserverLogs{}
			}
		} else {
			*(*EnvironmentSpecLoggingConfigurationWebserverLogs)(ptr) = EnvironmentSpecLoggingConfigurationWebserverLogs{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecLoggingConfigurationWebserverLogs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationWebserverLogs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecLoggingConfigurationWebserverLogs)(ptr) = obj
		} else {
			*(*EnvironmentSpecLoggingConfigurationWebserverLogs)(ptr) = EnvironmentSpecLoggingConfigurationWebserverLogs{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecLoggingConfigurationWebserverLogs", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecLoggingConfigurationWorkerLogsCodec struct {
}

func (EnvironmentSpecLoggingConfigurationWorkerLogsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecLoggingConfigurationWorkerLogs)(ptr) == nil
}

func (EnvironmentSpecLoggingConfigurationWorkerLogsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecLoggingConfigurationWorkerLogs)(ptr)
	var objs []EnvironmentSpecLoggingConfigurationWorkerLogs
	if obj != nil {
		objs = []EnvironmentSpecLoggingConfigurationWorkerLogs{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationWorkerLogs{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecLoggingConfigurationWorkerLogsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecLoggingConfigurationWorkerLogs)(ptr) = EnvironmentSpecLoggingConfigurationWorkerLogs{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecLoggingConfigurationWorkerLogs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationWorkerLogs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecLoggingConfigurationWorkerLogs)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecLoggingConfigurationWorkerLogs)(ptr) = EnvironmentSpecLoggingConfigurationWorkerLogs{}
			}
		} else {
			*(*EnvironmentSpecLoggingConfigurationWorkerLogs)(ptr) = EnvironmentSpecLoggingConfigurationWorkerLogs{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecLoggingConfigurationWorkerLogs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecLoggingConfigurationWorkerLogs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecLoggingConfigurationWorkerLogs)(ptr) = obj
		} else {
			*(*EnvironmentSpecLoggingConfigurationWorkerLogs)(ptr) = EnvironmentSpecLoggingConfigurationWorkerLogs{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecLoggingConfigurationWorkerLogs", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type EnvironmentSpecNetworkConfigurationCodec struct {
}

func (EnvironmentSpecNetworkConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*EnvironmentSpecNetworkConfiguration)(ptr) == nil
}

func (EnvironmentSpecNetworkConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*EnvironmentSpecNetworkConfiguration)(ptr)
	var objs []EnvironmentSpecNetworkConfiguration
	if obj != nil {
		objs = []EnvironmentSpecNetworkConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecNetworkConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (EnvironmentSpecNetworkConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*EnvironmentSpecNetworkConfiguration)(ptr) = EnvironmentSpecNetworkConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []EnvironmentSpecNetworkConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecNetworkConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*EnvironmentSpecNetworkConfiguration)(ptr) = objs[0]
			} else {
				*(*EnvironmentSpecNetworkConfiguration)(ptr) = EnvironmentSpecNetworkConfiguration{}
			}
		} else {
			*(*EnvironmentSpecNetworkConfiguration)(ptr) = EnvironmentSpecNetworkConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj EnvironmentSpecNetworkConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(EnvironmentSpecNetworkConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*EnvironmentSpecNetworkConfiguration)(ptr) = obj
		} else {
			*(*EnvironmentSpecNetworkConfiguration)(ptr) = EnvironmentSpecNetworkConfiguration{}
		}
	default:
		iter.ReportError("decode EnvironmentSpecNetworkConfiguration", "unexpected JSON type")
	}
}
