/*
Copyright The KubeVault Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	catalogv1alpha1 "kubevault.dev/operator/apis/catalog/v1alpha1"
	versioned "kubevault.dev/operator/client/clientset/versioned"
	internalinterfaces "kubevault.dev/operator/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubevault.dev/operator/client/listers/catalog/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VaultServerVersionInformer provides access to a shared informer and lister for
// VaultServerVersions.
type VaultServerVersionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.VaultServerVersionLister
}

type vaultServerVersionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewVaultServerVersionInformer constructs a new informer for VaultServerVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVaultServerVersionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVaultServerVersionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredVaultServerVersionInformer constructs a new informer for VaultServerVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVaultServerVersionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CatalogV1alpha1().VaultServerVersions().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CatalogV1alpha1().VaultServerVersions().Watch(context.TODO(), options)
			},
		},
		&catalogv1alpha1.VaultServerVersion{},
		resyncPeriod,
		indexers,
	)
}

func (f *vaultServerVersionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVaultServerVersionInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *vaultServerVersionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&catalogv1alpha1.VaultServerVersion{}, f.defaultInformer)
}

func (f *vaultServerVersionInformer) Lister() v1alpha1.VaultServerVersionLister {
	return v1alpha1.NewVaultServerVersionLister(f.Informer().GetIndexer())
}
