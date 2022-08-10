package web

import (
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/log"
	"go.uber.org/atomic"

	"git.harmonycloud.cn/yeyazhou/go-httpserver/pkg/config"
	"git.harmonycloud.cn/yeyazhou/go-httpserver/pkg/core/business"
)

// Options for the web Handler.
type Options struct {
	ListenAddress   string
	EnableLifecycle bool
	CertFile        string
	KeyFile         string
	// Flags           map[string]string
}

type Handler struct {
	mtx    sync.RWMutex
	logger log.Logger

	business *business.API

	router   chi.Router
	reloadCh chan chan error
	options  *Options
	config   *config.Config
	birth    time.Time
	cwd      string

	ready atomic.Bool // ready is uint32 rather than boolean to be able to use atomic functions.
}
