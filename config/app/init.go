package app

func InitApp() IApp {
	externalServices := createExternalServices()
	internalServices := createInternalService(externalServices)

	return struct {
		iAppExternalServices
		iAppInternalServices
	}{
		externalServices,
		internalServices,
	}
}
