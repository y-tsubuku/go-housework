/*
 * go-housework
 *
 * 家事タスクを管理するAPIサーバです。
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"errors"
)

// DefaultApiService is a service that implents the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API. 
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// CreateFamily - 世帯登録API
func (s *DefaultApiService) CreateFamily(requestCreateFamily RequestCreateFamily) (interface{}, error) {
	// TODO - update CreateFamily with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'CreateFamily' not implemented")
}

// CreateTask - 家事タスク登録API
func (s *DefaultApiService) CreateTask(requestCreateTask RequestCreateTask) (interface{}, error) {
	// TODO - update CreateTask with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'CreateTask' not implemented")
}

// CreateUser - ユーザ登録API
func (s *DefaultApiService) CreateUser(requestCreateUser RequestCreateUser) (interface{}, error) {
	// TODO - update CreateUser with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'CreateUser' not implemented")
}

// DeleteFamily - 世帯情報削除API
func (s *DefaultApiService) DeleteFamily() (interface{}, error) {
	// TODO - update DeleteFamily with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'DeleteFamily' not implemented")
}

// DeleteFamilyMember - 世帯メンバ削除API
func (s *DefaultApiService) DeleteFamilyMember(memberId int64) (interface{}, error) {
	// TODO - update DeleteFamilyMember with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'DeleteFamilyMember' not implemented")
}

// DeleteTask - 家事タスク削除API
func (s *DefaultApiService) DeleteTask(taskId int64) (interface{}, error) {
	// TODO - update DeleteTask with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'DeleteTask' not implemented")
}

// DeleteUser - ユーザ削除API
func (s *DefaultApiService) DeleteUser() (interface{}, error) {
	// TODO - update DeleteUser with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'DeleteUser' not implemented")
}

// GetHealth - 死活監視API
func (s *DefaultApiService) GetHealth() (interface{}, error) {
	// TODO - update GetHealth with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetHealth' not implemented")
}

// ListFamilyMembers - 世帯メンバ一覧取得API
func (s *DefaultApiService) ListFamilyMembers() (interface{}, error) {
	// TODO - update ListFamilyMembers with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ListFamilyMembers' not implemented")
}

// ListTasks - 家事タスク一覧取得API
func (s *DefaultApiService) ListTasks(date string) (interface{}, error) {
	// TODO - update ListTasks with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ListTasks' not implemented")
}

// Login - ログインAPI
func (s *DefaultApiService) Login(userName string, password string) (interface{}, error) {
	// TODO - update Login with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'Login' not implemented")
}

// RegisterFamilyMember - 世帯メンバ登録API
func (s *DefaultApiService) RegisterFamilyMember(requestRegisterFamilyMember RequestRegisterFamilyMember) (interface{}, error) {
	// TODO - update RegisterFamilyMember with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'RegisterFamilyMember' not implemented")
}

// ResetPassword - パスワードリセットAPI
func (s *DefaultApiService) ResetPassword(requestResetPassword RequestResetPassword) (interface{}, error) {
	// TODO - update ResetPassword with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ResetPassword' not implemented")
}

// ShowFamily - 世帯詳細情報取得API
func (s *DefaultApiService) ShowFamily() (interface{}, error) {
	// TODO - update ShowFamily with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'ShowFamily' not implemented")
}

// UpdateFamily - 世帯情報更新API
func (s *DefaultApiService) UpdateFamily(requestUpdateFamily RequestUpdateFamily) (interface{}, error) {
	// TODO - update UpdateFamily with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'UpdateFamily' not implemented")
}

// UpdateTask - 家事タスク更新API
func (s *DefaultApiService) UpdateTask(requestUpdateTask RequestUpdateTask) (interface{}, error) {
	// TODO - update UpdateTask with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'UpdateTask' not implemented")
}
