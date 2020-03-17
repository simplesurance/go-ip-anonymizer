package goanonymizer

import (
	"testing"
)

func TestAnonymize_AnonymizeIp(t *testing.T) {

	tests := []struct {
		name    string
		V4Mask  string
		V6Mask  string
		ipStr   string
		want    string
		wantErr bool
	}{
		{name: "success_v4", V4Mask: "255.255.0.0", V6Mask: "ffff:ffff:ffff:fff:0000", ipStr: "123.123.123.123", want: "123.123.0.0"},
		{name: "success_v6", V4Mask: "255.255.0.0", V6Mask: "ffff:ffff:ffff:0::", ipStr: "9681:93f0:c8a9:bbb2:c347:93fc:85ee:ea56", want: "9681:93f0:c8a9::"},
		{name: "success_v4_default", V4Mask: "", V6Mask: "", ipStr: "123.123.123.123", want: "123.123.123.0"},
		{name: "success_v6_default", V4Mask: "", V6Mask: "", ipStr: "9681:93f0:c8a9:bbb2:c347:93fc:85ee:ea56", want: "9681:93f0:c8a9:bbb2::"},
		{name: "error_v4", V4Mask: "", V6Mask: "", ipStr: "test", want: "", wantErr: true},
		{name: "error_v6", V4Mask: "", V6Mask: "", ipStr: "test", want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewAnonymize(tt.V4Mask, tt.V6Mask)
			got, err := a.AnonymizeIp(tt.ipStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("AnonymizeIp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AnonymizeIp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
