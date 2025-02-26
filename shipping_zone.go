package goshopify

import (
	"github.com/shopspring/decimal"
)

// ShippingZoneService is an interface for interfacing with the shipping zones endpoint
// of the Shopify API.
// See: https://help.shopify.com/api/reference/store-properties/shippingzone
type ShippingZoneService interface {
	List() ([]ShippingZone, error)
}

// ShippingZoneServiceOp handles communication with the shipping zone related methods
// of the Shopify API.
type ShippingZoneServiceOp struct {
	client *Client
}

// ShippingZone represents a Shopify shipping zone
type ShippingZone struct {
	ID                           int64                         `json:"id,omitempty" bson:"id,omitempty"`
	Name                         string                        `json:"name,omitempty" bson:"name,omitempty"`
	ProfileID                    string                        `json:"profile_id,omitempty" bson:"profile_id,omitempty"`
	LocationGroupID              string                        `json:"location_group_id,omitempty" bson:"location_group_id,omitempty"`
	AdminGraphqlAPIID            string                        `json:"admin_graphql_api_id,omitempty" bson:"admin_graphql_api_id,omitempty"`
	Countries                    []ShippingCountry             `json:"countries,omitempty" bson:"countries,omitempty"`
	WeightBasedShippingRates     []WeightBasedShippingRate     `json:"weight_based_shipping_rates,omitempty" bson:"weight_based_shipping_rates,omitempty"`
	PriceBasedShippingRates      []PriceBasedShippingRate      `json:"price_based_shipping_rates,omitempty" bson:"price_based_shipping_rates,omitempty"`
	CarrierShippingRateProviders []CarrierShippingRateProvider `json:"carrier_shipping_rate_providers,omitempty" bson:"carrier_shipping_rate_providers,omitempty"`
}

// ShippingCountry represents a Shopify shipping country
type ShippingCountry struct {
	ID             int64              `json:"id,omitempty" bson:"id,omitempty"`
	ShippingZoneID int64              `json:"shipping_zone_id,omitempty" bson:"shipping_zone_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Tax            *decimal.Decimal   `json:"tax,omitempty" bson:"tax,omitempty"`
	Code           string             `json:"code,omitempty" bson:"code,omitempty"`
	TaxName        string             `json:"tax_name,omitempty" bson:"tax_name,omitempty"`
	Provinces      []ShippingProvince `json:"provinces,omitempty" bson:"provinces,omitempty"`
}

// ShippingProvince represents a Shopify shipping province
type ShippingProvince struct {
	ID             int64            `json:"id,omitempty" bson:"id,omitempty"`
	CountryID      int64            `json:"country_id,omitempty" bson:"country_id,omitempty"`
	ShippingZoneID int64            `json:"shipping_zone_id,omitempty" bson:"shipping_zone_id,omitempty"`
	Name           string           `json:"name,omitempty" bson:"name,omitempty"`
	Code           string           `json:"code,omitempty" bson:"code,omitempty"`
	Tax            *decimal.Decimal `json:"tax,omitempty" bson:"tax,omitempty"`
	TaxName        string           `json:"tax_name,omitempty" bson:"tax_name,omitempty"`
	TaxType        string           `json:"tax_type,omitempty" bson:"tax_type,omitempty"`
	TaxPercentage  *decimal.Decimal `json:"tax_percentage,omitempty" bson:"tax_percentage,omitempty"`
}

// WeightBasedShippingRate represents a Shopify weight-constrained shipping rate
type WeightBasedShippingRate struct {
	ID             int64            `json:"id,omitempty" bson:"id,omitempty"`
	ShippingZoneID int64            `json:"shipping_zone_id,omitempty" bson:"shipping_zone_id,omitempty"`
	Name           string           `json:"name,omitempty" bson:"name,omitempty"`
	Price          *decimal.Decimal `json:"price,omitempty" bson:"price,omitempty"`
	WeightLow      *decimal.Decimal `json:"weight_low,omitempty" bson:"weight_low,omitempty"`
	WeightHigh     *decimal.Decimal `json:"weight_high,omitempty" bson:"weight_high,omitempty"`
}

// PriceBasedShippingRate represents a Shopify subtotal-constrained shipping rate
type PriceBasedShippingRate struct {
	ID               int64            `json:"id,omitempty" bson:"id,omitempty"`
	ShippingZoneID   int64            `json:"shipping_zone_id,omitempty" bson:"shipping_zone_id,omitempty"`
	Name             string           `json:"name,omitempty" bson:"name,omitempty"`
	Price            *decimal.Decimal `json:"price,omitempty" bson:"price,omitempty"`
	MinOrderSubtotal *decimal.Decimal `json:"min_order_subtotal,omitempty" bson:"min_order_subtotal,omitempty"`
	MaxOrderSubtotal *decimal.Decimal `json:"max_order_subtotal,omitempty" bson:"max_order_subtotal,omitempty"`
}

// CarrierShippingRateProvider represents a Shopify carrier-constrained shipping rate
type CarrierShippingRateProvider struct {
	ID               int64             `json:"id,omitempty" bson:"id,omitempty"`
	CarrierServiceID int64             `json:"carrier_service_id,omitempty" bson:"carrier_service_id,omitempty"`
	ShippingZoneID   int64             `json:"shipping_zone_id,omitempty" bson:"shipping_zone_id,omitempty"`
	FlatModifier     *decimal.Decimal  `json:"flat_modifier,omitempty" bson:"flat_modifier,omitempty"`
	PercentModifier  *decimal.Decimal  `json:"percent_modifier,omitempty" bson:"percent_modifier,omitempty"`
	ServiceFilter    map[string]string `json:"service_filter,omitempty" bson:"service_filter,omitempty"`
}

// Represents the result from the shipping_zones.json endpoint
type ShippingZonesResource struct {
	ShippingZones []ShippingZone `json:"shipping_zones" bson:"shipping_zones"`
}

// List shipping zones
func (s *ShippingZoneServiceOp) List() ([]ShippingZone, error) {
	resource := new(ShippingZonesResource)
	err := s.client.Get("shipping_zones.json", resource, nil)
	return resource.ShippingZones, err
}
