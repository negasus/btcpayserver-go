package btcpayserver

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/negasus/btcpayserver-go/models"
)

type WebhookHandler func(event *models.WebhookEvent)

func CheckSigHeader(sig string, secret string, body []byte) bool {
	if !strings.HasPrefix(sig, "sha256=") {
		return false
	}

	sig = strings.TrimPrefix(sig, "sha256=")

	h := hmac.New(sha256.New, []byte(secret))
	h.Write(body)

	return fmt.Sprintf("%x", h.Sum(nil)) == sig
}

func (c *Webhook) handler(secret string, handler WebhookHandler) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		sig := req.Header.Get("Btcpay-Sig")
		if sig == "" {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		body, errBody := io.ReadAll(req.Body)
		if errBody != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !CheckSigHeader(sig, secret, body) {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}

		e := models.WebhookEvent{}
		err := json.Unmarshal(body, &e)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		handler(&e)
	}
}

func (c *Webhook) Serve(ctx context.Context, ln net.Listener, secret string, handler WebhookHandler) error {
	server := http.Server{
		Handler: c.handler(secret, handler),
	}

	go func() {
		<-ctx.Done()
		server.Shutdown(context.Background()) // nolint: errcheck
	}()

	errServe := server.Serve(ln)
	if errServe != nil {
		if !errors.Is(errServe, http.ErrServerClosed) {
			return errServe
		}
	}

	return nil
}
