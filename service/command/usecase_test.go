package command

// func Test_commandUsecaseImpl_New(t *testing.T) {
// 	b := board.NewBoardUsecaseImpl()
// 	r := robot.NewRobotUsecaseImpl(b)
// 	c := commandUsecaseImpl{r}

// 	tests := []struct {
// 		Command     string
// 		ExpectedErr error
// 	}{
// 		{"PLACE 0,1,NORTH", nil},
// 		{"PLACE 0,1,NORTH MOVE LEFT LEFT MOVE REPORT", nil},
// 		{"PLACES 0,1,NORTH", model.ErrInvalidCommand},
// 		{"PLACE MOVE", model.ErrInvalidCommand},
// 		{"1,2,NORTH PLACE", model.ErrInvalidCommand},
// 		{"PLACE 1,2,NORTHS", model.ErrInvalidCommand},
// 		{"REPORT PLACE 1,2,NORTH", model.ErrInvalidCommand},
// 		{"", model.ErrInvalidCommand},
// 		{"PLACE", model.ErrInvalidCommand},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.Command, func(t *testing.T) {
// 			_, err := c.New(tt.Command)

// 			assert.Equal(t, tt.ExpectedErr, err)
// 		})
// 	}
// }
