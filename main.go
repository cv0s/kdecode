package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Get kubernetes config from local config
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error getting user home dir: %v\n", err)
		os.Exit(1)
	}
	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		fmt.Printf("error getting Kubernetes config: %v\n", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		fmt.Printf("error getting Kubernetes clientset: %v\n", err)
		os.Exit(1)
	}

	// Define Flags
	var (
		namespace  string
		secretName string
	)

	flag.StringVar(&namespace, "namespace", "", "defines the namespace of the secret")
	flag.StringVar(&secretName, "secret", "", "Define the secret to be decoded")

	// Parsen der Flags
	flag.Parse()

	secrets, err := clientset.CoreV1().Secrets(namespace).Get(context.WithValue(context.Background(), ".secret-file", ".secret-file"), secretName, v1.GetOptions{})

	//print secrets
	fmt.Printf("Secret Data:\n")
	for key, value := range secrets.Data {
		fmt.Printf("%s: %s\n", key, string(value))
	}

}
