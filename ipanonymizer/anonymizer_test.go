package ipanonymizer

import (
	"net"
	"reflect"
	"testing"
)

func TestIPString(t *testing.T) {
	tests := []struct {
		name    string
		V4Mask  net.IPMask
		V6Mask  net.IPMask
		ip      string
		want    string
		wantErr bool
	}{
		{
			name:   "ipv4_16mask",
			V4Mask: net.IPv4Mask(255, 255, 0, 0),
			V6Mask: defaultIPv6Mask,
			ip:     "8.8.8.8",
			want:   "8.8.0.0",
		},

		{
			name:   "ipv4_30mask",
			V4Mask: net.CIDRMask(30, 32),
			V6Mask: defaultIPv6Mask,
			ip:     "192.168.3.129",
			want:   "192.168.3.128",
		},

		{
			name:   "ipv6_64mask",
			V4Mask: defaultIPv4Mask,
			V6Mask: net.CIDRMask(64, 128),
			ip:     "2a00:1450:4001:820::200e",
			want:   "2a00:1450:4001:820::",
		},
		{
			name:   "ipv6_127mask",
			V4Mask: defaultIPv4Mask,
			V6Mask: net.CIDRMask(127, 128),
			ip:     "2a00:1450:4001:a20::100f",
			want:   "2a00:1450:4001:a20::100e",
		},

		{
			name:    "invalid_ipstr",
			V4Mask:  defaultIPv4Mask,
			V6Mask:  defaultIPv6Mask,
			ip:      "hello",
			wantErr: true,
		},

		{
			name:    "invalid_ipv4_addr",
			V4Mask:  defaultIPv4Mask,
			V6Mask:  defaultIPv6Mask,
			ip:      "127.0.0.512",
			wantErr: true,
		},

		{
			name:    "invalid_ipv6_addr",
			V4Mask:  defaultIPv4Mask,
			V6Mask:  defaultIPv6Mask,
			ip:      "::fg",
			wantErr: true,
		},

		{
			name:    "empty_ipstr",
			V4Mask:  defaultIPv4Mask,
			V6Mask:  defaultIPv6Mask,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			anonymizer := NewWithMask(tt.V4Mask, tt.V6Mask)

			anonIP, err := anonymizer.IPString(tt.ip)
			if (err != nil) != tt.wantErr {
				t.Fatalf("AnonymizeIp() error = %v, wantErr %v", err, tt.wantErr)
			}

			if anonIP != tt.want {
				t.Errorf("AnonymizeIp() got = %v, want %v", anonIP, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	defAnonymizer := New()
	anonymizer := NewWithMask(defaultIPv4Mask, defaultIPv6Mask)

	if !reflect.DeepEqual(defAnonymizer, anonymizer) {
		t.Errorf("New() got = %+v, want %+v", anonymizer, defAnonymizer)
	}
}
