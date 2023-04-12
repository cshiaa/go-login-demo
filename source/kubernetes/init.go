package kubernetes

import (
	"path/filepath"

	"go.uber.org/zap"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/apimachinery/pkg/version"

	"github.com/cshiaa/go-login-demo/global"
	"github.com/cshiaa/go-login-demo/utils"
)


func GetKubernetesClientSet() (clientset *kubernetes.Clientset, err error) {
	var kubeconfig string
	if utils.IsFileExist(global.RY_CONFIG.Kubernetes.Path){
		kubeconfig = global.RY_CONFIG.Kubernetes.Path
	} else if home := homedir.HomeDir(); home != "" {
		//kubernetes default path
		global.RY_LOG.Info("使用默认的配置文件: ${HOME}/.kube/config", zap.String("home", home))
		kubeconfig = filepath.Join(home, ".kube", "config")
	} else {
		kubeconfig = ""
	}
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func InitKubernetes() (version *version.Info, err error) {

	global.RY_CLIENTSET, err = GetKubernetesClientSet()
	if err != nil {
		global.RY_LOG.Error("Kubernetes 初始化失败")
		return nil, err
	}
	version, err = global.RY_CLIENTSET.ServerVersion()
	// pods, err := global.RY_CLIENTSET.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	// if err != nil {
	// 	global.RY_LOG.Error("获取Kubernetes Pods 失败")
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	if err != nil {
		global.RY_LOG.Error("获取Kubernetes 版本失败，请检查连接信息")
		return nil, err		
	}
	return version, err
}