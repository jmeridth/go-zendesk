package zendesk

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// Organization is struct for organization payload
// https://developer.zendesk.com/rest_api/docs/support/organizations
type Organization struct {
	ID                int64                  `json:"id,omitempty"`
	URL               string                 `json:"url,omitempty"`
	Name              string                 `json:"name"`
	DomainNames       []string               `json:"domain_names"`
	GroupID           int64                  `json:"group_id"`
	SharedTickets     bool                   `json:"shared_tickets"`
	SharedComments    bool                   `json:"shared_comments"`
	Tags              []string               `json:"tags"`
	CreatedAt         time.Time              `json:"created_at,omitempty"`
	UpdatedAt         time.Time              `json:"updated_at,omitempty"`
	Organizationields map[string]interface{} `json:"organization_fields,omitempty"`
}

// OrganizationAPI an interface containing all methods associated with zendesk organizations
type OrganizationAPI interface {
	CreateOrganization(ctx context.Context, org Organization) (Organization, error)
	GetOrganization(ctx context.Context, orgID int64) (Organization, error)
	UpdateOrganization(ctx context.Context, orgID int64, org Organization) (Organization, error)
	DeleteOrganization(ctx context.Context, orgID int64) error
}

// CreateOrganization creates new organization
// https://developer.zendesk.com/rest_api/docs/support/organizations#create-organization
func (z *Client) CreateOrganization(ctx context.Context, org Organization) (Organization, error) {
	var data, result struct {
		Organization Organization `json:"organization"`
	}

	data.Organization = org

	body, err := z.post(ctx, "/organizations.json", data)
	if err != nil {
		return Organization{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return Organization{}, err
	}

	return result.Organization, nil
}

// GetOrganization gets a specified organization
// ref: https://developer.zendesk.com/rest_api/docs/support/organizations#show-organization
func (z *Client) GetOrganization(ctx context.Context, orgID int64) (Organization, error) {
	var result struct {
		Organization Organization `json:"organization"`
	}

	body, err := z.get(ctx, fmt.Sprintf("/organizations/%d.json", orgID))

	if err != nil {
		return Organization{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return Organization{}, err
	}

	return result.Organization, err
}

// UpdateOrganization updates a organization with the specified organization
// ref: https://developer.zendesk.com/rest_api/docs/support/organizations#update-organization
func (z *Client) UpdateOrganization(ctx context.Context, orgID int64, org Organization) (Organization, error) {
	var result, data struct {
		Organization Organization `json:"organization"`
	}

	data.Organization = org

	body, err := z.put(ctx, fmt.Sprintf("/organizations/%d.json", orgID), data)

	if err != nil {
		return Organization{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return Organization{}, err
	}

	return result.Organization, err
}

// DeleteOrganization deletes the specified organization
// ref: https://developer.zendesk.com/rest_api/docs/support/organizations#delete-organization
func (z *Client) DeleteOrganization(ctx context.Context, orgID int64) error {
	err := z.delete(ctx, fmt.Sprintf("/organizations/%d.json", orgID))

	if err != nil {
		return err
	}

	return nil
}
