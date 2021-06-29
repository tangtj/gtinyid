package service

import (
	"github.com/tangtj/gtinyid/server/dao"
	"github.com/tangtj/gtinyid/server/model"
	"log"
	"sync"
	"time"
)

type IdTokenService struct {
	tokenMap map[string]*model.IdToken
	locker   sync.Locker
	once     *sync.Once
}

func NewIdTokenService() *IdTokenService {
	service := IdTokenService{tokenMap: map[string]*model.IdToken{}, locker: &sync.Mutex{}, once: &sync.Once{}}
	service._init()
	return &service
}

func (s *IdTokenService) _init() {

	getTokenMap := func() (map[string]*model.IdToken, error) {
		ids, err := dao.GetAllIdToken()
		if err != nil {
			log.Println("查询所有异常")
			return nil, err
		}
		sm := make(map[string]*model.IdToken)
		for _, id := range ids {
			sm[id.BizType] = id
		}
		return sm, nil
	}

	fresh := func() {
		defer s.locker.Unlock()
		s.locker.Lock()
		tokens, err := getTokenMap()
		if err == nil {
			s.tokenMap = tokens
		}
	}

	go func() {

		timer := time.NewTicker(time.Second * 5)

		for {
			select {
			case <-timer.C:
				dao.GetAllIdToken()
				fresh()
				timer.Reset(time.Second * 5)
			}
		}
	}()
	s.once.Do(fresh)
}

func (s *IdTokenService) CanGenerate(bizType string, token string) bool {
	t, ok := s.tokenMap[bizType]
	if !ok {
		return false
	}
	if t.Token != token {
		return false
	}
	return true
}
