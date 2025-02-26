package goshopify

import (
	"fmt"
	"time"
)

const webhooksBasePath = "webhooks"

type WebhookCountOptions struct {
	Address *string `json:"address,omitempty" url:"address,omitempty"`
	Topic   *string `json:"topic,omitempty" url:"topic,omitempty"`
}

type WebhookListOptions struct {
	Address      *string    `json:"address,omitempty" url:"address,omitempty"`
	CreatedAtMax *time.Time `json:"created_at_max,omitempty" url:"created_at_max,omitempty"`
	CreatedAtMin *time.Time `json:"created_at_min,omitempty" url:"created_at_min,omitempty"`
	Fields       []string   `json:"fields,omitempty" url:"fields,omitempty"`
	Limit        *int       `json:"limit,omitempty" url:"limit,omitempty"`
	SinceId      *int64     `json:"since_id,omitempty" url:"since_id,omitempty"`
	Topic        *string    `json:"topic,omitempty" url:"topic,omitempty"`
	UpdatedAtMax *time.Time `json:"updated_at_max,omitempty" url:"updated_at_max,omitempty"`
	UpdatedAtMin *time.Time `json:"updated_at_min,omitempty" url:"updated_at_min,omitempty"`
}

// WebhookService is an interface for interfacing with the webhook endpoints of
// the Shopify API.
// See: https://help.shopify.com/api/reference/webhook
type WebhookService interface {
	List(interface{}) ([]Webhook, error)
	Count(interface{}) (int, error)
	Get(int64, interface{}) (*Webhook, error)
	Create(Webhook) (*Webhook, error)
	Update(Webhook) (*Webhook, error)
	Delete(int64) error
}

// WebhookServiceOp handles communication with the webhook-related methods of
// the Shopify API.
type WebhookServiceOp struct {
	client *Client
}

// Webhook represents a Shopify webhook
type Webhook struct {
	ID                        int64      `json:"id,omitempty" bson:"id,omitempty"`
	ApiVersion                string     `json:"api_version,omitempty" bson:"api_version,omitempty"`
	Address                   string     `json:"address,omitempty" bson:"address,omitempty"`
	Topic                     string     `json:"topic,omitempty" bson:"topic,omitempty"`
	Format                    string     `json:"format,omitempty" bson:"format,omitempty"`
	CreatedAt                 *time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt                 *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Fields                    []string   `json:"fields,omitempty" bson:"fields,omitempty"`
	MetafieldNamespaces       []string   `json:"metafield_namespaces,omitempty" bson:"metafield_namespaces,omitempty"`
	PrivateMetafieldNamespace []string   `json:"private_metafield_namespace,omitempty" bson:"private_metafield_namespace,omitempty"`
}

// WebhookOptions can be used for filtering webhooks on a List request.
type WebhookOptions struct {
	Address string `url:"address,omitempty" bson:"address,omitempty"`
	Topic   string `url:"topic,omitempty" bson:"topic,omitempty"`
}

// WebhookResource represents the result from the admin/webhooks.json endpoint
type WebhookResource struct {
	Webhook *Webhook `json:"webhook" bson:"webhook"`
}

// WebhooksResource is the root object for a webhook get request.
type WebhooksResource struct {
	Webhooks []Webhook `json:"webhooks" bson:"webhooks"`
}

// List webhooks
func (s *WebhookServiceOp) List(options interface{}) ([]Webhook, error) {
	path := fmt.Sprintf("%s.json", webhooksBasePath)
	resource := new(WebhooksResource)
	err := s.client.Get(path, resource, options)
	return resource.Webhooks, err
}

// Count webhooks
func (s *WebhookServiceOp) Count(options interface{}) (int, error) {
	path := fmt.Sprintf("%s/count.json", webhooksBasePath)
	return s.client.Count(path, options)
}

// Get individual webhook
func (s *WebhookServiceOp) Get(webhookdID int64, options interface{}) (*Webhook, error) {
	path := fmt.Sprintf("%s/%d.json", webhooksBasePath, webhookdID)
	resource := new(WebhookResource)
	err := s.client.Get(path, resource, options)
	return resource.Webhook, err
}

// Create a new webhook
func (s *WebhookServiceOp) Create(webhook Webhook) (*Webhook, error) {
	path := fmt.Sprintf("%s.json", webhooksBasePath)
	wrappedData := WebhookResource{Webhook: &webhook}
	resource := new(WebhookResource)
	err := s.client.Post(path, wrappedData, resource)
	return resource.Webhook, err
}

// Update an existing webhook.
func (s *WebhookServiceOp) Update(webhook Webhook) (*Webhook, error) {
	path := fmt.Sprintf("%s/%d.json", webhooksBasePath, webhook.ID)
	wrappedData := WebhookResource{Webhook: &webhook}
	resource := new(WebhookResource)
	err := s.client.Put(path, wrappedData, resource)
	return resource.Webhook, err
}

// Delete an existing webhooks
func (s *WebhookServiceOp) Delete(ID int64) error {
	return s.client.Delete(fmt.Sprintf("%s/%d.json", webhooksBasePath, ID))
}
