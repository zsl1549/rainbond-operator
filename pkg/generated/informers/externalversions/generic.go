// RAINBOND, Application Management Platform
// Copyright (C) 2014-2020 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/goodrain/rainbond-operator/pkg/apis/rainbond/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=rainbond.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("rainbondclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rainbond().V1alpha1().RainbondClusters().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("rainbondpackages"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rainbond().V1alpha1().RainbondPackages().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("rainbondvolumes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rainbond().V1alpha1().RainbondVolumes().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("rbdcomponents"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Rainbond().V1alpha1().RbdComponents().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
