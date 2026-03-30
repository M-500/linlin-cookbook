package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	clientset, err := getK8sClient()
	if err != nil {
		panic(err)
	}

	// 获取 Node 列表
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("=== Kubernetes Nodes ===")

	for _, node := range nodes.Items {
		printNode(node)
	}
}

func getK8sClient() (*kubernetes.Clientset, error) {
	// 优先使用集群内配置
	config, err := rest.InClusterConfig()
	if err == nil {
		return kubernetes.NewForConfig(config)
	}

	// 否则使用本地 kubeconfig
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

func printNode(node v1.Node) {
	fmt.Printf("Name: %s\n", node.Name)

	// 标签
	fmt.Printf("Labels: %v\n", node.Labels)

	// 状态
	for _, condition := range node.Status.Conditions {
		if condition.Type == v1.NodeReady {
			fmt.Printf("Status: %s\n", condition.Status)
		}
	}

	// CPU / 内存
	cpu := node.Status.Capacity.Cpu()
	mem := node.Status.Capacity.Memory()

	fmt.Printf("CPU: %s\n", cpu.String())
	fmt.Printf("Memory: %s\n", mem.String())

	fmt.Println("------------------------")
}
