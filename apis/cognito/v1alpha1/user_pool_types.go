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
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type UserPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserPoolSpec   `json:"spec,omitempty"`
	Status            UserPoolStatus `json:"status,omitempty"`
}

type UserPoolSpecAccountRecoverySettingRecoveryMechanism struct {
	Name     *string `json:"name" tf:"name"`
	Priority *int64  `json:"priority" tf:"priority"`
}

type UserPoolSpecAccountRecoverySetting struct {
	// +kubebuilder:validation:MinItems=1
	RecoveryMechanism []UserPoolSpecAccountRecoverySettingRecoveryMechanism `json:"recoveryMechanism" tf:"recovery_mechanism"`
}

type UserPoolSpecAdminCreateUserConfigInviteMessageTemplate struct {
	// +optional
	EmailMessage *string `json:"emailMessage,omitempty" tf:"email_message"`
	// +optional
	EmailSubject *string `json:"emailSubject,omitempty" tf:"email_subject"`
	// +optional
	SmsMessage *string `json:"smsMessage,omitempty" tf:"sms_message"`
}

type UserPoolSpecAdminCreateUserConfig struct {
	// +optional
	AllowAdminCreateUserOnly *bool `json:"allowAdminCreateUserOnly,omitempty" tf:"allow_admin_create_user_only"`
	// +optional
	InviteMessageTemplate *UserPoolSpecAdminCreateUserConfigInviteMessageTemplate `json:"inviteMessageTemplate,omitempty" tf:"invite_message_template"`
}

type UserPoolSpecDeviceConfiguration struct {
	// +optional
	ChallengeRequiredOnNewDevice *bool `json:"challengeRequiredOnNewDevice,omitempty" tf:"challenge_required_on_new_device"`
	// +optional
	DeviceOnlyRememberedOnUserPrompt *bool `json:"deviceOnlyRememberedOnUserPrompt,omitempty" tf:"device_only_remembered_on_user_prompt"`
}

type UserPoolSpecEmailConfiguration struct {
	// +optional
	ConfigurationSet *string `json:"configurationSet,omitempty" tf:"configuration_set"`
	// +optional
	EmailSendingAccount *string `json:"emailSendingAccount,omitempty" tf:"email_sending_account"`
	// +optional
	FromEmailAddress *string `json:"fromEmailAddress,omitempty" tf:"from_email_address"`
	// +optional
	ReplyToEmailAddress *string `json:"replyToEmailAddress,omitempty" tf:"reply_to_email_address"`
	// +optional
	SourceArn *string `json:"sourceArn,omitempty" tf:"source_arn"`
}

type UserPoolSpecLambdaConfigCustomEmailSender struct {
	LambdaArn     *string `json:"lambdaArn" tf:"lambda_arn"`
	LambdaVersion *string `json:"lambdaVersion" tf:"lambda_version"`
}

type UserPoolSpecLambdaConfigCustomSmsSender struct {
	LambdaArn     *string `json:"lambdaArn" tf:"lambda_arn"`
	LambdaVersion *string `json:"lambdaVersion" tf:"lambda_version"`
}

type UserPoolSpecLambdaConfig struct {
	// +optional
	CreateAuthChallenge *string `json:"createAuthChallenge,omitempty" tf:"create_auth_challenge"`
	// +optional
	CustomEmailSender *UserPoolSpecLambdaConfigCustomEmailSender `json:"customEmailSender,omitempty" tf:"custom_email_sender"`
	// +optional
	CustomMessage *string `json:"customMessage,omitempty" tf:"custom_message"`
	// +optional
	CustomSmsSender *UserPoolSpecLambdaConfigCustomSmsSender `json:"customSmsSender,omitempty" tf:"custom_sms_sender"`
	// +optional
	DefineAuthChallenge *string `json:"defineAuthChallenge,omitempty" tf:"define_auth_challenge"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	PostAuthentication *string `json:"postAuthentication,omitempty" tf:"post_authentication"`
	// +optional
	PostConfirmation *string `json:"postConfirmation,omitempty" tf:"post_confirmation"`
	// +optional
	PreAuthentication *string `json:"preAuthentication,omitempty" tf:"pre_authentication"`
	// +optional
	PreSignUp *string `json:"preSignUp,omitempty" tf:"pre_sign_up"`
	// +optional
	PreTokenGeneration *string `json:"preTokenGeneration,omitempty" tf:"pre_token_generation"`
	// +optional
	UserMigration *string `json:"userMigration,omitempty" tf:"user_migration"`
	// +optional
	VerifyAuthChallengeResponse *string `json:"verifyAuthChallengeResponse,omitempty" tf:"verify_auth_challenge_response"`
}

type UserPoolSpecPasswordPolicy struct {
	// +optional
	MinimumLength *int64 `json:"minimumLength,omitempty" tf:"minimum_length"`
	// +optional
	RequireLowercase *bool `json:"requireLowercase,omitempty" tf:"require_lowercase"`
	// +optional
	RequireNumbers *bool `json:"requireNumbers,omitempty" tf:"require_numbers"`
	// +optional
	RequireSymbols *bool `json:"requireSymbols,omitempty" tf:"require_symbols"`
	// +optional
	RequireUppercase *bool `json:"requireUppercase,omitempty" tf:"require_uppercase"`
	// +optional
	TemporaryPasswordValidityDays *int64 `json:"temporaryPasswordValidityDays,omitempty" tf:"temporary_password_validity_days"`
}

type UserPoolSpecSchemaNumberAttributeConstraints struct {
	// +optional
	MaxValue *string `json:"maxValue,omitempty" tf:"max_value"`
	// +optional
	MinValue *string `json:"minValue,omitempty" tf:"min_value"`
}

type UserPoolSpecSchemaStringAttributeConstraints struct {
	// +optional
	MaxLength *string `json:"maxLength,omitempty" tf:"max_length"`
	// +optional
	MinLength *string `json:"minLength,omitempty" tf:"min_length"`
}

type UserPoolSpecSchema struct {
	AttributeDataType *string `json:"attributeDataType" tf:"attribute_data_type"`
	// +optional
	DeveloperOnlyAttribute *bool `json:"developerOnlyAttribute,omitempty" tf:"developer_only_attribute"`
	// +optional
	Mutable *bool   `json:"mutable,omitempty" tf:"mutable"`
	Name    *string `json:"name" tf:"name"`
	// +optional
	NumberAttributeConstraints *UserPoolSpecSchemaNumberAttributeConstraints `json:"numberAttributeConstraints,omitempty" tf:"number_attribute_constraints"`
	// +optional
	Required *bool `json:"required,omitempty" tf:"required"`
	// +optional
	StringAttributeConstraints *UserPoolSpecSchemaStringAttributeConstraints `json:"stringAttributeConstraints,omitempty" tf:"string_attribute_constraints"`
}

type UserPoolSpecSmsConfiguration struct {
	ExternalID   *string `json:"externalID" tf:"external_id"`
	SnsCallerArn *string `json:"snsCallerArn" tf:"sns_caller_arn"`
}

type UserPoolSpecSoftwareTokenMfaConfiguration struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
}

type UserPoolSpecUserPoolAddOns struct {
	AdvancedSecurityMode *string `json:"advancedSecurityMode" tf:"advanced_security_mode"`
}

type UserPoolSpecUsernameConfiguration struct {
	CaseSensitive *bool `json:"caseSensitive" tf:"case_sensitive"`
}

type UserPoolSpecVerificationMessageTemplate struct {
	// +optional
	DefaultEmailOption *string `json:"defaultEmailOption,omitempty" tf:"default_email_option"`
	// +optional
	EmailMessage *string `json:"emailMessage,omitempty" tf:"email_message"`
	// +optional
	EmailMessageByLink *string `json:"emailMessageByLink,omitempty" tf:"email_message_by_link"`
	// +optional
	EmailSubject *string `json:"emailSubject,omitempty" tf:"email_subject"`
	// +optional
	EmailSubjectByLink *string `json:"emailSubjectByLink,omitempty" tf:"email_subject_by_link"`
	// +optional
	SmsMessage *string `json:"smsMessage,omitempty" tf:"sms_message"`
}

type UserPoolSpec struct {
	State *UserPoolSpecResource `json:"state,omitempty" tf:"-"`

	Resource UserPoolSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type UserPoolSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccountRecoverySetting *UserPoolSpecAccountRecoverySetting `json:"accountRecoverySetting,omitempty" tf:"account_recovery_setting"`
	// +optional
	AdminCreateUserConfig *UserPoolSpecAdminCreateUserConfig `json:"adminCreateUserConfig,omitempty" tf:"admin_create_user_config"`
	// +optional
	AliasAttributes []string `json:"aliasAttributes,omitempty" tf:"alias_attributes"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AutoVerifiedAttributes []string `json:"autoVerifiedAttributes,omitempty" tf:"auto_verified_attributes"`
	// +optional
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date"`
	// +optional
	CustomDomain *string `json:"customDomain,omitempty" tf:"custom_domain"`
	// +optional
	DeviceConfiguration *UserPoolSpecDeviceConfiguration `json:"deviceConfiguration,omitempty" tf:"device_configuration"`
	// +optional
	Domain *string `json:"domain,omitempty" tf:"domain"`
	// +optional
	EmailConfiguration *UserPoolSpecEmailConfiguration `json:"emailConfiguration,omitempty" tf:"email_configuration"`
	// +optional
	EmailVerificationMessage *string `json:"emailVerificationMessage,omitempty" tf:"email_verification_message"`
	// +optional
	EmailVerificationSubject *string `json:"emailVerificationSubject,omitempty" tf:"email_verification_subject"`
	// +optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint"`
	// +optional
	EstimatedNumberOfUsers *int64 `json:"estimatedNumberOfUsers,omitempty" tf:"estimated_number_of_users"`
	// +optional
	LambdaConfig *UserPoolSpecLambdaConfig `json:"lambdaConfig,omitempty" tf:"lambda_config"`
	// +optional
	LastModifiedDate *string `json:"lastModifiedDate,omitempty" tf:"last_modified_date"`
	// +optional
	MfaConfiguration *string `json:"mfaConfiguration,omitempty" tf:"mfa_configuration"`
	Name             *string `json:"name" tf:"name"`
	// +optional
	PasswordPolicy *UserPoolSpecPasswordPolicy `json:"passwordPolicy,omitempty" tf:"password_policy"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	// +kubebuilder:validation:MinItems=1
	Schema []UserPoolSpecSchema `json:"schema,omitempty" tf:"schema"`
	// +optional
	SmsAuthenticationMessage *string `json:"smsAuthenticationMessage,omitempty" tf:"sms_authentication_message"`
	// +optional
	SmsConfiguration *UserPoolSpecSmsConfiguration `json:"smsConfiguration,omitempty" tf:"sms_configuration"`
	// +optional
	SmsVerificationMessage *string `json:"smsVerificationMessage,omitempty" tf:"sms_verification_message"`
	// +optional
	SoftwareTokenMfaConfiguration *UserPoolSpecSoftwareTokenMfaConfiguration `json:"softwareTokenMfaConfiguration,omitempty" tf:"software_token_mfa_configuration"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	UserPoolAddOns *UserPoolSpecUserPoolAddOns `json:"userPoolAddOns,omitempty" tf:"user_pool_add_ons"`
	// +optional
	UsernameAttributes []string `json:"usernameAttributes,omitempty" tf:"username_attributes"`
	// +optional
	UsernameConfiguration *UserPoolSpecUsernameConfiguration `json:"usernameConfiguration,omitempty" tf:"username_configuration"`
	// +optional
	VerificationMessageTemplate *UserPoolSpecVerificationMessageTemplate `json:"verificationMessageTemplate,omitempty" tf:"verification_message_template"`
}

type UserPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// UserPoolList is a list of UserPools
type UserPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of UserPool CRD objects
	Items []UserPool `json:"items,omitempty"`
}
