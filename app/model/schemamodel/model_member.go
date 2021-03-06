/*
 * go-housework
 *
 * 家事タスクを管理するAPIサーバです。
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package schemamodel

// Member - Memberスキーマ
type Member struct {

	// 家族メンバID
	MemberId int64 `json:"member_id"`

	// 家族メンバ名。ユーザ名と一致する。
	MemberName string `json:"member_name"`
}
