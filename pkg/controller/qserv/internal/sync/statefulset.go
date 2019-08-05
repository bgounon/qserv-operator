package sync

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	qservv1alpha1 "github.com/lsst/qserv-operator/pkg/apis/qserv/v1alpha1"
	"github.com/lsst/qserv-operator/pkg/scheme/qserv"
	"github.com/lsst/qserv-operator/pkg/staging/syncer"
)

// NewCzarStatefulSetSyncer returns a new sync.Interface for reconciling Qserv Czar StatefulSet
func NewCzarStatefulSetSyncer(r *qservv1alpha1.Qserv, c client.Client, scheme *runtime.Scheme) syncer.Interface {
	statefulSet := qserv.GenerateCzarStatefulSet(r, controllerLabels)
	return syncer.NewObjectSyncer("CzarStatefulSet", r, statefulSet, c, scheme, noFunc)
}

// NewWorkerStatefulSetSyncer returns a new sync.Interface for reconciling Qserv Worker StatefulSet
func NewWorkerStatefulSetSyncer(r *qservv1alpha1.Qserv, c client.Client, scheme *runtime.Scheme) syncer.Interface {
	statefulSet := qserv.GenerateWorkerStatefulSet(r, controllerLabels)
	return syncer.NewObjectSyncer("WorkerStatefulSet", r, statefulSet, c, scheme, noFunc)
}

// NewXrootdStatefulSetSyncer returns a new sync.Interface for reconciling xrootd redirectors cluster StatefulSet
func NewXrootdStatefulSetSyncer(r *qservv1alpha1.Qserv, c client.Client, scheme *runtime.Scheme) syncer.Interface {
	statefulSet := qserv.GenerateXrootdStatefulSet(r, controllerLabels)
	return syncer.NewObjectSyncer("XrootdStatefulSet", r, statefulSet, c, scheme, noFunc)
}
