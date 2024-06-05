// Code generated by go generate; DO NOT EDIT.
package obcluster

import ttypes "github.com/oceanbase/ob-operator/pkg/task/types"

const (
	tWaitOBZoneTopologyMatch    ttypes.TaskName = "wait obzone topology match"
	tWaitOBZoneDeleted          ttypes.TaskName = "wait obzone deleted"
	tModifyOBZoneReplica        ttypes.TaskName = "modify obzone replica"
	tDeleteOBZone               ttypes.TaskName = "delete obzone"
	tCreateOBZone               ttypes.TaskName = "create obzone"
	tBootstrap                  ttypes.TaskName = "bootstrap"
	tCreateUsers                ttypes.TaskName = "create users"
	tMaintainOBParameter        ttypes.TaskName = "maintain obparameter"
	tValidateUpgradeInfo        ttypes.TaskName = "validate upgrade info"
	tUpgradeCheck               ttypes.TaskName = "upgrade check"
	tBackupEssentialParameters  ttypes.TaskName = "backup essential parameters"
	tBeginUpgrade               ttypes.TaskName = "begin upgrade"
	tRollingUpgradeByZone       ttypes.TaskName = "rolling upgrade by zone"
	tFinishUpgrade              ttypes.TaskName = "finish upgrade"
	tModifySysTenantReplica     ttypes.TaskName = "modify sys tenant replica"
	tCreateServiceForMonitor    ttypes.TaskName = "create service for monitor"
	tRestoreEssentialParameters ttypes.TaskName = "restore essential parameters"
	tCheckAndCreateUserSecrets  ttypes.TaskName = "check and create user secrets"
	tCreateOBClusterService     ttypes.TaskName = "create obcluster service"
	tCheckImageReady            ttypes.TaskName = "check image ready"
	tCheckClusterMode           ttypes.TaskName = "check cluster mode"
	tCheckMigration             ttypes.TaskName = "check migration"
	tScaleOBZonesVertically     ttypes.TaskName = "scale obzones vertically"
	tExpandPVC                  ttypes.TaskName = "expand pvc"
	tModifyServerTemplate       ttypes.TaskName = "modify server template"
	tWaitOBZoneBootstrapReady   ttypes.TaskName = "wait obzone bootstrap ready"
	tWaitOBZoneRunning          ttypes.TaskName = "wait obzone running"
	tRollingUpdateOBZones       ttypes.TaskName = "rolling update obzones"
	tCheckEnvironment           ttypes.TaskName = "check environment"
	tAnnotateOBCluster          ttypes.TaskName = "annotate obcluster"
	tAdjustParameters           ttypes.TaskName = "adjust parameters"
)
