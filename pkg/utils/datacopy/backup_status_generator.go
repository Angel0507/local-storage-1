package datacopy

//type backupStatusGenerator struct {
//	// map[backupInstanceName]currentStage
//	stageMap map[string]string
//	bm       *BackupManager
//	lock     *sync.Mutex
//}
//
//func (bsg *backupStatusGenerator) SetCurrentStage(backupInstance *policyv1alpha1.Backup, stage string) {
//	bsg.lock.Lock()
//	defer bsg.lock.Unlock()
//
//	bsg.stageMap[uniqueName(backupInstance)] = stage
//}
//
//func (bsg *backupStatusGenerator) GetCurrentStage(backupInstance *policyv1alpha1.Backup) string {
//	if currentStage, exists := bsg.stageMap[uniqueName(backupInstance)]; exists && len(currentStage) > 0 {
//		return currentStage
//	}
//	return policyv1alpha1.BackupStageUnknown
//}
//
//func (bsg *backupStatusGenerator) getCurrentStageByUID(UID string) string {
//	if currentStage, exists := bsg.stageMap[UID]; exists && len(currentStage) > 0 {
//		return currentStage
//	}
//	return policyv1alpha1.BackupStageUnknown
//}
//
//func (bsg *backupStatusGenerator) sendPhase(backupInstance *policyv1alpha1.Backup, phase, msg string) {
//	bsg.bm.backupStatusCh <- &policyv1alpha1.BackupJobStatus{
//		BackupCRName:      backupInstance.Name,
//		BackupCRNamespace: backupInstance.Namespace,
//		Stage:             bsg.GetCurrentStage(backupInstance),
//		Message:           msg,
//		Phase:             phase,
//	}
//}
//
//func (bsg *backupStatusGenerator) DeleteStage(backupInstance *policyv1alpha1.Backup) {
//	delete(bsg.stageMap, uniqueName(backupInstance))
//}
//
//// Bypass backup job status to backup status channel,
//// so it can merge into main backup status
//func (bsg *backupStatusGenerator) watchDataCopyPhase() {
//	for {
//		select {
//		case ds := <-bsg.bm.dataCopyStatusCh:
//			userData := &policyv1alpha1.DataCopyUserStatus{}
//			json.Unmarshal([]byte(ds.UserData), userData)
//			bsg.bm.backupStatusCh <- &policyv1alpha1.BackupJobStatus{
//				BackupCRName:      userData.BackupCRName,
//				BackupCRNamespace: userData.BackupCRNamespace,
//				BackupCRUID:       userData.BackupCRUID,
//				Stage:             bsg.getCurrentStageByUID(userData.BackupCRUID),
//				JobName:           ds.JobName,
//				Phase:             ds.Phase,
//				Event:             ds.Event,
//				Message:           ds.Message,
//				Progress:          ds.Progress,
//			}
//		case <-bsg.bm.Ctx.Done():
//			return
//		}
//	}
//}
//
//func GetDataCopyUserData(backupInstance *policyv1alpha1.Backup) string {
//	dataCopyUserData := &policyv1alpha1.DataCopyUserStatus{
//		BackupCRName:      backupInstance.Name,
//		BackupCRNamespace: backupInstance.Namespace,
//		BackupCRUID:       string(backupInstance.UID),
//	}
//	jStr, _ := json.Marshal(dataCopyUserData)
//	return string(jStr)
//}
//
//func uniqueName(backupInstance *policyv1alpha1.Backup) string {
//	return string(backupInstance.UID)
//}
