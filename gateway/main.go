package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/olzh2102/toll-calculator/aggregator/client"
	"github.com/sirupsen/logrus"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func main() {
	listenAddr := flag.String("listenAddr", ":6000", "")
	client := client.NewHTTPClient("http://localhost:3000")
	invoiceHandler := newInvoiceHandler(client)
	http.HandleFunc("/invoice", makeAPIFunc(invoiceHandler.handleGetInvoice))
	logrus.Info("gateway HTTP server running on port: ", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

type InvoiceHandler struct {
	client client.Client
}

func newInvoiceHandler(c client.Client) *InvoiceHandler {
	return &InvoiceHandler{
		client: c,
	}
}

func (h *InvoiceHandler) handleGetInvoice(w http.ResponseWriter, r *http.Request) error {
	// * access to agg client
	inv, err := h.client.GetInvoice(context.Background(), 8352150199983364412)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, inv)
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func makeAPIFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}
