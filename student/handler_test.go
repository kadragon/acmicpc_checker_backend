package student

// var userJSON = `{"email":"kadragon@sasa.hs.kr", "acmicpc_id":"kadragon", "rname":"강동욱"}`

// func TestCreateStudent(t *testing.T) {
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	if assert.NoError(t, CreateHandler(c)) {
// 		assert.Equal(t, http.StatusCreated, rec.Code)

// 		rq := model.Student{}
// 		json.Unmarshal([]byte(userJSON), &rq)
// 		rs := model.Student{}
// 		json.Unmarshal(rec.Body.Bytes(), &rs)

// 		assert.Equal(t, rq.Email, rs.Email)
// 		assert.Equal(t, rq.AcmicpcID, rs.AcmicpcID)
// 		assert.Equal(t, rq.Rname, rs.Rname)
// 	}
// }

// func TestCreateUser(t *testing.T) {
// 	// Setup
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	if assert.NoError(t, createUser(c)) {
// 		assert.Equal(t, http.StatusCreated, rec.Code)
// 		assert.Equal(t, userJSON, rec.Body.String())
// 	}
// }
