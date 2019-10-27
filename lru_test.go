package lrucache

import (
	"github.com/miguelbemartin/lru-cache/models"
	"reflect"
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name    string
		args    args
		want    *client
		wantErr bool
	}{
		{
			name: "error size negative",
			args: args{
				size: -1,
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "error size 0",
			args: args{
				size: 0,
			},
			want: nil,
			wantErr: true,
		},
		{
			name: "correct",
			args: args{
				size: 1,
			},
			want: &client{
				size: 1,
				stream: &models.Stream{
					Length: 0,
					Limit:  1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLRUCache(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLRUCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLRUCache() got = %v, want %v", got, tt.want)
			}
		})
	}
}