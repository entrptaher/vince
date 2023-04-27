/*
Licensed under the GNU AFFERO GENERAL PUBLIC LICENSE Version 3
*/
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gernest/vince/pkg/apis/site/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SiteLister helps list Sites.
// All objects returned here must be treated as read-only.
type SiteLister interface {
	// List lists all Sites in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Site, err error)
	// Sites returns an object that can list and get Sites.
	Sites(namespace string) SiteNamespaceLister
	SiteListerExpansion
}

// siteLister implements the SiteLister interface.
type siteLister struct {
	indexer cache.Indexer
}

// NewSiteLister returns a new SiteLister.
func NewSiteLister(indexer cache.Indexer) SiteLister {
	return &siteLister{indexer: indexer}
}

// List lists all Sites in the indexer.
func (s *siteLister) List(selector labels.Selector) (ret []*v1alpha1.Site, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Site))
	})
	return ret, err
}

// Sites returns an object that can list and get Sites.
func (s *siteLister) Sites(namespace string) SiteNamespaceLister {
	return siteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SiteNamespaceLister helps list and get Sites.
// All objects returned here must be treated as read-only.
type SiteNamespaceLister interface {
	// List lists all Sites in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Site, err error)
	// Get retrieves the Site from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Site, error)
	SiteNamespaceListerExpansion
}

// siteNamespaceLister implements the SiteNamespaceLister
// interface.
type siteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Sites in the indexer for a given namespace.
func (s siteNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Site, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Site))
	})
	return ret, err
}

// Get retrieves the Site from the indexer for a given namespace and name.
func (s siteNamespaceLister) Get(name string) (*v1alpha1.Site, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("site"), name)
	}
	return obj.(*v1alpha1.Site), nil
}
