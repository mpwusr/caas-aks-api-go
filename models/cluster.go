package models

type CreateClusterRequest struct {
	SubscriptionID string `json:"subscription_id"`
	ResourceGroup  string `json:"resource_group"`
	ClusterName    string `json:"cluster_name"`
	Location       string `json:"location"`
	NodeCount      int32  `json:"node_count"`
}
