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
	"fmt"
	"strings"

	base "kubeform.dev/apimachinery/api/v1alpha1"
	"kubeform.dev/apimachinery/pkg/util"

	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *ComputeEnvironment) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-batch-aws-kubeform-com-v1alpha1-computeenvironment,mutating=false,failurePolicy=fail,groups=batch.aws.kubeform.com,resources=computeenvironments,versions=v1alpha1,name=computeenvironment.batch.aws.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &ComputeEnvironment{}

var computeenvironmentForceNewList = map[string]bool{
	"/compute_environment_name":                                   true,
	"/compute_environment_name_prefix":                            true,
	"/compute_resources/*/allocation_strategy":                    true,
	"/compute_resources/*/bid_percentage":                         true,
	"/compute_resources/*/ec2_configuration/*/image_id_override":  true,
	"/compute_resources/*/ec2_configuration/*/image_type":         true,
	"/compute_resources/*/ec2_key_pair":                           true,
	"/compute_resources/*/image_id":                               true,
	"/compute_resources/*/instance_role":                          true,
	"/compute_resources/*/instance_type":                          true,
	"/compute_resources/*/launch_template/*/launch_template_id":   true,
	"/compute_resources/*/launch_template/*/launch_template_name": true,
	"/compute_resources/*/launch_template/*/version":              true,
	"/compute_resources/*/spot_iam_fleet_role":                    true,
	"/compute_resources/*/tags":                                   true,
	"/compute_resources/*/type":                                   true,
	"/type":                                                       true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ComputeEnvironment) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ComputeEnvironment) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*ComputeEnvironment)
	oldObj := res.Spec.Resource

	jsnitr := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		TagKey:                 "tf",
		ValidateJsonRawMessage: true,
		TypeEncoders:           GetEncoder(),
		TypeDecoders:           GetDecoder(),
	}.Froze()

	byteNew, err := jsnitr.Marshal(newObj)
	if err != nil {
		return err
	}
	tempNew := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteNew, &tempNew)
	if err != nil {
		return err
	}

	byteOld, err := jsnitr.Marshal(oldObj)
	if err != nil {
		return err
	}
	tempOld := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteOld, &tempOld)
	if err != nil {
		return err
	}

	for key, _ := range computeenvironmentForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`computeenvironment "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ComputeEnvironment) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`computeenvironment "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
