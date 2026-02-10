package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var gvr = schema.GroupVersionResource{
	Group:    "security.demo",
	Version:  "v1",
	Resource: "imagescans",
}

func getConfig() (*rest.Config, error) {
	kubeconfig := filepath.Join(os.Getenv("USERPROFILE"), ".kube", "config")
	return clientcmd.BuildConfigFromFlags("", kubeconfig)
}

func main() {
	// create logs folder
	os.MkdirAll("logs", 0755)

	// create log file
	logFile, err := os.OpenFile("logs/crd.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)

	// kube config
	config, err := getConfig()
	if err != nil {
		panic(err)
	}
    // connecting to kubernetes
	// dynamic client
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	resource := client.Resource(gvr).Namespace("default")

	fmt.Println("Watcher started... waiting for ImageScan CRDs")

	// ✅ correct watcher
	watcher, err := resource.Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for event := range watcher.ResultChan() {
		obj := event.Object.(*unstructured.Unstructured)

		data, _ := json.MarshalIndent(obj.Object, "", "  ")

		logger.Println("NEW CRD CREATED:")
		logger.Println(string(data))
		logger.Println("------------------------")

		fmt.Println("New ImageScan detected and logged")

		time.Sleep(1 * time.Second)
	}
}
