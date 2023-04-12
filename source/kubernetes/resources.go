package kubernetes

import (
	"context"
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cshiaa/go-login-demo/global"
	"github.com/cshiaa/go-login-demo/source/config"
)


func GetDeploymentList() (list []config.ResourcesDeployment, err error){

	deploymentList, err := global.RY_CLIENTSET.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		global.RY_LOG.Error("获取Kubernetes Deployment List 失败")
		return nil, err
	}
	for _, dep := range deploymentList.Items {
		var deployment config.ResourcesDeployment

		depJSON, _ := json.Marshal(dep)
		if err := json.Unmarshal(depJSON, &deployment); err != nil {
			panic(err)
		}
		list = append(list, deployment)
	}
	return list, err
}

