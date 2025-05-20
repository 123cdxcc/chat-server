// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// User is the golang structure for table user.
type User struct {
	Id       int64  `json:"id"       orm:"id"       description:"用户ID"` // 用户ID
	Username string `json:"username" orm:"username" description:"用户名"`  // 用户名
	Online   int    `json:"online"   orm:"online"   description:"是否在线"` // 是否在线
}
