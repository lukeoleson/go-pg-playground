package examples

const ACTOR_NUM_ROWS int = 200

//func TestSelectAllIntoSelect(t *testing.T) {
//	db := dal.Connect()
//	defer db.Close()
//
//	testCases := []struct {
//		name string
//	}{
//		{
//			name: "SELECT * FROM Actor",
//		},
//	}
//	var actors []models.Actor
//	err := db.Model(&models.Actor{}).Select(&actors)
//	assert.NoError(t, err)
//	assert.Equal(t, len(actors), ACTOR_NUM_ROWS)
//}
//
//func TestSelectAllIntoModel(t *testing.T) {
//	db := dal.Connect()
//	defer db.Close()
//
//	var actors []models.Actor
//	err := db.Model(&actors).Select()
//	assert.NoError(t, err)
//	assert.Equal(t, len(actors), ACTOR_NUM_ROWS)
//}
//
//func TestSelectAllError(t *testing.T) {
//	db := dal.Connect()
//	defer db.Close()
//
//	var actors []models.Actor
//	err := db.Model(actors).Select()
//	assert.Error(t, err)
//}
