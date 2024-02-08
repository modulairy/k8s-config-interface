package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/modulairy/k8s-configmap-api-server/internal/confighandler/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Get(ctx *model.Context) {

	currentConfigMap, err := ctx.Clientset.CoreV1().ConfigMaps(ctx.Configuration.Namespace).Get(context.TODO(), ctx.Configuration.ConfigName, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Error getting ConfigMap: %v\n", err)
		http.Error(ctx.Writer, fmt.Sprintf("Internal Server Error: %w", err), http.StatusInternalServerError)
		return
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(currentConfigMap.Data[ctx.Configuration.ConfigKey]), "", "  "); err != nil {
		http.Error(ctx.Writer, fmt.Sprintf("Error formatting JSON: %w", err), http.StatusBadRequest)
		return
	}
	ctx.Writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.Writer.Write([]byte(currentConfigMap.Data[ctx.Configuration.ConfigKey]))
}
