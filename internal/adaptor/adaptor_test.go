package adaptor

import (
	"context"
	"errors"
	"testing"
	"time"
)

var (
	mockErr = errors.New("error")
)

type errAdaptor struct {
}

func (_ errAdaptor) Run(_ context.Context) error {
	return mockErr
}

type blockingAdaptor struct {
}

func (_ blockingAdaptor) Run(ctx context.Context) error {
	select {
	case <-ctx.Done():
	}
	return nil
}

func TestRun(t *testing.T) {
	type args struct {
		ctx      context.Context
		adaptors []Adaptor
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error occurred",
			args: args{
				ctx:      context.Background(),
				adaptors: []Adaptor{blockingAdaptor{}, blockingAdaptor{}, errAdaptor{}, blockingAdaptor{}},
			},
			wantErr: true,
		},
		{
			name: "no error occurred",
			args: args{
				ctx: func() context.Context {
					ctx := context.Background()
					ctx, _ = context.WithTimeout(ctx, time.Millisecond*100)
					return ctx
				}(),
				adaptors: []Adaptor{blockingAdaptor{}, blockingAdaptor{}, blockingAdaptor{}, blockingAdaptor{}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Run(tt.args.ctx, tt.args.adaptors...); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
