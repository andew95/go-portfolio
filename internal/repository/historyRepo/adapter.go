package historyRepo

type HistoryAdapterRepository interface {
}

type historyAdapterRepository struct {
}

func NewHistoryRepository() HistoryAdapterRepository {
	return new(historyAdapterRepository)
}

func (repo *historyAdapterRepository) Exercute() {

}
