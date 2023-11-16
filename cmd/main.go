package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/eleniums/mining-post/mem"
	"github.com/eleniums/mining-post/server"
	"github.com/go-chi/chi"

	log "github.com/sirupsen/logrus"
)

var (
	serviceName = "mining-post"
	version     = "dev"
)

var (
	httpHost  string
	httpPort  string
	dbURI     string
	certFile  string
	keyFile   string
	logLevel  string
	logFormat string
)

func main() {
	// parse configuration - flags have highest priority, then env vars, and then defaults
	flag.StringVar(&httpHost, "http-host", getEnvStr("HTTP_HOST", "127.0.0.1"), "HTTP_HOST: host to serve endpoint on")
	flag.StringVar(&httpPort, "http-port", getEnvStr("HTTP_PORT", "9090"), "HTTP_PORT: port to serve endpoint on")
	flag.StringVar(&dbURI, "db-uri", getEnvStr("DB_URI", ""), "DB_URI: connection string to database")
	flag.StringVar(&certFile, "tls-cert-file", getEnvStr("TLS_CERT_FILE", ""), "TLS_CERT_FILE: cert file for enabling a TLS connection")
	flag.StringVar(&keyFile, "tls-key-file", getEnvStr("TLS_KEY_FILE", ""), "TLS_KEY_FILE: key file for enabling a TLS connection")
	flag.StringVar(&logLevel, "log-level", getEnvStr("LOG_LEVEL", "info"), "LOG_LEVEL: level to use for logs (debug|info|warn|error|fatal|panic)")
	flag.StringVar(&logFormat, "log-format", getEnvStr("LOG_FORMAT", "text"), "LOG_FORMAT: format to use for logs (text|json)")
	flag.Parse()

	initLog(logLevel, logFormat)

	log.Infof("%s - version: %s", serviceName, version)

	// debug print all flags with values
	flag.VisitAll(func(f *flag.Flag) {
		log.Debugf("%s=%s", f.Name, f.Value.String())
	})

	// create item storage accessor
	log.Info("Using in-memory cache for data storage")
	db := mem.NewCache()

	// create the server
	srv := server.NewServer(db)

	// serve the endpoint
	serve(srv)
}

func serve(srv *server.Server) {
	// create router
	r := chi.NewRouter()

	// add middleware
	r.Use(server.Audit)

	// register handlers
	r.Get("/ping", srv.Ping)
	r.Route("/items", func(r chi.Router) {
		r.Get("/", srv.GetItems)
		r.Post("/", srv.InsertItem)
		r.Route("/{itemID}", func(r chi.Router) {
			r.Get("/", srv.GetItemByID)
			r.Put("/", srv.UpdateItem)
			r.Delete("/", srv.DeleteItem)
		})
	})

	// create http server
	s := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
		Handler:      r,
	}

	// determine if TLS should be used
	var lis net.Listener
	if certFile != "" && keyFile != "" {
		log.Info("Adding TLS credentials (HTTPS)")

		tlsConfig, err := createTLSConfig(certFile, keyFile)
		if err != nil {
			log.WithError(err).Fatal("Failed to load TLS credentials")
		}

		// start listening
		lis, err = tls.Listen("tcp", fmt.Sprintf("%s:%s", httpHost, httpPort), tlsConfig)
		if err != nil {
			log.WithError(err).Fatal("Failed to listen")
		}

		log.Infof("Listening for HTTPS requests at: %v", lis.Addr().String())
	} else {
		log.Info("Using insecure endpoint (HTTP - no TLS)")

		// start listening
		var err error
		lis, err = net.Listen("tcp", fmt.Sprintf("%s:%s", httpHost, httpPort))
		if err != nil {
			log.WithError(err).Fatal("Failed to listen")
		}

		log.Infof("Listening for HTTP requests at: %v", lis.Addr().String())
	}

	var wg sync.WaitGroup
	onShutdown(func() {
		wg.Add(1)
		defer wg.Done()
		log.Infof("Shutting down the server...")
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(60*time.Second))
		defer cancel()
		err := s.Shutdown(ctx)
		if err != nil {
			log.WithError(err).Error("Error occurred while shutting down server")
			return
		}
		log.Infof("Server shutdown successfully")
	})

	// start the server
	err := s.Serve(lis)
	if err != nil && err != http.ErrServerClosed {
		log.WithError(err).Fatal("Error occurred while serving")
	}

	// wait for shutdown to finish
	wg.Wait()
}

// initLog will initialize the logger.
func initLog(level, format string) {
	// use stdout since default for logrus is stderr
	log.SetOutput(os.Stdout)

	// set log level
	l, err := log.ParseLevel(level)
	if err != nil {
		l = log.InfoLevel
	}
	log.SetLevel(l)

	// set log format
	switch strings.ToLower(format) {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	case "text":
	default:
		log.SetFormatter(&log.TextFormatter{})
	}
}

// createTLSConfig will load a cert/key pair and return a TLS config.
func createTLSConfig(certFile, keyFile string) (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		MinVersion: tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		Certificates: []tls.Certificate{cert},
	}, nil
}

// onShutdown will perform an action when the service is terminated.
func onShutdown(action func()) {
	go func() {
		shutdown := make(chan os.Signal, 1)
		// https://www.gnu.org/software/libc/manual/html_node/Termination-Signals.html
		signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
		<-shutdown
		action()
	}()
}

// getEnvStr will return the value for the requested environment variable or a default value.
func getEnvStr(envVar string, defaultValue string) string {
	v := os.Getenv(envVar)
	if v == "" {
		return defaultValue
	}
	return v
}
