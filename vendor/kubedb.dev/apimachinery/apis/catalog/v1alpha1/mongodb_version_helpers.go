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
	"fmt"

	"kubedb.dev/apimachinery/api/crds"
	"kubedb.dev/apimachinery/apis"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

func (_ MongoDBVersion) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourcePluralMongoDBVersion))
}

var _ apis.ResourceInfo = &MongoDBVersion{}

func (m MongoDBVersion) ResourceShortCode() string {
	return ResourceCodeMongoDBVersion
}

func (m MongoDBVersion) ResourceKind() string {
	return ResourceKindMongoDBVersion
}

func (m MongoDBVersion) ResourceSingular() string {
	return ResourceSingularMongoDBVersion
}

func (m MongoDBVersion) ResourcePlural() string {
	return ResourcePluralMongoDBVersion
}

func (m MongoDBVersion) ValidateSpecs() error {
	if m.Spec.Version == "" ||
		m.Spec.DB.Image == "" ||
		m.Spec.Tools.Image == "" ||
		m.Spec.Exporter.Image == "" ||
		m.Spec.InitContainer.Image == "" {
		return fmt.Errorf(`atleast one of the following specs is not set for mongodbVersion "%v":
spec.version,
spec.db.image,
spec.tools.image,
spec.exporter.image,
spec.initContainer.image.`, m.Name)
	}
	return nil
}
