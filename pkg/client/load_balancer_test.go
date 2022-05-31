package client

import (
	"context"
	"reflect"
	"testing"

	"github.com/HewlettPackard/hpegl-vmaas-cmp-go-sdk/pkg/models"
	gomock "github.com/golang/mock/gomock"
)

func Test_loadBalancerAPIService_CreateLoadBalancer(t *testing.T) {
	ctx1 := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Client APIClientHandler
		Cfg    Configuration
	}
	type args struct {
		ctx     context.Context
		request models.CreateLoadBalancerRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.CreateNetworkLoadBalancerResp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Normal test case 1: Create Router",
			args: args{
				ctx: ctx1,
				request: models.CreateLoadBalancerRequest{
					NetworkLoadBalancer: models.CreateNetworkLoadBalancerRequest{
						Name:            "tf_LB",
						Type:            "nsx-t",
						Description:     "creating LB",
						NetworkServerID: 1,
					},
				},
			},
			want: models.CreateNetworkLoadBalancerResp{
				Success: true,
				NetworkLoadBalancerResp: models.NetworkLoadBalancerResp{
					ID:   1,
					Name: "tf_LB",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &loadBalancerAPIService{
				Client: tt.fields.Client,
				Cfg:    tt.fields.Cfg,
			}
			got, err := lb.CreateLoadBalancer(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadBalancerAPIService.CreateLoadBalancer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadBalancerAPIService.CreateLoadBalancer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadBalancerAPIService_DeleteLoadBalancer(t *testing.T) {
	ctx1 := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Client APIClientHandler
		Cfg    Configuration
	}
	type args struct {
		ctx  context.Context
		lbID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.SuccessOrErrorMessage
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Normal Test case 1: Delete a  router",
			args: args{ctx: ctx1,
				lbID: 1},
			want: models.SuccessOrErrorMessage{
				Success: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &loadBalancerAPIService{
				Client: tt.fields.Client,
				Cfg:    tt.fields.Cfg,
			}
			got, err := lb.DeleteLoadBalancer(tt.args.ctx, tt.args.lbID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadBalancerAPIService.DeleteLoadBalancer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadBalancerAPIService.DeleteLoadBalancer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadBalancerAPIService_GetSpecificLoadBalancers(t *testing.T) {
	ctx1 := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Client APIClientHandler
		Cfg    Configuration
	}
	type args struct {
		ctx  context.Context
		lbID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.GetSpecificNetworkLoadBalancer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get specific LB",
			args: args{
				ctx:  ctx1,
				lbID: 1,
			},
			want: models.GetSpecificNetworkLoadBalancer{
				GetSpecificNetworkLoadBalancerResp: models.GetSpecificNetworkLoadBalancerResp{
					ID:   1,
					Name: "LB-1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &loadBalancerAPIService{
				Client: tt.fields.Client,
				Cfg:    tt.fields.Cfg,
			}
			got, err := lb.GetSpecificLoadBalancers(tt.args.ctx, tt.args.lbID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadBalancerAPIService.GetSpecificLoadBalancers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadBalancerAPIService.GetSpecificLoadBalancers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadBalancerAPIService_CreateLBMonitor(t *testing.T) {
	ctx1 := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Client APIClientHandler
		Cfg    Configuration
	}
	type args struct {
		ctx     context.Context
		request models.CreateLBMonitor
		lbID    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.CreateLBMonitorResp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "create LB monitor",
			args: args{
				ctx: ctx1,
				request: models.CreateLBMonitor{
					CreateLBMonitorReq: models.CreateLBMonitorReq{
						Name:            "Test-Monitor1",
						Description:     "a test monitor",
						MonitorType:     "LBHttpsMonitorProfile",
						MonitorTimeout:  15,
						MonitorInterval: 30,
					},
				},
				lbID: 1,
			},
			want: models.CreateLBMonitorResp{
				Success: true,
				LBMonitorResp: models.LBMonitorResp{
					ID:              1,
					Name:            "Test-Monitor1",
					Description:     "a test monitor",
					MonitorType:     "LBHttpsMonitorProfile",
					MonitorTimeout:  15,
					MonitorInterval: 30,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &loadBalancerAPIService{
				Client: tt.fields.Client,
				Cfg:    tt.fields.Cfg,
			}
			got, err := lb.CreateLBMonitor(tt.args.ctx, tt.args.request, tt.args.lbID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadBalancerAPIService.CreateLBMonitor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadBalancerAPIService.CreateLBMonitor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadBalancerAPIService_DeleteLBMonitor(t *testing.T) {
	ctx1 := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Client APIClientHandler
		Cfg    Configuration
	}
	type args struct {
		ctx         context.Context
		lbID        int
		lbMonitorID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.SuccessOrErrorMessage
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Delete LB monitor",
			args: args{
				ctx:         ctx1,
				lbID:        1,
				lbMonitorID: 2,
			},
			want: models.SuccessOrErrorMessage{
				Success: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &loadBalancerAPIService{
				Client: tt.fields.Client,
				Cfg:    tt.fields.Cfg,
			}
			got, err := lb.DeleteLBMonitor(tt.args.ctx, tt.args.lbID, tt.args.lbMonitorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadBalancerAPIService.DeleteLBMonitor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadBalancerAPIService.DeleteLBMonitor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadBalancerAPIService_GetSpecificLBMonitor(t *testing.T) {
	ctx1 := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Client APIClientHandler
		Cfg    Configuration
	}
	type args struct {
		ctx         context.Context
		lbID        int
		lbmonitorID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.GetSpecificNetworkLoadBalancer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GET LB monitor",
			args: args{
				ctx:         ctx1,
				lbID:        1,
				lbmonitorID: 2,
			},
			want: models.GetSpecificNetworkLoadBalancer{
				GetSpecificNetworkLoadBalancerResp: models.GetSpecificNetworkLoadBalancerResp{
					ID:          1,
					Name:        "LB monitor",
					Description: "LB",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &loadBalancerAPIService{
				Client: tt.fields.Client,
				Cfg:    tt.fields.Cfg,
			}
			got, err := lb.GetSpecificLBMonitor(tt.args.ctx, tt.args.lbID, tt.args.lbmonitorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadBalancerAPIService.GetSpecificLBMonitor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadBalancerAPIService.GetSpecificLBMonitor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadBalancerAPIService_CreateLBProfile(t *testing.T) {
	ctx1 := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Client APIClientHandler
		Cfg    Configuration
	}
	type args struct {
		ctx     context.Context
		request models.CreateLBProfile
		lbID    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.CreateLBProfileResp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Create LB profile",
			args: args{
				ctx: ctx1,
				request: models.CreateLBProfile{
					CreateLBProfileReq: models.CreateLBProfileReq{
						Name:        "LB profile",
						Description: "creating",
					},
				},
			},
			want: models.CreateLBProfileResp{
				Success: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &loadBalancerAPIService{
				Client: tt.fields.Client,
				Cfg:    tt.fields.Cfg,
			}
			got, err := lb.CreateLBProfile(tt.args.ctx, tt.args.request, tt.args.lbID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadBalancerAPIService.CreateLBProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadBalancerAPIService.CreateLBProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadBalancerAPIService_DeleteLBProfile(t *testing.T) {
	ctx1 := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Client APIClientHandler
		Cfg    Configuration
	}
	type args struct {
		ctx         context.Context
		lbID        int
		lbProfileID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.SuccessOrErrorMessage
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Delete LB Profile",
			args: args{ctx: ctx1,
				lbID:        1,
				lbProfileID: 2},
			want: models.SuccessOrErrorMessage{
				Success: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &loadBalancerAPIService{
				Client: tt.fields.Client,
				Cfg:    tt.fields.Cfg,
			}
			got, err := lb.DeleteLBProfile(tt.args.ctx, tt.args.lbID, tt.args.lbProfileID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadBalancerAPIService.DeleteLBProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadBalancerAPIService.DeleteLBProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadBalancerAPIService_GetSpecificLBProfile(t *testing.T) {
	ctx1 := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Client APIClientHandler
		Cfg    Configuration
	}
	type args struct {
		ctx         context.Context
		lbID        int
		lbProfileID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.GetLBSpecificProfile
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GET LB Profile",
			args: args{ctx: ctx1,
				lbID:        1,
				lbProfileID: 2},
			want: models.GetLBSpecificProfile{
				GetLBSpecificProfilesResp: models.GetLBSpecificProfilesResp{
					ID:   2,
					Name: "LB Profile",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &loadBalancerAPIService{
				Client: tt.fields.Client,
				Cfg:    tt.fields.Cfg,
			}
			got, err := lb.GetSpecificLBProfile(tt.args.ctx, tt.args.lbID, tt.args.lbProfileID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadBalancerAPIService.GetSpecificLBProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadBalancerAPIService.GetSpecificLBProfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadBalancerAPIService_CreateLBPool(t *testing.T) {
	ctx1 := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Client APIClientHandler
		Cfg    Configuration
	}
	type args struct {
		ctx     context.Context
		request models.CreateLBPool
		lbID    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.CreateLBPoolResp
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "LB Pool",
			args: args{ctx: ctx1,
				request: models.CreateLBPool{
					CreateLBPoolReq: models.CreateLBPoolReq{
						Name:        "LB Pool",
						Description: "creating Pool",
					},
				},
			},
			want: models.CreateLBPoolResp{
				Success: true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &loadBalancerAPIService{
				Client: tt.fields.Client,
				Cfg:    tt.fields.Cfg,
			}
			got, err := lb.CreateLBPool(tt.args.ctx, tt.args.request, tt.args.lbID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadBalancerAPIService.CreateLBPool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadBalancerAPIService.CreateLBPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadBalancerAPIService_DeleteLBPool(t *testing.T) {
	ctx1 := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Client APIClientHandler
		Cfg    Configuration
	}
	type args struct {
		ctx      context.Context
		lbID     int
		lbPoolID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.SuccessOrErrorMessage
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Delete LB Pool",
			args: args{ctx: ctx1,
				lbID:     1,
				lbPoolID: 2,
			},
			want: models.SuccessOrErrorMessage{
				Success: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &loadBalancerAPIService{
				Client: tt.fields.Client,
				Cfg:    tt.fields.Cfg,
			}
			got, err := lb.DeleteLBPool(tt.args.ctx, tt.args.lbID, tt.args.lbPoolID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadBalancerAPIService.DeleteLBPool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadBalancerAPIService.DeleteLBPool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadBalancerAPIService_GetSpecificLBPool(t *testing.T) {
	ctx1 := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type fields struct {
		Client APIClientHandler
		Cfg    Configuration
	}
	type args struct {
		ctx      context.Context
		lbID     int
		lbPoolID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.GetSpecificLBPool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GET LB Pool",
			args: args{
				ctx:      ctx1,
				lbID:     1,
				lbPoolID: 2,
			},
			want: models.GetSpecificLBPool{
				GetSpecificLBPoolResp: models.GetSpecificLBPoolResp{
					ID:   2,
					Name: "GET LB POOL",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lb := &loadBalancerAPIService{
				Client: tt.fields.Client,
				Cfg:    tt.fields.Cfg,
			}
			got, err := lb.GetSpecificLBPool(tt.args.ctx, tt.args.lbID, tt.args.lbPoolID)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadBalancerAPIService.GetSpecificLBPool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadBalancerAPIService.GetSpecificLBPool() = %v, want %v", got, tt.want)
			}
		})
	}
}
