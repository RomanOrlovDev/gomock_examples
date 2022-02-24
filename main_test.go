package main

import (
	"AwesomeProject/item"
	"fmt"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestIterateThroughItems(t *testing.T) {
	type args struct {
		is item.Isnterface
	}
	tests := []struct {
		name string
		args func(ctrl *gomock.Controller) args
		want bool
	}{
		{
			name: "gen err",
			args: func(ctrl *gomock.Controller) args {
				//is := item.NewItems()
				//is.Add(item.NewItem(1, "Some item"))
				is := item.NewMockIsnterface(ctrl)
				i := item.NewMockFacadeItem(ctrl)
				i.EXPECT().GetNameWithValidation().Return("", fmt.Errorf("validation error occurred"))

				is.EXPECT().Elements().Return([]item.FacadeItem{i}).Times(1)

				return args{
					is: is,
				}
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			args := tt.args(ctrl)
			if got := IterateThroughItems(&args.is); got != tt.want {
				t.Errorf("IterateThroughItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyStruct_CreateMyStruct(t *testing.T) {
	type fields struct {
		a string
	}
	tests := []struct {
		name   string
		fields fields
		want   MyStruct
	}{
		{
			name: "ok",
			fields: fields{
				a: "some value",
			},
			want: MyStruct{
				a: "some value",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MyStruct{
				a: tt.fields.a,
			}
			if got := m.CreateMyStruct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMyStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
