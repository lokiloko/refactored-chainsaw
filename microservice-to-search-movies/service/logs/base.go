package logs

type LogsService interface {
	WriteLog(request interface{}, response interface{}) error
}

type logsService struct {
}

func NewService() LogsService {
	return &logsService{}
}
