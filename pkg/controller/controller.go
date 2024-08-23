package controller

import (
	"github.com/ComplianceAsCode/compliance-operator/pkg/controller/metrics"
	"github.com/ComplianceAsCode/compliance-operator/pkg/utils"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// AddToManagerFuncs is a list of functions to add all Controllers to the Manager
var AddToManagerFuncs []func(manager.Manager, *metrics.Metrics, utils.CtlplaneSchedulingInfo, *kubernetes.Clientset) error

// AddToManager adds all Controllers to the Manager
func AddToManager(m manager.Manager,
	met *metrics.Metrics,
	si utils.CtlplaneSchedulingInfo,
	kubeClient *kubernetes.Clientset,
) error {
	setupLog := ctrl.Log.WithName("controller").WithName("AddToManager")

	// Add metrics Startup to the manager
	setupLog.Info("Adding metrics to manager")
	if err := m.Add(met); err != nil {
		setupLog.Error(err, "Failed to add metrics to manager")
		return err
	}
	setupLog.Info("Metrics added to manager successfully")

	// Add controllers to manager
	setupLog.Info("Adding controllers to manager")
	for _, f := range AddToManagerFuncs {
		setupLog.Info("Invoking AddToManager function", "function", f)
		if err := f(m, met, si, kubeClient); err != nil {
			setupLog.Error(err, "Failed to add controller to manager", "function", f)
			return err
		}
		setupLog.Info("Controller added to manager successfully", "function", f)
	}

	return nil
}
