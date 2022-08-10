package business

import (
	"io"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"git.harmonycloud.cn/yeyazhou/go-httpserver/pkg/chilog"
	"git.harmonycloud.cn/yeyazhou/go-httpserver/pkg/config"
)

type API struct {
	// Protect against config, template and http client
	mtx sync.RWMutex

	conf       *config.Config
	httpClient *http.Client
	logger     log.Logger
}

func NewAPI(logger log.Logger) *API {
	return &API{
		logger: logger,
	}
}

func (api *API) Update(conf *config.Config) {
	api.mtx.Lock()
	defer api.mtx.Unlock()

	api.conf = conf
	api.httpClient = &http.Client{
		Transport: &http.Transport{
			Proxy:             http.ProxyFromEnvironment,
			DisableKeepAlives: true,
		},
	}
}

func (api *API) Routes() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestLogger(&chilog.KitLogger{Logger: api.logger}))
	router.Use(middleware.Recoverer)
	router.Get("/hello", api.helloWorld)
	return router
}

func (api *API) helloWorld(w http.ResponseWriter, r *http.Request) {
	// api.mtx.RLock()
	// conf := api.conf
	// httpClient := api.httpClient
	// api.mtx.RUnlock()

	logger := log.With(api.logger, "api", "helloWorld")

	level.Error(logger).Log("msg", "Received a Get request from /hello")

	io.WriteString(w, "OK")
}
