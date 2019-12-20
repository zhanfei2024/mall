package models

import (
	"errors"
	"strconv"
	"time"
)

// Objects Objects
var (
	Objects map[string]*Object
)

// Object 结构体
type Object struct {
	ObjectId   string
	Score      int64
	PlayerName string
}

func init() {
	Objects = make(map[string]*Object)
	Objects["hjkhsbnmn123"] = &Object{"hjkhsbnmn123", 100, "astaxie"}
	Objects["mjjkxsxsaa23"] = &Object{"mjjkxsxsaa23", 101, "someone"}
}

// AddOne 新增
func AddOne(object Object) (ObjectId string) {
	object.ObjectId = "astaxie" + strconv.FormatInt(time.Now().UnixNano(), 10)
	Objects[object.ObjectId] = &object
	return object.ObjectId
}

// GetOne 查询指定
func GetOne(ObjectId string) (object *Object, err error) {
	if v, ok := Objects[ObjectId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectID Not Exist")
}

// GetAll 查询全部
func GetAll() map[string]*Object {
	return Objects
}

// Update 更新
func Update(ObjectId string, Score int64) (err error) {
	if v, ok := Objects[ObjectId]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("ObjectID Not Exist")
}

// Delete 删除
func Delete(ObjectId string) {
	delete(Objects, ObjectId)
}
