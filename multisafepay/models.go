package multisafepay

const (
	OrderStatusCancelled   OrderStatus = "cancelled"
	OrderStatusCompleted   OrderStatus = "completed"
	OrderStatusDeclined    OrderStatus = "declined"
	OrderStatusExpired     OrderStatus = "expired"
	OrderStatusInitialized OrderStatus = "initialized"
	OrderStatusRefunded    OrderStatus = "refunded"
	OrderStatusReserved    OrderStatus = "reserved"
	OrderStatusShipped     OrderStatus = "shipped"
	OrderStatusUncleared   OrderStatus = "uncleared"
	OrderStatusVoid        OrderStatus = "void"
)

// Response from the API, contains a boolean to indicate success
type Response struct {
	Success bool `json:"success"`
}

// ErrorResponse contains information about errors as reported by the API
type ErrorResponse struct {
	Response

	Data      interface{} `json:"data,omitempty"`
	ErrorCode int         `json:"error_code"`
	ErrorInfo string      `json:"error_info"`
}

// PaymentOptions structure, see: https://docs.multisafepay.com/api/#payment-option-object
type PaymentOptions struct {
	NotificationURL    string `json:"notification_url,omitempty"`
	NotificationMethod string `json:"notification_method,omitempty"`
	RedirectURL        string `json:"redirect_url,omitempty"`
	CancelURL          string `json:"cancel_url,omitempty"`
	CloseWindow        string `json:"close_window,omitempty"`
}

// Customer structure, see: https://docs.multisafepay.com/api/#customer-object
type Customer struct {
	Locale      string `json:"locale,omitempty"`
	IPAddress   string `json:"ip_address,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Company     string `json:"company,omitempty"`
	Address1    string `json:"address1,omitempty"`
	HouseNumber string `json:"house_number,omitempty"`
	ZIPCode     string `json:"zip_code,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
	Referrer    string `json:"referrer,omitempty"`
	UserAgent   string `json:"user_agent,omitempty"`
}

// SecondChange for in Order
type SecondChance struct {
	SendEmail bool `json:"send_email"`
}

// Order structure, see: https://docs.multisafepay.com/api/#orders
type Order struct {
	Type           string          `json:"type,omitempty"`
	OrderID        ID              `json:"order_id,omitempty"`
	Gateway        string          `json:"gateway,omitempty"`
	Currency       string          `json:"currency,omitempty"`
	Amount         int             `json:"amount,omitempty"`
	Description    string          `json:"description,omitempty"`
	PaymentOptions *PaymentOptions `json:"payment_options,omitempty"`
	Customer       *Customer       `json:"customer,omitempty"`
	SecondChance   *SecondChance   `json:"second_chance"`
}

// PostOrderResponseData is the Data field of PostOrderResponse
type PostOrderResponseData struct {
	OrderID    ID     `json:"order_id"`
	PaymentURL string `json:"payment_url,omitempty"`
}

// PostOrderResponse is a response to POST /orders
// Documentation: https://docs.multisafepay.com/api/#orders
type PostOrderResponse struct {
	Response
	Data PostOrderResponseData `json:"data"`
}

// Cost model (see GetOrderResponse)
type Cost struct {
	TransactionID ID      `json:"transaction_id"`
	Description   string  `json:"description"`
	Type          string  `json:"type"`
	Status        string  `json:"status"`
	Created       Time    `json:"created"`
	Amount        float64 `json:"amount"`
}

// RelatedTransaction (see GetOrderResponse)
type RelatedTransaction struct {
	Amount        int    `json:"amount"`
	Costs         []Cost `json:"costs"`
	Created       Time   `json:"created"`
	Currency      string `json:"currency"`
	Description   string `json:"description"`
	Modified      Time   `json:"modified"`
	Status        string `json:"status"`
	TransactionID ID     `json:"transaction_id"`
}

// GetOrderResponseData is the Data field of GetOrderResponse
type GetOrderResponseData struct {
	TransactionID       ID                       `json:"transaction_id"`
	OrderID             ID                       `json:"order_id"`
	Created             Time                     `json:"created"`
	Currency            string                   `json:"currency"`
	Amount              int                      `json:"amount"`
	Description         string                   `json:"description"`
	AmountRefunded      int                      `json:"amount_refunded"`
	Status              OrderStatus              `json:"status,omitempty"`
	FinancialStatus     string                   `json:"financial_status"`
	Reason              string                   `json:"reason"`
	ReasonCode          string                   `json:"reason_code"`
	FastCheckout        string                   `json:"fastcheckout"`
	Modified            Time                     `json:"modified"`
	Customer            *Customer                `json:"customer"`
	PaymentDetails      map[string]interface{}   `json:"payment_details"`
	Costs               []Cost                   `json:"costs"`
	RelatedTransactions []RelatedTransaction     `json:"related_transactions"`
	PaymentMethods      []map[string]interface{} `json:"payment_methods"`
}

// The OrderStatus type is used to represent the status of an order
// Documentation: https://docs.multisafepay.com/reference/getorder
type OrderStatus string

// GetOrderResponse is a response to GET /orders/{order_id}
// Documentation: https://docs.multisafepay.com/api/#retrieve-an-order
type GetOrderResponse struct {
	Response
	Data GetOrderResponseData
}
