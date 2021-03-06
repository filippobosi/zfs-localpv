/*
Copyright 2019 The OpenEBS Authors

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

package v1alpha1

import (
	v1alpha1 "github.com/openebs/zfs-localpv/pkg/apis/openebs.io/zfs/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ZFSSnapshotLister helps list ZFSSnapshots.
type ZFSSnapshotLister interface {
	// List lists all ZFSSnapshots in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ZFSSnapshot, err error)
	// ZFSSnapshots returns an object that can list and get ZFSSnapshots.
	ZFSSnapshots(namespace string) ZFSSnapshotNamespaceLister
	ZFSSnapshotListerExpansion
}

// zFSSnapshotLister implements the ZFSSnapshotLister interface.
type zFSSnapshotLister struct {
	indexer cache.Indexer
}

// NewZFSSnapshotLister returns a new ZFSSnapshotLister.
func NewZFSSnapshotLister(indexer cache.Indexer) ZFSSnapshotLister {
	return &zFSSnapshotLister{indexer: indexer}
}

// List lists all ZFSSnapshots in the indexer.
func (s *zFSSnapshotLister) List(selector labels.Selector) (ret []*v1alpha1.ZFSSnapshot, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ZFSSnapshot))
	})
	return ret, err
}

// ZFSSnapshots returns an object that can list and get ZFSSnapshots.
func (s *zFSSnapshotLister) ZFSSnapshots(namespace string) ZFSSnapshotNamespaceLister {
	return zFSSnapshotNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ZFSSnapshotNamespaceLister helps list and get ZFSSnapshots.
type ZFSSnapshotNamespaceLister interface {
	// List lists all ZFSSnapshots in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ZFSSnapshot, err error)
	// Get retrieves the ZFSSnapshot from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ZFSSnapshot, error)
	ZFSSnapshotNamespaceListerExpansion
}

// zFSSnapshotNamespaceLister implements the ZFSSnapshotNamespaceLister
// interface.
type zFSSnapshotNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ZFSSnapshots in the indexer for a given namespace.
func (s zFSSnapshotNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ZFSSnapshot, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ZFSSnapshot))
	})
	return ret, err
}

// Get retrieves the ZFSSnapshot from the indexer for a given namespace and name.
func (s zFSSnapshotNamespaceLister) Get(name string) (*v1alpha1.ZFSSnapshot, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("zfssnapshot"), name)
	}
	return obj.(*v1alpha1.ZFSSnapshot), nil
}
