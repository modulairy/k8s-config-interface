package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/modulairy/k8s-configmap-api-server/config"
	"github.com/modulairy/k8s-configmap-api-server/internal/confighandler/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Put(ctx *model.Context) {

	config.Mutex.Lock()
	defer config.Mutex.Unlock()

	currentConfigMap, err := ctx.Clientset.CoreV1().ConfigMaps(ctx.Configuration.Namespace).Get(context.TODO(), ctx.Configuration.ConfigName, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Error getting ConfigMap: %v\n", err)
		http.Error(ctx.Writer, fmt.Sprintf("Internal Server Error: %w", err), http.StatusInternalServerError)
		return
	}

	currentConfigMap.Data[ctx.Configuration.ConfigKey] = ctx.BodyAsString
	_, err = ctx.Clientset.CoreV1().ConfigMaps(ctx.Configuration.Namespace).Update(context.TODO(), currentConfigMap, metav1.UpdateOptions{})
	if err != nil {
		fmt.Printf("Error updating ConfigMap: %v\n", err)
		http.Error(ctx.Writer, fmt.Sprintf("Internal Server Error: %w", err), http.StatusInternalServerError)
		return
	}
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write([]byte("ConfigMap updated successfully"))
}
