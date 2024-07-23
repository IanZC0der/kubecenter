package metrics

type MetricsItem struct {
	Name  string `json:"title"`
	Label string `json:"label"`
	Value string `json:"value"`
	// in RGB, easier for the frontend to present the data
	Color string `json:"color"`
}
