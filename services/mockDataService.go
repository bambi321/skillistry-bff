package services

import "github.com/google/uuid"

type MockDataService struct{}

type MockSkills struct{}

type MockTasks struct{}

func (mockDataService *MockDataService) FetchSkills(userId uuid.UUID) *MockSkills {
	return &MockSkills{}
}

func (MockDataService *MockDataService) FetchTasks(userId uuid.UUID, skillId uuid.UUID) *MockTasks {
	return &MockTasks{}
}
