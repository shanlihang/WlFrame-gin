package model

import "gorm.io/gorm"

// Request
type Result struct {
	gorm.Model
	Blood    `json:"blood"`
	DeviceID string `json:"deviceID" gorm:"column:deviceID"`
	Dgc      string `json:"dgc" gorm:"column:dgc"`
	Ecg      `json:"ecg"`
	ExamNo   string `json:"examNo" gorm:"column:examNo"`
	Fat      `json:"fat"`
	Hw       `json:"hw"`
	NS       string `json:"ns" gorm:"column:ns"`
	Semang   string `json:"semang" gorm:"column:semang"`
	Sfz      `json:"sfz"`
	Shili    `json:"shili"`
	Spo2     `json:"spo2"`
	Tiwen    string `json:"tiwen" gorm:"column:tiwen"`
	Xlcp     `json:"xlcp"`
	Xt       `json:"xt"`
	Zybs     string `json:"zybs" gorm:"column:zybs"`
}

type Blood struct {
	High  string `json:"high" gorm:"column:blood_high"`
	Low   string `json:"low" gorm:"column:blood_low"`
	Rate  string `json:"rate" gorm:"column:blood_rate"`
	Rhigh string `json:"rhigh" gorm:"column:blood_rhigh"`
	Rlow  string `json:"rlow" gorm:"column:blood_rlow"`
}

type Ecg struct {
	Data   string `json:"data" gorm:"column:ecg_data"`
	Len    int64  `json:"len" gorm:"column:ecg_len"`
	Result string `json:"result" gorm:"column:ecg_result"`
}

type Fat struct {
	Dbz    string `json:"dbz" gorm:"column:fat_dbz"`
	Dbzlv  string `json:"dbzlv" gorm:"column:fat_dbzlv"`
	Gl     string `json:"gl" gorm:"column:fat_gl"`
	Gy     string `json:"gy" gorm:"column:fat_gy"`
	Jcdx   string `json:"jcdx" gorm:"column:fat_jcdx"`
	Jrl    string `json:"jrl" gorm:"column:fat_jrl"`
	Jrlv   string `json:"jrlv" gorm:"column:fat_jrlv"`
	Nzzf   string `json:"nzzf" gorm:"column:fat_nzzf"`
	Qztz   string `json:"qztz" gorm:"column:fat_qztz"`
	Tsfl   string `json:"tsfl" gorm:"column:fat_tsfl"`
	Tsflv  string `json:"tsflv" gorm:"column:fat_tsflv"`
	Xbnyl  string `json:"xbnyl" gorm:"column:fat_xbnyl"`
	Xbnylv string `json:"xbnylv" gorm:"column:fat_xbnylv"`
	Xbwyl  string `json:"xbwyl" gorm:"column:fat_xbwyl"`
	Xbwylv string `json:"xbwylv" gorm:"column:fat_xbwylv"`
	Zfl    string `json:"zfl" gorm:"column:fat_zfl"`
	Zflv   string `json:"zflv" gorm:"column:fat_zflv"`
}

type Hw struct {
	Bmi    string `json:"bmi" gorm:"column:hw_bmi"`
	Height string `json:"height" gorm:"column:hw_height"`
	Weight string `json:"weight" gorm:"column:hw_weight"`
}

type Sfz struct {
	Age      string `json:"age" gorm:"column:sfz_age"`
	Birthday string `json:"birthday" gorm:"column:sfz_birthday"`
	Idnumber string `json:"idnumber" gorm:"column:sfz_idnumber"`
	Name     string `json:"name" gorm:"column:sfz_name"`
	Nation   string `json:"nation" gorm:"column:sfz_nation"`
	Sex      string `json:"sex" gorm:"column:sfz_sex"`
}

type Shili struct {
	LeftEye  string `json:"left_eye" gorm:"column:shili_left_eye"`
	RightEye string `json:"right_eye" gorm:"column:shili_right_eye"`
}

type Spo2 struct {
	SP string `json:"sp" gorm:"column:spo2_sp"`
}

type Xlcp struct {
	Eq     string `json:"eq" gorm:"column:xlcp_eq"`
	Hfxx   string `json:"hfxx" gorm:"column:xlcp_hfxx"`
	Hmdjl  string `json:"hmdjl" gorm:"column:xlcp_hmdjl"`
	Lnyy   string `json:"lnyy" gorm:"column:xlcp_lnyy"`
	Pstr   string `json:"pstr" gorm:"column:xlcp_pstr"`
	Qxjkd  string `json:"qxjkd" gorm:"column:xlcp_qxjkd"`
	Rgza   string `json:"rgza" gorm:"column:xlcp_rgza"`
	Shmyd  string `json:"shmyd" gorm:"column:xlcp_shmyd"`
	Smzkpg string `json:"smzkpg" gorm:"column:xlcp_smzkpg"`
	Ucla   string `json:"ucla" gorm:"column:xlcp_ucla"`
	Zcjkpd string `json:"zcjkpd" gorm:"column:xlcp_zcjkpd"`
	Zpyy   string `json:"zpyy" gorm:"column:xlcp_zpyy"`
}

type Xt struct {
	Type  string `json:"type" gorm:"column:xt_type"`
	Value string `json:"value" gorm:"column:xt_value"`
}

func (r Result) TableName() string {
	return "result"
}
