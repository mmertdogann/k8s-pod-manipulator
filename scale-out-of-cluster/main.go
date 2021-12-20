package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func connectToK8s() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Fatal(err)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("Failed to create K8s clientset")
	}

	return clientset
}

func checkRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website!")
}

func scale(replicaCount int32) {
	clientset := connectToK8s()

	s, err := clientset.AppsV1().
		Deployments("default").
		GetScale(context.TODO(), "nodeapp-deployment", metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}

	sc := *s
	sc.Spec.Replicas = replicaCount

	us, err := clientset.AppsV1().
		Deployments("default").
		UpdateScale(context.TODO(),
			"nodeapp-deployment", &sc, metav1.UpdateOptions{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(*us)
}

func scaleController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	replicaCount := vars["replicaCount"]
	convertedReplica, _ := strconv.Atoi(replicaCount)
	scale(int32(convertedReplica))
	fmt.Fprintf(w, "Deployment has scaled to %s replica/s", replicaCount)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", checkRoot)

	r.HandleFunc("/scale/{replicaCount}", scaleController)

	http.ListenAndServe(":3000", r)
}
