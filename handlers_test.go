package letscloud

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/letscloud-community/letscloud-go/domains"
	"github.com/letscloud-community/letscloud-go/httpclient"
)

const (
	TEST_API_KEY = "newtesttoken"
)

func TestClient_CreateInstance(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return([]byte(`{"success": true, "message": "Instance successfully created"}`), nil)

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return(nil, errors.New("server not reachable"))

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	type args struct {
		request *domains.CreateInstanceRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "passing nil request",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			wantErr: true,
		},
		{
			name: "passing nothing",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args: args{
				request: &domains.CreateInstanceRequest{},
			},
			wantErr: true,
		},
		{
			name: "passing something",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args: args{
				request: &domains.CreateInstanceRequest{
					Label: "Label Example",
				},
			},
			wantErr: true,
		},
		{
			name: "passing full data",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args: args{
				request: &domains.CreateInstanceRequest{
					LocationSlug: "MIA1",
					PlanSlug:     "1vcpu-1gb-10ssd",
					Hostname:     "Hostname Example",
					Label:        "Label Example",
					ImageSlug:    "ubuntu-20.04-x86_64",
				},
			},
			wantErr: false,
		},
		{
			name: "not able to reach the server",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args: args{
				request: &domains.CreateInstanceRequest{
					LocationSlug: "MIA1",
					PlanSlug:     "1vcpu-1gb-10ssd",
					Hostname:     "Hostname Example",
					Label:        "Label Example",
					ImageSlug:    "ubuntu-20.04-x86_64",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			if err := c.CreateInstance(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("CreateInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_DeleteInstance(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return([]byte(`{"success": true, "message": "Instance successfully deleted"}`), nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	type args struct {
		identifier string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "passing no identifier",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			wantErr: true,
		},
		{
			name: "passing no id",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args: args{
				identifier: "identifier-example",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			if err := c.DeleteInstance(tt.args.identifier); (err != nil) != tt.wantErr {
				t.Errorf("DeleteInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_DeleteSSHKey(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return([]byte(`{"success": true, "message": "SSH Key was successfully deleted!"}`), nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "sending invalid slug",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args:    args{},
			wantErr: true,
		},
		{
			name: "sending valid slug",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args: args{
				slug: "your-ssh-key-slug",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			if err := c.DeleteSSHKey(tt.args.slug); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSSHKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Instance(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return([]byte(`{"success": true, "data": {"identifier": "identifier-example"}}`), nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	type args struct {
		identifier string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domains.Instance
		wantErr bool
	}{
		{
			name: "sending invalid id",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "sending valid id",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args: args{
				identifier: "identifier-example",
			},
			want: &domains.Instance{
				Identifier: "identifier-example",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			got, err := c.Instance(tt.args.identifier)
			if (err != nil) != tt.wantErr {
				t.Errorf("Instance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Instance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Instances(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	var instances = []domains.Instance{
		{
			Identifier:    "identifier-example",
			Booted:        true,
			Built:         true,
			Locked:        false,
			Suspended:     false,
			Memory:        1024,
			TotalDiskSize: 10,
			CPUS:          1,
			Label:         "Label Example",
			IPAddresses: []domains.IPAddress{
				{
					Address: "192.168.3.5",
				},
				{
					Address: "172.168.3.5",
				},
			},
			TemplateLabel: "Ubuntu 20.04 x64",
			Hostname:      "Hostname Example",
			RootPassword:  "testpassword",
			Location: domains.Location{
				Slug:    "MIA1",
				Country: "United States",
				City:    "Miami",
			},
		},
		{
			Identifier:    "4514254",
			Booted:        true,
			Built:         true,
			Locked:        false,
			Suspended:     false,
			Memory:        2048,
			TotalDiskSize: 20,
			CPUS:          2,
			Label:         "Label Example 2",
			IPAddresses: []domains.IPAddress{
				{
					Address: "192.168.3.6",
				},
				{
					Address: "172.168.3.6",
				},
			},
			TemplateLabel: "Ubuntu 20.04 x64",
			Hostname:      "Hostname Example 2",
			RootPassword:  "testpassword",
			Location: domains.Location{
				Slug:    "MIA1",
				Country: "United States",
				City:    "Miami",
			},
		},
	}

	b, _ := json.Marshal(domains.GetInstancesResponse{
		CommonResponse: domains.CommonResponse{
			Success: true,
		},
		Data: instances,
	})

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return(b, nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domains.Instance
		wantErr bool
	}{
		{
			name: "getting all thhe instances",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			want:    instances,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			got, err := c.Instances()
			if (err != nil) != tt.wantErr {
				t.Errorf("Instances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Instances() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_LocationImages(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	var images = []domains.Image{
		{
			Slug:   "rocky-9.3-x86_64",
			Distro: "Rocky 9.3 x86_64",
			OS:     "rocky",
		},
		{
			Slug:   "ubuntu-22.04-x86_64",
			Distro: "Ubuntu 22.04 x86_64",
			OS:     "ubuntu",
		},
	}

	imagesResp, _ := json.Marshal(domains.GetLocationImagesResponse{
		CommonResponse: domains.CommonResponse{
			Success: true,
		},
		Data: images,
	})

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return(imagesResp, nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domains.Image
		wantErr bool
	}{
		{
			name: "get all location images",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args: args{
				slug: "MIA1",
			},
			want:    images,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			got, err := c.LocationImages(tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocationImages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocationImages() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_LocationPlans(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	var plans = []domains.LocationPlanWrapper{
		{
			Country: "United States",
			City:    "Miami",
			Slug:    "MIA1",
			Plans: []domains.Plan{
				{
					Slug:         "1vcpu-1gb-10ssd",
					CurrencyCode: "USD",
					Shortcode:    "$",
					Core:         1,
					Memory:       1024,
					Disk:         10,
					Bandwidth:    1000,
					MonthlyValue: "5.00",
				},
				{
					Slug:         "1vcpu-2gb-20ssd",
					CurrencyCode: "USD",
					Shortcode:    "$",
					Core:         1,
					Memory:       2048,
					Disk:         20,
					Bandwidth:    1500,
					MonthlyValue: "10.00",
				},
			},
		},
	}

	plansResp, _ := json.Marshal(domains.GetLocationPlansResponse{
		CommonResponse: domains.CommonResponse{
			Success: true,
		},
		Data: plans,
	})

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return(plansResp, nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domains.Plan
		wantErr bool
	}{
		{
			name: "get all location plans with invalid",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "get all location plans",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args: args{
				slug: "MIA1",
			},
			want:    plans[0].Plans,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			got, err := c.LocationPlans(tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("LocationPlans() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LocationPlans() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Locations(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	var locations = []domains.Location{
		{
			Slug:      "MIA1",
			Country:   "United States",
			City:      "Miami",
			Available: true,
		},
		{
			Slug:      "DAL1",
			Country:   "United States",
			City:      "Dallas",
			Available: false,
		},
	}

	locationsResp, _ := json.Marshal(domains.GetLocationsResponse{
		CommonResponse: domains.CommonResponse{
			Success: true,
		},
		Data: locations,
	})

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return(locationsResp, nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domains.Location
		wantErr bool
	}{
		{
			name: "getting all the locations",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			want:    locations,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			got, err := c.Locations()
			if (err != nil) != tt.wantErr {
				t.Errorf("Locations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Locations() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_NewSSHKey(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	var newKey = domains.SSHKey{
		Slug:       "my-ssh-key",
		Title:      "My SSH Key",
		PublicKey:  "<random-generated>",
		PrivateKey: "<random-generated>",
	}

	sshCreateResp, _ := json.Marshal(domains.CreateOrGetSSHKeysResponse{
		CommonResponse: domains.CommonResponse{
			Success: true,
		},
		Data: newKey,
	})

	var newKeyV2 = domains.SSHKey{
		Slug:       "my-ssh-key",
		Title:      "My SSH Key",
		PublicKey:  "<this-is-a-test-key>",
		PrivateKey: "<random-generated>",
	}

	sshCreateRespV2, _ := json.Marshal(domains.CreateOrGetSSHKeysResponse{
		CommonResponse: domains.CommonResponse{
			Success: true,
		},
		Data: newKeyV2,
	})

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil).Times(2)
	mclient.EXPECT().SendRequest(gomock.Any()).Return(sshCreateResp, nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return(sshCreateRespV2, nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	type args struct {
		title string
		key   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domains.SSHKey
		wantErr bool
	}{
		{
			name: "sending invalid data",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "sending valid data without the key",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args: args{
				title: "my-ssh",
			},
			want:    &newKey,
			wantErr: false,
		},
		{
			name: "sending valid data with the key",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args: args{
				title: "my-ssh",
				key:   "<this-is-a-test-key>",
			},
			want:    &newKeyV2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			got, err := c.NewSSHKey(tt.args.title, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSSHKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSSHKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_PowerOffInstance(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	powerOffResp, _ := json.Marshal(domains.CommonResponse{
		Success: true,
		Message: "your instance has been turned off",
	})

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return(powerOffResp, nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	type args struct {
		identifier string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "sending power off signal without id",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			wantErr: true,
		},
		{
			name: "sending power off signal with id",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args:    args{identifier: "identifier-example"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			if err := c.PowerOffInstance(tt.args.identifier); (err != nil) != tt.wantErr {
				t.Errorf("PowerOffInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_PowerOnInstance(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	powerOnResp, _ := json.Marshal(domains.CommonResponse{
		Success: true,
		Message: "your instance has been turned on",
	})

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return(powerOnResp, nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	type args struct {
		identifier string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "sending power on signal without id",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			wantErr: true,
		},
		{
			name: "sending power on signal with id",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args:    args{identifier: "identifier-example"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			if err := c.PowerOnInstance(tt.args.identifier); (err != nil) != tt.wantErr {
				t.Errorf("PowerOnInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Profile(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	var profile = domains.Profile{
		Name:        "John Doe",
		CompanyName: "Company Name",
		Email:       "johndoe@test.com",
		Currency:    "$",
		Balance:     "999.99",
	}

	getProfileResp, _ := json.Marshal(domains.GetProfileResponse{
		CommonResponse: domains.CommonResponse{
			Success: true,
		},
		Data: profile,
	})

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return(getProfileResp, nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	tests := []struct {
		name    string
		fields  fields
		want    *domains.Profile
		wantErr bool
	}{
		{
			name: "getting profile info",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			want:    &profile,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			got, err := c.Profile()
			if (err != nil) != tt.wantErr {
				t.Errorf("Profile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Profile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_RebootInstance(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	rebootResp, _ := json.Marshal(domains.CommonResponse{
		Success: true,
		Message: "your instance has been rebooted",
	})

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return(rebootResp, nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	type args struct {
		identifier string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "sending reboot signal without id",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			wantErr: true,
		},
		{
			name: "sending reboot signal with id",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args:    args{identifier: "asasasas"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			if err := c.RebootInstance(tt.args.identifier); (err != nil) != tt.wantErr {
				t.Errorf("RebootInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_ResetPasswordInstance(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	resetPassResp, _ := json.Marshal(domains.CommonResponse{
		Success: true,
		Message: "your password reset successful",
	})

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return(resetPassResp, nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	type args struct {
		identifier  string
		newPassword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "sending reset password invalid data",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			wantErr: true,
		},
		{
			name: "sending reset password valid data",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args: args{
				identifier:  "identifier-example",
				newPassword: "<new-password>",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			if err := c.ResetPasswordInstance(tt.args.identifier, tt.args.newPassword); (err != nil) != tt.wantErr {
				t.Errorf("ResetPasswordInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_SSHKey(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	sshKey := domains.SSHKey{

		Slug:       "new-ssh-key",
		Title:      "New SSH Key",
		PublicKey:  "<this-is-a-public-key>",
		PrivateKey: "<this-is-a-purivate-key>",
	}

	getSSHResp, _ := json.Marshal(domains.CreateOrGetSSHKeysResponse{
		CommonResponse: domains.CommonResponse{
			Success: true,
		},
		Data: sshKey,
	})

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return(getSSHResp, nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domains.SSHKey
		wantErr bool
	}{
		{
			name: "sending invalid ssh data",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "sending valid ssh data",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			args:    args{slug: "new-ssh-key"},
			want:    &sshKey,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			got, err := c.SSHKey(tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("SSHKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SSHKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SSHKeys(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	sshKeys := []domains.SSHKey{
		{
			Slug:       "new-ssh-key",
			Title:      "New SSH Key",
			PublicKey:  "<this-is-a-public-key>",
			PrivateKey: "<this-is-a-purivate-key>",
		},
	}

	getSSHResp, _ := json.Marshal(domains.GetSSHKeysResponse{
		CommonResponse: domains.CommonResponse{
			Success: true,
		},
		Data: sshKeys,
	})

	mclient.EXPECT().NewRequest(gomock.Any(), gomock.Any(), gomock.Any()).Return(new(http.Request), nil)
	mclient.EXPECT().SendRequest(gomock.Any()).Return(getSSHResp, nil)

	type fields struct {
		token     string
		debug     bool
		requester Requester
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domains.SSHKey
		wantErr bool
	}{
		{
			name: "get all ssh keys",
			fields: fields{
				token:     TEST_API_KEY,
				debug:     false,
				requester: mclient,
			},
			want:    sshKeys,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			got, err := c.SSHKeys()
			if (err != nil) != tt.wantErr {
				t.Errorf("SSHKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SSHKeys() got = %v, want %v", got, tt.want)
			}
		})
	}
}
