package genpass

import (
	"os"
	"testing"
)

func GenerateDirTest(t *testing.T) {
	err := GenerateDir()
	if err != nil {
		t.Errorf("unexpected error")
	}
}

func GenerateFileCSVTest(t *testing.T) {
	err := GenerateFileCSV()
	if err != nil {
		t.Errorf("unexpected error")
	}
}

func WriteCSVTest(t *testing.T) {
	type args struct {
		title string
		n     int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{"Test1", 16}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteCSV(tt.args.title, tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("WriteCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	os.RemoveAll("test-will-delete")

}
