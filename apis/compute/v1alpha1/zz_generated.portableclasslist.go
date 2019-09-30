/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import resource "github.com/crossplaneio/crossplane-runtime/pkg/resource"

// GetPortableClassItems of this KubernetesClusterClassList.
func (csl *KubernetesClusterClassList) GetPortableClassItems() []resource.PortableClass {
	items := make([]resource.PortableClass, len(csl.Items))
	for i := range csl.Items {
		items[i] = resource.PortableClass(&csl.Items[i])
	}
	return items
}

// SetPortableClassItems of this KubernetesClusterClassList.
func (csl *KubernetesClusterClassList) SetPortableClassItems(i []resource.PortableClass) {
	csl.Items = make([]KubernetesClusterClass, 0, len(i))
	for j := range i {
		if actual, ok := i[j].(*KubernetesClusterClass); ok {
			csl.Items = append(csl.Items, *actual)
		}
	}
}
