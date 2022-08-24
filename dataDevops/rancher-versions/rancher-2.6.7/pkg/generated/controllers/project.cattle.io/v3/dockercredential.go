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

package v3

import (
	"context"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	v3 "github.com/rancher/rancher/pkg/apis/project.cattle.io/v3"
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

type DockerCredentialHandler func(string, *v3.DockerCredential) (*v3.DockerCredential, error)

type DockerCredentialController interface {
	generic.ControllerMeta
	DockerCredentialClient

	OnChange(ctx context.Context, name string, sync DockerCredentialHandler)
	OnRemove(ctx context.Context, name string, sync DockerCredentialHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() DockerCredentialCache
}

type DockerCredentialClient interface {
	Create(*v3.DockerCredential) (*v3.DockerCredential, error)
	Update(*v3.DockerCredential) (*v3.DockerCredential, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v3.DockerCredential, error)
	List(namespace string, opts metav1.ListOptions) (*v3.DockerCredentialList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.DockerCredential, err error)
}

type DockerCredentialCache interface {
	Get(namespace, name string) (*v3.DockerCredential, error)
	List(namespace string, selector labels.Selector) ([]*v3.DockerCredential, error)

	AddIndexer(indexName string, indexer DockerCredentialIndexer)
	GetByIndex(indexName, key string) ([]*v3.DockerCredential, error)
}

type DockerCredentialIndexer func(obj *v3.DockerCredential) ([]string, error)

type dockerCredentialController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewDockerCredentialController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) DockerCredentialController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &dockerCredentialController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromDockerCredentialHandlerToHandler(sync DockerCredentialHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.DockerCredential
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.DockerCredential))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *dockerCredentialController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.DockerCredential))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateDockerCredentialDeepCopyOnChange(client DockerCredentialClient, obj *v3.DockerCredential, handler func(obj *v3.DockerCredential) (*v3.DockerCredential, error)) (*v3.DockerCredential, error) {
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

func (c *dockerCredentialController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *dockerCredentialController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *dockerCredentialController) OnChange(ctx context.Context, name string, sync DockerCredentialHandler) {
	c.AddGenericHandler(ctx, name, FromDockerCredentialHandlerToHandler(sync))
}

func (c *dockerCredentialController) OnRemove(ctx context.Context, name string, sync DockerCredentialHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromDockerCredentialHandlerToHandler(sync)))
}

func (c *dockerCredentialController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *dockerCredentialController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *dockerCredentialController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *dockerCredentialController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *dockerCredentialController) Cache() DockerCredentialCache {
	return &dockerCredentialCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *dockerCredentialController) Create(obj *v3.DockerCredential) (*v3.DockerCredential, error) {
	result := &v3.DockerCredential{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *dockerCredentialController) Update(obj *v3.DockerCredential) (*v3.DockerCredential, error) {
	result := &v3.DockerCredential{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *dockerCredentialController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *dockerCredentialController) Get(namespace, name string, options metav1.GetOptions) (*v3.DockerCredential, error) {
	result := &v3.DockerCredential{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *dockerCredentialController) List(namespace string, opts metav1.ListOptions) (*v3.DockerCredentialList, error) {
	result := &v3.DockerCredentialList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *dockerCredentialController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *dockerCredentialController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v3.DockerCredential, error) {
	result := &v3.DockerCredential{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type dockerCredentialCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *dockerCredentialCache) Get(namespace, name string) (*v3.DockerCredential, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.DockerCredential), nil
}

func (c *dockerCredentialCache) List(namespace string, selector labels.Selector) (ret []*v3.DockerCredential, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.DockerCredential))
	})

	return ret, err
}

func (c *dockerCredentialCache) AddIndexer(indexName string, indexer DockerCredentialIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.DockerCredential))
		},
	}))
}

func (c *dockerCredentialCache) GetByIndex(indexName, key string) (result []*v3.DockerCredential, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.DockerCredential, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.DockerCredential))
	}
	return result, nil
}
