package getHistoryService

type Response struct {
	Position    string    `json:"position"`
	Company     string    `json:"company"`
	StartDate   string    `json:"startDate"`
	EndDate     string    `json:"endDate"`
	Description *[]string `json:"description"`
}
