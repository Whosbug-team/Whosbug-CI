package crypto

import (
	"fmt"
	"testing"

	"git.woa.com/bkdevops/whosbug/config"
)

func TestBase64Decrypt(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "test-decrypt",
			args:    args{text: "+9MI3HCSnUgg0Q=="},
			want:    "to_encrypt",
			wantErr: false,
		},
	}

	config.WhosbugConfig.ProjectID = "test-project"
	config.WhosbugConfig.CryptoKey = ""

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Base64Decrypt(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Base64Decrypt() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(fmt.Sprintf("Base64Decrypt() = %v, want %v", got, tt.want))
			}
		})
	}
}

func TestBase64Encrypt(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-encrypt",
			args: args{text: "to_encrypt"},
			want: "+9MI3HCSnUgg0Q==",
		},
	}

	config.WhosbugConfig.ProjectID = "test-project"
	config.WhosbugConfig.CryptoKey = ""

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base64Encrypt(tt.args.text); got != tt.want {
				t.Errorf("Base64Encrypt() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(fmt.Sprintf("Base64Encrypt() = %v, want %v", got, tt.want))
			}
		})
	}
}
