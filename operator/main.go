package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "", "Path to a kubeconfig file")
	namespace := flag.String("namespace", "default", "Namespace to list Pods from")

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Erro ao criar configuração do cliente: %v", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error to create client k8s: %v", err)
		os.Exit(1)
	}

	pods, err := clientset.CoreV1().Pods(*namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error to list pods: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Pods in namespace %s:\n", *namespace)
	for _, pod := range pods.Items {
		fmt.Printf("- %s\n", pod.Name)
	}
}
