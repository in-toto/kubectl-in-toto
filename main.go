package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	_ "github.com/golang/glog"
	"github.com/santiagotorres/kubectl-in-toto/pkg/in_toto"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/kubectl/pkg/pluginutils"
)

func init() {
	flag.CommandLine.Set("logtostderr", "true")
	flag.CommandLine.Set("v", os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_V"))
}

const (
	usage   = "usage: kubectl in-toto [pod|deployment|statefulset|daemonset]/name"
	unknown = "unknown type must be pod, deployment, statefulset or daemonset"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	resource := os.Args[1]
	parts := strings.Split(resource, "/")

	if len(parts) != 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	client, ns := loadConfig()

	fmt.Println("scanning pod", parts[1])
	pod, err := client.CoreV1().Pods(ns).Get(parts[1], metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
    }
    pod.TypeMeta = metav1.TypeMeta{
        Kind:       "Pod",
        APIVersion: "v1",
    }
    for i := range pod.Spec.Containers {
        fmt.Printf("%v", pod.Spec.Containers[i].Image);
        result, _ := in_toto.NewClient().ScanContainer(pod.Spec.Containers[i].Image)
	    result.Dump(os.Stdout)
    }
}

func loadConfig() (*kubernetes.Clientset, string) {
	restConfig, kubeConfig, err := pluginutils.InitClientAndConfig()
	if err != nil {
		panic(err)
	}
	c := kubernetes.NewForConfigOrDie(restConfig)
	ns, _, _ := kubeConfig.Namespace()
	return c, ns
}
