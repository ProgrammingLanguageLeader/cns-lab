package main

type TrafficParams struct {
	// Номер сетевой службы получателя (http, ssh, rdp, https)
	ServiceNumber uint

	// Длительность сессии в секундах
	SessionDuration uint

	// Количество пакетов в сессии
	PackageCount uint

	// Количество запросов в сессии
	RequestCount uint

	// Количество ответов в сессии
	ResponseCount uint

	// Количество запросов с подозрительными последовательностями символов
	SuspiciousRequestCount uint

	// Количество попыток авторизации в минуту
	AuthAttemptCount uint

	// Количество ошибочных пакетов в сессии
	ErrorPackagesCount uint

	// Состояние соединения (active, established, closed)
	ConnectionState uint

	// Идентичность портов отправителя и получателя (идентичные/различные)
	PortsIdentity uint
}

func GenerateTrafficParams(variant uint, studentsNumber uint) *TrafficParams {
	trafficParams := TrafficParams{}
	serviceNumber := variant % 5
	if serviceNumber == 0 {
		serviceNumber = 5
	}
	trafficParams.ServiceNumber = serviceNumber
	trafficParams.ConnectionState = serviceNumber

	trafficParams.SessionDuration = 5000 / studentsNumber * variant
	trafficParams.PackageCount = 3750 / studentsNumber * variant
	trafficParams.RequestCount = 200 / studentsNumber * variant
	trafficParams.ResponseCount = 200 / studentsNumber * variant
	trafficParams.SuspiciousRequestCount = 50 / studentsNumber * variant
	trafficParams.AuthAttemptCount = 5000 / studentsNumber * variant
	trafficParams.ErrorPackagesCount = 150 / studentsNumber * variant

	trafficParams.PortsIdentity = (variant + 1) % 2 + 1
	return &trafficParams
}

func (trafficParams *TrafficParams) GenerateNext() *TrafficParams {
	trafficParams.ServiceNumber = (trafficParams.ServiceNumber + 1) % 4 + 1
	trafficParams.SessionDuration += 10
	trafficParams.PackageCount += 7
	trafficParams.RequestCount += 3
	trafficParams.ResponseCount += 2
	trafficParams.SuspiciousRequestCount += 5
	trafficParams.AuthAttemptCount += 13
	trafficParams.ErrorPackagesCount += 4
	trafficParams.ConnectionState = (trafficParams.ConnectionState + 1) % 3 + 1
	trafficParams.PortsIdentity = (trafficParams.PortsIdentity + 1) % 2 + 1
	return trafficParams
}

func (trafficParams TrafficParams) ToArray() []uint {
	return []uint{
		trafficParams.ServiceNumber,
		trafficParams.SessionDuration,
		trafficParams.PackageCount,
		trafficParams.RequestCount,
		trafficParams.ResponseCount,
		trafficParams.SuspiciousRequestCount,
		trafficParams.AuthAttemptCount,
		trafficParams.ErrorPackagesCount,
		trafficParams.ConnectionState,
		trafficParams.PortsIdentity,
	}
}
