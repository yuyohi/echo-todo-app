package app

type taskService struct {
	r *taskRepository
}

func NewService(r *taskRepository) *taskService {
	return &taskService{r: r}
}

func (s *taskService) CreateTask(res TaskRequest) (int, error) {
	task := NewTask(res.Title, res.Description, res.Status)
	id, err := s.r.InsertTask(task)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (s *taskService) GetTasks() ([]Task, error) {
	tasks, err := s.r.getTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TaskResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
}
