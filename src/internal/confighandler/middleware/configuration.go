package middleware

import (
	"github.com/modulairy/k8s-configmap-api-server/config"
	"github.com/modulairy/k8s-configmap-api-server/internal/confighandler/model"
)

func InvokeConfigure(ctx *model.Context) (*model.Context, error) {

	if len(config.PERMIT_NAMESPACE) == 0 {
		ctx.Configuration.Namespace = ctx.Request.URL.Query().Get("subscriptionId")
	} else {
		ctx.Configuration.Namespace = config.PERMIT_NAMESPACE
	}

	if len(config.PERMIT_CONFIG_NAME) == 0 {
		ctx.Configuration.ConfigName = ctx.Request.URL.Query().Get("configName")
	} else {
		ctx.Configuration.ConfigName = config.PERMIT_CONFIG_NAME
	}

	if len(config.PERMIT_CONFIG_KEY) == 0 {
		ctx.Configuration.ConfigKey = ctx.Request.URL.Query().Get("configKey")
	} else {
		ctx.Configuration.ConfigKey = config.PERMIT_CONFIG_KEY
	}
	return ctx, nil
}
