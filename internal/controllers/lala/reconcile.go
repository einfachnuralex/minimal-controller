package lala

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	lalav1alpha1 "github.com/einfachnuralex/minimal-controller/pkg/apis/lala/v1alpha1"
)

type Reconciler struct {
	client client.Client
	Name   string
}

// Reconcile implements reconcile.TypedReconciler.
func (r *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	logr := logf.FromContext(ctx)
	logr.Info("recieved")

	entry := &lalav1alpha1.Lala{}
	if err := r.client.Get(ctx, req.NamespacedName, entry); err != nil {
		return reconcile.Result{}, client.IgnoreNotFound(err)
	}

	// Set the instance ID if it's not already set
	if entry.Status.InstanceID == "" {
		entry.Status.InstanceID = r.Name

		if err := r.patchStatus(ctx, entry); err != nil {
			return reconcile.Result{}, err
		}
		logr.Info("patched status")
	}

	return reconcile.Result{}, nil
}

func (r *Reconciler) patchStatus(ctx context.Context, entry *lalav1alpha1.Lala) error {
	key := client.ObjectKeyFromObject(entry)
	latest := &lalav1alpha1.Lala{}
	if err := r.client.Get(ctx, key, latest); err != nil {
		return err
	}
	return r.client.Status().Patch(ctx, entry, client.MergeFrom(latest))
}
