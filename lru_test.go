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
			want:    nil,
			wantErr: true,
		},
		{
			name: "error size 0",
			args: args{
				size: 0,
			},
			want:    nil,
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

func Test_client_Set(t *testing.T) {
	type fields struct {
		size   int
		stream *models.Stream
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantErr      bool
		sizeExpected int
	}{
		{
			name: "correct",
			fields: fields{
				size: 1,
				stream: &models.Stream{
					Length: 0,
					Limit:  1,
				},
			},
			args: args{
				key:   "key",
				value: "value",
			},
			wantErr:      false,
			sizeExpected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				size:   tt.fields.size,
				stream: tt.fields.stream,
			}
			if err := c.Set(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if c.stream.Length != tt.sizeExpected {
				t.Errorf("size expected %v, size found %v", tt.sizeExpected, c.stream.Length)
				return
			}
		})
	}
}

func Test_client_Get(t *testing.T) {
	type fields struct {
		size   int
		stream *models.Stream
	}
	type args struct {
		key string
	}
	tests := []struct {
		name          string
		fields        fields
		key           string
		valueExpected int
		wantErr       bool
	}{
		{
			name: "correct",
			fields: fields{
				size: 2,
				stream: &models.Stream{
					Length: 0,
					Limit:  2,
				},
			},
			key:           "key-1",
			valueExpected: 1,
			wantErr:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				size:   tt.fields.size,
				stream: tt.fields.stream,
			}
			err := c.Set(tt.key, tt.valueExpected)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			valueFound, err := c.Get(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if valueFound != tt.valueExpected {
				t.Errorf("value expected %v, value found %v", tt.valueExpected, valueFound)
				return
			}

		})
	}
}
