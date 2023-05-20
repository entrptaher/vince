/*
Licensed under the GNU AFFERO GENERAL PUBLIC LICENSE Version 3
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1alpha1 "github.com/gernest/vince/pkg/apis/vince/v1alpha1"
	vincev1alpha1 "github.com/gernest/vince/pkg/gen/client/vince/applyconfiguration/vince/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=staples, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("Container"):
		return &vincev1alpha1.ContainerApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Site"):
		return &vincev1alpha1.SiteApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SiteSpec"):
		return &vincev1alpha1.SiteSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("SiteStatus"):
		return &vincev1alpha1.SiteStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Target"):
		return &vincev1alpha1.TargetApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Vince"):
		return &vincev1alpha1.VinceApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("VinceSpec"):
		return &vincev1alpha1.VinceSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Volume"):
		return &vincev1alpha1.VolumeApplyConfiguration{}

	}
	return nil
}
