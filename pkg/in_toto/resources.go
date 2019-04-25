package in_toto

import (
    "fmt"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/* A small helper struct to hold container image information */
type ContainerSpec struct {
    Imagename string
    ImageID string
}

type resourceTypeHandler func(client *kubernetes.Clientset, name string, namespace string) ([]ContainerSpec)
/*
ResolveResourceTypeHandler

Resolves a function handler fo a specific type of resource name. The function
handler will have a signature of:
    func(client kubernetes.Clientset, name string, namespace string) ([]ContainerSpec)

*/
func ResolveResourceTypeHandler(resourceType string) resourceTypeHandler {

    if resourceType == "pod" {
        return ResolvePod
    }

    return nil
}

func ResolvePod(client *kubernetes.Clientset, name string, namespace string) ([]ContainerSpec) {

    result := make([]ContainerSpec, 0)

    pod, err := client.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})
    if err != nil {
        fmt.Println(err)
    }

    pod.TypeMeta = metav1.TypeMeta{
        Kind:       "Pod",
        APIVersion: "v1",
    }
    for _, status := range pod.Status.ContainerStatuses {
        result = append(result,
                        ContainerSpec{
                            Imagename: status.Image,
                            ImageID: status.ImageID,
                        })
    }

    return result

}
