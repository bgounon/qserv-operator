package sync

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	qservv1alpha1 "github.com/lsst/qserv-operator/pkg/apis/qserv/v1alpha1"
	"github.com/lsst/qserv-operator/pkg/scheme/qserv"
	"github.com/lsst/qserv-operator/pkg/staging/syncer"
)

// NewCzarServiceSyncer returns a new sync.Interface for reconciling Czar Service
func NewCzarServiceSyncer(r *qservv1alpha1.Qserv, c client.Client, scheme *runtime.Scheme) syncer.Interface {
	svc := qserv.GenerateCzarService(r, controllerLabels)
	return syncer.NewObjectSyncer("CzarService", r, svc, c, scheme, noFunc)
}

// NewCzarServiceSyncer returns a new sync.Interface for reconciling Czar Service
func NewWorkerServiceSyncer(r *qservv1alpha1.Qserv, c client.Client, scheme *runtime.Scheme) syncer.Interface {
	svc := qserv.GenerateWorkerService(r, controllerLabels)
	return syncer.NewObjectSyncer("WorkerService", r, svc, c, scheme, noFunc)
}

// NewXrootdRedirectorServiceSyncer returns a new sync.Interface for reconciling Xrootd Redirector Service
func NewXrootdRedirectorServiceSyncer(r *qservv1alpha1.Qserv, c client.Client, scheme *runtime.Scheme) syncer.Interface {
	svc := qserv.GenerateXrootdRedirectorService(r, controllerLabels)
	return syncer.NewObjectSyncer("XrootdRedirectorService", r, svc, c, scheme, noFunc)
}
