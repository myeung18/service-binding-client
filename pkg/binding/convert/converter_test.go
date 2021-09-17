package convert

import (
	"github.com/myeung18/service-binding-client/pkg/binding/internal/fileconfig"
	"testing"
)

func TestMongoDBConverter_Convert(t *testing.T) {
	type args struct {
		binding fileconfig.ServiceBinding
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Correct connection string returned",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "mongodb",
					Provider:    "atlas",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"password": "password",
						"srv":      "true",
						"options":  "some-db-options",
						"database": "local-db",
					},
				},
			},
			want: "mongodb+srv://a-db-user:password@example.com:10011/local-db?some-db-options",
		},
		{
			name: "Correct connection string returned - no options",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "mongodb",
					Provider:    "atlas",
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
			want: "mongodb+srv://a-db-user:password@example.com:10011/local-db",
		},
		{
			name: "Correct connection string returned - no database",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "mongodb",
					Provider:    "atlas",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"password": "password",
						"srv":      "true",
						"options":  "some-db-options",
						"database": "",
					},
				},
			},
			want: "mongodb+srv://a-db-user:password@example.com:10011/?some-db-options",
		},
		{
			name: "Correct connection string returned - no database and options",
			args: args{
				binding: fileconfig.ServiceBinding{
					Name:        "local",
					BindingType: "mongodb",
					Provider:    "atlas",
					Properties: map[string]string{
						"host":     "example.com:10011",
						"username": "a-db-user",
						"password": "password",
						"srv":      "true",
						"options":  "",
						"database": "",
					},
				},
			},
			want: "mongodb+srv://a-db-user:password@example.com:10011",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MongoDBConverter{}
			if got := m.Convert(tt.args.binding); got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}