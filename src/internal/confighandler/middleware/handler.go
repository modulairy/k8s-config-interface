package middleware

import (
	"net/http"

	"github.com/modulairy/k8s-configmap-api-server/internal/confighandler/handler"
	"github.com/modulairy/k8s-configmap-api-server/internal/confighandler/model"
)

func InvokeHandler(ctx *model.Context) {
	switch ctx.Request.Method {
	case http.MethodGet:
		handler.Get(ctx)
	case http.MethodPut:
		handler.Put(ctx)
	default:
		http.Error(ctx.Writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
