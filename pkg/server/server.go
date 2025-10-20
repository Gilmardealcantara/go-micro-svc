package server

//nolint:staticcheck
import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/log"
)

type ShutdownFunc func() error

type Server struct {
	logger  log.Loggable
	svr     *http.Server
	errC    chan error
	cleanup []func() error
}

// Graceful shutdown will need to be implemented if you want your service handle abrupt shutdown
// and perform cleanup tasks before exiting (ie. closing DB connections, flushing sentry, etc.)
// K8s pods will send a SIGTERM signal to the container when it is being terminated, so this listener
// will catch that signal and perform the cleanup tasks before exiting.

// You can read more on that k8s termination best practices here: https://cloud.google.com/blog/products/gcp/kubernetes-best-practices-terminating-with-grace

// You can read about graceful shutdown for Go here:
// https://mariocarrion.com/2021/05/21/golang-microservices-graceful-shutdown.html
// https://www.rudderstack.com/blog/implementing-graceful-shutdown-in-go/

// This is a sample implementation, so feel free to modify it to fit your needs
func (s *Server) ShutdownListener(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ctx.Done()
	s.logger.Info(ctx, "Shutdown signal received, initiating graceful shutdown")
	// listen to messages on error channel
	if err := s.err(); err != nil {
		s.svcCleanup(ctx)

		s.logger.Fatal(ctx, "Server failed to shutdown gracefully", "err", err)
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// This handles the case where the server shutdown
	// exceeds the deadline and _hard_ exits the program if so.
	// Note that If the provided context expires before the shutdown is complete,
	// shutdown returns the context's error.
	err := s.svr.Shutdown(ctxTimeout)
	if err != nil {
		s.logger.Error(ctxTimeout, fmt.Sprintf("Server failed to shutdown gracefully: %s", err))
		s.errC <- err
	}

	// Blocking until either the context is done or the timeout is reached.
	// Context would be done if srv.Shutdown occurs before the timeout.
	//nolint:gosimple // S1000 ignore this!
	select {
	case <-ctxTimeout.Done():
		if errors.Is(ctxTimeout.Err(), context.DeadlineExceeded) {
			s.svcCleanup(ctxTimeout)
			s.logger.Fatal(ctxTimeout, "graceful shutdown timed out, forcing exit")
		}
	}

	s.logger.Info(ctx, "Server shutdown gracefully")
	s.svcCleanup(ctx)
}

func (s *Server) Start(ctx context.Context) {
	s.logger.Info(ctx, fmt.Sprintf("Listening on %s", s.svr.Addr))
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go s.ShutdownListener(ctx, wg)

	if err := s.svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		s.logger.Error(ctx, fmt.Sprintf("Server failed to start: %s", err))
		// It's ok to exit early here since the server failed to start.
		// Note that the ShutdownListener above will continue running, but since the
		// healthchecks will fail, k8s will send SIGTERM to the pod,
		// which will initiate the ShutdownListener.
		return
	}

	wg.Wait()
}

func (s *Server) svcCleanup(ctx context.Context) {
	for _, cleaner := range s.cleanup {
		err := cleaner()
		if err != nil {
			s.logger.Error(ctx, fmt.Sprintf("Failed to cleanup service: %s", err))
		}
	}
}

func (s *Server) err() error {
	select {
	case err := <-s.errC:
		return err
	default:
		return nil
	}
}

type ServiceShutdownFuncOpts func(*Server)

func WithAdditionalShutdown(f ShutdownFunc) ServiceShutdownFuncOpts {
	return func(s *Server) {
		s.cleanup = append(s.cleanup, f)
	}
}

func New(
	logger *log.Logger,
	svr *http.Server,
	otherShutdownFuncs ...ServiceShutdownFuncOpts,
) Server {
	s := Server{
		svr:    svr,
		logger: logger,
	}

	for _, f := range otherShutdownFuncs {
		f(&s)
	}

	return s
}
