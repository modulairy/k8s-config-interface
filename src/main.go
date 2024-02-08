package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	cruntimeconfig "sigs.k8s.io/controller-runtime/pkg/client/config"
)

var (
	clientset *kubernetes.Clientset
	mutex     sync.Mutex
	configName string
	namespace string
	configKey string
)

func init(){
	var (
		err error
		config *rest.Config
	)
	_, ok := os.LookupEnv("KUBERNETES_SERVICE_HOST")
	if	!ok {
		config = cruntimeconfig.GetConfigOrDie()
	}else{
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err)
		}
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	configName = strings.Trim(os.Getenv("PERMIT_CONFIG_NAME")," ")
	namespace = strings.Trim(os.Getenv("PERMIT_NAMESPACE")," ")
	configKey = strings.Trim(os.Getenv("PERMIT_CONFIG_KEY")," ")
}

func updateConfigMap(w http.ResponseWriter, r *http.Request) {
	var(
		_namespace string
		_key string
		_configName string
	)
	if(len(namespace)==0){
		_namespace = r.URL.Query().Get("subscriptionId")
	}else{
		_namespace= namespace
	}

	if(len(configName)==0){
		_configName = r.URL.Query().Get("configName")
	}else{
		_configName= configName
	}

	if(len(configKey)==0){
		_key = r.URL.Query().Get("configKey")
	}else{
		_key= configKey
	}

	var request interface{}
	data, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &request)
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		http.Error(w,fmt.Sprintf("Error formatting JSON: %w" , err), http.StatusBadRequest)
		return
	}
	currentConfigMap, err := clientset.CoreV1().ConfigMaps(_namespace).Get(context.TODO(), _configName, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Error getting ConfigMap: %v\n", err)
		http.Error(w,fmt.Sprintf("Internal Server Error: %w" , err), http.StatusInternalServerError)
		return
	}

	currentConfigMap.Data[_key]=prettyJSON.String()
	_, err = clientset.CoreV1().ConfigMaps(_namespace).Update(context.TODO(), currentConfigMap, metav1.UpdateOptions{})
	if err != nil {
		fmt.Printf("Error updating ConfigMap: %v\n", err)
		http.Error(w,fmt.Sprintf("Internal Server Error: %w" , err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ConfigMap updated successfully"))
}

func getConfigMap(w http.ResponseWriter, r *http.Request) {
	var(
		_namespace string
		_key string
		_configName string
	)
	if(len(namespace)==0){
		_namespace = r.URL.Query().Get("subscriptionId")
	}else{
		_namespace= namespace
	}

	if(len(configName)==0){
		_configName = r.URL.Query().Get("configName")
	}else{
		_configName= configName
	}

	if(len(configKey)==0){
		_key = r.URL.Query().Get("configKey")
	}else{
		_key= configKey
	}
	currentConfigMap, err := clientset.CoreV1().ConfigMaps(_namespace).Get(context.TODO(), _configName, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Error getting ConfigMap: %v\n", err)
		http.Error(w,fmt.Sprintf("Internal Server Error: %w" , err), http.StatusInternalServerError)
		return
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(currentConfigMap.Data[_key]), "", "  "); err != nil {
		http.Error(w,fmt.Sprintf("Error formatting JSON: %w" , err), http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(currentConfigMap.Data[_key]))	

}

func main() {
	fmt.Printf("starting server...")
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ok"))
	})
	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			getConfigMap(w,r)
		case http.MethodPut:
			updateConfigMap(w,r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
