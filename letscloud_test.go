package letscloud

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/letscloud-community/letscloud-go/httpclient"
)

func TestLetsCloud_APIKey(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	mclient.EXPECT().APIKey().Return(TEST_API_KEY)

	type fields struct {
		debug     bool
		requester Requester
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get api key",
			fields: fields{
				debug:     false,
				requester: mclient,
			},
			want: TEST_API_KEY,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			if got := c.APIKey(); got != tt.want {
				t.Errorf("APIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLetsCloud_SetAPIKey(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	mclient.EXPECT().SetAPIKey(gomock.Any())

	type fields struct {
		debug     bool
		requester Requester
	}
	type args struct {
		ak string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "set empty api key",
			fields: fields{
				debug:     false,
				requester: mclient,
			},
			wantErr: true,
		},
		{
			name: "set api key",
			fields: fields{
				debug:     false,
				requester: mclient,
			},
			args:    args{ak: TEST_API_KEY},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LetsCloud{
				debug:     tt.fields.debug,
				requester: tt.fields.requester,
			}
			if err := c.SetAPIKey(tt.args.ak); (err != nil) != tt.wantErr {
				t.Errorf("SetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLetsCloud_SetTimeout(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	mclient.EXPECT().SetTimeout(time.Duration(10))

	type fields struct {
		debug     bool
		requester Requester
	}
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "set timeout",
			fields: fields{
				debug:     false,
				requester: mclient,
			},
			args: args{
				d: 10,
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
			if err := c.SetTimeout(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("SetTimeout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	mc := gomock.NewController(t)
	defer mc.Finish()

	mclient := httpclient.NewMockRequester(mc)

	type args struct {
		ak string
	}
	tests := []struct {
		name    string
		args    args
		want    *LetsCloud
		wantErr bool
	}{
		{
			name:    "new letscloud instance without api key",
			want:    nil,
			wantErr: true,
		},
		{
			name: "new letscloud instance with api key",
			want: &LetsCloud{
				debug:     false,
				requester: mclient,
			},
			args: args{
				ak: TEST_API_KEY,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.ak)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				if got != nil {
					return
				}
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}
