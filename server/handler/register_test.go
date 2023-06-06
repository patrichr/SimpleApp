package handler

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"errors"

	pb "github.com/patrichr/SimpleApp/grpc"
	"github.com/patrichr/SimpleApp/server/model"
)

func TestServer_Register(t *testing.T) {
	type fields struct {
		repo Repo
	}
	type args struct {
		ctx context.Context
		in  *pb.RegRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.RegResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				repo: &RepoMock{
					RegisterFunc: func(ctx context.Context, req model.Register) (int, error) {
						return 0, nil
					},
				},
			},
			args: args{
				ctx: context.Background(),
				in: &pb.RegRequest{
					TaskName:    "Make it",
					OwnerName:   "Me",
					Description: "create simple to do app",
					Status:      "0",
				},
			},
			want: &pb.RegResponse{
				HttpStatus: http.StatusCreated,
				Message:    "successfully registered",
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				repo: &RepoMock{
					RegisterFunc: func(ctx context.Context, req model.Register) (int, error) {
						return 0, errors.New("error registration")
					},
				},
			},
			args: args{
				ctx: context.Background(),
				in: &pb.RegRequest{
					TaskName:    "Make it",
					OwnerName:   "Me",
					Description: "create simple to do app",
					Status:      "0",
				},
			},
			want: &pb.RegResponse{
				HttpStatus: http.StatusInternalServerError,
				Message:    "error registration",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				repo: tt.fields.repo,
			}
			got, err := s.Register(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
