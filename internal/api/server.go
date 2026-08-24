package api

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"
)

// NewMux builds the HTTP router.
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	NewRoutes(mux)
	return mux
}

// NewServer builds an HTTP server with sane timeouts.
func NewServer(addr string) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           NewMux(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
}

// Listen starts the blocking HTTP service.
func Listen(addr string) error {
	srv := NewServer(addr)
	return srv.ListenAndServe()
}

// ListenWithContext starts the service and stops on context cancellation.
func ListenWithContext(ctx context.Context, addr string) error {
	srv := NewServer(addr)
	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.ListenAndServe()
	}()
	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdownCtx)
		return nil
	case err := <-errCh:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}

// IsPortAvailable reports whether an address can be bound.
func IsPortAvailable(addr string) bool {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return false
	}
	_ = l.Close()
	return true
}
