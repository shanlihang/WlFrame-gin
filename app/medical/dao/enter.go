package dao

import (
	"WlFrame-gin/app/medical/model"
	"WlFrame-gin/utils/global"

	"gorm.io/gorm"
)

// 检测结果
func InsertResult(res *model.Result) *gorm.DB {
	result := global.DB.Create(res)
	return result
}
func SelectResultsList(examNo string, deviceID string, name string, idnumber string) ([]model.Result, error) {
	var results []model.Result
	tx := global.DB.Model(model.Result{})
	if examNo != "" {
		tx.Where("examNo LIKE ?", "%"+examNo+"%")
	}
	if deviceID != "" {
		tx.Where("deviceID = ?", deviceID)
	}
	if name != "" {
		tx.Where("sfz_name LIKE ?", "%"+name+"%")
	}
	if idnumber != "" {
		tx.Where("sfz_idnumber LIKE ?", "%"+idnumber+"%")
	}
	err := tx.Find(&results).Error
	return results, err
}
func SelectResultById(ID int64) (model.Result, *gorm.DB) {
	var res model.Result
	result := global.DB.Model(model.Result{}).Where("id = ?", ID).First(&res)
	return res, result
}
func DeleteResult(ID int64) *gorm.DB {
	result := global.DB.Delete(&model.Result{}, ID)
	return result
}

// 社区
func InsertCommunity(comm *model.Community) *gorm.DB {
	result := global.DB.Create(comm)
	return result
}
func SelectCommunityList(name string, district string, address string) ([]model.Community, error) {
	var comms []model.Community
	tx := global.DB.Model(model.Community{})
	if name != "" {
		tx.Where("POI_name LIKE ?", "%"+name+"%")
	}
	if district != "" {
		tx.Where("POI_district LIKE ?", "%"+district+"%")
	}
	if address != "" {
		tx.Where("POI_address LIKE ?", "%"+address+"%")
	}
	err := tx.Find(&comms).Error
	return comms, err
}
func DeleteCommunity(ID int64) *gorm.DB {
	result := global.DB.Delete(&model.Community{}, ID)
	return result
}

// 物品
func InsertGoods(good *model.Goods) *gorm.DB {
	result := global.DB.Create(good)
	return result
}
func SelectGoodsList(name string, remark string) ([]model.Goods, error) {
	var goods []model.Goods
	tx := global.DB.Model(model.Goods{})
	if name != "" {
		tx.Where("name Like ?", "%"+name+"%")
	}
	if remark != "" {
		tx.Where("remark Like ?", "%"+remark+"%")
	}
	err := tx.Find(&goods).Error
	return goods, err
}
func SelectGoodsById(id int64) (model.Goods, *gorm.DB) {
	var good model.Goods
	result := global.DB.Where("ID = ?", id).First(&good)
	return good, result
}
func UpdateGood(good *model.Goods) *gorm.DB {
	result := global.DB.Model(good).Updates(good)
	return result
}
func DeleteGoods(ID int64) *gorm.DB {
	result := global.DB.Delete(&model.Goods{}, ID)
	return result
}

// 推送
func InsertMsg(msg *model.PushMsg) *gorm.DB {
	result := global.DB.Create(msg)
	return result
}
func UpdateMsg(msg *model.PushMsg) *gorm.DB {
	result := global.DB.Updates(msg)
	return result
}
func SelectMsgList(title string, content string) ([]model.PushMsg, error) {
	var msg []model.PushMsg
	tx := global.DB.Model(model.PushMsg{})
	if title != "" {
		tx.Where("title LIKE ?", "%"+title+"%")
	}
	if content != "" {
		tx.Where("content LIKE ?", "%"+content+"%")
	}
	err := tx.Find(&msg).Error
	return msg, err
}
func SelectMsgById(ID int64) (model.PushMsg, *gorm.DB) {
	var msg model.PushMsg
	result := global.DB.Model(model.PushMsg{}).Where("id = ?", ID).First(&msg)
	return msg, result
}
func DeleteMsg(ID int64) *gorm.DB {
	result := global.DB.Delete(&model.PushMsg{}, ID)
	return result
}

// 居民
func InsertPeople(people *model.People) *gorm.DB {
	result := global.DB.Create(people)
	return result
}
func SelectPeoplesList(name string, email string, phone string, idnumber string) ([]model.People, error) {
	var peoples []model.People
	tx := global.DB.Model(model.People{})
	if name != "" {
		tx.Where("name Like ?", "%"+name+"%")
	}
	if email != "" {
		tx.Where("email Like ?", "%"+email+"%")
	}
	if phone != "" {
		tx.Where("phone Like ?", "%"+phone+"%")
	}
	if idnumber != "" {
		tx.Where("idnumber Like ?", "%"+idnumber+"%")
	}
	err := tx.Find(&peoples).Error
	return peoples, err
}
func DeletePeople(ID int64) *gorm.DB {
	result := global.DB.Delete(&model.People{}, ID)
	return result
}
func UpdatePeople(people model.People) *gorm.DB {
	result := global.DB.Updates(people)
	return result
}

// 反馈
func InsertFeedback(feedback *model.Feedback) *gorm.DB {
	result := global.DB.Create(feedback)
	return result
}
func SelectFeedbacksList(content string, status string) ([]model.Feedback, error) {
	var feedbacks []model.Feedback
	tx := global.DB.Model(model.Feedback{})
	if content != "" {
		tx.Where("content LIKE ?", "%"+content+"%")
	}
	if status != "" {
		tx.Where("status = ?", status)
	}
	err := tx.Preload("People").Find(&feedbacks).Error
	return feedbacks, err
}
func DeleteFeedback(ID int64) *gorm.DB {
	result := global.DB.Delete(&model.Feedback{}, ID)
	return result
}
func UpdateFeedback(ID int64) *gorm.DB {
	result := global.DB.Model(&model.Feedback{}).Where("ID = ?", ID).Update("status", 1)
	return result
}
