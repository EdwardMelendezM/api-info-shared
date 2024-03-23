create_clock_mocks:
	mockery --dir=clock --name=Clock --filename=clock_mock.go --output=clock/mocks --outpkg=clock