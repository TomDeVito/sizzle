package service

type Service interface {
	SmokeService() SmokeServicer
}

type SmokeServicer interface {
}
