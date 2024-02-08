package config

import (
	"os"
	"strings"
	"sync"
)

var (
	PERMIT_CONFIG_KEY  string
	PERMIT_NAMESPACE   string
	PERMIT_CONFIG_NAME string
	Mutex              sync.Mutex
)

func init() {
	PERMIT_CONFIG_NAME = strings.Trim(os.Getenv("PERMIT_CONFIG_NAME"), " ")
	PERMIT_NAMESPACE = strings.Trim(os.Getenv("PERMIT_NAMESPACE"), " ")
	PERMIT_CONFIG_KEY = strings.Trim(os.Getenv("PERMIT_CONFIG_KEY"), " ")
}
