package service

import (
	"context"

	"caas-aks-api-go/models"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

func CreateCluster(req models.CreateClusterRequest) (*armcontainerservice.ManagedCluster, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	client, err := armcontainerservice.NewManagedClustersClient(req.SubscriptionID, cred, nil)
	if err != nil {
		return nil, err
	}

	params := armcontainerservice.ManagedCluster{
		Location: &req.Location,
		Properties: &armcontainerservice.ManagedClusterProperties{
			DNSPrefix: &req.ClusterName,
			AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
				{
					Name:   toPtr("agentpool"),
					Count:  toPtr(req.NodeCount),
					VMSize: toPtr("Standard_DS2_v2"),
					Mode:   toPtr(armcontainerservice.AgentPoolModeSystem),
				},
			},
		},
	}

	pollerResp, err := client.BeginCreateOrUpdate(context.TODO(), req.ResourceGroup, req.ClusterName, params, nil)
	if err != nil {
		return nil, err
	}

	result, err := pollerResp.PollUntilDone(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return &result.ManagedCluster, nil
}

func GetCluster(subscriptionID, resourceGroup, name string) (*armcontainerservice.ManagedCluster, error) {
	cred, _ := azidentity.NewDefaultAzureCredential(nil)
	client, _ := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, nil)
	resp, err := client.Get(context.TODO(), resourceGroup, name, nil)
	if err != nil {
		return nil, err
	}
	return &resp.ManagedCluster, nil
}

func ListClusters(subscriptionID, resourceGroup string) ([]*armcontainerservice.ManagedCluster, error) {
	cred, _ := azidentity.NewDefaultAzureCredential(nil)
	client, _ := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, nil)

	pager := client.NewListByResourceGroupPager(resourceGroup, nil)
	var clusters []*armcontainerservice.ManagedCluster
	for pager.More() {
		page, err := pager.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}
		clusters = append(clusters, page.Value...)
	}
	return clusters, nil
}

func toPtr[T any](v T) *T {
	return &v
}
