package crypt

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestCrypt(t *testing.T) {
	type args struct {
		password string
		salt     string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"case 1", args{"nutmeg", "Mi"}, "MiqkFWCm1fNJI", false},
		{"case 2", args{"ellen1", "ri"}, "ri79kNd7V6.Sk", false},
		{"case 3", args{"Sharon", "./"}, "./UY9Q7TvYJDg", false},
		{"case 4", args{"norahs", "am"}, "amfIADT2iqjA.", false},
		{"case 5", args{"norash", "7a"}, "7aWAMMtfbjEJo", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Crypt(tt.args.password, tt.args.salt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Crypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Crypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrypt2(t *testing.T) {
	bytesRead, _ := ioutil.ReadFile("resources/passwords.txt")
	file_content := string(bytesRead)
	lines := strings.Split(file_content, "\n")
	count := 0
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Split(line, " ")
		password := parts[0]
		salt := parts[1][:2]
		hash, _ := Crypt(password, salt)
		if hash != parts[1] {
			t.Errorf("Crypt() = %v, want %v", hash, parts[1])
		} else {
			count++
		}
	}

	if count < len(lines)-1 {
		t.Errorf("In correct count")
	}
}
