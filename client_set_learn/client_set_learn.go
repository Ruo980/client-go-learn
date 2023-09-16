package client_set_learn

import (
	"context"
	"fmt"

	"log"

	"client-go-learn/config"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// getPodsDefault
//
//	@Description: 使用 ClientSet 获得 Default 命名空间所有 Pod
//	@param clientset
//	@return *v1.PodList
func getPodsTest(clientset *kubernetes.Clientset) *v1.PodList {
	pods, err := clientset.
		CoreV1().
		Pods("test").
		List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	return pods
}

// ClientSetLearn
//
//	@Description: 主函数：执行 ClientSet 操作的入口
func ClientSetLearn() {
	// 利用 k8sConfig 获取 ClientSet对象
	k8sConfig := config.NewK8sConfig()
	client := k8sConfig.InitClient()

	// 得到 test 空间全部 pod 信息
	podList := getPodsTest(client)
	for _, item := range podList.Items {
		fmt.Printf("namespace:%v,name:%v\n", item.Namespace, item.Name)
	}
}
