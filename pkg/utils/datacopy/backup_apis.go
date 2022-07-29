package datacopy

import (
	"time"
)

const (
	BackupAnnotationName          = "uds.dce.daocloud.io/backup-running-status"
	BackupJobAnnotationName       = "uds.dce.daocloud.io/backup-job-running-status"
	BackupTargetConfigAnnNameOnCM = "uds.dce.daocloud.io/backup-target-config"
	BackupConfigKeyNFS            = "nfs.config"

	BackupTmpPVCCreationTimeout = time.Minute * 5

	BackupSelfDeploymentName                    = "dce-uds-policy-controller"
	BackupSelfDeploymentNamespace               = "kube-system"
	BackupSelfDeploymentControllerContainerName = "dce-uds-policy-controller"

	BackupTargetTypeNFS = "NFS"
	BackupTargetTypeS3  = "S3"

	BackupWorkloadReplicasAnnName = "uds.dce.daocloud.io/backup-workload-replicas"
	BackupWorkloadControlTimeout  = time.Minute * 5

	BackendAbilityDirect = iota
	BackendAbilitySnapshot
	BackendAbilityNotSupport

	BackupStageStartBackup              = "Start backup"
	BackupStageFinishBackup             = "Finish backup"
	BackupStageInitRclone               = "Init rclone"
	BackupStageDataToTmpPVC             = "Copy data to tmp PVC"
	BackupStageDataToRemote             = "Copy data to Remote"
	BackupStageDeleteRemoteBackup       = "Deleting remote backup"
	BackupStageFinishDeleteRemoteBackup = "Finish delete remote backup"
	BackupStageUnknown                  = "Unknown"

	BackupPhasePending = "pending"
	BackupPhaseSuccess = "success"
	BackupPhaseRunning = "running"
	BackupPhaseFailed  = "failed"

	BackupFinalizerWaitForRemoteDeleted = "daocloud.io/wait-for-remote-deleted"

	FakeCapacityForNFS = int64(1024 * 1024 * 1024 * 10)
)

type BackupJobStatus struct {
	BackupCRName      string
	BackupCRNamespace string
	BackupCRUID       string
	Stage             string
	JobName           string
	Phase             string
	Event             string
	Message           string
	Progress          *Progress
}

type DataCopyUserStatus struct {
	BackupCRName      string `json:"backupCRName"`
	BackupCRNamespace string `json:"backupCRNamespace"`
	BackupCRUID       string `json:"BackupCRUID"`
}
