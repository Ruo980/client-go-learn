package config

import (
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

// 配置路径常量： 配置文件的绝对路径
const kubeConfigFilePath = "./kube/conf"

// K8sConfig
//
//	@Description: client-go 设置对象：生成配置对象、构造函数、提供初始化各客户端对象的方法
type K8sConfig struct {
}

// NewK8sConfig
//
//	@Description: 无参构造函数
//	@return *K8sConfig
func NewK8sConfig() *K8sConfig {
	return &K8sConfig{}
}

// K8sRestConfig
//
//	@Description: 读取配置文件，生成一个配置对象
//	@receiver this
//	@return *rest.conf
func (this *K8sConfig) K8sRestConfig() *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigFilePath)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

// InitClient
//
//	@Description: 生成一种客户端对象：ClientSet
//	@receiver this
//	@return *kubernetes.Clientset
func (this *K8sConfig) InitClient() *kubernetes.Clientset {
	c, err := kubernetes.NewForConfig(this.K8sRestConfig())

	if err != nil {
		log.Fatal(err)
	}

	return c
}

// InitDynamicClient
//
//	@Description: 生成一种客户端对象：DynamicClient
//	@receiver this
//	@return dynamic.Interface
func (this *K8sConfig) InitDynamicClient() dynamic.Interface {
	c, err := dynamic.NewForConfig(this.K8sRestConfig())

	if err != nil {
		log.Fatal(err)
	}

	return c
}

// InitDiscoveryClient
//
//	@Description: 生成一种客户端对象：DiscoveryClient
//	@receiver this
//	@return *discovery.DiscoveryClient
func (this *K8sConfig) InitDiscoveryClient() *discovery.DiscoveryClient {
	return discovery.NewDiscoveryClient(this.InitClient().RESTClient())
}
