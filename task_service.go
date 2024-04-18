package app

type taskService struct {
	r *taskRepository
}

func NewService(r *taskRepository) *taskService {
	return &taskService{r: r}
}

func (s *taskService) CreateTask() (string, error) {

}