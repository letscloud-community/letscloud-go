package letscloud

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/letscloud-community/letscloud-go/domains"
)

//Profile retrieves the profile of current user
func (c *LetsCloud) Profile() (*domains.Profile, error) {
	req, err := c.requester.NewRequest(http.MethodGet, baseURL+"/profile", nil)
	if err != nil {
		return nil, err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var out domains.GetProfileResponse

	err = processResponse(b, &out)
	if err != nil {
		return nil, err
	}

	if !out.Success {
		return nil, errors.New(out.Message)
	}

	return &out.Data, nil
}

//Locations fetches all the locations of letscloud
func (c *LetsCloud) Locations() ([]domains.Location, error) {
	req, err := c.requester.NewRequest(http.MethodGet, baseURL+"/locations", nil)
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

//LocationPlans fetches all the pricing plans of the given location
func (c *LetsCloud) LocationPlans(name string) ([]domains.Plan, error) {
	if name == "" {
		return nil, errors.New("Please provide a valid location slug")
	}

	req, err := c.requester.NewRequest(http.MethodGet, baseURL+fmt.Sprintf("/locations/%s/plans", name), nil)
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

	return out.Data.Plans, nil
}

//LocationImages fetches all the VM images of the given location
func (c *LetsCloud) LocationImages(name string) ([]domains.Image, error) {
	req, err := c.requester.NewRequest(http.MethodGet, baseURL+fmt.Sprintf("/locations/%s/images", name), nil)
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

//NewSSHKey creates a new SSH key
func (c *LetsCloud) NewSSHKey(title, key string) (*domains.SSHKey, error) {
	payload := domains.SSHKeyCreateRequest{Title: title}

	if err := validateStruct(payload); err != nil {
		return nil, err
	}

	if key != "" {
		payload.Key = key
	}

	req, err := c.requester.NewRequest(http.MethodPost, baseURL+"/sshkeys", payload)
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

//SSHKeys returns all the SSH key of current user
func (c *LetsCloud) SSHKeys() ([]domains.SSHKey, error) {
	req, err := c.requester.NewRequest(http.MethodGet, baseURL+"/sshkeys", nil)
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

//SSHKey retrieves details of a given SSH key of current user
func (c *LetsCloud) SSHKey(title string) (*domains.SSHKey, error) {
	if title == "" {
		return nil, errors.New("Please provide a valid ssh key title")
	}

	req, err := c.requester.NewRequest(http.MethodGet, baseURL+"/sshkeys/"+title, nil)
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

//DeleteSSHKey deletes an existing SSH key of current user
func (c *LetsCloud) DeleteSSHKey(slug string) error {
	if slug == "" {
		return errors.New("Please provide a valid slug")
	}

	req, err := c.requester.NewRequest(http.MethodDelete, baseURL+"/sshkeys", domains.SSHKeyDelRequest{Slug: slug})
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

//Instances fetches all the instances created by the current user
func (c *LetsCloud) Instances() ([]domains.Instance, error) {
	req, err := c.requester.NewRequest(http.MethodGet, baseURL+"/instances", nil)
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

//CreateInstance creates a new instance
func (c *LetsCloud) CreateInstance(request *domains.CreateInstanceRequest) error {
	if request == nil {
		return errors.New("Please provide valid data in order to create instance")
	}

	if *request == (domains.CreateInstanceRequest{}) {
		return errors.New("Please provide valid data in order to create instance")
	}

	if err := validateStruct(*request); err != nil {
		return err
	}

	req, err := c.requester.NewRequest(http.MethodPost, baseURL+"/instances", request)
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

//Instance gets details about a particular instance of the current user
func (c *LetsCloud) Instance(id string) (*domains.Instance, error) {
	if id == "" {
		return nil, errors.New("Please provide a valid instance identifier")
	}

	req, err := c.requester.NewRequest(http.MethodGet, baseURL+"/instances/"+id, nil)
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

//DeleteInstance deletes any existing instance of the user
func (c *LetsCloud) DeleteInstance(id string) error {
	if id == "" {
		return errors.New("Please provide a valid instance identifier")
	}

	req, err := c.requester.NewRequest(http.MethodDelete, baseURL+"/instances/"+id, nil)
	if err != nil {
		return err
	}

	b, err := c.requester.SendRequest(req)
	if err != nil {
		return err
	}

	fmt.Println(b)

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

//PowerOnInstance turns on any existing instance of the current user
func (c *LetsCloud) PowerOnInstance(id string) error {
	if id == "" {
		return errors.New("Please provide a valid instance identifier")
	}

	req, err := c.requester.NewRequest(http.MethodPut, baseURL+"/instances/"+id+"/power-on", nil)
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

//PowerOffInstance turns off any existing instance of the current user
func (c *LetsCloud) PowerOffInstance(id string) error {
	if id == "" {
		return errors.New("Please provide a valid instance identifier")
	}

	req, err := c.requester.NewRequest(http.MethodPut, baseURL+"/instances/"+id+"/power-off", nil)
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

//RebootInstance as the name suggests, it reboots the instance
func (c *LetsCloud) RebootInstance(id string) error {
	if id == "" {
		return errors.New("Please provide a valid instance identifier")
	}

	req, err := c.requester.NewRequest(http.MethodPut, baseURL+"/instances/"+id+"/reboot", nil)
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

//ResetPasswordInstance is used for resetting the forgotten password of any instance
func (c *LetsCloud) ResetPasswordInstance(id, newPassword string) error {
	if id == "" || newPassword == "" {
		return errors.New("Please provide a valid instance identifier and new password")
	}

	req, err := c.requester.NewRequest(http.MethodPut, baseURL+"/instances/"+id+"/reset-password",
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
