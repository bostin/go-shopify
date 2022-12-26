package goshopify

import (
	"fmt"
	"time"
)

const collectsBasePath = "collects"

type CollectListOptions struct {
	Fields       *string `json:"fields,omitempty" url:"fields,omitempty" bson:"fields,omitempty"`
	Limit        *int    `json:"limit,omitempty" url:"limit,omitempty" bson:"limit,omitempty"`
	SinceID      *int64  `json:"since_id,omitempty" url:"since_id,omitempty" bson:"since_id,omitempty"`
	ProductID    *string `json:"product_id,omitempty" url:"product_id,omitempty" bson:"product_id,omitempty"`
	CollectionID *string `json:"collection_id,omitempty" url:"collection_id,omitempty" bson:"collection_id,omitempty"`
}

// CollectService is an interface for interfacing with the collect endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/products/collect
type CollectService interface {
	List(interface{}) ([]Collect, error)
	Count(interface{}) (int, error)
}

// CollectServiceOp handles communication with the collect related methods of
// the Shopify API.
type CollectServiceOp struct {
	client *Client
}

// Collect represents a Shopify collect
type Collect struct {
	ID           int64      `json:"id,omitempty" bson:"id,omitempty"`
	CollectionID int64      `json:"collection_id,omitempty" bson:"collection_id,omitempty"`
	ProductID    int64      `json:"product_id,omitempty" bson:"product_id,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Position     int        `json:"position,omitempty" bson:"position,omitempty"`
	SortValue    string     `json:"sort_value,omitempty" bson:"sort_value,omitempty"`
	// Featured     bool       `json:"featured,omitempty" bson:"featured,omitempty"`
}

// Represents the result from the collects/X.json endpoint
type CollectResource struct {
	Collect *Collect `json:"collect" bson:"collect"`
}

// Represents the result from the collects.json endpoint
type CollectsResource struct {
	Collects []Collect `json:"collects" bson:"collects"`
}

// List collects
func (s *CollectServiceOp) List(options interface{}) ([]Collect, error) {
	path := fmt.Sprintf("%s.json", collectsBasePath)
	resource := new(CollectsResource)
	err := s.client.Get(path, resource, options)
	return resource.Collects, err
}

// Count collects
func (s *CollectServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", collectsBasePath)
	return s.client.Count(path, options)
}
