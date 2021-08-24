package srv

import (
	"database/sql"
	"github.com/Shopify/sarama"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/domain/dto"
	"time"
)

const MAX_COUNT = 99

type Service struct {
	dao *domain.DAO
}

func NewService(dao *domain.DAO) *Service {
	return &Service{dao: dao}
}

func (s *Service) Do(msg *sarama.ConsumerMessage) error {

	if logger.DebugEnabled() {
		logger.Debugf(
			"Received message Topic(%s) | Key(%s) | Message(%s) \n", msg.Topic, string(msg.Key), string(msg.Value),
		)
	}
	switch key := string(msg.Key); key {
	case "hl.user":
		return s.doKafkaUser(msg.Value)
	case "hl.dialog_message":
		return s.doKafkaDialogMessage(msg.Value)
	default:
		logger.Debugf("default")
	}
	return nil
}

func (s *Service) doKafkaUser(value []byte) error {

	ku, err := dto.UnmarshalKafkaUser(value)
	if err != nil {
		return err
	}
	user, err := s.dao.User.Create(ku.User())
	if err != nil {
		return err
	}
	if logger.DebugEnabled() {
		logger.Debugf("hl.user: %s", user)
	}
	return nil
}

func (s *Service) doKafkaDialogMessage(value []byte) error {

	dm, err := dto.UnmarshalKafkaDialogMessage(value)
	if err != nil {
		return err
	}
	go s.upsertCounter(dm)
	return nil
}

func (s *Service) upsertCounter(dm *dto.KafkaDialogMessage) {

	for i := 0; i < MAX_COUNT; i++ {
		user, err := s.dao.User.ReadUserById(dm.To_user)
		if err != nil {
			if err == sql.ErrNoRows {
				time.Sleep(8 * time.Second)
			} else {
				logger.Errorf("upsertCounter: 1: %s", err)
				return
			}
		} else {
			counter, err := s.dao.Counter.ReadByUserId(dm.To_user)
			if err != nil {
				if err == sql.ErrNoRows {
					counter := &domain.Counter{Username: user.Username, Total: 1, Unread: 1}
					s.insertCounter(dm, counter)
					return
				} else {
					logger.Errorf("upsertCounter: 2: %s", err)
					return
				}
			} else {
				s.updateCounter(dm, counter)
				return
			}
		}
	}
}

func (s *Service) insertCounter(dm *dto.KafkaDialogMessage, counter *domain.Counter) {

	if dm.Operation == "c" && dm.Already_read == 0 {
		_, err := s.dao.Counter.Create(counter)
		if err != nil {
			logger.Errorf("insertCounter: 1: %s", err)
		}
	}
}

func (s *Service) updateCounter(dm *dto.KafkaDialogMessage, counter *domain.Counter) {

	if dm.Operation == "c" {
		counter.Total += 1
		if dm.Already_read == 0 {
			counter.Unread += 1
		}
		_, err := s.dao.Counter.Update(counter)
		if err != nil {
			logger.Errorf("updateCounter: 1: %s", err)
		}
	} else if dm.Operation == "u" {
		if dm.Already_read == 1 {
			counter.Unread -= 1
			_, err := s.dao.Counter.Update(counter)
			if err != nil {
				logger.Errorf("updateCounter: 2: %s", err)
			}
		}
	}
}
