package model

import (
	"net/http"

	"k8s.io/client-go/kubernetes"
)

type Context struct {
	Clientset     *kubernetes.Clientset
	Configuration Configuration
	Writer        http.ResponseWriter
	Request       *http.Request
	BodyAsString  string
}
