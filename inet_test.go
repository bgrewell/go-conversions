package conversions

import (
	"reflect"
	"testing"
)

func TestInet4_aton(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name       string
		args       args
		wantIp_int uint32
	}{
		{
			name: "localhost",
			args: args{ip: "1.0.0.127"},
			wantIp_int: 16777343,
		},
		{
			name: "192.168.253.1",
			args: args{ip: "1.253.168.192"},
			wantIp_int: 33401024,
		},
	}
	for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
		if gotIp_int := Inet4_aton(tt.args.ip); gotIp_int != tt.wantIp_int {
			t.Errorf("Inet4_aton() = %v, want %v", gotIp_int, tt.wantIp_int)
		}
	})
	}
}

func TestInet4_haton(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name       string
		args       args
		wantIp_int uint32
	}{
		{
			name: "localhost",
			args: args{ip: "127.0.0.1"},
			wantIp_int: 16777343,
		},
		{
			name: "192.168.253.1",
			args: args{ip: "192.168.253.1"},
			wantIp_int: 33401024,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIp_int := Inet4_haton(tt.args.ip); gotIp_int != tt.wantIp_int {
				t.Errorf("Inet4_haton() = %v, want %v", gotIp_int, tt.wantIp_int)
			}
		})
	}
}

func TestInet4_ntoa(t *testing.T) {
	type args struct {
		ip uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "localhost",
			args: args{ip: 16777343},
			want: "1.0.0.127",
		},
		{
			name: "192.168.253.1",
			args: args{ip: 33401024},
			want: "1.253.168.192",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Inet4_ntoa(tt.args.ip); got != tt.want {
				t.Errorf("Inet4_ntoa() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInet4_ntoha(t *testing.T) {
	type args struct {
		ip uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "localhost",
			args: args{ip: 16777343},
			want: "127.0.0.1",
		},
		{
			name: "192.168.253.1",
			args: args{ip: 33401024},
			want: "192.168.253.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Inet4_ntoha(tt.args.ip); got != tt.want {
				t.Errorf("Inet4_ntoha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIpv4MaskString(t *testing.T) {
	type args struct {
		m []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "class A",
			args: args{m: []byte{0xFF, 0x00, 0x00, 0x00}},
			want: "255.0.0.0",
		},
		{
			name: "class C",
			args: args{m: []byte{0xFF, 0xFF, 0xFF, 0x00}},
			want: "255.255.255.0",
		},
		{
			name: "/27",
			args: args{m: []byte{255, 255, 255, 224}},
			want: "255.255.255.224",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ipv4MaskString(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ipv4MaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}