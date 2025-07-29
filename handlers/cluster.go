package handlers

import (
	"caas-aks-api-go/models"
	"caas-aks-api-go/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// @Summary Create AKS Cluster
// @Tags AKS
// @Accept json
// @Produce json
// @Param cluster body models.CreateClusterRequest true "AKS Cluster Config"
// @Success 200 {object} interface{}
// @Failure 500 {object} map[string]string
// @Router /clusters [post]
func CreateCluster(w http.ResponseWriter, r *http.Request) {
	var req models.CreateClusterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", 400)
		return
	}
	cluster, err := service.CreateCluster(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create AKS cluster: %v", err), 500)
		return
	}
	json.NewEncoder(w).Encode(cluster)
}

func GetCluster(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	group := r.URL.Query().Get("resourceGroup")
	sub := r.URL.Query().Get("subscriptionId")
	cluster, err := service.GetCluster(sub, group, name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get cluster: %v", err), 500)
		return
	}
	json.NewEncoder(w).Encode(cluster)
}

func ListClusters(w http.ResponseWriter, r *http.Request) {
	group := r.URL.Query().Get("resourceGroup")
	sub := r.URL.Query().Get("subscriptionId")
	clusters, err := service.ListClusters(sub, group)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to list clusters: %v", err), 500)
		return
	}
	json.NewEncoder(w).Encode(clusters)
}
