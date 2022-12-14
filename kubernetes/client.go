package kubernetes

import (
	"context"

	"github.com/gKits/easykube/utils"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
    Clientset   *kubernetes.Clientset
    kubeconfig  string
}

/*
creates a new kubernetes client

:param: kubeconfig: string: path to config file of kubernetes cluster
:return: error: on fail
:return: Client: a kubernetes client with a kubernetes Clientset
*/
func NewClient(kubeconfig string) (Client, error) {
    client := Client{}

    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        return Client{}, err
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return Client{}, err
    }

    client.kubeconfig = kubeconfig
    client.Clientset = clientset

    return client, nil
}

/*
creates a new deployment on the kubernetes cluster

:param: deploymentYaml: string: path to yaml file defining deployment
:return: error: on fail
*/
func (c *Client) CreateDeployment(deploymentYaml string) error{
    var deployment appsV1.Deployment
    if err := utils.UnmarshalYaml(deploymentYaml, &deployment); err != nil {
        return err
    }
    if _, err := c.Clientset.AppsV1().Deployments(deployment.Namespace).Create(context.TODO(), &deployment, metaV1.CreateOptions{}); err != nil {
        return err
    }
    return nil
}

/*
creates a new secret on the kubernetes cluster

:param: secretYaml: string: path to yaml file defining secret
:return: error: on fail
*/
func (c *Client) CreateSecret(secretYaml string) error{
    var secret coreV1.Secret
    if err := utils.UnmarshalYaml(secretYaml, &secret); err != nil {
        return err
    }

    if _, err := c.Clientset.CoreV1().Secrets(secret.Namespace).Create(context.TODO(), &secret, metaV1.CreateOptions{}); err != nil {
        return err
    }
    return nil
}

/*
creates a new service on the kubernetes cluster

:param: serviceYaml: string: path to yaml file defining service
:return: error: on fail
*/
func (c *Client) CreateService(serviceYaml string) error{
    var service coreV1.Service
    if err := utils.UnmarshalYaml(serviceYaml, &service); err != nil {
        return err
    }

    if _, err := c.Clientset.CoreV1().Services(service.Namespace).Create(context.TODO(), &service, metaV1.CreateOptions{}); err != nil {
        return err
    }
    return nil
}

/*
creates a new pod on the kubernetes cluster

:param: podYaml: string: path to yaml file defining pod
:return: error: on fail
*/
func (c *Client) CreatePod(podYaml string) error{
    var pod coreV1.Pod
    if err := utils.UnmarshalYaml(podYaml, &pod); err != nil {
        return err
    }

    if _, err := c.Clientset.CoreV1().Pods(pod.Namespace).Create(context.TODO(), &pod, metaV1.CreateOptions{}); err != nil {
        return err
    }
    return nil
}

