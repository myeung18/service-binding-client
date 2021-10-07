package convert

import (
	"fmt"
	"github.com/myeung18/service-binding-client/internal/fileconfig"
	"testing"
)

func TestPostgreSQLConnectionStringConverter_Convert(t *testing.T) {
	type args struct {
		binding fileconfig.ServiceBinding
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Correct connection string returned with empty options",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "postgresql",
					Provider:    "Crunchy Bridges",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"password": "password",
						"srv":      "true",
						"options":  "",
						"database": "local-db",
					},
				},
			},
			want: "host=example.com:10011 user=a-db-user password=password dbname=local-db",
		},
		{
			name: "Correct connection string returned without password",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "postgresql",
					Provider:    "Crunchy Bridges",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"srv":      "true",
						"options":  "sslmode=disable&connect_timeout=10",
						"database": "local-db",
					},
				},
			},
			want: "host=example.com:10011 user=a-db-user dbname=local-db sslmode=disable connect_timeout=10",
		},
		{
			name: "Correct connection string returned with special char escaping",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "postgresql",
					Provider:    "Crunchy Bridges",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"password": "password'",
						"srv":      "true",
						"options":  "some-db-options_that_is_invalid",
						"database": "local-db",
					},
				},
			},
			want: "host=example.com:10011 user=a-db-user password=password\\' dbname=local-db",
		},
		{
			name: "Invalid_or_incomplete options are ignored",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "postgresql",
					Provider:    "Crunchy Bridges",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"password": "password'",
						"srv":      "true",
						"options":  "option1=value1&option2",
						"database": "local-db",
					},
				},
			},
			want: "host=example.com:10011 user=a-db-user password=password\\' dbname=local-db option1=value1",
		},
		{
			name: "Options contain an invalid option - with no key",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "postgresql",
					Provider:    "Crunchy Bridges",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"password": "password'",
						"srv":      "true",
						"options":  "=value1&option2=",
						"database": "local-db",
					},
				},
			},
			want: "host=example.com:10011 user=a-db-user password=password\\' dbname=local-db",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &PostgreSQLConnectionStringConverter{}
			if got := m.Convert(tt.args.binding); got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
			fmt.Println(m.Convert(tt.args.binding))
		})
	}
}
