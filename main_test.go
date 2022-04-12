package main

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"my_module/my_pckg"
	"testing"
)

func TestNeedTest(t *testing.T) {

	tests := []struct {
		name    string
		want    Simpler
		prepare func(ctrl *gomock.Controller) Simpler
	}{
		{
			name: "ok",
			want: nil,
			prepare: func(ctrl *gomock.Controller) Simpler {
				m := my_pckg.NewMockSimpler(ctrl)
				m.EXPECT().SimpleMethod()
				m.EXPECT().ConcurrentMethod()
				return m
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			m := tt.prepare(ctrl)
			assert.Equalf(t, tt.want, NeedTest(m), "NeedTest()")
		})
	}
}
