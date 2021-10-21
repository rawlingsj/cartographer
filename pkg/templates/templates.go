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

import (
	"fmt"
	v1alpha12 "github.com/vmware-tanzu/cartographer/pkg/apis/carto/v1alpha1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/vmware-tanzu/cartographer/pkg/eval"
)

type Template interface {
	GetResourceTemplate() v1alpha12.TemplateSpec
	GetDefaultParams() v1alpha12.DefaultParams
	GetOutput(stampedObject *unstructured.Unstructured) (*Output, error)
	GetName() string
	GetKind() string
}

func NewModelFromAPI(template client.Object) (Template, error) {
	switch v := template.(type) {

	case *v1alpha12.ClusterSourceTemplate:
		return NewClusterSourceTemplateModel(v, eval.EvaluatorBuilder()), nil
	case *v1alpha12.ClusterImageTemplate:
		return NewClusterImageTemplateModel(v, eval.EvaluatorBuilder()), nil
	case *v1alpha12.ClusterConfigTemplate:
		return NewClusterConfigTemplateModel(v, eval.EvaluatorBuilder()), nil
	case *v1alpha12.ClusterDeploymentTemplate:
		return NewClusterDeploymentTemplateModel(v, eval.EvaluatorBuilder()), nil
	case *v1alpha12.ClusterTemplate:
		return NewClusterTemplateModel(v), nil
	}
	return nil, fmt.Errorf("component does not match a known template")
}
