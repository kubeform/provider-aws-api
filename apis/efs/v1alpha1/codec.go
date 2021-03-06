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
		jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecPosixUser{}).Type1()):                 AccessPointSpecPosixUserCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecRootDirectory{}).Type1()):             AccessPointSpecRootDirectoryCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecRootDirectoryCreationInfo{}).Type1()): AccessPointSpecRootDirectoryCreationInfoCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BackupPolicySpecBackupPolicy{}).Type1()):             BackupPolicySpecBackupPolicyCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecPosixUser{}).Type1()):                 AccessPointSpecPosixUserCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecRootDirectory{}).Type1()):             AccessPointSpecRootDirectoryCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecRootDirectoryCreationInfo{}).Type1()): AccessPointSpecRootDirectoryCreationInfoCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(BackupPolicySpecBackupPolicy{}).Type1()):             BackupPolicySpecBackupPolicyCodec{},
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
type AccessPointSpecPosixUserCodec struct {
}

func (AccessPointSpecPosixUserCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AccessPointSpecPosixUser)(ptr) == nil
}

func (AccessPointSpecPosixUserCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AccessPointSpecPosixUser)(ptr)
	var objs []AccessPointSpecPosixUser
	if obj != nil {
		objs = []AccessPointSpecPosixUser{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecPosixUser{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AccessPointSpecPosixUserCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AccessPointSpecPosixUser)(ptr) = AccessPointSpecPosixUser{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AccessPointSpecPosixUser

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecPosixUser{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AccessPointSpecPosixUser)(ptr) = objs[0]
			} else {
				*(*AccessPointSpecPosixUser)(ptr) = AccessPointSpecPosixUser{}
			}
		} else {
			*(*AccessPointSpecPosixUser)(ptr) = AccessPointSpecPosixUser{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj AccessPointSpecPosixUser

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecPosixUser{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*AccessPointSpecPosixUser)(ptr) = obj
		} else {
			*(*AccessPointSpecPosixUser)(ptr) = AccessPointSpecPosixUser{}
		}
	default:
		iter.ReportError("decode AccessPointSpecPosixUser", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type AccessPointSpecRootDirectoryCodec struct {
}

func (AccessPointSpecRootDirectoryCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AccessPointSpecRootDirectory)(ptr) == nil
}

func (AccessPointSpecRootDirectoryCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AccessPointSpecRootDirectory)(ptr)
	var objs []AccessPointSpecRootDirectory
	if obj != nil {
		objs = []AccessPointSpecRootDirectory{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecRootDirectory{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AccessPointSpecRootDirectoryCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AccessPointSpecRootDirectory)(ptr) = AccessPointSpecRootDirectory{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AccessPointSpecRootDirectory

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecRootDirectory{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AccessPointSpecRootDirectory)(ptr) = objs[0]
			} else {
				*(*AccessPointSpecRootDirectory)(ptr) = AccessPointSpecRootDirectory{}
			}
		} else {
			*(*AccessPointSpecRootDirectory)(ptr) = AccessPointSpecRootDirectory{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj AccessPointSpecRootDirectory

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecRootDirectory{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*AccessPointSpecRootDirectory)(ptr) = obj
		} else {
			*(*AccessPointSpecRootDirectory)(ptr) = AccessPointSpecRootDirectory{}
		}
	default:
		iter.ReportError("decode AccessPointSpecRootDirectory", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type AccessPointSpecRootDirectoryCreationInfoCodec struct {
}

func (AccessPointSpecRootDirectoryCreationInfoCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*AccessPointSpecRootDirectoryCreationInfo)(ptr) == nil
}

func (AccessPointSpecRootDirectoryCreationInfoCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*AccessPointSpecRootDirectoryCreationInfo)(ptr)
	var objs []AccessPointSpecRootDirectoryCreationInfo
	if obj != nil {
		objs = []AccessPointSpecRootDirectoryCreationInfo{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecRootDirectoryCreationInfo{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (AccessPointSpecRootDirectoryCreationInfoCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*AccessPointSpecRootDirectoryCreationInfo)(ptr) = AccessPointSpecRootDirectoryCreationInfo{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []AccessPointSpecRootDirectoryCreationInfo

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecRootDirectoryCreationInfo{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*AccessPointSpecRootDirectoryCreationInfo)(ptr) = objs[0]
			} else {
				*(*AccessPointSpecRootDirectoryCreationInfo)(ptr) = AccessPointSpecRootDirectoryCreationInfo{}
			}
		} else {
			*(*AccessPointSpecRootDirectoryCreationInfo)(ptr) = AccessPointSpecRootDirectoryCreationInfo{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj AccessPointSpecRootDirectoryCreationInfo

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(AccessPointSpecRootDirectoryCreationInfo{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*AccessPointSpecRootDirectoryCreationInfo)(ptr) = obj
		} else {
			*(*AccessPointSpecRootDirectoryCreationInfo)(ptr) = AccessPointSpecRootDirectoryCreationInfo{}
		}
	default:
		iter.ReportError("decode AccessPointSpecRootDirectoryCreationInfo", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type BackupPolicySpecBackupPolicyCodec struct {
}

func (BackupPolicySpecBackupPolicyCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*BackupPolicySpecBackupPolicy)(ptr) == nil
}

func (BackupPolicySpecBackupPolicyCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*BackupPolicySpecBackupPolicy)(ptr)
	var objs []BackupPolicySpecBackupPolicy
	if obj != nil {
		objs = []BackupPolicySpecBackupPolicy{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BackupPolicySpecBackupPolicy{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (BackupPolicySpecBackupPolicyCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*BackupPolicySpecBackupPolicy)(ptr) = BackupPolicySpecBackupPolicy{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []BackupPolicySpecBackupPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BackupPolicySpecBackupPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*BackupPolicySpecBackupPolicy)(ptr) = objs[0]
			} else {
				*(*BackupPolicySpecBackupPolicy)(ptr) = BackupPolicySpecBackupPolicy{}
			}
		} else {
			*(*BackupPolicySpecBackupPolicy)(ptr) = BackupPolicySpecBackupPolicy{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj BackupPolicySpecBackupPolicy

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(BackupPolicySpecBackupPolicy{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*BackupPolicySpecBackupPolicy)(ptr) = obj
		} else {
			*(*BackupPolicySpecBackupPolicy)(ptr) = BackupPolicySpecBackupPolicy{}
		}
	default:
		iter.ReportError("decode BackupPolicySpecBackupPolicy", "unexpected JSON type")
	}
}
