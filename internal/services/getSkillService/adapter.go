package getSkillService

type GetSkillService interface {
	Execute() ([]Response, error)
}

type getSkillService struct{}

func NewGetSkillService() GetSkillService {
	return &getSkillService{}
}

func (s *getSkillService) Execute() ([]Response, error) {
	response := []Response{
		{
			Name:    "Golang",
			Picture: "Greenmoons",
			Score:   10,
		},
	}
	return response, nil

}
