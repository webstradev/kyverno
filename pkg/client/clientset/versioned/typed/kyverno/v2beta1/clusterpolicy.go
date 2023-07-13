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

// Code generated by client-gen. DO NOT EDIT.

package v2beta1

import (
	"context"
	"time"

	v2beta1 "github.com/kyverno/kyverno/api/kyverno/v2beta1"
	scheme "github.com/kyverno/kyverno/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterPoliciesGetter has a method to return a ClusterPolicyInterface.
// A group's client should implement this interface.
type ClusterPoliciesGetter interface {
	ClusterPolicies() ClusterPolicyInterface
}

// ClusterPolicyInterface has methods to work with ClusterPolicy resources.
type ClusterPolicyInterface interface {
	Create(ctx context.Context, clusterPolicy *v2beta1.ClusterPolicy, opts v1.CreateOptions) (*v2beta1.ClusterPolicy, error)
	Update(ctx context.Context, clusterPolicy *v2beta1.ClusterPolicy, opts v1.UpdateOptions) (*v2beta1.ClusterPolicy, error)
	UpdateStatus(ctx context.Context, clusterPolicy *v2beta1.ClusterPolicy, opts v1.UpdateOptions) (*v2beta1.ClusterPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2beta1.ClusterPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2beta1.ClusterPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta1.ClusterPolicy, err error)
	ClusterPolicyExpansion
}

// clusterPolicies implements ClusterPolicyInterface
type clusterPolicies struct {
	client rest.Interface
}

// newClusterPolicies returns a ClusterPolicies
func newClusterPolicies(c *KyvernoV2beta1Client) *clusterPolicies {
	return &clusterPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterPolicy, and returns the corresponding clusterPolicy object, and an error if there is any.
func (c *clusterPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2beta1.ClusterPolicy, err error) {
	result = &v2beta1.ClusterPolicy{}
	err = c.client.Get().
		Resource("clusterpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterPolicies that match those selectors.
func (c *clusterPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v2beta1.ClusterPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2beta1.ClusterPolicyList{}
	err = c.client.Get().
		Resource("clusterpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterPolicies.
func (c *clusterPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterPolicy and creates it.  Returns the server's representation of the clusterPolicy, and an error, if there is any.
func (c *clusterPolicies) Create(ctx context.Context, clusterPolicy *v2beta1.ClusterPolicy, opts v1.CreateOptions) (result *v2beta1.ClusterPolicy, err error) {
	result = &v2beta1.ClusterPolicy{}
	err = c.client.Post().
		Resource("clusterpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterPolicy and updates it. Returns the server's representation of the clusterPolicy, and an error, if there is any.
func (c *clusterPolicies) Update(ctx context.Context, clusterPolicy *v2beta1.ClusterPolicy, opts v1.UpdateOptions) (result *v2beta1.ClusterPolicy, err error) {
	result = &v2beta1.ClusterPolicy{}
	err = c.client.Put().
		Resource("clusterpolicies").
		Name(clusterPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterPolicies) UpdateStatus(ctx context.Context, clusterPolicy *v2beta1.ClusterPolicy, opts v1.UpdateOptions) (result *v2beta1.ClusterPolicy, err error) {
	result = &v2beta1.ClusterPolicy{}
	err = c.client.Put().
		Resource("clusterpolicies").
		Name(clusterPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterPolicy and deletes it. Returns an error if one occurs.
func (c *clusterPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterpolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterpolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterPolicy.
func (c *clusterPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta1.ClusterPolicy, err error) {
	result = &v2beta1.ClusterPolicy{}
	err = c.client.Patch(pt).
		Resource("clusterpolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}