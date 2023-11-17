package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/eleniums/mining-post/server"
	"github.com/go-chi/chi"
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
	flag.StringVar(&logLevel, "log-level", getEnvStr("LOG_LEVEL", "info"), "LOG_LEVEL: level to use for logs (debug|info|warn|error)")
	flag.StringVar(&logFormat, "log-format", getEnvStr("LOG_FORMAT", "text"), "LOG_FORMAT: format to use for logs (text|json)")
	flag.Parse()

	initLog(logLevel, logFormat)

	slog.Info("Starting service", "service", serviceName, "version", version)

	// debug print all flags with values
	flags := []any{}
	flag.VisitAll(func(f *flag.Flag) {
		flags = append(flags, slog.String(f.Name, f.Value.String()))
	})
	slog.Debug("flags", flags...)

	// create the server
	srv := server.NewServer()

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
	r.Route("/player", func(r chi.Router) {
		r.Get("/inventory", srv.PlayerListInventory)
	})
	r.Route("/market", func(r chi.Router) {
		r.Get("/stock", srv.MarketListStock)
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
		slog.Info("Adding TLS credentials (HTTPS)")

		tlsConfig, err := createTLSConfig(certFile, keyFile)
		if err != nil {
			slog.Error("Failed to load TLS credentials", server.ErrAttr(err))
			os.Exit(1)
		}

		// start listening
		lis, err = tls.Listen("tcp", fmt.Sprintf("%s:%s", httpHost, httpPort), tlsConfig)
		if err != nil {
			slog.Error("Failed to listen", server.ErrAttr(err))
			os.Exit(1)
		}

		slog.Info("Listening for HTTPS requests", "address", lis.Addr().String())
	} else {
		slog.Info("Using insecure endpoint (HTTP - no TLS)")

		// start listening
		var err error
		lis, err = net.Listen("tcp", fmt.Sprintf("%s:%s", httpHost, httpPort))
		if err != nil {
			slog.Error("Failed to listen", server.ErrAttr(err))
			os.Exit(1)
		}

		slog.Info("Listening for HTTP requests", "address", lis.Addr().String())
	}

	var wg sync.WaitGroup
	onShutdown(func() {
		wg.Add(1)
		defer wg.Done()
		slog.Info("Shutting down the server...")
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(60*time.Second))
		defer cancel()
		err := s.Shutdown(ctx)
		if err != nil {
			slog.Error("Error occurred while shutting down server", server.ErrAttr(err))
			return
		}
		slog.Info("Server shutdown successfully")
	})

	// start the server
	err := s.Serve(lis)
	if err != nil && err != http.ErrServerClosed {
		slog.Error("Error occurred while serving", server.ErrAttr(err))
		os.Exit(1)
	}

	// wait for shutdown to finish
	wg.Wait()
}

// initLog will initialize the logger.
func initLog(level string, format string) {
	// parse log level
	logLevel := &slog.LevelVar{}
	switch strings.ToLower(level) {
	case "debug":
		logLevel.Set(slog.LevelDebug)
	case "info":
		logLevel.Set(slog.LevelInfo)
	case "warn":
		logLevel.Set(slog.LevelWarn)
	case "error":
		logLevel.Set(slog.LevelError)
	default:
		logLevel.Set(slog.LevelInfo)
	}

	// parse log format
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	var logHandler slog.Handler
	switch strings.ToLower(format) {
	case "json":
		logHandler = slog.NewJSONHandler(os.Stdout, opts)
	case "text":
		logHandler = slog.NewTextHandler(os.Stdout, opts)
	default:
		logHandler = slog.NewTextHandler(os.Stdout, opts)
	}

	slog.SetDefault(slog.New(logHandler))
}

// createTLSConfig will load a cert/key pair and return a TLS config.
func createTLSConfig(certFile string, keyFile string) (*tls.Config, error) {
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
