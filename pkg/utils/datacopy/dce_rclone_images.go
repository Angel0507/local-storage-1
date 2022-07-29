package datacopy

import (
	"fmt"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	k8sruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func (dcm *DataCopyManager) GetDCERcloneImage(deployName, deployNamespace, containerName string) (string, error) {
	registryEp, version, err := dcm.getRegistryEpAndVersion(deployName, deployNamespace, containerName)
	if err != nil {
		logger.WithError(err).Error("Failed to get registry Endpoint")
		return "", err
	}

	return fmt.Sprintf("%s/dce-uds-rclone:%s", registryEp, version), nil
}

func (dcm *DataCopyManager) getRegistryEpAndVersion(deployName, deployNamespace, containerName string) (string, string, error) {
	selfDeployment := &appsv1.Deployment{}
	objKey := k8sruntimeclient.ObjectKey{
		Name:      deployName,
		Namespace: deployNamespace,
	}

	if err := dcm.k8sControllerClient.Get(dcm.ctx, objKey, selfDeployment); err != nil {
		logger.WithError(err).Errorf(
			"Failed to get self delplyment %s, %s",
			deployName,
			deployNamespace,
		)
		return "", "", err
	}

	var image string
	for _, container := range selfDeployment.Spec.Template.Spec.Containers {
		if container.Name == containerName {
			image = container.Image
			break
		}
	}

	sepWithHost := strings.Split(image, "/")

	if len(sepWithHost) < 2 {
		err := fmt.Errorf("Invalid container image name")
		logger.WithError(err).Error("Failed to parser URL from image")
		return "", "", err
	}

	versionIndex := strings.LastIndex(sepWithHost[len(sepWithHost)-1], ":")
	version := sepWithHost[len(sepWithHost)-1][versionIndex+1:]

	return strings.Join(sepWithHost[:len(sepWithHost)-1], "/"), version, nil
}
