/*
Copyright The KubeDB Authors.

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

package v1alpha1

import (
	"kubedb.dev/apimachinery/api/crds"
	"kubedb.dev/apimachinery/apis"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

func (_ ProxySQLVersion) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourcePluralProxySQLVersion))
}

var _ apis.ResourceInfo = &ProxySQLVersion{}

func (p ProxySQLVersion) ResourceShortCode() string {
	return ""
}

func (p ProxySQLVersion) ResourceKind() string {
	return ResourceKindProxySQLVersion
}

func (p ProxySQLVersion) ResourceSingular() string {
	return ResourceSingularProxySQLVersion
}

func (p ProxySQLVersion) ResourcePlural() string {
	return ResourcePluralProxySQLVersion
}
