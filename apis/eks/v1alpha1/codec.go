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
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEncryptionConfig{}).Type1()):         ClusterSpecEncryptionConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEncryptionConfigProvider{}).Type1()): ClusterSpecEncryptionConfigProviderCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecKubernetesNetworkConfig{}).Type1()):  ClusterSpecKubernetesNetworkConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecVpcConfig{}).Type1()):                ClusterSpecVpcConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(IdentityProviderConfigSpecOidc{}).Type1()):      IdentityProviderConfigSpecOidcCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecLaunchTemplate{}).Type1()):         NodeGroupSpecLaunchTemplateCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecRemoteAccess{}).Type1()):           NodeGroupSpecRemoteAccessCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecScalingConfig{}).Type1()):          NodeGroupSpecScalingConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecUpdateConfig{}).Type1()):           NodeGroupSpecUpdateConfigCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEncryptionConfig{}).Type1()):         ClusterSpecEncryptionConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEncryptionConfigProvider{}).Type1()): ClusterSpecEncryptionConfigProviderCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecKubernetesNetworkConfig{}).Type1()):  ClusterSpecKubernetesNetworkConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecVpcConfig{}).Type1()):                ClusterSpecVpcConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(IdentityProviderConfigSpecOidc{}).Type1()):      IdentityProviderConfigSpecOidcCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecLaunchTemplate{}).Type1()):         NodeGroupSpecLaunchTemplateCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecRemoteAccess{}).Type1()):           NodeGroupSpecRemoteAccessCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecScalingConfig{}).Type1()):          NodeGroupSpecScalingConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecUpdateConfig{}).Type1()):           NodeGroupSpecUpdateConfigCodec{},
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
type ClusterSpecEncryptionConfigCodec struct {
}

func (ClusterSpecEncryptionConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecEncryptionConfig)(ptr) == nil
}

func (ClusterSpecEncryptionConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecEncryptionConfig)(ptr)
	var objs []ClusterSpecEncryptionConfig
	if obj != nil {
		objs = []ClusterSpecEncryptionConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEncryptionConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecEncryptionConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecEncryptionConfig)(ptr) = ClusterSpecEncryptionConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecEncryptionConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEncryptionConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecEncryptionConfig)(ptr) = objs[0]
			} else {
				*(*ClusterSpecEncryptionConfig)(ptr) = ClusterSpecEncryptionConfig{}
			}
		} else {
			*(*ClusterSpecEncryptionConfig)(ptr) = ClusterSpecEncryptionConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecEncryptionConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEncryptionConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecEncryptionConfig)(ptr) = obj
		} else {
			*(*ClusterSpecEncryptionConfig)(ptr) = ClusterSpecEncryptionConfig{}
		}
	default:
		iter.ReportError("decode ClusterSpecEncryptionConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecEncryptionConfigProviderCodec struct {
}

func (ClusterSpecEncryptionConfigProviderCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecEncryptionConfigProvider)(ptr) == nil
}

func (ClusterSpecEncryptionConfigProviderCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecEncryptionConfigProvider)(ptr)
	var objs []ClusterSpecEncryptionConfigProvider
	if obj != nil {
		objs = []ClusterSpecEncryptionConfigProvider{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEncryptionConfigProvider{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecEncryptionConfigProviderCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecEncryptionConfigProvider)(ptr) = ClusterSpecEncryptionConfigProvider{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecEncryptionConfigProvider

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEncryptionConfigProvider{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecEncryptionConfigProvider)(ptr) = objs[0]
			} else {
				*(*ClusterSpecEncryptionConfigProvider)(ptr) = ClusterSpecEncryptionConfigProvider{}
			}
		} else {
			*(*ClusterSpecEncryptionConfigProvider)(ptr) = ClusterSpecEncryptionConfigProvider{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecEncryptionConfigProvider

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecEncryptionConfigProvider{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecEncryptionConfigProvider)(ptr) = obj
		} else {
			*(*ClusterSpecEncryptionConfigProvider)(ptr) = ClusterSpecEncryptionConfigProvider{}
		}
	default:
		iter.ReportError("decode ClusterSpecEncryptionConfigProvider", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecKubernetesNetworkConfigCodec struct {
}

func (ClusterSpecKubernetesNetworkConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecKubernetesNetworkConfig)(ptr) == nil
}

func (ClusterSpecKubernetesNetworkConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecKubernetesNetworkConfig)(ptr)
	var objs []ClusterSpecKubernetesNetworkConfig
	if obj != nil {
		objs = []ClusterSpecKubernetesNetworkConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecKubernetesNetworkConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecKubernetesNetworkConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecKubernetesNetworkConfig)(ptr) = ClusterSpecKubernetesNetworkConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecKubernetesNetworkConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecKubernetesNetworkConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecKubernetesNetworkConfig)(ptr) = objs[0]
			} else {
				*(*ClusterSpecKubernetesNetworkConfig)(ptr) = ClusterSpecKubernetesNetworkConfig{}
			}
		} else {
			*(*ClusterSpecKubernetesNetworkConfig)(ptr) = ClusterSpecKubernetesNetworkConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecKubernetesNetworkConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecKubernetesNetworkConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecKubernetesNetworkConfig)(ptr) = obj
		} else {
			*(*ClusterSpecKubernetesNetworkConfig)(ptr) = ClusterSpecKubernetesNetworkConfig{}
		}
	default:
		iter.ReportError("decode ClusterSpecKubernetesNetworkConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type ClusterSpecVpcConfigCodec struct {
}

func (ClusterSpecVpcConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*ClusterSpecVpcConfig)(ptr) == nil
}

func (ClusterSpecVpcConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*ClusterSpecVpcConfig)(ptr)
	var objs []ClusterSpecVpcConfig
	if obj != nil {
		objs = []ClusterSpecVpcConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecVpcConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (ClusterSpecVpcConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*ClusterSpecVpcConfig)(ptr) = ClusterSpecVpcConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []ClusterSpecVpcConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecVpcConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*ClusterSpecVpcConfig)(ptr) = objs[0]
			} else {
				*(*ClusterSpecVpcConfig)(ptr) = ClusterSpecVpcConfig{}
			}
		} else {
			*(*ClusterSpecVpcConfig)(ptr) = ClusterSpecVpcConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj ClusterSpecVpcConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(ClusterSpecVpcConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*ClusterSpecVpcConfig)(ptr) = obj
		} else {
			*(*ClusterSpecVpcConfig)(ptr) = ClusterSpecVpcConfig{}
		}
	default:
		iter.ReportError("decode ClusterSpecVpcConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type IdentityProviderConfigSpecOidcCodec struct {
}

func (IdentityProviderConfigSpecOidcCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*IdentityProviderConfigSpecOidc)(ptr) == nil
}

func (IdentityProviderConfigSpecOidcCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*IdentityProviderConfigSpecOidc)(ptr)
	var objs []IdentityProviderConfigSpecOidc
	if obj != nil {
		objs = []IdentityProviderConfigSpecOidc{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(IdentityProviderConfigSpecOidc{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (IdentityProviderConfigSpecOidcCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*IdentityProviderConfigSpecOidc)(ptr) = IdentityProviderConfigSpecOidc{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []IdentityProviderConfigSpecOidc

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(IdentityProviderConfigSpecOidc{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*IdentityProviderConfigSpecOidc)(ptr) = objs[0]
			} else {
				*(*IdentityProviderConfigSpecOidc)(ptr) = IdentityProviderConfigSpecOidc{}
			}
		} else {
			*(*IdentityProviderConfigSpecOidc)(ptr) = IdentityProviderConfigSpecOidc{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj IdentityProviderConfigSpecOidc

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(IdentityProviderConfigSpecOidc{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*IdentityProviderConfigSpecOidc)(ptr) = obj
		} else {
			*(*IdentityProviderConfigSpecOidc)(ptr) = IdentityProviderConfigSpecOidc{}
		}
	default:
		iter.ReportError("decode IdentityProviderConfigSpecOidc", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type NodeGroupSpecLaunchTemplateCodec struct {
}

func (NodeGroupSpecLaunchTemplateCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*NodeGroupSpecLaunchTemplate)(ptr) == nil
}

func (NodeGroupSpecLaunchTemplateCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*NodeGroupSpecLaunchTemplate)(ptr)
	var objs []NodeGroupSpecLaunchTemplate
	if obj != nil {
		objs = []NodeGroupSpecLaunchTemplate{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecLaunchTemplate{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (NodeGroupSpecLaunchTemplateCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*NodeGroupSpecLaunchTemplate)(ptr) = NodeGroupSpecLaunchTemplate{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []NodeGroupSpecLaunchTemplate

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecLaunchTemplate{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*NodeGroupSpecLaunchTemplate)(ptr) = objs[0]
			} else {
				*(*NodeGroupSpecLaunchTemplate)(ptr) = NodeGroupSpecLaunchTemplate{}
			}
		} else {
			*(*NodeGroupSpecLaunchTemplate)(ptr) = NodeGroupSpecLaunchTemplate{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj NodeGroupSpecLaunchTemplate

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecLaunchTemplate{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*NodeGroupSpecLaunchTemplate)(ptr) = obj
		} else {
			*(*NodeGroupSpecLaunchTemplate)(ptr) = NodeGroupSpecLaunchTemplate{}
		}
	default:
		iter.ReportError("decode NodeGroupSpecLaunchTemplate", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type NodeGroupSpecRemoteAccessCodec struct {
}

func (NodeGroupSpecRemoteAccessCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*NodeGroupSpecRemoteAccess)(ptr) == nil
}

func (NodeGroupSpecRemoteAccessCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*NodeGroupSpecRemoteAccess)(ptr)
	var objs []NodeGroupSpecRemoteAccess
	if obj != nil {
		objs = []NodeGroupSpecRemoteAccess{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecRemoteAccess{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (NodeGroupSpecRemoteAccessCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*NodeGroupSpecRemoteAccess)(ptr) = NodeGroupSpecRemoteAccess{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []NodeGroupSpecRemoteAccess

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecRemoteAccess{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*NodeGroupSpecRemoteAccess)(ptr) = objs[0]
			} else {
				*(*NodeGroupSpecRemoteAccess)(ptr) = NodeGroupSpecRemoteAccess{}
			}
		} else {
			*(*NodeGroupSpecRemoteAccess)(ptr) = NodeGroupSpecRemoteAccess{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj NodeGroupSpecRemoteAccess

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecRemoteAccess{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*NodeGroupSpecRemoteAccess)(ptr) = obj
		} else {
			*(*NodeGroupSpecRemoteAccess)(ptr) = NodeGroupSpecRemoteAccess{}
		}
	default:
		iter.ReportError("decode NodeGroupSpecRemoteAccess", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type NodeGroupSpecScalingConfigCodec struct {
}

func (NodeGroupSpecScalingConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*NodeGroupSpecScalingConfig)(ptr) == nil
}

func (NodeGroupSpecScalingConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*NodeGroupSpecScalingConfig)(ptr)
	var objs []NodeGroupSpecScalingConfig
	if obj != nil {
		objs = []NodeGroupSpecScalingConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecScalingConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (NodeGroupSpecScalingConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*NodeGroupSpecScalingConfig)(ptr) = NodeGroupSpecScalingConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []NodeGroupSpecScalingConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecScalingConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*NodeGroupSpecScalingConfig)(ptr) = objs[0]
			} else {
				*(*NodeGroupSpecScalingConfig)(ptr) = NodeGroupSpecScalingConfig{}
			}
		} else {
			*(*NodeGroupSpecScalingConfig)(ptr) = NodeGroupSpecScalingConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj NodeGroupSpecScalingConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecScalingConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*NodeGroupSpecScalingConfig)(ptr) = obj
		} else {
			*(*NodeGroupSpecScalingConfig)(ptr) = NodeGroupSpecScalingConfig{}
		}
	default:
		iter.ReportError("decode NodeGroupSpecScalingConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type NodeGroupSpecUpdateConfigCodec struct {
}

func (NodeGroupSpecUpdateConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*NodeGroupSpecUpdateConfig)(ptr) == nil
}

func (NodeGroupSpecUpdateConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*NodeGroupSpecUpdateConfig)(ptr)
	var objs []NodeGroupSpecUpdateConfig
	if obj != nil {
		objs = []NodeGroupSpecUpdateConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecUpdateConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (NodeGroupSpecUpdateConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*NodeGroupSpecUpdateConfig)(ptr) = NodeGroupSpecUpdateConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []NodeGroupSpecUpdateConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecUpdateConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*NodeGroupSpecUpdateConfig)(ptr) = objs[0]
			} else {
				*(*NodeGroupSpecUpdateConfig)(ptr) = NodeGroupSpecUpdateConfig{}
			}
		} else {
			*(*NodeGroupSpecUpdateConfig)(ptr) = NodeGroupSpecUpdateConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj NodeGroupSpecUpdateConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(NodeGroupSpecUpdateConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*NodeGroupSpecUpdateConfig)(ptr) = obj
		} else {
			*(*NodeGroupSpecUpdateConfig)(ptr) = NodeGroupSpecUpdateConfig{}
		}
	default:
		iter.ReportError("decode NodeGroupSpecUpdateConfig", "unexpected JSON type")
	}
}
