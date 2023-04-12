package controller

import (
	// "context"

	"net/http"

	"github.com/gin-gonic/gin"

	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/cshiaa/go-login-demo/global"
	rykubernetes "github.com/cshiaa/go-login-demo/source/kubernetes"
)

//获取Kubernetes版本 检查Kubernetes配置文件是否成功
func GetKubernetesVersion(c *gin.Context) {

	version, err := rykubernetes.InitKubernetes()
	if err != nil {
		global.RY_LOG.Error("Kubernetes 初始化失败")
		c.JSON(http.StatusBadRequest, gin.H{"message":"Kubernetes 连接失败，请检查配置文件", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"success","version": version.GitVersion})

}

//获取Kubernetes Resource
func GetKubernetesResource(c *gin.Context) {

	
	_, err := rykubernetes.InitKubernetes()
	if err != nil {
		global.RY_LOG.Error("Kubernetes 连接失败")
	}
	deploymentList, err := rykubernetes.GetDeploymentList()
	if err != nil {
		global.RY_LOG.Error("获取Kubernetes resource失败")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"success","DeploymentList": deploymentList})
}