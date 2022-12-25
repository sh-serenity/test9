package main

import (
//    "bufio"
    "context"
    "flag"
    "fmt"
//    "os"
    "path/filepath"

    appsv1 "k8s.io/api/apps/v1"
    apiv1 "k8s.io/api/core/v1"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/util/homedir"
    "k8s.io/klog"
    apierrors "k8s.io/apimachinery/pkg/api/errors"
//    "k8s.io/client-go/util/retry"
    //
    // Uncomment to load all auth plugins
    // _ "k8s.io/client-go/plugin/pkg/client/auth"
    //
    // Or uncomment to load specific auth plugins
    // _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
    // _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
    // _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {
    var kubeconfig *string
    if home := homedir.HomeDir(); home != "" {
	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
    } else {
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
    }
    flag.Parse()

    config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
	panic(err)
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
	panic(err)
    }

    deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

    deployment := &appsv1.Deployment{
	ObjectMeta: metav1.ObjectMeta{
	    Name: "demo",
	},
	Spec: appsv1.DeploymentSpec{
//	    Replicas: int32Ptr(2),
	    Selector: &metav1.LabelSelector{
		MatchLabels: map[string]string{
		    "app": "demo",
		},
	    },
	    Template: apiv1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
		    Labels: map[string]string{
			"app": "demo",
		    },
		},
		Spec: apiv1.PodSpec{
		    Containers: []apiv1.Container{
			{
			    Name:  "name",
			    Image: "docker-dev-kaas-local.docker.mirantis.net/mirantis/ceph/golang:1.18.5",
			    Command: []string{
				"sleep",
			    },
			    Args: []string{
				"10000",
			    },			    },
			},
		    },
		},
	    },
	}
    

    // Create Deployment
    fmt.Println("Creating deployment...")
    result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
    if err != nil {
	panic(err)
    }
    fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
    result, err = clientset.AppsV1().Deployments("default").Get(context.Background(), "demo", metav1.GetOptions{})
    if err != nil {	
 	if apierrors.IsNotFound(err) {
	    // create if there's no ceph-tools deployment
	    klog.V(1).Infof("create name  deployment %s/%s", "default", "demo")
	    _, err = deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	    if err != nil {
		fmt.Println(err)
	    }
	} else {
	    klog.Errorf("failed to get name deployment %s/%s: %v", "default", "demo", err)
	    if err != nil {
		fmt.Println(err)
	    }

	}
    }
}