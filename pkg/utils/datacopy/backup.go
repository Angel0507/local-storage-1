package datacopy

//import (
//	"context"
//	"sync"
//
//	log "github.com/sirupsen/logrus"
//	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
//)
//
//var (
//	logger = log.WithField("module", "Backup")
//)
//
//type BackupManager struct {
//	Identifier          string
//	DataCopyManager     *DataCopyManager
//	backupStatusCh      chan *policyv1alpha1.BackupJobStatus
//	dataCopyStatusCh    chan *DataCopyStatus
//	K8sControllerClient k8sclient.Client
//	Ctx                 context.Context
//	backendAbility      int
//	Bsg                 *backupStatusGenerator
//}
//
//func New(ctx context.Context,
//	client k8sclient.Client,
//	backupStatusCh chan *policyv1alpha1.BackupJobStatus,
//	identifier string) (*BackupManager, error) {
//	dataCopyStatusCh := make(chan *datacopyutil.DataCopyStatus, 100)
//	dcm, err := datacopyutil.NewDataCopyManager(ctx, policyv1alpha1.BackupJobAnnotationName, client, dataCopyStatusCh)
//	if err != nil {
//		return nil, err
//	}
//
//	wcm := workloadsutil.NewWorkloadControlManager(ctx, client, identifier)
//	pm := pvcutil.NewPVCManager(ctx, client, identifier)
//	bm := &BackupManager{
//		Identifier:          identifier,
//		DataCopyManager:     dcm,
//		workloadsManager:    wcm,
//		PVCManager:          pm,
//		K8sControllerClient: client,
//		Ctx:                 ctx,
//		backupStatusCh:      backupStatusCh,
//		dataCopyStatusCh:    dataCopyStatusCh,
//	}
//	bsg := &backupStatusGenerator{
//		stageMap: make(map[string]string),
//		lock:     &sync.Mutex{},
//		bm:       bm,
//	}
//	bm.Bsg = bsg
//	return bm, nil
//}
//
//func (bm *BackupManager) Run() {
//	go bm.Bsg.watchDataCopyPhase()
//	bm.DataCopyManager.Run()
//}
//
//func (bm *BackupManager) DeleteBackup(backupInstance *policyv1alpha1.Backup) error {
//	defer bm.Bsg.DeleteStage(backupInstance)
//
//	var backupper Backupper
//	switch backupInstance.Spec.BackupTargetType {
//	case policyv1alpha1.BackupTargetTypeS3:
//		backupper = NewS3Backupper(backupInstance, bm)
//	case policyv1alpha1.BackupTargetTypeNFS:
//		backupper = NewNFSBackupper(backupInstance, bm)
//	}
//
//	if err := backupper.processDeleteBackup(backupInstance); err != nil {
//		bm.Bsg.sendPhase(backupInstance, policyv1alpha1.BackupPhaseFailed, err.Error())
//		return err
//	}
//
//	newBackupInstance := backupInstance.DeepCopy()
//	for i, finalizer := range newBackupInstance.Finalizers {
//		if finalizer == policyv1alpha1.BackupFinalizerWaitForRemoteDeleted {
//			newBackupInstance.Finalizers = append(newBackupInstance.Finalizers[:i], newBackupInstance.Finalizers[i+1:]...)
//		}
//	}
//	if err := bm.K8sControllerClient.Update(bm.Ctx, newBackupInstance); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (bm *BackupManager) GetBackendAbility(backupInstance *policyv1alpha1.Backup) int {
//	return policyv1alpha1.BackendAbilityNotSupport
//}
