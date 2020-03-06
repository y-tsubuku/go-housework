/*
 * go-housework
 *
 * 家事タスクを管理するAPIサーバです。
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package schemamodel

// ResponseDeleteUser - ユーザ削除APIの正常系レスポンススキーマ
type ResponseDeleteUser struct {

	User User `json:"user"`

	// HTTPレスポンスメッセージ
	Message string `json:"message"`
}