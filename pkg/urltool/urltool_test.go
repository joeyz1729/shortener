package urltool

import "testing"

func TestGetBasePath(t *testing.T) {
	type args struct {
		URL string
	}
	tests := []struct {
		name         string
		args         args
		wantBasePath string
		wantErr      bool
	}{
		// TODO: Add test cases.
		{ 
			name: "base example",
			args: args{URL:"https://www.baidu.com/11"},
			wantBasePath: "11",
			wantErr: false,
		},
		{ 
			name: "invalid relative path example",
			args: args{URL:"/xxxx/1234"},
			wantBasePath: "",
			wantErr: true,
		},
		{ 
			name: "null string",
			args: args{URL:""},
			wantBasePath: "",
			wantErr: true,
		},
		{ 
			name: "url with query params",
			args: args{URL:"https://www.bilibili.com/video?vid=111"},
			wantBasePath: "video",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBasePath, err := GetBasePath(tt.args.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBasePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotBasePath != tt.wantBasePath {
				t.Errorf("GetBasePath() = %v, want %v", gotBasePath, tt.wantBasePath)
			}
		})
	}
}
