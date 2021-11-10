// Copyright 2021 VMware
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package templates

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"fmt"
	v1alpha12 "github.com/vmware-tanzu/cartographer/pkg/apis/carto/v1alpha1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type clusterImageTemplate struct {
	template  *v1alpha12.ClusterImageTemplate
	evaluator evaluator
}

func (t clusterImageTemplate) GetKind() string {
	return t.template.Kind
}

func NewClusterImageTemplateModel(template *v1alpha12.ClusterImageTemplate, eval evaluator) *clusterImageTemplate {
	return &clusterImageTemplate{template: template, evaluator: eval}
}

func (t clusterImageTemplate) GetName() string {
	return t.template.Name
}

func (t clusterImageTemplate) GetOutput(stampedObject *unstructured.Unstructured) (*Output, error) {
	image, err := t.evaluator.EvaluateJsonPath(t.template.Spec.ImagePath, stampedObject.UnstructuredContent())
	if err != nil {
		return nil, &JsonPathError{
			Err:        fmt.Errorf("evaluate image json path: %w", err),
			expression: t.template.Spec.ImagePath,
		}
	}

	return &Output{
		Image: image,
	}, nil
}

func (t clusterImageTemplate) GetResourceTemplate() v1alpha12.TemplateSpec {
	return t.template.Spec.TemplateSpec
}

func (t clusterImageTemplate) GetDefaultParams() v1alpha12.DefaultParams {
	return t.template.Spec.Params
}
