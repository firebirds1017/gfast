// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package blog_comment

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
)

// Entity is the golang structure for table blog_comment.
type Entity struct {
	CommentId       uint   `orm:"comment_id,primary" json:"comment_id"`      //
	CommentUserId   uint   `orm:"comment_user_id"    json:"comment_user_id"` // 评论用户的用户id
	CommentContent  string `orm:"comment_content"    json:"comment_content"` // 评论内容
	CommentPid      uint   `orm:"comment_pid"        json:"comment_pid"`     // 当前评论所回复的父评论的id
	CommentNum      uint   `orm:"comment_num"        json:"comment_num"`     // 当前评论的回复数（下一级子评论数）
	CommentStatus   uint   `orm:"comment_status"     json:"comment_status"`  // 此评论的状态，0隐藏，1发布
	CommentLogId    uint   `orm:"comment_log_id"     json:"comment_log_id"`  // 当前评论所属日志id
	CreateTime      uint   `orm:"create_time"        json:"create_time"`     // 评论创建时间
	CommentNickname string `orm:"comment_nickname" json:"comment_nickname"`  // 评论用户昵称
	ReplyName       string `orm:"reply_name" json:"reply_name"`              // 当前所回复的对象昵称
	ReplyId         uint   `orm:"reply_id" json:"reply_id"`                  // 当前回复对象的id
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}
