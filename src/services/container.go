package services

type Container struct {
	MockDataService *MockDataService
}

func CreateContainer() *Container {

	mockDataService := &MockDataService{}

	return &Container{
		MockDataService: mockDataService,
	}
}
