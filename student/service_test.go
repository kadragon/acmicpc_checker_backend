package student

import (
	"acmicpc_checker_v2_backend/db"
	"acmicpc_checker_v2_backend/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUnique(t *testing.T) {
	db := db.TestDbconnect()

	if assert.NoError(t, db.AutoMigrate(&model.Student{})) {
		// 중복되는 데이터가 삽입되는것이 방지되는지 확인
		assert.NoError(t, db.Create(&model.Student{Email: "kangdongouk@gmail.com", AcmicpcID: "kadragon", Rname: "강동욱"}).Error)
		assert.EqualError(t, createUnique(db, &model.Student{Email: "kangdongouk@gmail.com", AcmicpcID: "kadragon", Rname: "강동욱"}), "student duplicated")

		// 데이터 삽입이 정상적으로 되는지 확인
		assert.NoError(t, createUnique(db, &model.Student{Email: "kadragon@sasa.hs.kr", AcmicpcID: "sasa", Rname: "강동욱"}))

		var testStudent model.Student
		assert.NoError(t, db.Where("acmicpc_id = ?", "sasa").Find(&testStudent).Error)
		assert.Equal(t, "kadragon@sasa.hs.kr", testStudent.Email)
	}
}

func TestGet(t *testing.T) {
	db := db.TestDbconnect()

	if assert.NoError(t, db.AutoMigrate(&model.Student{})) {
		// 데이터 읽기 확인을 위한 기초 데이터 삽입
		assert.NoError(t, db.Create(&model.Student{Email: "kangdongouk@gmail.com", AcmicpcID: "kadragon", Rname: "강동욱"}).Error)
		assert.NoError(t, db.Create(&model.Student{Email: "kadragon@sasa.hs.kr", AcmicpcID: "sasa", Rname: "강동욱"}).Error)

		act, err := get(db, 1)
		assert.NoError(t, err)
		assert.Equal(t, "kangdongouk@gmail.com", act.Email)
		assert.Equal(t, "kadragon", act.AcmicpcID)
		assert.Equal(t, "강동욱", act.Rname)
	}
}

func TestUpdate(t *testing.T) {
	db := db.TestDbconnect()

	if assert.NoError(t, db.AutoMigrate(&model.Student{})) {
		data := model.Student{Email: "kangdongouk@gmail.com", AcmicpcID: "kadragon", Rname: "강동욱"}

		// 데이터 읽기 확인을 위한 기초 데이터 삽입
		assert.NoError(t, db.Create(&data).Error)
		assert.NoError(t, db.Where("id = ?", 1).Find(&data).Error)

		testAcmicpcID := "sasa"
		data.AcmicpcID = testAcmicpcID

		assert.NoError(t, update(db, &data))

		var test model.Student
		assert.NoError(t, db.Where("email = ? and rname = ?", data.Email, data.Rname).Find(&test).Error)
		assert.Equal(t, testAcmicpcID, test.AcmicpcID)
	}
}
