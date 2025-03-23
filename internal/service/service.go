package service

import (
	"fmt"
	"time"

	"github.com/aAmer0neee/Bot_Study-Helper/internal/cache"
)

type (
	Service struct {
		Cache *cache.Cache
	}
)

func InitService(cache *cache.Cache) *Service {
	return &Service{Cache: cache}
}

func (s *Service)StartService(userID int64){
	
	s.Cache.AddRecord(fmt.Sprintf("user:%d",userID),"new_user",0)
}

func (s *Service)CreateTaskService(userID int64){
	status, _ := s.Cache.GetRecord(fmt.Sprintf("user:%d",userID))
	if status == "wait"{

		s.Cache.AddRecord(fmt.Sprintf("user:%d",userID),"wait",time.Minute * 10)
	
	}
}

