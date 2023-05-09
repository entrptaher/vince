/*
Licensed under the GNU AFFERO GENERAL PUBLIC LICENSE Version 3
*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/gernest/vince/pkg/apis/vince/v1alpha1"
	vincev1alpha1 "github.com/gernest/vince/pkg/gen/client/vince/applyconfiguration/vince/v1alpha1"
	scheme "github.com/gernest/vince/pkg/gen/client/vince/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SitesGetter has a method to return a SiteInterface.
// A group's client should implement this interface.
type SitesGetter interface {
	Sites(namespace string) SiteInterface
}

// SiteInterface has methods to work with Site resources.
type SiteInterface interface {
	Create(ctx context.Context, site *v1alpha1.Site, opts v1.CreateOptions) (*v1alpha1.Site, error)
	Update(ctx context.Context, site *v1alpha1.Site, opts v1.UpdateOptions) (*v1alpha1.Site, error)
	UpdateStatus(ctx context.Context, site *v1alpha1.Site, opts v1.UpdateOptions) (*v1alpha1.Site, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Site, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SiteList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Site, err error)
	Apply(ctx context.Context, site *vincev1alpha1.SiteApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Site, err error)
	ApplyStatus(ctx context.Context, site *vincev1alpha1.SiteApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Site, err error)
	SiteExpansion
}

// sites implements SiteInterface
type sites struct {
	client rest.Interface
	ns     string
}

// newSites returns a Sites
func newSites(c *StaplesV1alpha1Client, namespace string) *sites {
	return &sites{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the site, and returns the corresponding site object, and an error if there is any.
func (c *sites) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Site, err error) {
	result = &v1alpha1.Site{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sites").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Sites that match those selectors.
func (c *sites) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SiteList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SiteList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sites").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sites.
func (c *sites) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sites").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a site and creates it.  Returns the server's representation of the site, and an error, if there is any.
func (c *sites) Create(ctx context.Context, site *v1alpha1.Site, opts v1.CreateOptions) (result *v1alpha1.Site, err error) {
	result = &v1alpha1.Site{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sites").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(site).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a site and updates it. Returns the server's representation of the site, and an error, if there is any.
func (c *sites) Update(ctx context.Context, site *v1alpha1.Site, opts v1.UpdateOptions) (result *v1alpha1.Site, err error) {
	result = &v1alpha1.Site{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sites").
		Name(site.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(site).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *sites) UpdateStatus(ctx context.Context, site *v1alpha1.Site, opts v1.UpdateOptions) (result *v1alpha1.Site, err error) {
	result = &v1alpha1.Site{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sites").
		Name(site.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(site).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the site and deletes it. Returns an error if one occurs.
func (c *sites) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sites").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sites) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sites").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched site.
func (c *sites) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Site, err error) {
	result = &v1alpha1.Site{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sites").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied site.
func (c *sites) Apply(ctx context.Context, site *vincev1alpha1.SiteApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Site, err error) {
	if site == nil {
		return nil, fmt.Errorf("site provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(site)
	if err != nil {
		return nil, err
	}
	name := site.Name
	if name == nil {
		return nil, fmt.Errorf("site.Name must be provided to Apply")
	}
	result = &v1alpha1.Site{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("sites").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *sites) ApplyStatus(ctx context.Context, site *vincev1alpha1.SiteApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Site, err error) {
	if site == nil {
		return nil, fmt.Errorf("site provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(site)
	if err != nil {
		return nil, err
	}

	name := site.Name
	if name == nil {
		return nil, fmt.Errorf("site.Name must be provided to Apply")
	}

	result = &v1alpha1.Site{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("sites").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
