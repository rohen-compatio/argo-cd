// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplications implements ApplicationInterface
type FakeApplications struct {
	Fake *FakeArgoprojV1alpha1
	ns   string
}

var applicationsResource = v1alpha1.SchemeGroupVersion.WithResource("applications")

var applicationsKind = v1alpha1.SchemeGroupVersion.WithKind("Application")

// Get takes name of the application, and returns the corresponding application object, and an error if there is any.
func (c *FakeApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Application, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationsResource, c.ns, name), &v1alpha1.Application{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Application), err
}

// List takes label and field selectors, and returns the list of Applications that match those selectors.
func (c *FakeApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationsResource, applicationsKind, c.ns, opts), &v1alpha1.ApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApplicationList{ListMeta: obj.(*v1alpha1.ApplicationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applications.
func (c *FakeApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationsResource, c.ns, opts))

}

// Create takes the representation of a application and creates it.  Returns the server's representation of the application, and an error, if there is any.
func (c *FakeApplications) Create(ctx context.Context, application *v1alpha1.Application, opts v1.CreateOptions) (result *v1alpha1.Application, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationsResource, c.ns, application), &v1alpha1.Application{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Application), err
}

// Update takes the representation of a application and updates it. Returns the server's representation of the application, and an error, if there is any.
func (c *FakeApplications) Update(ctx context.Context, application *v1alpha1.Application, opts v1.UpdateOptions) (result *v1alpha1.Application, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationsResource, c.ns, application), &v1alpha1.Application{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Application), err
}

// Delete takes name of the application and deletes it. Returns an error if one occurs.
func (c *FakeApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(applicationsResource, c.ns, name, opts), &v1alpha1.Application{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApplicationList{})
	return err
}

// Patch applies the patch and returns the patched application.
func (c *FakeApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Application, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Application{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Application), err
}
