package entity

type DistanceMatrixResponse struct {
	Rows []struct {
		Elements []struct {
			Distance struct {
				Text string `json:"text"`
			} `json:"distance"`
			Duration struct {
				Text string `json:"text"`
			} `json:"duration"`
		} `json:"elements"`
	} `json:"rows"`
	Status string `json:"status"`
}
