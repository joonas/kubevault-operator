/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package framework

import (
	"context"

	api "kubevault.dev/apimachinery/apis/policy/v1alpha1"

	. "github.com/onsi/gomega"
	"gomodules.xyz/x/crypto/rand"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	meta_util "kmodules.xyz/client-go/meta"
)

func (f *Invocation) VaultPolicyBinding(policies, saNames, saNamespaces []string) *api.VaultPolicyBinding {
	return &api.VaultPolicyBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      rand.WithUniqSuffix("v-policy-binding"),
			Namespace: f.namespace,
			Labels: map[string]string{
				"test": f.app,
			},
		},
		Spec: api.VaultPolicyBindingSpec{
			VaultRef: core.LocalObjectReference{
				Name: f.VaultAppRef.Name,
			},
			SubjectRef: api.SubjectRef{
				Kubernetes: &api.KubernetesSubjectRef{
					ServiceAccountNames:      saNames,
					ServiceAccountNamespaces: saNamespaces,
				},
			},
			Policies: []api.PolicyIdentifier{
				{
					Ref: policies[0],
				},
			},
		},
	}
}

func (f *Framework) CreateVaultPolicyBinding(obj *api.VaultPolicyBinding) (*api.VaultPolicyBinding, error) {
	return f.CSClient.PolicyV1alpha1().VaultPolicyBindings(obj.Namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
}

func (f *Framework) GetVaultPolicyBinding(obj *api.VaultPolicyBinding) (*api.VaultPolicyBinding, error) {
	return f.CSClient.PolicyV1alpha1().VaultPolicyBindings(obj.Namespace).Get(context.TODO(), obj.Name, metav1.GetOptions{})
}

func (f *Framework) UpdateVaultPolicyBinding(obj *api.VaultPolicyBinding) (*api.VaultPolicyBinding, error) {
	return f.CSClient.PolicyV1alpha1().VaultPolicyBindings(obj.Namespace).Update(context.TODO(), obj, metav1.UpdateOptions{})
}

func (f *Framework) DeleteVaultPolicyBinding(meta metav1.ObjectMeta) error {
	return f.CSClient.PolicyV1alpha1().VaultPolicyBindings(meta.Namespace).Delete(context.TODO(), meta.Name, meta_util.DeleteInBackground())
}

func (f *Framework) EventuallyVaultPolicyBinding(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(func() *api.VaultPolicyBinding {
		obj, err := f.CSClient.PolicyV1alpha1().VaultPolicyBindings(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred())
		return obj
	})
}
