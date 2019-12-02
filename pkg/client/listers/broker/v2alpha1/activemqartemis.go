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

// Code generated by lister-gen. DO NOT EDIT.

package v2alpha1

import (
	v2alpha1 "github.com/rh-messaging/activemq-artemis-operator/pkg/apis/broker/v2alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ActiveMQArtemisLister helps list ActiveMQArtemises.
type ActiveMQArtemisLister interface {
	// List lists all ActiveMQArtemises in the indexer.
	List(selector labels.Selector) (ret []*v2alpha1.ActiveMQArtemis, err error)
	// ActiveMQArtemises returns an object that can list and get ActiveMQArtemises.
	ActiveMQArtemises(namespace string) ActiveMQArtemisNamespaceLister
	ActiveMQArtemisListerExpansion
}

// activeMQArtemisLister implements the ActiveMQArtemisLister interface.
type activeMQArtemisLister struct {
	indexer cache.Indexer
}

// NewActiveMQArtemisLister returns a new ActiveMQArtemisLister.
func NewActiveMQArtemisLister(indexer cache.Indexer) ActiveMQArtemisLister {
	return &activeMQArtemisLister{indexer: indexer}
}

// List lists all ActiveMQArtemises in the indexer.
func (s *activeMQArtemisLister) List(selector labels.Selector) (ret []*v2alpha1.ActiveMQArtemis, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.ActiveMQArtemis))
	})
	return ret, err
}

// ActiveMQArtemises returns an object that can list and get ActiveMQArtemises.
func (s *activeMQArtemisLister) ActiveMQArtemises(namespace string) ActiveMQArtemisNamespaceLister {
	return activeMQArtemisNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ActiveMQArtemisNamespaceLister helps list and get ActiveMQArtemises.
type ActiveMQArtemisNamespaceLister interface {
	// List lists all ActiveMQArtemises in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v2alpha1.ActiveMQArtemis, err error)
	// Get retrieves the ActiveMQArtemis from the indexer for a given namespace and name.
	Get(name string) (*v2alpha1.ActiveMQArtemis, error)
	ActiveMQArtemisNamespaceListerExpansion
}

// activeMQArtemisNamespaceLister implements the ActiveMQArtemisNamespaceLister
// interface.
type activeMQArtemisNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ActiveMQArtemises in the indexer for a given namespace.
func (s activeMQArtemisNamespaceLister) List(selector labels.Selector) (ret []*v2alpha1.ActiveMQArtemis, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.ActiveMQArtemis))
	})
	return ret, err
}

// Get retrieves the ActiveMQArtemis from the indexer for a given namespace and name.
func (s activeMQArtemisNamespaceLister) Get(name string) (*v2alpha1.ActiveMQArtemis, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2alpha1.Resource("activemqartemis"), name)
	}
	return obj.(*v2alpha1.ActiveMQArtemis), nil
}