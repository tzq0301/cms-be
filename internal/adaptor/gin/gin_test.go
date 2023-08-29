package ginadaptor

import (
	"context"
	"testing"
	"time"
)

func TestAdaptor_Run(t *testing.T) {
	type fields struct {
		port int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "cancel by error",
			fields: fields{
				port: 808000,
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "cancel by timeout",
			fields: fields{
				port: 9000,
			},
			args: args{
				ctx: func() context.Context {
					ctx := context.Background()
					ctx, _ = context.WithTimeout(ctx, time.Millisecond*100)
					return ctx
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := New(context.Background(), tt.fields.port, "")
			if err := a.Run(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
