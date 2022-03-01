//go:build generate
// +build generate

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

// NOTE(negz): See the below link for details on what is happening here.
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

// Remove existing CRDs
//go:generate rm -rf ../cluster/crds

// Generate deepcopy methodsets and CRD manifests
//go:generate go run -tags generate sigs.k8s.io/controller-tools/cmd/controller-gen object:headerFile=../hack/boilerplate.go.txt paths=./pkg/v1alpha1;./pkg/v1beta1;./pkg/v1;./apiextensions/...;./secrets/... crd:crdVersions=v1 output:artifacts:config=../cluster/crds

// NOTE(hasheddan): we generate the meta.pkg.crossplane.io types separately as
// the generated CRDs are never installed, only used for API documentation.
//go:generate go run -tags generate sigs.k8s.io/controller-tools/cmd/controller-gen object:headerFile=../hack/boilerplate.go.txt paths=./pkg/meta/... crd:crdVersions=v1 output:artifacts:config=../docs/api-docs/crds

// Generate webhook manifests
//go:generate go run -tags generate sigs.k8s.io/controller-tools/cmd/controller-gen webhook paths=./pkg/v1alpha1;./pkg/v1beta1;./pkg/v1;./apiextensions/... output:artifacts:config=../cluster/webhookconfigurations

// Generate clientset for types.
//go:generate rm -rf ../internal/client
//go:generate go run -tags generate k8s.io/code-generator/cmd/client-gen --clientset-name "versioned" --build-tag="ignore_autogenerated" --go-header-file "../hack/boilerplate.go.txt" --output-package "github.com/crossplane/crossplane/internal/client/clientset" --input-base "github.com/crossplane/crossplane/apis" --output-base "../tmp-clientgen" --input "pkg/v1alpha1,pkg/v1beta1,pkg/v1,apiextensions/v1,secrets/v1alpha1"
//go:generate cp -r ../tmp-clientgen/github.com/crossplane/crossplane/internal/client ../internal/client
//go:generate rm -rf ../tmp-clientgen

package apis

import (
	_ "k8s.io/code-generator"                           //nolint:typecheck
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen" //nolint:typecheck
)
