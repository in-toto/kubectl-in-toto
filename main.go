package main

import (
	"flag"
	"fmt"
	"os"
    "strings"

	_ "github.com/golang/glog"
	"github.com/in-toto/kubectl-in-toto/pkg/in_toto"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/kubectl/pkg/pluginutils"
)

const (
	unknown = "unknown type, must be pod"
)

func parseTarget(target []string) (string, string, error) {
    if len(target) < 1 {
        return "", "", fmt.Errorf("")
    }

    parts := strings.SplitN(target[0], "/", 2)
    targetType := parts[0]
    targetName := parts[1]
    if targetType != "pod" {
        return "", "" , fmt.Errorf(unknown)
    }
    return targetType, targetName, nil
}

func parseArgs() (*in_toto.VerificationSetup, error) {

    setup := new(in_toto.VerificationSetup)

    flag.StringVar(&setup.KeyPath, "key", "root.pub",
        "the pathname to the root pubkey (root.pub)")
    flag.StringVar(&setup.KeyPath, "k", "root.pub",
        "the pathname to the root pubkey (root.pub)")
    flag.StringVar(&setup.LayoutPath,
        "layout", "root.layout", "the name of the root layout (root.layout")
    flag.StringVar(&setup.LayoutPath,
        "l", "root.layout", "the name of the root layout (root.layout")
    flag.Parse()
    targetType, targetName, err := parseTarget(flag.Args())

    if err != nil {
        return nil, err
    }
    setup.TargetType = targetType
    setup.Name = targetName

    return setup, nil
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

func main() {

    inTotoConfig, err := parseArgs();
    if err != nil {
        fmt.Println(err.Error())
        flag.Usage()
        os.Exit(1)
    }
	client, ns := loadConfig()

    fmt.Printf("[resolve] scanning pod: %s\n", inTotoConfig.Name)
    handler := in_toto.ResolveResourceTypeHandler(inTotoConfig.TargetType)
    if handler == nil {
        flag.Usage()
        os.Exit(1)
    }

    containers := handler(client, inTotoConfig.Name, ns)
    for _, container := range containers {
        fmt.Printf("[scan] resolved pod container as: %v (%v). \n\tIn-toto output follows:\n",
            container.ImageID, container.Imagename)
        result := in_toto.ScanContainer(inTotoConfig, container.ImageID)
        if result != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    }
}
