/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta3

import (
	"context"

	v1beta3 "k8s.io/api/flowcontrol/v1beta3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	flowcontrolv1beta3 "k8s.io/client-go/applyconfigurations/flowcontrol/v1beta3"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// PriorityLevelConfigurationsGetter has a method to return a PriorityLevelConfigurationInterface.
// A group's client should implement this interface.
type PriorityLevelConfigurationsGetter interface {
	PriorityLevelConfigurations() PriorityLevelConfigurationInterface
}

// PriorityLevelConfigurationInterface has methods to work with PriorityLevelConfiguration resources.
type PriorityLevelConfigurationInterface interface {
	Create(ctx context.Context, priorityLevelConfiguration *v1beta3.PriorityLevelConfiguration, opts v1.CreateOptions) (*v1beta3.PriorityLevelConfiguration, error)
	Update(ctx context.Context, priorityLevelConfiguration *v1beta3.PriorityLevelConfiguration, opts v1.UpdateOptions) (*v1beta3.PriorityLevelConfiguration, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, priorityLevelConfiguration *v1beta3.PriorityLevelConfiguration, opts v1.UpdateOptions) (*v1beta3.PriorityLevelConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta3.PriorityLevelConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta3.PriorityLevelConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.PriorityLevelConfiguration, err error)
	Apply(ctx context.Context, priorityLevelConfiguration *flowcontrolv1beta3.PriorityLevelConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *v1beta3.PriorityLevelConfiguration, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, priorityLevelConfiguration *flowcontrolv1beta3.PriorityLevelConfigurationApplyConfiguration, opts v1.ApplyOptions) (result *v1beta3.PriorityLevelConfiguration, err error)
	PriorityLevelConfigurationExpansion
}

// priorityLevelConfigurations implements PriorityLevelConfigurationInterface
type priorityLevelConfigurations struct {
	*gentype.ClientWithListAndApply[*v1beta3.PriorityLevelConfiguration, *v1beta3.PriorityLevelConfigurationList, *flowcontrolv1beta3.PriorityLevelConfigurationApplyConfiguration]
}

// newPriorityLevelConfigurations returns a PriorityLevelConfigurations
func newPriorityLevelConfigurations(c *FlowcontrolV1beta3Client) *priorityLevelConfigurations {
	return &priorityLevelConfigurations{
		gentype.NewClientWithListAndApply[*v1beta3.PriorityLevelConfiguration, *v1beta3.PriorityLevelConfigurationList, *flowcontrolv1beta3.PriorityLevelConfigurationApplyConfiguration](
			"prioritylevelconfigurations",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1beta3.PriorityLevelConfiguration { return &v1beta3.PriorityLevelConfiguration{} },
			func() *v1beta3.PriorityLevelConfigurationList { return &v1beta3.PriorityLevelConfigurationList{} }),
	}
}
