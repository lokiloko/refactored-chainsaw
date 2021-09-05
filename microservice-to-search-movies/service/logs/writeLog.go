package logs

import (
	"fmt"
	"time"
)

func (l logsService) WriteLog(request interface{}, response interface{}) error {
	fmt.Println(fmt.Sprintf("Request:%v \n Response:%v \n DateTime:%v", request, response, time.Now()))

	return nil
}
