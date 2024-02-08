package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/modulairy/k8s-configmap-api-server/config"
	"github.com/modulairy/k8s-configmap-api-server/internal/confighandler/model"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	cruntimeconfig "sigs.k8s.io/controller-runtime/pkg/client/config"
)

var clientSet *kubernetes.Clientset

func init() {

	var err error
	var config *rest.Config

	if _, ok := os.LookupEnv("KUBERNETES_SERVICE_HOST"); !ok {
		if config, err = cruntimeconfig.GetConfig(); err != nil {
			panic(err)
		}
	} else {
		if config, err = rest.InClusterConfig(); err != nil {
			panic(err)
		}
	}

	if config == nil {
		panic("config is not found.")
	}

	if clientSet, err = kubernetes.NewForConfig(config); err != nil {
		panic(err)
	}

}

func InvokeContext(writer http.ResponseWriter, request *http.Request) (*model.Context, error) {

	var prettyJSON bytes.Buffer

	if request.Method == http.MethodPatch ||
		request.Method == http.MethodPut ||
		request.Method == http.MethodPost {
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, fmt.Sprintf("Error formatting JSON: %w", err), http.StatusBadRequest)
			return nil, err
		}
		json.Unmarshal(data, &request)
		if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
			http.Error(writer, fmt.Sprintf("Error formatting JSON: %w", err), http.StatusBadRequest)
			return nil, err
		}
	}
	return &model.Context{
		Writer:       writer,
		Request:      request,
		Clientset:    clientSet,
		BodyAsString: prettyJSON.String(),
		Configuration: model.Configuration{
			ConfigName: config.PERMIT_CONFIG_NAME,
			ConfigKey:  config.PERMIT_CONFIG_KEY,
			Namespace:  config.PERMIT_NAMESPACE,
		},
	}, nil
}
