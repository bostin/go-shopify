package goshopify

import (
	"time"
)

// ShopService is an interface for interfacing with the shop endpoint of the
// Shopify API.
// See: https://help.shopify.com/api/reference/shop
type ShopService interface {
	Get(options interface{}) (*Shop, error)
}

// ShopServiceOp handles communication with the shop related methods of the
// Shopify API.
type ShopServiceOp struct {
	client *Client
}

// Shop represents a Shopify shop
type Shop struct {
	ID                              int64      `json:"id" bson:"id"`
	Name                            string     `json:"name" bson:"name"`
	ShopOwner                       string     `json:"shop_owner" bson:"shop_owner"`
	Email                           string     `json:"email" bson:"email"`
	CustomerEmail                   string     `json:"customer_email" bson:"customer_email"`
	CreatedAt                       *time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt                       *time.Time `json:"updated_at" bson:"updated_at"`
	Address1                        string     `json:"address1" bson:"address1"`
	Address2                        string     `json:"address2" bson:"address2"`
	City                            string     `json:"city" bson:"city"`
	Country                         string     `json:"country" bson:"country"`
	CountryCode                     string     `json:"country_code" bson:"country_code"`
	CountryName                     string     `json:"country_name" bson:"country_name"`
	Currency                        string     `json:"currency" bson:"currency"`
	Domain                          string     `json:"domain" bson:"domain"`
	Latitude                        float64    `json:"latitude" bson:"latitude"`
	Longitude                       float64    `json:"longitude" bson:"longitude"`
	Phone                           string     `json:"phone" bson:"phone"`
	Province                        string     `json:"province" bson:"province"`
	ProvinceCode                    string     `json:"province_code" bson:"province_code"`
	Zip                             string     `json:"zip" bson:"zip"`
	MoneyFormat                     string     `json:"money_format" bson:"money_format"`
	MoneyInEmailsFormat             string     `json:"money_in_emails_format" bson:"money_in_emails_format"`
	MoneyWithCurrencyFormat         string     `json:"money_with_currency_format" bson:"money_with_currency_format"`
	MoneyWithCurrencyInEmailsFormat string     `json:"money_with_currency_in_emails_format" bson:"money_with_currency_in_emails_format"`
	WeightUnit                      string     `json:"weight_unit" bson:"weight_unit"`
	MyshopifyDomain                 string     `json:"myshopify_domain" bson:"myshopify_domain"`
	PlanName                        string     `json:"plan_name" bson:"plan_name"`
	PlanDisplayName                 string     `json:"plan_display_name" bson:"plan_display_name"`
	PasswordEnabled                 bool       `json:"password_enabled" bson:"password_enabled"`
	PrimaryLocale                   string     `json:"primary_locale" bson:"primary_locale"`
	PrimaryLocationId               int64      `json:"primary_location_id" bson:"primary_location_id"`
	Timezone                        string     `json:"timezone" bson:"timezone"`
	IanaTimezone                    string     `json:"iana_timezone" bson:"iana_timezone"`
	Finances                        bool       `json:"finances,omitempty" bson:"finances,omitempty"` // @deprecated
	ForceSSL                        bool       `json:"force_ssl" bson:"force_ssl"`                   // @deprecated
	TaxShipping                     bool       `json:"tax_shipping" bson:"tax_shipping"`
	TaxesIncluded                   bool       `json:"taxes_included" bson:"taxes_included"`
	HasStorefront                   bool       `json:"has_storefront" bson:"has_storefront"`
	HasDiscounts                    bool       `json:"has_discounts" bson:"has_discounts"`
	HasGiftCards                    bool       `json:"has_gift_cards" bson:"has_gift_cards"`
	SetupRequired                   bool       `json:"setup_required" bson:"setup_require"`
	CountyTaxes                     bool       `json:"county_taxes" bson:"county_taxes"`
	CheckoutAPISupported            bool       `json:"checkout_api_supported" bson:"checkout_api_supported"`
	Source                          string     `json:"source" bson:"source"`
	GoogleAppsDomain                string     `json:"google_apps_domain" bson:"google_apps_domain"`
	GoogleAppsLoginEnabled          bool       `json:"google_apps_login_enabled" bson:"google_apps_login_enabled"`
	EligibleForPayments             bool       `json:"eligible_for_payments" bson:"eligible_for_payments"`
	EligibleForCarReaderGiveaway    bool       `json:"eligible_for_car_reader_giveaway,omitempty" bson:"eligible_for_car_reader_giveaway,omitempty"`
	RequiresExtraPaymentsAgreement  bool       `json:"requires_extra_payments_agreement" bson:"requires_extra_payments_agreement"`
	PreLaunchEnabled                bool       `json:"pre_launch_enabled" bson:"pre_launch_enabled"`
	EnabledPresentmentCurrencies    []string   `json:"enabled_presentment_currencies,omitempty" bson:"enabled_presentment_currencies,omitempty"`
	MultiLocationEnabled            bool       `json:"multi_location_enabled,omitempty" bson:"multi_location_enabled,omitempty"` // @deprecated
	CookieConsentLevel              string     `json:"cookie_consent_level,omitempty" bson:"cookie_consent_level,omitempty"`
	TransactionalSMSDisabled        bool       `json:"transactional_sms_disabled,omitempty" bson:"transactional_sms_disabled,omitempty"`
}

// Represents the result from the admin/shop.json endpoint
type ShopResource struct {
	Shop *Shop `json:"shop" bson:"shop"`
}

// Get shop
func (s *ShopServiceOp) Get(options interface{}) (*Shop, error) {
	resource := new(ShopResource)
	err := s.client.Get("shop.json", resource, options)
	return resource.Shop, err
}
