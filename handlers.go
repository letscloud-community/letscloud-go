package letscloud

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/letscloud-community/letscloud-go/domains"
)

// Profile retrieves the profile of current user
func (c *LetsCloud) Profile() (*domains.Profile, error) {
	c.debugLog("Fetching profile")

	req, err := c.requester.NewRequest(http.MethodGet, "/profile", nil)
	if err != nil {
		return nil, err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return nil, err
	}

	c.debugLog("Profile response: " + string(b))

	var out domains.GetProfileResponse

	err = processResponse(b, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, errors.New(out.Message)
	}

	c.debugLog("Profile fetched successfully")
	return &out.Data, nil
}

// Locations fetches all the locations of letscloud
func (c *LetsCloud) Locations() ([]domains.Location, error) {
	req, err := c.requester.NewRequest(http.MethodGet, "/locations", nil)
	if err != nil {
		return nil, err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var out domains.GetLocationsResponse

	err = processResponse(b, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, errors.New(out.Message)
	}

	return out.Data, nil
}

// LocationPlans fetches all the pricing plans of the given location
func (c *LetsCloud) LocationPlans(slug string) ([]domains.Plan, error) {
	if slug == "" {
		return nil, errors.New("please provide a valid location slug")
	}

	req, err := c.requester.NewRequest(http.MethodGet, fmt.Sprintf("/locations/%s/plans", slug), nil)
	if err != nil {
		return nil, err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var out domains.GetLocationPlansResponse

	err = processResponse(b, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, errors.New(out.Message)
	}

	var allPlans []domains.Plan
	for _, location := range out.Data {
		allPlans = append(allPlans, location.Plans...)
	}

	return allPlans, nil
}

// LocationImages fetches all the VM images of the given location
func (c *LetsCloud) LocationImages(slug string) ([]domains.Image, error) {
	if slug == "" {
		return nil, errors.New("please provide a valid location slug")
	}

	req, err := c.requester.NewRequest(http.MethodGet, fmt.Sprintf("/locations/%s/images", slug), nil)
	if err != nil {
		return nil, err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var out domains.GetLocationImagesResponse

	err = processResponse(b, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, errors.New(out.Message)
	}

	return out.Data, nil
}

// NewSSHKey creates a new SSH key
func (c *LetsCloud) NewSSHKey(title, key string) (*domains.SSHKey, error) {
	payload := domains.SSHKeyCreateRequest{Title: title}

	if err := validateStruct(payload); err != nil {
		return nil, err
	}

	if key != "" {
		payload.Key = key
	}

	req, err := c.requester.NewRequest(http.MethodPost, "/sshkeys", payload)
	if err != nil {
		return nil, err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var out domains.CreateOrGetSSHKeysResponse

	err = processResponse(b, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, errors.New(out.Message)
	}

	return &out.Data, nil
}

// SSHKeys returns all the SSH key of current user
func (c *LetsCloud) SSHKeys() ([]domains.SSHKey, error) {
	req, err := c.requester.NewRequest(http.MethodGet, "/sshkeys", nil)
	if err != nil {
		return nil, err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var out domains.GetSSHKeysResponse

	err = processResponse(b, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, errors.New(out.Message)
	}

	return out.Data, nil
}

// SSHKey retrieves details of a given SSH key of current user
func (c *LetsCloud) SSHKey(title string) (*domains.SSHKey, error) {
	if title == "" {
		return nil, errors.New("please provide a valid ssh key title")
	}

	req, err := c.requester.NewRequest(http.MethodGet, "/sshkeys/"+title, nil)
	if err != nil {
		return nil, err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var out domains.CreateOrGetSSHKeysResponse

	err = processResponse(b, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, errors.New(out.Message)
	}

	return &out.Data, nil
}

// DeleteSSHKey deletes an existing SSH key of current user
func (c *LetsCloud) DeleteSSHKey(slug string) error {
	if slug == "" {
		return errors.New("please provide a valid slug")
	}

	req, err := c.requester.NewRequest(http.MethodDelete, "/sshkeys", domains.SSHKeyDelRequest{Slug: slug})
	if err != nil {
		return err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return err
	}

	var out domains.CreateOrGetSSHKeysResponse

	err = processResponse(b, &out)
	if err != nil {
		return err
	}

	if !out.Success {
		return errors.New(out.Message)
	}

	return nil
}

// Instances fetches all the instances created by the current user
func (c *LetsCloud) Instances() ([]domains.Instance, error) {
	req, err := c.requester.NewRequest(http.MethodGet, "/instances", nil)
	if err != nil {
		return nil, err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var out domains.GetInstancesResponse

	err = processResponse(b, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, errors.New(out.Message)
	}

	return out.Data, nil
}

// CreateInstance creates a new instance
func (c *LetsCloud) CreateInstance(request *domains.CreateInstanceRequest) error {
	if request == nil {
		return errors.New("please provide valid data in order to create instance")
	}

	if *request == (domains.CreateInstanceRequest{}) {
		return errors.New("please provide valid data in order to create instance")
	}

	if err := validateStruct(*request); err != nil {
		return err
	}

	req, err := c.requester.NewRequest(http.MethodPost, "/instances", request)
	if err != nil {
		return err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return err
	}

	var out domains.GetInstanceResponse

	err = processResponse(b, &out)
	if err != nil {
		return err
	}

	if !out.Success {
		return errors.New(out.Message)
	}

	return nil
}

// Instance gets details about a particular instance of the current user
func (c *LetsCloud) Instance(identifier string) (*domains.Instance, error) {
	if identifier == "" {
		return nil, errors.New("please provide a valid instance identifier")
	}

	req, err := c.requester.NewRequest(http.MethodGet, "/instances/"+identifier, nil)
	if err != nil {
		return nil, err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var out domains.GetInstanceResponse

	err = processResponse(b, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, errors.New(out.Message)
	}

	return &out.Data, nil
}

// DeleteInstance deletes any existing instance of the user
func (c *LetsCloud) DeleteInstance(identifier string) error {
	if identifier == "" {
		return errors.New("please provide a valid instance identifier")
	}

	req, err := c.requester.NewRequest(http.MethodDelete, "/instances/"+identifier, nil)
	if err != nil {
		return err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return err
	}

	var out domains.CommonResponse

	err = processResponse(b, &out)
	if err != nil {
		return err
	}

	if !out.Success {
		return errors.New(out.Message)
	}

	return nil
}

// PowerOnInstance turns on any existing instance of the current user
func (c *LetsCloud) PowerOnInstance(identifier string) error {
	if identifier == "" {
		return errors.New("please provide a valid instance identifier")
	}

	req, err := c.requester.NewRequest(http.MethodPut, "/instances/"+identifier+"/power-on", nil)
	if err != nil {
		return err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return err
	}

	var out domains.CommonResponse

	err = processResponse(b, &out)
	if err != nil {
		return err
	}

	if !out.Success {
		return errors.New(out.Message)
	}

	return nil
}

// PowerOffInstance turns off any existing instance of the current user
func (c *LetsCloud) PowerOffInstance(identifier string) error {
	if identifier == "" {
		return errors.New("please provide a valid instance identifier")
	}

	req, err := c.requester.NewRequest(http.MethodPut, "/instances/"+identifier+"/power-off", nil)
	if err != nil {
		return err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return err
	}

	var out domains.CommonResponse

	err = processResponse(b, &out)
	if err != nil {
		return err
	}

	if !out.Success {
		return errors.New(out.Message)
	}

	return nil
}

// RebootInstance as the name suggests, it reboots the instance
func (c *LetsCloud) RebootInstance(identifier string) error {
	if identifier == "" {
		return errors.New("please provide a valid instance identifier")
	}

	req, err := c.requester.NewRequest(http.MethodPut, "/instances/"+identifier+"/reboot", nil)
	if err != nil {
		return err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return err
	}

	var out domains.CommonResponse

	err = processResponse(b, &out)
	if err != nil {
		return err
	}

	if !out.Success {
		return errors.New(out.Message)
	}

	return nil
}

// ResetPasswordInstance is used for resetting the forgotten password of any instance
func (c *LetsCloud) ResetPasswordInstance(identifier, newPassword string) error {
	if identifier == "" || newPassword == "" {
		return errors.New("please provide a valid instance identifier and new password")
	}

	req, err := c.requester.NewRequest(http.MethodPut, "/instances/"+identifier+"/reset-password",
		domains.InstanceResetPasswordRequest{Password: newPassword})
	if err != nil {
		return err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return err
	}

	var out domains.CommonResponse

	err = processResponse(b, &out)
	if err != nil {
		return err
	}

	if !out.Success {
		return errors.New(out.Message)
	}

	return nil
}

// NewSnapshot creates a new snapshot of the instance
func (c *LetsCloud) NewSnapshot(label, identifier string) (*domains.CreateOrGetSnapshotResponse, error) {
	if identifier == "" || label == "" {
		return nil, errors.New("please provide a valid instance identifier and label")
	}

	req, err := c.requester.NewRequest(http.MethodPost, "/instances/"+identifier+"/snapshots",
		domains.SnapshotCreateRequest{Label: label})
	if err != nil {
		return nil, err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var out domains.CreateOrGetSnapshotResponse

	err = processResponse(b, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, errors.New(out.Message)
	}

	return &out, nil
}

// Snapshots fetches all the snapshots of the current user
func (c *LetsCloud) Snapshots() ([]domains.Snapshot, error) {
	req, err := c.requester.NewRequest(http.MethodGet, "/snapshots", nil)
	if err != nil {
		return nil, err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var out domains.GetSnapshotsResponse

	err = processResponse(b, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, errors.New(out.Message)
	}

	return out.Data, nil
}

// Snapshot gets details about a particular snapshot of the current user
func (c *LetsCloud) Snapshot(slug string) (*domains.Snapshot, error) {
	if slug == "" {
		return nil, errors.New("please provide a valid snapshot slug")
	}

	req, err := c.requester.NewRequest(http.MethodGet, "/snapshots/"+slug, nil)
	if err != nil {
		return nil, err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var out domains.CreateOrGetSnapshotResponse

	err = processResponse(b, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, errors.New(out.Message)
	}

	return &out.Data, nil
}

// UpdateSnapshot updates an existing snapshot of the current user
func (c *LetsCloud) UpdateSnapshot(slug, label string) error {
	if slug == "" || label == "" {
		return errors.New("please provide a valid snapshot slug and label")
	}

	req, err := c.requester.NewRequest(http.MethodPut, "/snapshots/"+slug,
		domains.SnapshotUpdateRequest{Label: label})
	if err != nil {
		return err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return err
	}

	var out domains.CommonResponse

	err = processResponse(b, &out)
	if err != nil {
		return err
	}

	if !out.Success {
		return errors.New(out.Message)
	}

	return nil
}

// DeleteSnapshot deletes an existing snapshot of the current user
func (c *LetsCloud) DeleteSnapshot(slug string) error {
	if slug == "" {
		return errors.New("please provide a valid snapshot slug")
	}

	req, err := c.requester.NewRequest(http.MethodDelete, "/snapshots/"+slug, nil)
	if err != nil {
		return err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return err
	}

	var out domains.CommonResponse

	err = processResponse(b, &out)
	if err != nil {
		return err
	}

	if !out.Success {
		return errors.New(out.Message)
	}

	return nil
}
