/*
Copyright 2023 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type BundleDeploymentHandler func(string, *v1alpha1.BundleDeployment) (*v1alpha1.BundleDeployment, error)

type BundleDeploymentController interface {
	generic.ControllerMeta
	BundleDeploymentClient

	OnChange(ctx context.Context, name string, sync BundleDeploymentHandler)
	OnRemove(ctx context.Context, name string, sync BundleDeploymentHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() BundleDeploymentCache
}

type BundleDeploymentClient interface {
	Create(*v1alpha1.BundleDeployment) (*v1alpha1.BundleDeployment, error)
	Update(*v1alpha1.BundleDeployment) (*v1alpha1.BundleDeployment, error)
	UpdateStatus(*v1alpha1.BundleDeployment) (*v1alpha1.BundleDeployment, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1alpha1.BundleDeployment, error)
	List(namespace string, opts metav1.ListOptions) (*v1alpha1.BundleDeploymentList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BundleDeployment, err error)
}

type BundleDeploymentCache interface {
	Get(namespace, name string) (*v1alpha1.BundleDeployment, error)
	List(namespace string, selector labels.Selector) ([]*v1alpha1.BundleDeployment, error)

	AddIndexer(indexName string, indexer BundleDeploymentIndexer)
	GetByIndex(indexName, key string) ([]*v1alpha1.BundleDeployment, error)
}

type BundleDeploymentIndexer func(obj *v1alpha1.BundleDeployment) ([]string, error)

type bundleDeploymentController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewBundleDeploymentController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) BundleDeploymentController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &bundleDeploymentController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromBundleDeploymentHandlerToHandler(sync BundleDeploymentHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1alpha1.BundleDeployment
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1alpha1.BundleDeployment))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *bundleDeploymentController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1alpha1.BundleDeployment))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateBundleDeploymentDeepCopyOnChange(client BundleDeploymentClient, obj *v1alpha1.BundleDeployment, handler func(obj *v1alpha1.BundleDeployment) (*v1alpha1.BundleDeployment, error)) (*v1alpha1.BundleDeployment, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *bundleDeploymentController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *bundleDeploymentController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *bundleDeploymentController) OnChange(ctx context.Context, name string, sync BundleDeploymentHandler) {
	c.AddGenericHandler(ctx, name, FromBundleDeploymentHandlerToHandler(sync))
}

func (c *bundleDeploymentController) OnRemove(ctx context.Context, name string, sync BundleDeploymentHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromBundleDeploymentHandlerToHandler(sync)))
}

func (c *bundleDeploymentController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *bundleDeploymentController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *bundleDeploymentController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *bundleDeploymentController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *bundleDeploymentController) Cache() BundleDeploymentCache {
	return &bundleDeploymentCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *bundleDeploymentController) Create(obj *v1alpha1.BundleDeployment) (*v1alpha1.BundleDeployment, error) {
	result := &v1alpha1.BundleDeployment{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *bundleDeploymentController) Update(obj *v1alpha1.BundleDeployment) (*v1alpha1.BundleDeployment, error) {
	result := &v1alpha1.BundleDeployment{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *bundleDeploymentController) UpdateStatus(obj *v1alpha1.BundleDeployment) (*v1alpha1.BundleDeployment, error) {
	result := &v1alpha1.BundleDeployment{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *bundleDeploymentController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *bundleDeploymentController) Get(namespace, name string, options metav1.GetOptions) (*v1alpha1.BundleDeployment, error) {
	result := &v1alpha1.BundleDeployment{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *bundleDeploymentController) List(namespace string, opts metav1.ListOptions) (*v1alpha1.BundleDeploymentList, error) {
	result := &v1alpha1.BundleDeploymentList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *bundleDeploymentController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *bundleDeploymentController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1alpha1.BundleDeployment, error) {
	result := &v1alpha1.BundleDeployment{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type bundleDeploymentCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *bundleDeploymentCache) Get(namespace, name string) (*v1alpha1.BundleDeployment, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1alpha1.BundleDeployment), nil
}

func (c *bundleDeploymentCache) List(namespace string, selector labels.Selector) (ret []*v1alpha1.BundleDeployment, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BundleDeployment))
	})

	return ret, err
}

func (c *bundleDeploymentCache) AddIndexer(indexName string, indexer BundleDeploymentIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1alpha1.BundleDeployment))
		},
	}))
}

func (c *bundleDeploymentCache) GetByIndex(indexName, key string) (result []*v1alpha1.BundleDeployment, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1alpha1.BundleDeployment, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1alpha1.BundleDeployment))
	}
	return result, nil
}

type BundleDeploymentStatusHandler func(obj *v1alpha1.BundleDeployment, status v1alpha1.BundleDeploymentStatus) (v1alpha1.BundleDeploymentStatus, error)

type BundleDeploymentGeneratingHandler func(obj *v1alpha1.BundleDeployment, status v1alpha1.BundleDeploymentStatus) ([]runtime.Object, v1alpha1.BundleDeploymentStatus, error)

func RegisterBundleDeploymentStatusHandler(ctx context.Context, controller BundleDeploymentController, condition condition.Cond, name string, handler BundleDeploymentStatusHandler) {
	statusHandler := &bundleDeploymentStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromBundleDeploymentHandlerToHandler(statusHandler.sync))
}

func RegisterBundleDeploymentGeneratingHandler(ctx context.Context, controller BundleDeploymentController, apply apply.Apply,
	condition condition.Cond, name string, handler BundleDeploymentGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &bundleDeploymentGeneratingHandler{
		BundleDeploymentGeneratingHandler: handler,
		apply:                             apply,
		name:                              name,
		gvk:                               controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterBundleDeploymentStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type bundleDeploymentStatusHandler struct {
	client    BundleDeploymentClient
	condition condition.Cond
	handler   BundleDeploymentStatusHandler
}

func (a *bundleDeploymentStatusHandler) sync(key string, obj *v1alpha1.BundleDeployment) (*v1alpha1.BundleDeployment, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type bundleDeploymentGeneratingHandler struct {
	BundleDeploymentGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *bundleDeploymentGeneratingHandler) Remove(key string, obj *v1alpha1.BundleDeployment) (*v1alpha1.BundleDeployment, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1alpha1.BundleDeployment{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *bundleDeploymentGeneratingHandler) Handle(obj *v1alpha1.BundleDeployment, status v1alpha1.BundleDeploymentStatus) (v1alpha1.BundleDeploymentStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.BundleDeploymentGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
