package goshopify

import (
	"fmt"
	"time"
)

const smartCollectionsBasePath = "smart_collections"
const smartCollectionsResourceName = "collections"

type SmartCollectionCountOptions struct {
	ProductId       *int       `json:"product_id,omitempty"  url:"product_id,omitempty"`
	PublishedAtMax  *time.Time `json:"published_at_max,omitempty"  url:"published_at_max,omitempty"`
	PublishedAtMin  *time.Time `json:"published_at_min,omitempty"  url:"published_at_min,omitempty"`
	PublishedStatus *string    `json:"published_status,omitempty"  url:"published_status,omitempty"`
	Title           *string    `json:"title,omitempty"  url:"title,omitempty"`
	UpdatedAtMax    *time.Time `json:"updated_at_max"  url:"updated_at_max,omitempty"`
	UpdatedAtMin    *time.Time `json:"updated_at_min"  url:"updated_at_min,omitempty"`
}

type SmartCollectionListOptions struct {
	Fields          *string    `json:"fields,omitempty" url:"fields,omitempty"`
	Handle          *string    `json:"handle,omitempty" url:"handle,omitempty"`
	Ids             *string    `json:"ids,omitempty" url:"ids,omitempty"`
	Limit           *int       `json:"limit,omitempty" url:"limit,omitempty"`
	ProductId       *int       `json:"product_id,omitempty" url:"product_id,omitempty"`
	PublishedAtMax  *time.Time `json:"published_at_max,omitempty" url:"published_at_max,omitempty"`
	PublishedAtMin  *time.Time `json:"published_at_min,omitempty" url:"published_at_min,omitempty"`
	PublishedStatus *string    `json:"published_status,omitempty" url:"published_status,omitempty"`
	SinceId         *int64     `json:"since_id,omitempty" url:"since_id,omitempty"`
	Title           *string    `json:"title,omitempty" url:"title,omitempty"`
	UpdatedAtMax    *time.Time `json:"updated_at_max,omitempty" url:"updated_at_max,omitempty"`
	UpdatedAtMin    *time.Time `json:"updated_at_min,omitempty" url:"updated_at_min,omitempty"`
}

// SmartCollectionService is an interface for interacting with the smart
// collection endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/smartcollection
type SmartCollectionService interface {
	List(interface{}) ([]SmartCollection, error)
	Count(interface{}) (int, error)
	Get(int64, interface{}) (*SmartCollection, error)
	Create(SmartCollection) (*SmartCollection, error)
	Update(SmartCollection) (*SmartCollection, error)
	Delete(int64) error

	// MetafieldsService used for SmartCollection resource to communicate with Metafields resource
	MetafieldsService
}

// SmartCollectionServiceOp handles communication with the smart collection
// related methods of the Shopify API.
type SmartCollectionServiceOp struct {
	client *Client
}

type Rule struct {
	Column    string `json:"column" bson:"column"`
	Relation  string `json:"relation" bson:"relation"`
	Condition string `json:"condition" bson:"condition"`
}

// SmartCollection represents a Shopify smart collection.
type SmartCollection struct {
	ID             int64       `json:"id" bson:"id"`
	Handle         string      `json:"handle" bson:"handle"`
	Title          string      `json:"title" bson:"title"`
	UpdatedAt      *time.Time  `json:"updated_at" bson:"updated_at"`
	BodyHTML       string      `json:"body_html" bson:"body_html"`
	SortOrder      string      `json:"sort_order" bson:"sort_order"`
	TemplateSuffix string      `json:"template_suffix" bson:"template_suffix"`
	Image          Image       `json:"image" bson:"image"`
	Published      bool        `json:"published" bson:"published"`
	PublishedAt    *time.Time  `json:"published_at" bson:"published_at"`
	PublishedScope string      `json:"published_scope" bson:"published_scope"`
	Rules          []Rule      `json:"rules" bson:"rules"`
	Disjunctive    bool        `json:"disjunctive" bson:"disjunctive"`
	Metafields     []Metafield `json:"metafields,omitempty" bson:"metafields"`
}

// SmartCollectionResource represents the result from the smart_collections/X.json endpoint
type SmartCollectionResource struct {
	Collection *SmartCollection `json:"smart_collection" bson:"smart_collection"`
}

// SmartCollectionsResource represents the result from the smart_collections.json endpoint
type SmartCollectionsResource struct {
	Collections []SmartCollection `json:"smart_collections" bson:"smart_collections"`
}

// List smart collections
func (s *SmartCollectionServiceOp) List(options interface{}) ([]SmartCollection, error) {
	path := fmt.Sprintf("%s.json", smartCollectionsBasePath)
	resource := new(SmartCollectionsResource)
	err := s.client.Get(path, resource, options)
	return resource.Collections, err
}

// Count smart collections
func (s *SmartCollectionServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", smartCollectionsBasePath)
	return s.client.Count(path, options)
}

// Get individual smart collection
func (s *SmartCollectionServiceOp) Get(collectionID int64, options interface{}) (*SmartCollection, error) {
	path := fmt.Sprintf("%s/%d.json", smartCollectionsBasePath, collectionID)
	resource := new(SmartCollectionResource)
	err := s.client.Get(path, resource, options)
	return resource.Collection, err
}

// Create a new smart collection
// See Image for the details of the Image creation for a collection.
func (s *SmartCollectionServiceOp) Create(collection SmartCollection) (*SmartCollection, error) {
	path := fmt.Sprintf("%s.json", smartCollectionsBasePath)
	wrappedData := SmartCollectionResource{Collection: &collection}
	resource := new(SmartCollectionResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Collection, err
}

// Update an existing smart collection
func (s *SmartCollectionServiceOp) Update(collection SmartCollection) (*SmartCollection, error) {
	path := fmt.Sprintf("%s/%d.json", smartCollectionsBasePath, collection.ID)
	wrappedData := SmartCollectionResource{Collection: &collection}
	resource := new(SmartCollectionResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Collection, err
}

// Delete an existing smart collection.
func (s *SmartCollectionServiceOp) Delete(collectionID int64) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", smartCollectionsBasePath, collectionID))
}

// List metafields for a smart collection
func (s *SmartCollectionServiceOp) ListMetafields(smartCollectionID int64, options interface{}) ([]Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: smartCollectionsResourceName, resourceID: smartCollectionID}
	return metafieldService.List(options)
}

// Count metafields for a smart collection
func (s *SmartCollectionServiceOp) CountMetafields(smartCollectionID int64, options interface{}) (int, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: smartCollectionsResourceName, resourceID: smartCollectionID}
	return metafieldService.Count(options)
}

// Get individual metafield for a smart collection
func (s *SmartCollectionServiceOp) GetMetafield(smartCollectionID int64, metafieldID int64, options interface{}) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: smartCollectionsResourceName, resourceID: smartCollectionID}
	return metafieldService.Get(metafieldID, options)
}

// Create a new metafield for a smart collection
func (s *SmartCollectionServiceOp) CreateMetafield(smartCollectionID int64, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: smartCollectionsResourceName, resourceID: smartCollectionID}
	return metafieldService.Create(metafield)
}

// Update an existing metafield for a smart collection
func (s *SmartCollectionServiceOp) UpdateMetafield(smartCollectionID int64, metafield Metafield) (*Metafield, error) {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: smartCollectionsResourceName, resourceID: smartCollectionID}
	return metafieldService.Update(metafield)
}

// // Delete an existing metafield for a smart collection
func (s *SmartCollectionServiceOp) DeleteMetafield(smartCollectionID int64, metafieldID int64) error {
	metafieldService := &MetafieldServiceOp{client: s.client, resource: smartCollectionsResourceName, resourceID: smartCollectionID}
	return metafieldService.Delete(metafieldID)
}
