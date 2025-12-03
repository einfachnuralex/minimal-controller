package lala

import (
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	lalav1alpha1 "github.com/einfachnuralex/minimal-controller/pkg/apis/lala/v1alpha1"
)

func (r *Reconciler) AddToManager(mgr manager.Manager) error {
	if r.client == nil {
		r.client = mgr.GetClient()
	}
	err := builder.
		ControllerManagedBy(mgr).
		For(&lalav1alpha1.Lala{}).
		Named(r.Name).
		Complete(r)
	if err != nil {
		return fmt.Errorf("could not create controller for lala: %w", err)
	}
	return nil
}
