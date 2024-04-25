package app

type taskService struct {
	r *taskRepository
}

func NewService(r *taskRepository) *taskService {
	return &taskService{r: r}
}

func (s *taskService) CreateTask(req *TaskRequest) (int, error) {
	task := NewTask(req.Title, req.Description, req.Status)
	id, err := s.r.InsertTask(task)
	if err != nil {
		return 0, err
	}

	return id, err
}

func (s *taskService) GetTasks() ([]Task, error) {
	tasks, err := s.r.GetTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *taskService) UpdateTask(res *TaskRequest, id int) error {
	task := NewTask(res.Title, res.Description, res.Status)
	task.ID = id
	err := s.r.UpdateTask(task)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) DeleteTask(id int) error {
	err := s.r.DeleteTask(id)
	if err != nil {
		return err
	}

	return nil
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
