/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "k8s.io/metacontroller/apis/metacontroller/v1alpha1"
)

// CompositeControllerLister helps list CompositeControllers.
type CompositeControllerLister interface {
	// List lists all CompositeControllers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.CompositeController, err error)
	// CompositeControllers returns an object that can list and get CompositeControllers.
	CompositeControllers(namespace string) CompositeControllerNamespaceLister
	CompositeControllerListerExpansion
}

// compositeControllerLister implements the CompositeControllerLister interface.
type compositeControllerLister struct {
	indexer cache.Indexer
}

// NewCompositeControllerLister returns a new CompositeControllerLister.
func NewCompositeControllerLister(indexer cache.Indexer) CompositeControllerLister {
	return &compositeControllerLister{indexer: indexer}
}

// List lists all CompositeControllers in the indexer.
func (s *compositeControllerLister) List(selector labels.Selector) (ret []*v1alpha1.CompositeController, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CompositeController))
	})
	return ret, err
}

// CompositeControllers returns an object that can list and get CompositeControllers.
func (s *compositeControllerLister) CompositeControllers(namespace string) CompositeControllerNamespaceLister {
	return compositeControllerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CompositeControllerNamespaceLister helps list and get CompositeControllers.
type CompositeControllerNamespaceLister interface {
	// List lists all CompositeControllers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.CompositeController, err error)
	// Get retrieves the CompositeController from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.CompositeController, error)
	CompositeControllerNamespaceListerExpansion
}

// compositeControllerNamespaceLister implements the CompositeControllerNamespaceLister
// interface.
type compositeControllerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CompositeControllers in the indexer for a given namespace.
func (s compositeControllerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CompositeController, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CompositeController))
	})
	return ret, err
}

// Get retrieves the CompositeController from the indexer for a given namespace and name.
func (s compositeControllerNamespaceLister) Get(name string) (*v1alpha1.CompositeController, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("compositecontroller"), name)
	}
	return obj.(*v1alpha1.CompositeController), nil
}
