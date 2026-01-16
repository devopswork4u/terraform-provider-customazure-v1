package client

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
)

type GraphClient struct {
	Client *msgraphsdk.GraphServiceClient
}

func NewGraphClient(
	tenantID string,
	clientID string,
	clientSecret string,
) (*GraphClient, error) {

	cred, err := azidentity.NewClientSecretCredential(
		tenantID,
		clientID,
		clientSecret,
		nil,
	)
	if err != nil {
		return nil, err
	}

	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(
		cred,
		[]string{"https://graph.microsoft.com/.default"},
	)
	if err != nil {
		return nil, err
	}

	return &GraphClient{
		Client: client,
	}, nil
}
