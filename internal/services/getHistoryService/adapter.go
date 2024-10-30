package getHistoryService

type GetHistoryService interface {
	Execute() ([]Response, error)
}

type getHistoryService struct{}

func NewGetHistory() GetHistoryService {
	return &getHistoryService{}
}

func (s *getHistoryService) Execute() ([]Response, error) {
	response := []Response{
		{
			Position:    "Backend Developer",
			Company:     "Greenmoons",
			StartDate:   "Aug 2023",
			EndDate:     "Aug 2024",
			Description: nil,
		},
	}
	return response, nil

}
