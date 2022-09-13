package goshopify

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const applicationChargesBasePath = "application_charges"

// ApplicationChargeService is an interface for interacting with the
// ApplicationCharge endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/billing/applicationcharge
type ApplicationChargeService interface {
	Create(ApplicationCharge) (*ApplicationCharge, error)
	Get(int64, interface{}) (*ApplicationCharge, error)
	List(interface{}) ([]ApplicationCharge, error)
	Activate(ApplicationCharge) (*ApplicationCharge, error)
}

type ApplicationChargeServiceOp struct {
	client *Client
}

type ApplicationCharge struct {
	ID                 int64            `json:"id" bson:"id"`
	Name               string           `json:"name" bson:"name"`
	APIClientID        int64            `json:"api_client_id" bson:"api_client_id"`
	Price              *decimal.Decimal `json:"price" bson:"price"`
	Status             string           `json:"status" bson:"status"`
	ReturnURL          string           `json:"return_url" bson:"return_url"`
	Test               *bool            `json:"test" bson:"test"`
	CreatedAt          *time.Time       `json:"created_at" bson:"created_at"`
	UpdatedAt          *time.Time       `json:"updated_at" bson:"updated_at"`
	ChargeType         *string          `json:"charge_type" bson:"charge_type"`
	DecoratedReturnURL string           `json:"decorated_return_url" bson:"decorated_return_url"`
	ConfirmationURL    string           `json:"confirmation_url" bson:"confirmation_url"`
}

// ApplicationChargeResource represents the result from the
// admin/application_charges{/X{/activate.json}.json}.json endpoints.
type ApplicationChargeResource struct {
	Charge *ApplicationCharge `json:"application_charge" bson:"application_charge"`
}

// ApplicationChargesResource represents the result from the
// admin/application_charges.json endpoint.
type ApplicationChargesResource struct {
	Charges []ApplicationCharge `json:"application_charges" bson:"application_charges"`
}

// Create creates new application charge.
func (a ApplicationChargeServiceOp) Create(charge ApplicationCharge) (*ApplicationCharge, error) {
	path := fmt.Sprintf("%s.json", applicationChargesBasePath)
	resource := &ApplicationChargeResource{}
	return resource.Charge, a.client.Post(path, ApplicationChargeResource{Charge: &charge}, resource)
}

// Get gets individual application charge.
func (a ApplicationChargeServiceOp) Get(chargeID int64, options interface{}) (*ApplicationCharge, error) {
	path := fmt.Sprintf("%s/%d.json", applicationChargesBasePath, chargeID)
	resource := &ApplicationChargeResource{}
	return resource.Charge, a.client.Get(path, resource, options)
}

// List gets all application charges.
func (a ApplicationChargeServiceOp) List(options interface{}) ([]ApplicationCharge, error) {
	path := fmt.Sprintf("%s.json", applicationChargesBasePath)
	resource := &ApplicationChargesResource{}
	return resource.Charges, a.client.Get(path, resource, options)
}

// Activate activates application charge.
func (a ApplicationChargeServiceOp) Activate(charge ApplicationCharge) (*ApplicationCharge, error) {
	path := fmt.Sprintf("%s/%d/activate.json", applicationChargesBasePath, charge.ID)
	resource := &ApplicationChargeResource{}
	return resource.Charge, a.client.Post(path, ApplicationChargeResource{Charge: &charge}, resource)
}
