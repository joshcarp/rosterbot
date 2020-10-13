package secrets

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

const projectID = `joshcarp-installer`

func CreateSecret(name string, payload []byte) error {
	ctx := context.Background()
	secretClinet, _ := secretmanager.NewClient(context.Background())
	createSecretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", projectID),
		SecretId: name,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}
	_, b := secretClinet.CreateSecret(ctx, createSecretReq)
	if b != nil{
		return b
	}
	return UpdateSecret(name, payload)
}

func GetSecret(name string) (*secretmanagerpb.Secret, error) {
	secretClinet, _ := secretmanager.NewClient(context.Background())
	return secretClinet.GetSecret(context.Background(), &secretmanagerpb.GetSecretRequest{Name: fmt.Sprintf("projects/%s/secrets/%s", projectID, name)})
}

func GetSecretData(name string) ([]byte, error) {
	secretClinet, _ := secretmanager.NewClient(context.Background())
	s, err := secretClinet.AccessSecretVersion(context.Background(), &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s/versions/latest", projectID, name),
	})
	if err != nil {
		return nil, err
	}
	return s.Payload.Data, nil
}

func UpdateSecret(name string, payload []byte ) error {
	secretClinet, _ := secretmanager.NewClient(context.Background())
	secret, err := GetSecret(name)
	if err != nil {
		return err
	}
	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: payload,
		},
	}
	_, err = secretClinet.AddSecretVersion(context.Background(), addSecretVersionReq)
	return err
}
