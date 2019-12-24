package reconciler

import (
	"context"
	"fmt"
	// TODO: zap

	/*
		corev1 "k8s.io/api/core/v1"
		mtav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	*/
	"k8s.io/apimachinery/pkg/runtime"

	controller "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ObservationReconciler tores the information needed to reconcile Pods to
// create Observations.
type ObservationReconciler struct {
	client.Client
	scheme *runtime.Scheme
}

// Reconcile implements controller-runtime/pkg/reconcile.Reconciler
func (r *ObservationReconciler) Reconcile(req controller.Request) (controller.Result, error) {
	ctx := context.Background()

	fmt.Print(ctx)

	return controller.Result{}, nil
}
