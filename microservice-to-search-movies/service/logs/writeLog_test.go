package logs

func (s *LogsSuite) Test_GetByID() {
	s.Run("when success", func() {
		err := s.service.WriteLog("asd", "asd")

		s.Nil(err)
	})
}

