/*
Licensed under the GNU AFFERO GENERAL PUBLIC LICENSE Version 3
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VinceSpecApplyConfiguration represents an declarative configuration of the VinceSpec type for use
// with apply.
type VinceSpecApplyConfiguration struct {
	*VolumeApplyConfiguration    `json:"volume,omitempty"`
	*ContainerApplyConfiguration `json:"container,omitempty"`
}

// VinceSpecApplyConfiguration constructs an declarative configuration of the VinceSpec type for use with
// apply.
func VinceSpec() *VinceSpecApplyConfiguration {
	return &VinceSpecApplyConfiguration{}
}

// WithSelector sets the Selector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Selector field is set to the value of the last call.
func (b *VinceSpecApplyConfiguration) WithSelector(value v1.LabelSelector) *VinceSpecApplyConfiguration {
	b.ensureVolumeApplyConfigurationExists()
	b.Selector = &value
	return b
}

// WithSize sets the Size field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Size field is set to the value of the last call.
func (b *VinceSpecApplyConfiguration) WithSize(value resource.Quantity) *VinceSpecApplyConfiguration {
	b.ensureVolumeApplyConfigurationExists()
	b.Size = &value
	return b
}

// WithStorageClass sets the StorageClass field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageClass field is set to the value of the last call.
func (b *VinceSpecApplyConfiguration) WithStorageClass(value string) *VinceSpecApplyConfiguration {
	b.ensureVolumeApplyConfigurationExists()
	b.StorageClass = &value
	return b
}

// WithSubPath sets the SubPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SubPath field is set to the value of the last call.
func (b *VinceSpecApplyConfiguration) WithSubPath(value string) *VinceSpecApplyConfiguration {
	b.ensureVolumeApplyConfigurationExists()
	b.SubPath = &value
	return b
}

func (b *VinceSpecApplyConfiguration) ensureVolumeApplyConfigurationExists() {
	if b.VolumeApplyConfiguration == nil {
		b.VolumeApplyConfiguration = &VolumeApplyConfiguration{}
	}
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *VinceSpecApplyConfiguration) WithImage(value string) *VinceSpecApplyConfiguration {
	b.ensureContainerApplyConfigurationExists()
	b.Image = &value
	return b
}

// WithEnv adds the given value to the Env field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Env field.
func (b *VinceSpecApplyConfiguration) WithEnv(values ...corev1.EnvVar) *VinceSpecApplyConfiguration {
	b.ensureContainerApplyConfigurationExists()
	for i := range values {
		b.Env = append(b.Env, values[i])
	}
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *VinceSpecApplyConfiguration) WithResources(value corev1.ResourceRequirements) *VinceSpecApplyConfiguration {
	b.ensureContainerApplyConfigurationExists()
	b.Resources = &value
	return b
}

func (b *VinceSpecApplyConfiguration) ensureContainerApplyConfigurationExists() {
	if b.ContainerApplyConfiguration == nil {
		b.ContainerApplyConfiguration = &ContainerApplyConfiguration{}
	}
}
