package whosbugAssigns

import "testing"


func Test__genToken(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"test empty", "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := _genToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("_genToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("_genToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLatestRelease(t *testing.T) {
	type args struct {
		projectId string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"no found", args{projectId: "whosbug_test_1"}, "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLatestRelease(tt.args.projectId)
			if (err != nil) != tt.wantErr {
				t.Errorf("getLatestRelease() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getLatestRelease() got = %v, want %v", got, tt.want)
			}
		})
	}
}