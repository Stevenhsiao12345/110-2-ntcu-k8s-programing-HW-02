package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path"
	
	//corev1 "k8s.io/api/core/v1"
	apiv1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/utils/pointer"
	
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	config, err := clientcmd.BuildConfigFromFlags("", path.Join(home, ".kube/config"))
	if err != nil {
		panic(err.Error())
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}


	
	namespace := "default"

	//desired := corev1.ConfigMap{Data: map[string]string{"foo": "bar"}}
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "apa000dep1",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pointer.Int32Ptr(3),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"ntcu-k8s": "hw2",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"ntcu-k8s": "hw2",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "apa000ex91",
							Image: "nginx",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
	
	
	
	deployment.Namespace = namespace
	deployment.GenerateName = "crud-typed-simple-"

	// List
	cmList, err := client.
		AppsV1().
		Deployments(namespace).
		List(
			context.TODO(),
			metav1.ListOptions{},
		)
	if err != nil {
		panic(err.Error())
	}

	for _, c := range cmList.Items {
		fmt.Printf("Existing Deployment name: %s\n", c.Name)
	}
	prompt()

	// Create
	created, err := client.
		AppsV1().
		Deployments(namespace).
		Create(
			context.TODO(),
			deployment,
			metav1.CreateOptions{},
		)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Created Deployment %s/%s\n", namespace, created.GetObjectMeta().GetName())

	//if !reflect.DeepEqual(created.Data, desired.Data) {
	//	panic("Created Deployment has unexpected data")
	//}

	prompt()

	// Read
	read, err := client.
		AppsV1().
		Deployments(namespace).
		Get(
			context.Background(),
			created.GetObjectMeta().GetName(),
			metav1.GetOptions{},
		)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Read Deployment %s/%s\n", namespace, read.GetObjectMeta().GetName())
	//fmt.Println(read.ObjectMeta.Name["ntcu-k8s"])

	//if !reflect.DeepEqual(read.Data, desired.Data) {
	//	panic("Read Deployment has unexpected data")
	//}
	prompt()

	// Update
	/*
	read.ObjectMeta.Name["ntcu-k8s"] = "qux"
	updated, err := client.
		CoreV1().
		ConfigMaps(namespace).
		Update(
    		context.Background(),
			read,
			metav1.UpdateOptions{},
		)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Updated ConfigMap %s/%s\n", namespace, updated.GetName())
	fmt.Println(read.Data["foo"])

	if !reflect.DeepEqual(updated.Data, read.Data) {
		panic("Updated ConfigMap has unexpected data")
	}
	prompt()
     */

	 namespace2 := "default"

	 //desired := corev1.ConfigMap{Data: map[string]string{"foo": "bar"}}
	 service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "api000ser1",
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"ntcu-k8s": "hw2",
			},
			Type: apiv1.ServiceTypeNodePort,
			Ports: []apiv1.ServicePort{
				{
					Port: 80,
					NodePort: 30100,
					Protocol: apiv1.ProtocolTCP,

				},
			},
		},
	}
	 
	 
	 
	 service.Namespace = namespace2
	 service.GenerateName = "crud-typed-simple-"
    
     //List
	 
	 cmList2, err2 := client.
	 CoreV1().
	 Services(namespace2).
	 List(
		 context.TODO(),
		 metav1.ListOptions{},
	 )
   if err2 != nil {
	 panic(err2.Error())
    }

    for _, c1 := range cmList2.Items {
	 fmt.Printf("Existing Service name: %s\n", c1.Name)
    }
    prompt()
	
     //Create
	 created2, err2 := client.
	 CoreV1().
	 Services(namespace2).
	 Create(
		 context.TODO(),
		 service,
		 metav1.CreateOptions{},
	 )
	if err2 != nil {
		panic(err2.Error())
	}
    fmt.Printf("Created Service %s/%s\n", namespace2, created2.GetObjectMeta().GetName())

	prompt()

	// Read
	read2, err2 := client.
		CoreV1().
		Services(namespace2).
		Get(
			context.Background(),
			created2.GetObjectMeta().GetName(),
			metav1.GetOptions{},
		)
	if err2 != nil {
		panic(err2.Error())
	}

	fmt.Printf("Read Service %s/%s\n", namespace, read2.GetObjectMeta().GetName())
	//fmt.Println(read.ObjectMeta.Name["ntcu-k8s"])

	//if !reflect.DeepEqual(read.Data, desired.Data) {
	//	panic("Read Deployment has unexpected data")
	//}
	prompt() 



	// Delete
	err = client.
		AppsV1().
		Deployments(namespace).
		Delete(
			context.TODO(),
			created.GetObjectMeta().GetName(),
			metav1.DeleteOptions{},
		)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Deleted Deployment %s/%s\n", namespace, created.GetObjectMeta().GetName())

	err2 = client.
	CoreV1().
	Services(namespace2).
	Delete(
		context.TODO(),
		created2.GetObjectMeta().GetName(),
		metav1.DeleteOptions{},
	)
	if err != nil {
		panic(err2.Error())
	}

	fmt.Printf("Deleted Service %s/%s\n", namespace, created2.GetObjectMeta().GetName())

}

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}

