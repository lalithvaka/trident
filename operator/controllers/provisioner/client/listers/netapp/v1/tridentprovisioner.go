// Copyright 2020 NetApp, Inc. All Rights Reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/netapp/trident/operator/controllers/provisioner/apis/netapp/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TridentProvisionerLister helps list TridentProvisioners.
type TridentProvisionerLister interface {
	// List lists all TridentProvisioners in the indexer.
	List(selector labels.Selector) (ret []*v1.TridentProvisioner, err error)
	// TridentProvisioners returns an object that can list and get TridentProvisioners.
	TridentProvisioners(namespace string) TridentProvisionerNamespaceLister
	TridentProvisionerListerExpansion
}

// tridentProvisionerLister implements the TridentProvisionerLister interface.
type tridentProvisionerLister struct {
	indexer cache.Indexer
}

// NewTridentProvisionerLister returns a new TridentProvisionerLister.
func NewTridentProvisionerLister(indexer cache.Indexer) TridentProvisionerLister {
	return &tridentProvisionerLister{indexer: indexer}
}

// List lists all TridentProvisioners in the indexer.
func (s *tridentProvisionerLister) List(selector labels.Selector) (ret []*v1.TridentProvisioner, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentProvisioner))
	})
	return ret, err
}

// TridentProvisioners returns an object that can list and get TridentProvisioners.
func (s *tridentProvisionerLister) TridentProvisioners(namespace string) TridentProvisionerNamespaceLister {
	return tridentProvisionerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TridentProvisionerNamespaceLister helps list and get TridentProvisioners.
type TridentProvisionerNamespaceLister interface {
	// List lists all TridentProvisioners in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.TridentProvisioner, err error)
	// Get retrieves the TridentProvisioner from the indexer for a given namespace and name.
	Get(name string) (*v1.TridentProvisioner, error)
	TridentProvisionerNamespaceListerExpansion
}

// tridentProvisionerNamespaceLister implements the TridentProvisionerNamespaceLister
// interface.
type tridentProvisionerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TridentProvisioners in the indexer for a given namespace.
func (s tridentProvisionerNamespaceLister) List(selector labels.Selector) (ret []*v1.TridentProvisioner, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TridentProvisioner))
	})
	return ret, err
}

// Get retrieves the TridentProvisioner from the indexer for a given namespace and name.
func (s tridentProvisionerNamespaceLister) Get(name string) (*v1.TridentProvisioner, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("tridentprovisioner"), name)
	}
	return obj.(*v1.TridentProvisioner), nil
}
