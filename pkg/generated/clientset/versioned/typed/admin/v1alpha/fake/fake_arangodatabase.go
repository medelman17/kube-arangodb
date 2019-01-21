//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha "github.com/arangodb/kube-arangodb/pkg/apis/admin/v1alpha"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeArangoDatabases implements ArangoDatabaseInterface
type FakeArangoDatabases struct {
	Fake *FakeDatabaseadminV1alpha
	ns   string
}

var arangodatabasesResource = schema.GroupVersionResource{Group: "databaseadmin.arangodb.com", Version: "v1alpha", Resource: "arangodatabases"}

var arangodatabasesKind = schema.GroupVersionKind{Group: "databaseadmin.arangodb.com", Version: "v1alpha", Kind: "ArangoDatabase"}

// Get takes name of the arangoDatabase, and returns the corresponding arangoDatabase object, and an error if there is any.
func (c *FakeArangoDatabases) Get(name string, options v1.GetOptions) (result *v1alpha.ArangoDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(arangodatabasesResource, c.ns, name), &v1alpha.ArangoDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoDatabase), err
}

// List takes label and field selectors, and returns the list of ArangoDatabases that match those selectors.
func (c *FakeArangoDatabases) List(opts v1.ListOptions) (result *v1alpha.ArangoDatabaseList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(arangodatabasesResource, arangodatabasesKind, c.ns, opts), &v1alpha.ArangoDatabaseList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha.ArangoDatabaseList{ListMeta: obj.(*v1alpha.ArangoDatabaseList).ListMeta}
	for _, item := range obj.(*v1alpha.ArangoDatabaseList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested arangoDatabases.
func (c *FakeArangoDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(arangodatabasesResource, c.ns, opts))

}

// Create takes the representation of a arangoDatabase and creates it.  Returns the server's representation of the arangoDatabase, and an error, if there is any.
func (c *FakeArangoDatabases) Create(arangoDatabase *v1alpha.ArangoDatabase) (result *v1alpha.ArangoDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(arangodatabasesResource, c.ns, arangoDatabase), &v1alpha.ArangoDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoDatabase), err
}

// Update takes the representation of a arangoDatabase and updates it. Returns the server's representation of the arangoDatabase, and an error, if there is any.
func (c *FakeArangoDatabases) Update(arangoDatabase *v1alpha.ArangoDatabase) (result *v1alpha.ArangoDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(arangodatabasesResource, c.ns, arangoDatabase), &v1alpha.ArangoDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoDatabase), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeArangoDatabases) UpdateStatus(arangoDatabase *v1alpha.ArangoDatabase) (*v1alpha.ArangoDatabase, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(arangodatabasesResource, "status", c.ns, arangoDatabase), &v1alpha.ArangoDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoDatabase), err
}

// Delete takes name of the arangoDatabase and deletes it. Returns an error if one occurs.
func (c *FakeArangoDatabases) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(arangodatabasesResource, c.ns, name), &v1alpha.ArangoDatabase{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeArangoDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(arangodatabasesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha.ArangoDatabaseList{})
	return err
}

// Patch applies the patch and returns the patched arangoDatabase.
func (c *FakeArangoDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha.ArangoDatabase, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(arangodatabasesResource, c.ns, name, data, subresources...), &v1alpha.ArangoDatabase{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha.ArangoDatabase), err
}