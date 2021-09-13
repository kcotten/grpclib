package widgetserver

import (
	"context"
	pb "grpclib/widget"
	"reflect"
	"testing"
	"time"
)

func Test_server_Update(t *testing.T) {
	tctx, cancel := context.WithTimeout(context.Background(), time.Second)
	name := "item1"
	quantity := 7
	request := &pb.UpdateRequest{Name: name, Quantity: int64(quantity)}
	defer cancel()
	type fields struct {
		UnimplementedProcessServer pb.UnimplementedProcessServer
	}
	type args struct {
		ctx context.Context
		in  *pb.UpdateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UpdateResponse
		wantErr bool
	}{
		// Test cases
		{
			name: "Test1",
			args: args{
				ctx: tctx,
				in:  request,
			},
			want: &pb.UpdateResponse{
				Message: "ok",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedProcessServer: tt.fields.UnimplementedProcessServer,
			}
			got, err := s.Update(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
