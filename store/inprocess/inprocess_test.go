package inprocess_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/telia-oss/sidecred"
	secretstore "github.com/telia-oss/sidecred/store/inprocess"
)

func TestInProcessStore(t *testing.T) {
	var (
		teamName     = "team-name"
		secret       = &sidecred.Credential{Name: "secret-name", Value: "secret-value"}
		pathTemplate = "/concourse/{{ .Namespace }}/{{ .Name }}"
		secretPath   = "/concourse/team-name/secret-name"
	)

	tests := []struct {
		description    string
		config         json.RawMessage
		secretTemplate string
		secretPath     string
	}{
		{
			description:    "works as expected",
			secretTemplate: pathTemplate,
			secretPath:     secretPath,
		},
		{
			description:    "supports arbitrary path templates",
			secretTemplate: "concourse.{{ .Namespace }}.{{ .Name }}",
			secretPath:     "concourse.team-name.secret-name",
		},
		{
			description:    "supports setting secret template from config",
			config:         []byte(`{"secret_template":"{{ .Namespace }}!?!{{ .Name }}"}`),
			secretTemplate: "concourse.{{ .Namespace }}.{{ .Name }}",
			secretPath:     "team-name!?!secret-name",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			store := secretstore.New(secretstore.WithSecretTemplate(tc.secretTemplate))

			path, err := store.Write(context.TODO(), teamName, secret, tc.config)
			assert.NoError(t, err)
			assert.Equal(t, tc.secretPath, path)

			actual, found, err := store.Read(context.TODO(), path, nil)
			assert.Nil(t, err)
			assert.Equal(t, true, found)
			assert.Equal(t, secret.Value, actual)

			err = store.Delete(context.TODO(), path, nil)
			assert.Nil(t, err)
		})
	}
}
