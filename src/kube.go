package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"k8s.io/client-go/tools/clientcmd"
)

var kube_eventsCmd = &cobra.Command{
	Use:   "kube-events",
	Short: "kube-events pod",
	Long:  "kube-events example custom action with k8s based on event",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("kube-events podname is required")
			return
		}
		fmt.Println("args:", args)
		kube_events(args)
	},
}

func init() {
	RootCmd.AddCommand(kube_eventsCmd)
}

func kube_events(args []string) {
	fmt.Println("running kube_events", args)
	client, err := createClient("/Users/chris/.kube/config")
	if err != nil {
		fmt.Println("ERROR: kube createClient verify access to k8s cluster")
		return
	}
	events, _ := client.CoreV1().Events("default").List(context.TODO(), metav1.ListOptions{FieldSelector: "involvedObject.name=" + args[0], TypeMeta: metav1.TypeMeta{Kind: "Pod"}})
	for _, item := range events.Items {
		fmt.Println(item)
	}

}

func createClient(kubeconfigPath string) (kubernetes.Interface, error) {
	var kubeconfig *rest.Config

	if kubeconfigPath != "" {
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
		if err != nil {
			return nil, fmt.Errorf("unable to load kubeconfig from %s: %v", kubeconfigPath, err)
		}
		kubeconfig = config
	} else {
		config, err := rest.InClusterConfig()
		if err != nil {
			return nil, fmt.Errorf("unable to load in-cluster config: %v", err)
		}
		kubeconfig = config
	}

	client, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create a client: %v", err)
	}

	return client, nil
}
