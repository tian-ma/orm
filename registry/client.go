/*
Copyright 2022.

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

package registry

import (
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
)

type Client struct {
	dynamic.Interface
}

func (c *Client) GetResourceWithGVK(gvk schema.GroupVersionKind, req types.NamespacedName) (*unstructured.Unstructured, error) {

	var err error
	gvr := r.findGVRfromGVK(gvk)
	if gvr == nil {
		return nil, errors.New("Operator " + gvk.String() + "is not installed")
	}

	obj := &unstructured.Unstructured{}
	obj, err = r.Resource(*gvr).Namespace(req.Namespace).Get(r.ctx, req.Name, metav1.GetOptions{})

	return obj, err
}

func (c *Client) UpdateResourceWithGVK(gvk schema.GroupVersionKind, obj *unstructured.Unstructured) error {
	var err error

	gvr := r.findGVRfromGVK(gvk)
	if gvr == nil {
		return errors.New("Operator " + gvk.String() + "is not installed")
	}

	_, err = c.Resource(*gvr).Namespace(obj.GetNamespace()).Update(r.ctx, obj, metav1.UpdateOptions{})

	return err

}
