// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kfsoftware/hlf-operator/api/hlf.kungfusoftware.es/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FabricOrdererNodeLister helps list FabricOrdererNodes.
// All objects returned here must be treated as read-only.
type FabricOrdererNodeLister interface {
	// List lists all FabricOrdererNodes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricOrdererNode, err error)
	// FabricOrdererNodes returns an object that can list and get FabricOrdererNodes.
	FabricOrdererNodes(namespace string) FabricOrdererNodeNamespaceLister
	FabricOrdererNodeListerExpansion
}

// fabricOrdererNodeLister implements the FabricOrdererNodeLister interface.
type fabricOrdererNodeLister struct {
	indexer cache.Indexer
}

// NewFabricOrdererNodeLister returns a new FabricOrdererNodeLister.
func NewFabricOrdererNodeLister(indexer cache.Indexer) FabricOrdererNodeLister {
	return &fabricOrdererNodeLister{indexer: indexer}
}

// List lists all FabricOrdererNodes in the indexer.
func (s *fabricOrdererNodeLister) List(selector labels.Selector) (ret []*v1alpha1.FabricOrdererNode, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricOrdererNode))
	})
	return ret, err
}

// FabricOrdererNodes returns an object that can list and get FabricOrdererNodes.
func (s *fabricOrdererNodeLister) FabricOrdererNodes(namespace string) FabricOrdererNodeNamespaceLister {
	return fabricOrdererNodeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FabricOrdererNodeNamespaceLister helps list and get FabricOrdererNodes.
// All objects returned here must be treated as read-only.
type FabricOrdererNodeNamespaceLister interface {
	// List lists all FabricOrdererNodes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FabricOrdererNode, err error)
	// Get retrieves the FabricOrdererNode from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FabricOrdererNode, error)
	FabricOrdererNodeNamespaceListerExpansion
}

// fabricOrdererNodeNamespaceLister implements the FabricOrdererNodeNamespaceLister
// interface.
type fabricOrdererNodeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FabricOrdererNodes in the indexer for a given namespace.
func (s fabricOrdererNodeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FabricOrdererNode, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FabricOrdererNode))
	})
	return ret, err
}

// Get retrieves the FabricOrdererNode from the indexer for a given namespace and name.
func (s fabricOrdererNodeNamespaceLister) Get(name string) (*v1alpha1.FabricOrdererNode, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fabricorderernode"), name)
	}
	return obj.(*v1alpha1.FabricOrdererNode), nil
}
