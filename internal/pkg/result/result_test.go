package result

import (
	"errors"
	"reflect"
	"testing"

	pkgerrors "github.com/pkg/errors"
)

func TestSuccess(t *testing.T) {
	type args[T any] struct {
		options []Option[int]
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want Result[int]
	}
	tests := []testCase[int]{
		{
			name: "no data",
			args: args[int]{},
			want: Result[int]{
				Code:    0,
				Message: "success",
				Data:    nil,
			},
		},
		{
			name: "has data",
			args: args[int]{
				[]Option[int]{WithData(100)},
			},
			want: Result[int]{
				Code:    0,
				Message: "success",
				Data: func() *int {
					i := 100
					return &i
				}(),
			},
		},
		{
			name: "has nil option",
			args: args[int]{
				[]Option[int]{WithData(100), nil},
			},
			want: Result[int]{
				Code:    0,
				Message: "success",
				Data: func() *int {
					i := 100
					return &i
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Success(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Success() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFail(t *testing.T) {
	type args[T any] struct {
		options []Option[int]
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want Result[int]
	}
	tests := []testCase[int]{
		{
			name: "no data",
			args: args[int]{},
			want: Result[int]{
				Code:    1,
				Message: "fail",
				Data:    nil,
			},
		},
		{
			name: "has data",
			args: args[int]{
				[]Option[int]{WithData(100)},
			},
			want: Result[int]{
				Code:    1,
				Message: "fail",
				Data: func() *int {
					i := 100
					return &i
				}(),
			},
		},
		{
			name: "has nil option",
			args: args[int]{
				[]Option[int]{WithData(100), nil},
			},
			want: Result[int]{
				Code:    1,
				Message: "fail",
				Data: func() *int {
					i := 100
					return &i
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fail(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFrom(t *testing.T) {
	type args[T any] struct {
		err     error
		options []Option[int]
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want Result[int]
	}
	tests := []testCase[int]{
		{
			name: "nil err",
			args: args[int]{
				err: nil,
			},
			want: Result[int]{
				Code:    0,
				Message: "success",
				Data:    nil,
			},
		},
		{
			name: "not defined err",
			args: args[int]{
				err: errors.New("1"),
			},
			want: Result[int]{
				Code:    1,
				Message: "fail",
				Data:    nil,
			},
		},
		{
			name: "is defined err",
			args: args[int]{
				err: ErrInvalidParams,
			},
			want: Result[int]{
				Code:    ErrInvalidParams.Code,
				Message: ErrInvalidParams.Message,
				Data:    nil,
			},
		},
		{
			name: "is joined defined err",
			args: args[int]{
				err: errors.Join(ErrInvalidParams, errors.New("hello")),
			},
			want: Result[int]{
				Code:    ErrInvalidParams.Code,
				Message: ErrInvalidParams.Message,
				Data:    nil,
			},
		},
		{
			name: "is wrapped defined err",
			args: args[int]{
				err: pkgerrors.Wrap(ErrInvalidParams, "world"),
			},
			want: Result[int]{
				Code:    ErrInvalidParams.Code,
				Message: ErrInvalidParams.Message,
				Data:    nil,
			},
		},
		{
			name: "is wrapped and wrapped defined err",
			args: args[int]{
				err: pkgerrors.Wrap(pkgerrors.Wrap(ErrInvalidParams, "world"), "hello"),
			},
			want: Result[int]{
				Code:    ErrInvalidParams.Code,
				Message: ErrInvalidParams.Message,
				Data:    nil,
			},
		},
		{
			name: "is wrapped and wrapped defined err with data",
			args: args[int]{
				err:     pkgerrors.Wrap(pkgerrors.Wrap(ErrInvalidParams, "world"), "hello"),
				options: []Option[int]{WithData(100)},
			},
			want: Result[int]{
				Code:    ErrInvalidParams.Code,
				Message: ErrInvalidParams.Message,
				Data: func() *int {
					i := 100
					return &i
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := From(tt.args.err, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("From() = %v, want %v", got, tt.want)
			}
		})
	}
}
