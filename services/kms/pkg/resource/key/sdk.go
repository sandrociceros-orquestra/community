// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package key

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/kms"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/kms/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.KMS{}
	_ = &svcapitypes.Key{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.DescribeKeyWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeKey", respErr)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "UNKNOWN" {
			return nil, ackerr.NotFound
		}
		return nil, respErr
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.KeyMetadata.AWSAccountId != nil {
		ko.Status.AWSAccountID = resp.KeyMetadata.AWSAccountId
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.KeyMetadata.Arn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.KeyMetadata.Arn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.KeyMetadata.CloudHsmClusterId != nil {
		ko.Status.CloudHsmClusterID = resp.KeyMetadata.CloudHsmClusterId
	}
	if resp.KeyMetadata.CreationDate != nil {
		ko.Status.CreationDate = &metav1.Time{*resp.KeyMetadata.CreationDate}
	}
	if resp.KeyMetadata.CustomKeyStoreId != nil {
		ko.Spec.CustomKeyStoreID = resp.KeyMetadata.CustomKeyStoreId
	}
	if resp.KeyMetadata.CustomerMasterKeySpec != nil {
		ko.Spec.CustomerMasterKeySpec = resp.KeyMetadata.CustomerMasterKeySpec
	}
	if resp.KeyMetadata.DeletionDate != nil {
		ko.Status.DeletionDate = &metav1.Time{*resp.KeyMetadata.DeletionDate}
	}
	if resp.KeyMetadata.Description != nil {
		ko.Spec.Description = resp.KeyMetadata.Description
	}
	if resp.KeyMetadata.Enabled != nil {
		ko.Status.Enabled = resp.KeyMetadata.Enabled
	}
	if resp.KeyMetadata.EncryptionAlgorithms != nil {
		f9 := []*string{}
		for _, f9iter := range resp.KeyMetadata.EncryptionAlgorithms {
			var f9elem string
			f9elem = *f9iter
			f9 = append(f9, &f9elem)
		}
		ko.Status.EncryptionAlgorithms = f9
	}
	if resp.KeyMetadata.ExpirationModel != nil {
		ko.Status.ExpirationModel = resp.KeyMetadata.ExpirationModel
	}
	if resp.KeyMetadata.KeyId != nil {
		ko.Status.KeyID = resp.KeyMetadata.KeyId
	}
	if resp.KeyMetadata.KeyManager != nil {
		ko.Status.KeyManager = resp.KeyMetadata.KeyManager
	}
	if resp.KeyMetadata.KeyState != nil {
		ko.Status.KeyState = resp.KeyMetadata.KeyState
	}
	if resp.KeyMetadata.KeyUsage != nil {
		ko.Spec.KeyUsage = resp.KeyMetadata.KeyUsage
	}
	if resp.KeyMetadata.Origin != nil {
		ko.Spec.Origin = resp.KeyMetadata.Origin
	}
	if resp.KeyMetadata.SigningAlgorithms != nil {
		f16 := []*string{}
		for _, f16iter := range resp.KeyMetadata.SigningAlgorithms {
			var f16elem string
			f16elem = *f16iter
			f16 = append(f16, &f16elem)
		}
		ko.Status.SigningAlgorithms = f16
	}
	if resp.KeyMetadata.ValidTo != nil {
		ko.Status.ValidTo = &metav1.Time{*resp.KeyMetadata.ValidTo}
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Status.KeyID == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeKeyInput, error) {
	res := &svcsdk.DescribeKeyInput{}

	if r.ko.Status.KeyID != nil {
		res.SetKeyId(*r.ko.Status.KeyID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateKeyWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateKey", respErr)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.KeyMetadata.AWSAccountId != nil {
		ko.Status.AWSAccountID = resp.KeyMetadata.AWSAccountId
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.KeyMetadata.Arn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.KeyMetadata.Arn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.KeyMetadata.CloudHsmClusterId != nil {
		ko.Status.CloudHsmClusterID = resp.KeyMetadata.CloudHsmClusterId
	}
	if resp.KeyMetadata.CreationDate != nil {
		ko.Status.CreationDate = &metav1.Time{*resp.KeyMetadata.CreationDate}
	}
	if resp.KeyMetadata.DeletionDate != nil {
		ko.Status.DeletionDate = &metav1.Time{*resp.KeyMetadata.DeletionDate}
	}
	if resp.KeyMetadata.Enabled != nil {
		ko.Status.Enabled = resp.KeyMetadata.Enabled
	}
	if resp.KeyMetadata.EncryptionAlgorithms != nil {
		f9 := []*string{}
		for _, f9iter := range resp.KeyMetadata.EncryptionAlgorithms {
			var f9elem string
			f9elem = *f9iter
			f9 = append(f9, &f9elem)
		}
		ko.Status.EncryptionAlgorithms = f9
	}
	if resp.KeyMetadata.ExpirationModel != nil {
		ko.Status.ExpirationModel = resp.KeyMetadata.ExpirationModel
	}
	if resp.KeyMetadata.KeyId != nil {
		ko.Status.KeyID = resp.KeyMetadata.KeyId
	}
	if resp.KeyMetadata.KeyManager != nil {
		ko.Status.KeyManager = resp.KeyMetadata.KeyManager
	}
	if resp.KeyMetadata.KeyState != nil {
		ko.Status.KeyState = resp.KeyMetadata.KeyState
	}
	if resp.KeyMetadata.SigningAlgorithms != nil {
		f16 := []*string{}
		for _, f16iter := range resp.KeyMetadata.SigningAlgorithms {
			var f16elem string
			f16elem = *f16iter
			f16 = append(f16, &f16elem)
		}
		ko.Status.SigningAlgorithms = f16
	}
	if resp.KeyMetadata.ValidTo != nil {
		ko.Status.ValidTo = &metav1.Time{*resp.KeyMetadata.ValidTo}
	}

	rm.setStatusDefaults(ko)

	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateKeyInput, error) {
	res := &svcsdk.CreateKeyInput{}

	if r.ko.Spec.BypassPolicyLockoutSafetyCheck != nil {
		res.SetBypassPolicyLockoutSafetyCheck(*r.ko.Spec.BypassPolicyLockoutSafetyCheck)
	}
	if r.ko.Spec.CustomKeyStoreID != nil {
		res.SetCustomKeyStoreId(*r.ko.Spec.CustomKeyStoreID)
	}
	if r.ko.Spec.CustomerMasterKeySpec != nil {
		res.SetCustomerMasterKeySpec(*r.ko.Spec.CustomerMasterKeySpec)
	}
	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.KeyUsage != nil {
		res.SetKeyUsage(*r.ko.Spec.KeyUsage)
	}
	if r.ko.Spec.Origin != nil {
		res.SetOrigin(*r.ko.Spec.Origin)
	}
	if r.ko.Spec.Policy != nil {
		res.SetPolicy(*r.ko.Spec.Policy)
	}
	if r.ko.Spec.Tags != nil {
		f7 := []*svcsdk.Tag{}
		for _, f7iter := range r.ko.Spec.Tags {
			f7elem := &svcsdk.Tag{}
			if f7iter.TagKey != nil {
				f7elem.SetTagKey(*f7iter.TagKey)
			}
			if f7iter.TagValue != nil {
				f7elem.SetTagValue(*f7iter.TagValue)
			}
			f7 = append(f7, f7elem)
		}
		res.SetTags(f7)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	// TODO(jaypipes): Figure this out...
	return nil

}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Key,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
			break
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else if terminalCondition != nil {
		terminalCondition.Status = corev1.ConditionFalse
		terminalCondition.Message = nil
	}
	if terminalCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
