package infoblox

import (
	"encoding/json"
	"fmt"
)

// RecordHost returns the HOST record resource
// https://192.168.2.200/wapidoc/objects/record.host.html
func (c *Client) RecordHost() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "record:host",
	}
}

// RecordHostObject defines the HOST record object's fields
type RecordHostObject struct {
	Object
	Aliases         []string       `json:"aliases"`
	Comment         string         `json:"comment"`
	ConfigureForDNS bool           `json:"configure_for_dns"`
	Ipv4Addrs       []HostIpv4Addr `json:"ipv4addrs,omitempty"`
	Ipv6Addrs       []HostIpv6Addr `json:"ipv6addrs,omitempty"`
	Name            string         `json:"name,omitempty"`
	Ttl             int            `json:"ttl,omitempty"`
	View            string         `json:"view,omitempty"`
	Zone            string         `json:"zone,omitempty"`
}

// HostIpv4Addr is an ipv4 address for a HOST record
type HostIpv4Addr struct {
	Object           `json:"-"`
	ConfigureForDHCP bool   `json:"configure_for_dhcp"`
	Host             string `json:"host,omitempty"`
	Ipv4Addr         string `json:"ipv4addr,omitempty"`
	MAC              string `json:"mac,omitempty"`
}

// HostIpv6Addr is an ipv6 address for a HOST record
type HostIpv6Addr struct {
	Object           `json:"-"`
	ConfigureForDHCP bool   `json:"configure_for_dhcp"`
	Host             string `json:"host,omitempty"`
	Ipv6Addr         string `json:"ipv6addr,omitempty"`
	MAC              string `json:"mac,omitempty"`
}

// RecordHostObject instantiates an HOST record object with a WAPI ref
func (c *Client) RecordHostObject(ref string) *RecordHostObject {
	host := RecordHostObject{}
	host.Object = Object{
		Ref: ref,
		r:   c.RecordHost(),
	}
	return &host
}

// GetRecordHost fetches a HOST record from the Infoblox WAPI by its ref
func (c *Client) GetRecordHost(ref string, opts *Options) (*RecordHostObject, error) {
	resp, err := c.RecordHostObject(ref).get(opts)
	if err != nil {
		return nil, fmt.Errorf("Could not get created host record: %s", err)
	}

	var out RecordHostObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) FindRecordHost(name string, view string) ([]RecordHostObject, error) {
	// FindRecordHost searches the Infoblox WAPI for the HOST record with the given
	field := "name"
	viewName := "view"
	// conditions := []Condition{Condition{Field: &field, Value: name}}
	conditions := []Condition{
		Condition{
			Field: &field,
			Value: name,
		},
		Condition{
			Field: &viewName,
			Value: view,
		},
	}
	resp, err := c.RecordHost().find(conditions, nil)
	if err != nil {
		return nil, err
	}

	var out []RecordHostObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) CreateRecordHost(recordHostObject RecordHostObject) (string, error) {
	d, _ := json.Marshal(recordHostObject)
	resp, err := c.RecordHost().CreateJson("record:host", nil, d)
	if err != nil {
		return "", err
	}
	return resp, nil
}
