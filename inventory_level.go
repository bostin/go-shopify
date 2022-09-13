package goshopify

import (
	"fmt"
	"time"
)

const inventoryLevelsBasePath = "inventory_levels"

type InventoryLevelListOptions struct {
	InventoryItemIds *string    `json:"inventory_item_ids"  url:"inventory_item_ids,omitempty"`
	LocationIds      *string    `json:"location_ids" url:"location_ids,omitempty"`
	Limit            *int       `json:"limit" url:"limit,omitempty"`
	UpdatedAtMin     *time.Time `json:"updated_at_min" url:"updated_at_min,omitempty"`
}

// InventoryLevelService is an interface for interacting with the
// inventory levels endpoints of the Shopify API
// See https://help.shopify.com/en/api/reference/inventory/inventorylevel
type InventoryLevelService interface {
	List(interface{}) ([]InventoryLevel, error)
	Adjust(adjust InventoryLevelAdjust) (*InventoryLevel, error)
	Connect(connect InventoryLevelConnect) (*InventoryLevel, error)
	Set(level InventoryLevel) (*InventoryLevel, error)
	Delete(connect InventoryLevelConnect) error
}

// InventoryLevelServiceOp is the default implementation of the InventoryLevelService interface
type InventoryLevelServiceOp struct {
	client *Client
}

// InventoryLevel represents a Shopify inventory level
type InventoryLevel struct {
	InventoryItemID   int64      `json:"inventory_item_id" bson:"inventory_item_id"`
	LocationID        int64      `json:"location_id" bson:"location_id"`
	Available         *int       `json:"available,omitempty" bson:"available,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	AdminGraphqlAPIID string     `json:"admin_graphql_api_id,omitempty" bson:"admin_graphql_api_id,omitempty"`
}

type InventoryLevelAdjust struct {
	LocationID          int64 `json:"location_id" bson:"location_id"`
	InventoryItemID     int64 `json:"inventory_item_id" bson:"inventory_item_id"`
	AvailableAdjustment int   `json:"available_adjustment" bson:"available_adjustment"`
}

type InventoryLevelConnect struct {
	LocationID      int64 `json:"location_id" bson:"location_id"`
	InventoryItemID int64 `json:"inventory_item_id" bson:"inventory_item_id"`
}

// InventoryLevelResource is used for handling single level requests and responses
type InventoryLevelResource struct {
	InventoryLevel *InventoryLevel `json:"inventory_level" bson:"inventory_level"`
}

// InventoryLevelsResource is used for handling multiple level responses
type InventoryLevelsResource struct {
	InventoryLevels []InventoryLevel `json:"inventory_levels" bson:"inventory_levels"`
}

// List inventory levels
func (s *InventoryLevelServiceOp) List(options interface{}) ([]InventoryLevel, error) {
	path := fmt.Sprintf("%s.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelsResource)
	err := s.client.Get(path, resource, options)
	return resource.InventoryLevels, err
}

// Adjust the inventory level of an inventory item at a location
func (s *InventoryLevelServiceOp) Adjust(adjust InventoryLevelAdjust) (*InventoryLevel, error) {
	path := fmt.Sprintf("%s/adjust.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelResource)
	err := s.client.Post(path, adjust, resource)
	return resource.InventoryLevel, err
}

// Connect an inventory item to a location
func (s *InventoryLevelServiceOp) Connect(connect InventoryLevelConnect) (*InventoryLevel, error) {
	path := fmt.Sprintf("%s/connect.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelResource)
	err := s.client.Post(path, connect, resource)
	return resource.InventoryLevel, err
}

// Set the inventory level for an inventory item at a location
func (s *InventoryLevelServiceOp) Set(level InventoryLevel) (*InventoryLevel, error) {
	path := fmt.Sprintf("%s/set.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelResource)
	err := s.client.Post(path, level, resource)
	return resource.InventoryLevel, err
}

// Delete an inventory level from a location
func (s *InventoryLevelServiceOp) Delete(connect InventoryLevelConnect) error {
	return s.client.Delete(fmt.Sprintf("%s.json?inventory_item_id=%d&location_id=%d", inventoryLevelsBasePath, connect.InventoryItemID, connect.LocationID))
}
