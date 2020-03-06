/*
 * go-housework
 *
 * 家事タスクを管理するAPIサーバです。
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package schemamodel

// ResponseRegisterFamilyMember - 世帯詳細情報取得APIの正常系レスポンススキーマ
type ResponseRegisterFamilyMember struct {

	Family Family `json:"family"`

	Member Member `json:"member"`
}