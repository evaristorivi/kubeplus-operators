/*
Copyright 2019 The Kubernetes Authors.

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

package v1

import (
	v1 "github.com/cloud-ark/kubeplus-operators/moodle/pkg/apis/moodlecontroller/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MoodleLister helps list Moodles.
type MoodleLister interface {
	// List lists all Moodles in the indexer.
	List(selector labels.Selector) (ret []*v1.Moodle, err error)
	// Moodles returns an object that can list and get Moodles.
	Moodles(namespace string) MoodleNamespaceLister
	MoodleListerExpansion
}

// moodleLister implements the MoodleLister interface.
type moodleLister struct {
	indexer cache.Indexer
}

// NewMoodleLister returns a new MoodleLister.
func NewMoodleLister(indexer cache.Indexer) MoodleLister {
	return &moodleLister{indexer: indexer}
}

// List lists all Moodles in the indexer.
func (s *moodleLister) List(selector labels.Selector) (ret []*v1.Moodle, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Moodle))
	})
	return ret, err
}

// Moodles returns an object that can list and get Moodles.
func (s *moodleLister) Moodles(namespace string) MoodleNamespaceLister {
	return moodleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MoodleNamespaceLister helps list and get Moodles.
type MoodleNamespaceLister interface {
	// List lists all Moodles in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Moodle, err error)
	// Get retrieves the Moodle from the indexer for a given namespace and name.
	Get(name string) (*v1.Moodle, error)
	MoodleNamespaceListerExpansion
}

// moodleNamespaceLister implements the MoodleNamespaceLister
// interface.
type moodleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Moodles in the indexer for a given namespace.
func (s moodleNamespaceLister) List(selector labels.Selector) (ret []*v1.Moodle, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Moodle))
	})
	return ret, err
}

// Get retrieves the Moodle from the indexer for a given namespace and name.
func (s moodleNamespaceLister) Get(name string) (*v1.Moodle, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("moodle"), name)
	}
	return obj.(*v1.Moodle), nil
}
