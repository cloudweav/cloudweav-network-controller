/*
Copyright 2022 Rancher Labs, Inc.

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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/cloudweav/cloudweav/pkg/apis/cloudweavhci.io/v1beta1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/generic"
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

type PreferenceHandler func(string, *v1beta1.Preference) (*v1beta1.Preference, error)

type PreferenceController interface {
	generic.ControllerMeta
	PreferenceClient

	OnChange(ctx context.Context, name string, sync PreferenceHandler)
	OnRemove(ctx context.Context, name string, sync PreferenceHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() PreferenceCache
}

type PreferenceClient interface {
	Create(*v1beta1.Preference) (*v1beta1.Preference, error)
	Update(*v1beta1.Preference) (*v1beta1.Preference, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1beta1.Preference, error)
	List(namespace string, opts metav1.ListOptions) (*v1beta1.PreferenceList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Preference, err error)
}

type PreferenceCache interface {
	Get(namespace, name string) (*v1beta1.Preference, error)
	List(namespace string, selector labels.Selector) ([]*v1beta1.Preference, error)

	AddIndexer(indexName string, indexer PreferenceIndexer)
	GetByIndex(indexName, key string) ([]*v1beta1.Preference, error)
}

type PreferenceIndexer func(obj *v1beta1.Preference) ([]string, error)

type preferenceController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewPreferenceController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) PreferenceController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &preferenceController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromPreferenceHandlerToHandler(sync PreferenceHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1beta1.Preference
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1beta1.Preference))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *preferenceController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1beta1.Preference))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdatePreferenceDeepCopyOnChange(client PreferenceClient, obj *v1beta1.Preference, handler func(obj *v1beta1.Preference) (*v1beta1.Preference, error)) (*v1beta1.Preference, error) {
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

func (c *preferenceController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *preferenceController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *preferenceController) OnChange(ctx context.Context, name string, sync PreferenceHandler) {
	c.AddGenericHandler(ctx, name, FromPreferenceHandlerToHandler(sync))
}

func (c *preferenceController) OnRemove(ctx context.Context, name string, sync PreferenceHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromPreferenceHandlerToHandler(sync)))
}

func (c *preferenceController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *preferenceController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *preferenceController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *preferenceController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *preferenceController) Cache() PreferenceCache {
	return &preferenceCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *preferenceController) Create(obj *v1beta1.Preference) (*v1beta1.Preference, error) {
	result := &v1beta1.Preference{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *preferenceController) Update(obj *v1beta1.Preference) (*v1beta1.Preference, error) {
	result := &v1beta1.Preference{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *preferenceController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *preferenceController) Get(namespace, name string, options metav1.GetOptions) (*v1beta1.Preference, error) {
	result := &v1beta1.Preference{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *preferenceController) List(namespace string, opts metav1.ListOptions) (*v1beta1.PreferenceList, error) {
	result := &v1beta1.PreferenceList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *preferenceController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *preferenceController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1beta1.Preference, error) {
	result := &v1beta1.Preference{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type preferenceCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *preferenceCache) Get(namespace, name string) (*v1beta1.Preference, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1beta1.Preference), nil
}

func (c *preferenceCache) List(namespace string, selector labels.Selector) (ret []*v1beta1.Preference, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Preference))
	})

	return ret, err
}

func (c *preferenceCache) AddIndexer(indexName string, indexer PreferenceIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1beta1.Preference))
		},
	}))
}

func (c *preferenceCache) GetByIndex(indexName, key string) (result []*v1beta1.Preference, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1beta1.Preference, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1beta1.Preference))
	}
	return result, nil
}
