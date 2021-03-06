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
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecValidity{}).Type1()):                                          CertificateSpecValidityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecCertificateAuthorityConfiguration{}).Type1()):        CertificateAuthoritySpecCertificateAuthorityConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecCertificateAuthorityConfigurationSubject{}).Type1()): CertificateAuthoritySpecCertificateAuthorityConfigurationSubjectCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecRevocationConfiguration{}).Type1()):                  CertificateAuthoritySpecRevocationConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecRevocationConfigurationCrlConfiguration{}).Type1()):  CertificateAuthoritySpecRevocationConfigurationCrlConfigurationCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecValidity{}).Type1()):                                          CertificateSpecValidityCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecCertificateAuthorityConfiguration{}).Type1()):        CertificateAuthoritySpecCertificateAuthorityConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecCertificateAuthorityConfigurationSubject{}).Type1()): CertificateAuthoritySpecCertificateAuthorityConfigurationSubjectCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecRevocationConfiguration{}).Type1()):                  CertificateAuthoritySpecRevocationConfigurationCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecRevocationConfigurationCrlConfiguration{}).Type1()):  CertificateAuthoritySpecRevocationConfigurationCrlConfigurationCodec{},
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
type CertificateSpecValidityCodec struct {
}

func (CertificateSpecValidityCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CertificateSpecValidity)(ptr) == nil
}

func (CertificateSpecValidityCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CertificateSpecValidity)(ptr)
	var objs []CertificateSpecValidity
	if obj != nil {
		objs = []CertificateSpecValidity{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecValidity{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CertificateSpecValidityCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CertificateSpecValidity)(ptr) = CertificateSpecValidity{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CertificateSpecValidity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecValidity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CertificateSpecValidity)(ptr) = objs[0]
			} else {
				*(*CertificateSpecValidity)(ptr) = CertificateSpecValidity{}
			}
		} else {
			*(*CertificateSpecValidity)(ptr) = CertificateSpecValidity{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CertificateSpecValidity

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateSpecValidity{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CertificateSpecValidity)(ptr) = obj
		} else {
			*(*CertificateSpecValidity)(ptr) = CertificateSpecValidity{}
		}
	default:
		iter.ReportError("decode CertificateSpecValidity", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CertificateAuthoritySpecCertificateAuthorityConfigurationCodec struct {
}

func (CertificateAuthoritySpecCertificateAuthorityConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CertificateAuthoritySpecCertificateAuthorityConfiguration)(ptr) == nil
}

func (CertificateAuthoritySpecCertificateAuthorityConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CertificateAuthoritySpecCertificateAuthorityConfiguration)(ptr)
	var objs []CertificateAuthoritySpecCertificateAuthorityConfiguration
	if obj != nil {
		objs = []CertificateAuthoritySpecCertificateAuthorityConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecCertificateAuthorityConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CertificateAuthoritySpecCertificateAuthorityConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CertificateAuthoritySpecCertificateAuthorityConfiguration)(ptr) = CertificateAuthoritySpecCertificateAuthorityConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CertificateAuthoritySpecCertificateAuthorityConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecCertificateAuthorityConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CertificateAuthoritySpecCertificateAuthorityConfiguration)(ptr) = objs[0]
			} else {
				*(*CertificateAuthoritySpecCertificateAuthorityConfiguration)(ptr) = CertificateAuthoritySpecCertificateAuthorityConfiguration{}
			}
		} else {
			*(*CertificateAuthoritySpecCertificateAuthorityConfiguration)(ptr) = CertificateAuthoritySpecCertificateAuthorityConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CertificateAuthoritySpecCertificateAuthorityConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecCertificateAuthorityConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CertificateAuthoritySpecCertificateAuthorityConfiguration)(ptr) = obj
		} else {
			*(*CertificateAuthoritySpecCertificateAuthorityConfiguration)(ptr) = CertificateAuthoritySpecCertificateAuthorityConfiguration{}
		}
	default:
		iter.ReportError("decode CertificateAuthoritySpecCertificateAuthorityConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CertificateAuthoritySpecCertificateAuthorityConfigurationSubjectCodec struct {
}

func (CertificateAuthoritySpecCertificateAuthorityConfigurationSubjectCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CertificateAuthoritySpecCertificateAuthorityConfigurationSubject)(ptr) == nil
}

func (CertificateAuthoritySpecCertificateAuthorityConfigurationSubjectCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CertificateAuthoritySpecCertificateAuthorityConfigurationSubject)(ptr)
	var objs []CertificateAuthoritySpecCertificateAuthorityConfigurationSubject
	if obj != nil {
		objs = []CertificateAuthoritySpecCertificateAuthorityConfigurationSubject{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecCertificateAuthorityConfigurationSubject{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CertificateAuthoritySpecCertificateAuthorityConfigurationSubjectCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CertificateAuthoritySpecCertificateAuthorityConfigurationSubject)(ptr) = CertificateAuthoritySpecCertificateAuthorityConfigurationSubject{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CertificateAuthoritySpecCertificateAuthorityConfigurationSubject

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecCertificateAuthorityConfigurationSubject{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CertificateAuthoritySpecCertificateAuthorityConfigurationSubject)(ptr) = objs[0]
			} else {
				*(*CertificateAuthoritySpecCertificateAuthorityConfigurationSubject)(ptr) = CertificateAuthoritySpecCertificateAuthorityConfigurationSubject{}
			}
		} else {
			*(*CertificateAuthoritySpecCertificateAuthorityConfigurationSubject)(ptr) = CertificateAuthoritySpecCertificateAuthorityConfigurationSubject{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CertificateAuthoritySpecCertificateAuthorityConfigurationSubject

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecCertificateAuthorityConfigurationSubject{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CertificateAuthoritySpecCertificateAuthorityConfigurationSubject)(ptr) = obj
		} else {
			*(*CertificateAuthoritySpecCertificateAuthorityConfigurationSubject)(ptr) = CertificateAuthoritySpecCertificateAuthorityConfigurationSubject{}
		}
	default:
		iter.ReportError("decode CertificateAuthoritySpecCertificateAuthorityConfigurationSubject", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CertificateAuthoritySpecRevocationConfigurationCodec struct {
}

func (CertificateAuthoritySpecRevocationConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CertificateAuthoritySpecRevocationConfiguration)(ptr) == nil
}

func (CertificateAuthoritySpecRevocationConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CertificateAuthoritySpecRevocationConfiguration)(ptr)
	var objs []CertificateAuthoritySpecRevocationConfiguration
	if obj != nil {
		objs = []CertificateAuthoritySpecRevocationConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecRevocationConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CertificateAuthoritySpecRevocationConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CertificateAuthoritySpecRevocationConfiguration)(ptr) = CertificateAuthoritySpecRevocationConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CertificateAuthoritySpecRevocationConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecRevocationConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CertificateAuthoritySpecRevocationConfiguration)(ptr) = objs[0]
			} else {
				*(*CertificateAuthoritySpecRevocationConfiguration)(ptr) = CertificateAuthoritySpecRevocationConfiguration{}
			}
		} else {
			*(*CertificateAuthoritySpecRevocationConfiguration)(ptr) = CertificateAuthoritySpecRevocationConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CertificateAuthoritySpecRevocationConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecRevocationConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CertificateAuthoritySpecRevocationConfiguration)(ptr) = obj
		} else {
			*(*CertificateAuthoritySpecRevocationConfiguration)(ptr) = CertificateAuthoritySpecRevocationConfiguration{}
		}
	default:
		iter.ReportError("decode CertificateAuthoritySpecRevocationConfiguration", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type CertificateAuthoritySpecRevocationConfigurationCrlConfigurationCodec struct {
}

func (CertificateAuthoritySpecRevocationConfigurationCrlConfigurationCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*CertificateAuthoritySpecRevocationConfigurationCrlConfiguration)(ptr) == nil
}

func (CertificateAuthoritySpecRevocationConfigurationCrlConfigurationCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*CertificateAuthoritySpecRevocationConfigurationCrlConfiguration)(ptr)
	var objs []CertificateAuthoritySpecRevocationConfigurationCrlConfiguration
	if obj != nil {
		objs = []CertificateAuthoritySpecRevocationConfigurationCrlConfiguration{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecRevocationConfigurationCrlConfiguration{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (CertificateAuthoritySpecRevocationConfigurationCrlConfigurationCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*CertificateAuthoritySpecRevocationConfigurationCrlConfiguration)(ptr) = CertificateAuthoritySpecRevocationConfigurationCrlConfiguration{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []CertificateAuthoritySpecRevocationConfigurationCrlConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecRevocationConfigurationCrlConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*CertificateAuthoritySpecRevocationConfigurationCrlConfiguration)(ptr) = objs[0]
			} else {
				*(*CertificateAuthoritySpecRevocationConfigurationCrlConfiguration)(ptr) = CertificateAuthoritySpecRevocationConfigurationCrlConfiguration{}
			}
		} else {
			*(*CertificateAuthoritySpecRevocationConfigurationCrlConfiguration)(ptr) = CertificateAuthoritySpecRevocationConfigurationCrlConfiguration{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj CertificateAuthoritySpecRevocationConfigurationCrlConfiguration

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(CertificateAuthoritySpecRevocationConfigurationCrlConfiguration{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*CertificateAuthoritySpecRevocationConfigurationCrlConfiguration)(ptr) = obj
		} else {
			*(*CertificateAuthoritySpecRevocationConfigurationCrlConfiguration)(ptr) = CertificateAuthoritySpecRevocationConfigurationCrlConfiguration{}
		}
	default:
		iter.ReportError("decode CertificateAuthoritySpecRevocationConfigurationCrlConfiguration", "unexpected JSON type")
	}
}
