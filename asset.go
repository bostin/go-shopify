package goshopify

import (
	"fmt"
	"time"
)

const assetsBasePath = "themes"

// AssetService is an interface for interfacing with the asset endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/asset
type AssetService interface {
	List(int64, interface{}) ([]Asset, error)
	Get(int64, string) (*Asset, error)
	Update(int64, Asset) (*Asset, error)
	Delete(int64, string) error
}

// AssetServiceOp handles communication with the asset related methods of
// the Shopify API.
type AssetServiceOp struct {
	client *Client
}

// Asset represents a Shopify asset
type Asset struct {
	Attachment  string     `json:"attachment" bson:"attachment"`
	ContentType string     `json:"content_type" bson:"content_type"`
	Key         string     `json:"key" bson:"key"`
	PublicURL   string     `json:"public_url" bson:"public_url"`
	Size        int        `json:"size" bson:"size"`
	SourceKey   string     `json:"source_key" bson:"source_key"`
	Src         string     `json:"src" bson:"src"`
	ThemeID     int64      `json:"theme_id" bson:"theme_id"`
	Value       string     `json:"value" bson:"value"`
	CreatedAt   *time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at" bson:"updated_at"`
}

// AssetResource is the result from the themes/x/assets.json?asset[key]= endpoint
type AssetResource struct {
	Asset *Asset `json:"asset" bson:"asset"`
}

// AssetsResource is the result from the themes/x/assets.json endpoint
type AssetsResource struct {
	Assets []Asset `json:"assets" bson:"assets"`
}

type assetGetOptions struct {
	Key     string `url:"asset[key]" bson:"key"`
	ThemeID int64  `url:"theme_id" bson:"theme_id"`
}

// List the metadata for all assets in the given theme
func (s *AssetServiceOp) List(themeID int64, options interface{}) ([]Asset, error) {
	path := fmt.Sprintf("%s/%d/assets.json", assetsBasePath, themeID)
	resource := new(AssetsResource)
	err := s.client.Get(path, resource, options)
	return resource.Assets, err
}

// Get an asset by key from the given theme
func (s *AssetServiceOp) Get(themeID int64, key string) (*Asset, error) {
	path := fmt.Sprintf("%s/%d/assets.json", assetsBasePath, themeID)
	options := assetGetOptions{
		Key:     key,
		ThemeID: themeID,
	}
	resource := new(AssetResource)
	err := s.client.Get(path, resource, options)
	return resource.Asset, err
}

// Update an asset
func (s *AssetServiceOp) Update(themeID int64, asset Asset) (*Asset, error) {
	path := fmt.Sprintf("%s/%d/assets.json", assetsBasePath, themeID)
	wrappedData := AssetResource{Asset: &asset}
	resource := new(AssetResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Asset, err
}

// Delete an asset
func (s *AssetServiceOp) Delete(themeID int64, key string) error {
	path := fmt.Sprintf("%s/%d/assets.json?asset[key]=%s", assetsBasePath, themeID, key)
	return s.client.Delete(path)
}
