package controllers

import "finder/service"

var findSrv *service.Service

func Init(s *service.Service) {
	findSrv = s
}
