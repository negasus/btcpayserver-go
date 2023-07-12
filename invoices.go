package btcpayserver

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/negasus/btcpayserver-go/models"
)

type Invoices struct {
	requestFunc requestFunc
}

type InvoiceListOptions struct {
	OrderID    []string
	TextSearch string
	Status     models.InvoiceStatus
	EndDate    time.Time
	Take       int
	Skip       int
	StartDate  time.Time
}

func (o *InvoiceListOptions) values() url.Values {
	v := url.Values{}

	for _, id := range o.OrderID {
		v.Add("orderId", id)
	}
	if o.TextSearch != "" {
		v.Add("textSearch", o.TextSearch)
	}
	if o.Status != "" {
		v.Add("status", string(o.Status))
	}
	if !o.EndDate.IsZero() {
		v.Add("endDate", strconv.Itoa(int(o.EndDate.Unix())))
	}
	if o.Take != 0 {
		v.Add("take", strconv.Itoa(o.Take))
	}
	if o.Skip != 0 {
		v.Add("skip", strconv.Itoa(o.Skip))
	}
	if !o.StartDate.IsZero() {
		v.Add("startDate", strconv.Itoa(int(o.StartDate.Unix())))
	}

	return v
}

// List https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Invoices_GetInvoices
func (c *Invoices) List(storeID string, opts *InvoiceListOptions) ([]*models.Invoice, error) {
	if storeID == "" {
		return nil, fmt.Errorf("storeID is required")
	}

	var res []*models.Invoice

	var query url.Values
	if opts != nil {
		query = opts.values()
	}

	err := c.requestFunc(http.MethodGet, "/api/v1/stores/"+storeID+"/invoices", query, nil, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Create https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Invoices_CreateInvoice
func (c *Invoices) Create(storeID string, invoice models.NewInvoice) (*models.Invoice, error) {
	if storeID == "" {
		return nil, fmt.Errorf("storeID is required")
	}

	res := models.Invoice{}

	err := c.requestFunc(http.MethodPost, "/api/v1/stores/"+storeID+"/invoices", nil, invoice, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Get https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Invoices_GetInvoice
func (c *Invoices) Get(storeID, invoiceID string) (*models.Invoice, error) {
	if storeID == "" {
		return nil, fmt.Errorf("storeID is required")
	}
	if invoiceID == "" {
		return nil, fmt.Errorf("invoiceID is required")
	}

	res := models.Invoice{}

	err := c.requestFunc(http.MethodGet, "/api/v1/stores/"+storeID+"/invoices/"+invoiceID, nil, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Archive https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Invoices_ArchiveInvoice
func (c *Invoices) Archive(storeID, invoiceID string) error {
	if storeID == "" {
		return fmt.Errorf("storeID is required")
	}
	if invoiceID == "" {
		return fmt.Errorf("invoiceID is required")
	}

	return c.requestFunc(http.MethodDelete, "/api/v1/stores/"+storeID+"/invoices/"+invoiceID, nil, nil, nil)
}

// todo: Update https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Invoices_UpdateInvoice

type InvoiceGetPaymentMethodsOptions struct {
	OnlyAccountedPayments bool `json:"onlyAccountedPayments"`
}

func (o *InvoiceGetPaymentMethodsOptions) values() url.Values {
	v := url.Values{}

	v.Add("onlyAccountedPayments", strconv.FormatBool(o.OnlyAccountedPayments))

	return v
}

// GetPaymentMethods https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Invoices_GetInvoicePaymentMethods
func (c *Invoices) GetPaymentMethods(storeID, invoiceID string, opts *InvoiceGetPaymentMethodsOptions) ([]*models.PaymentMethod, error) {
	if storeID == "" {
		return nil, fmt.Errorf("storeID is required")
	}
	if invoiceID == "" {
		return nil, fmt.Errorf("invoiceID is required")
	}

	var query url.Values
	if opts != nil {
		query = opts.values()
	}

	var res []*models.PaymentMethod

	err := c.requestFunc(http.MethodGet, "/api/v1/stores/"+storeID+"/invoices/"+invoiceID+"/payment-methods", query, nil, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// todo: https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Invoices_MarkInvoiceStatus
// todo: https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Invoices_UnarchiveInvoice
// todo: https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Invoices_ActivatePaymentMethod
// todo: https://docs.btcpayserver.org/API/Greenfield/v1/#operation/Invoices_Refund
