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

var object1 = ObjectInfoType{"1","2","3","4","5","6","7","8"}
var object2 = ObjectInfoType{"9","10","11","12","13","14","15","16"}


func Test_postObjects(t *testing.T) {
	type args struct {
		projectId      string
		releaseVersion string
		commitHash     string
		objects        []ObjectInfoType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test1", args{"whosbug_test_1", "1.0.0", "0eaa40f", []ObjectInfoType{object1, object2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := postObjects(tt.args.projectId, tt.args.releaseVersion, tt.args.commitHash, tt.args.objects); (err != nil) != tt.wantErr {
				t.Errorf("postObjects() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}