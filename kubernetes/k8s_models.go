package kubernetes

type KubernetesSecret struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Data map[string]string `json:"data"`
}

// type KubernetesDeployment struct {
// 	ApiVersion string `json:"apiVersion"`
// 	Kind       string `json:"kind"`
// 	Metadata   struct {
// 		Name string `json:"name"`
// 	} `json:"metadata"`
// 	Spec struct {
// 		Replicas int `json:"replicas"`
// 		Selector struct {
// 			MatchLabels map[string]string `json:"matchLabels"`
// 		} `json:"selector"`
// 		Template struct {
// 			Metadata struct {
// 				Labels map[string]string `json:"labels"`
// 			} `json:"metadata"`
// 			Spec struct {
// 				Containers []struct {
// 					Name  string `json:"name"`
