package goshopify

import (
	"fmt"
	"net/http"
	"time"
)

const collectionsBasePath = "collections"

// CollectionService is an interface for interfacing with the collection endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/products/collection
type CollectionService interface {
	Get(collectionID int64, options interface{}) (*Collection, error)
	ListProducts(collectionID int64, options interface{}) ([]Product, error)
	ListProductsWithPagination(collectionID int64, options interface{}) ([]Product, *Pagination, error)
}

// CollectionServiceOp handles communication with the collection related methods of
// the Shopify API.
type CollectionServiceOp struct {
	client *Client
}

// Collection represents a Shopify collection
// https://shopify.dev/api/admin-rest/2022-10/resources/collection#get-collections-collection-id
type Collection struct {
	ID             int64      `json:"id" bson:"id"`
	Handle         string     `json:"handle" bson:"handle"`
	Title          string     `json:"title" bson:"title"`
	UpdatedAt      *time.Time `json:"updated_at" bson:"updated_at"`
	BodyHTML       string     `json:"body_html" bson:"body_html"`
	SortOrder      string     `json:"sort_order" bson:"sort_order"`
	TemplateSuffix string     `json:"template_suffix" bson:"template_suffix"`
	Image          Image      `json:"image" bson:"image"`
	PublishedAt    *time.Time `json:"published_at" bson:"published_at"`
	PublishedScope string     `json:"published_scope" bson:"published_scope"`
	CollectionType string     `json:"collection_type" bson:"collection_type"`
}

// Represents the result from the collections/X.json endpoint
type CollectionResource struct {
	Collection *Collection `json:"collection" bson:"collection"`
}

// Get individual collection
func (s *CollectionServiceOp) Get(collectionID int64, options interface{}) (*Collection, error) {
	path := fmt.Sprintf("%s/%d.json", collectionsBasePath, collectionID)
	resource := new(CollectionResource)
	err := s.client.Get(path, resource, options)
	return resource.Collection, err
}

// List products for a collection
func (s *CollectionServiceOp) ListProducts(collectionID int64, options interface{}) ([]Product, error) {
	products, _, err := s.ListProductsWithPagination(collectionID, options)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// List products for a collection and return pagination to retrieve next/previous results.
func (s *CollectionServiceOp) ListProductsWithPagination(collectionID int64, options interface{}) ([]Product, *Pagination, error) {
	path := fmt.Sprintf("%s/%d/products.json", collectionsBasePath, collectionID)
	resource := new(ProductsResource)
	headers := http.Header{}

	headers, err := s.client.createAndDoGetHeaders("GET", path, nil, options, resource)
	if err != nil {
		return nil, nil, err
	}

	// Extract pagination info from header
	linkHeader := headers.Get("Link")

	pagination, err := extractPagination(linkHeader)
	if err != nil {
		return nil, nil, err
	}

	return resource.Products, pagination, nil
}
