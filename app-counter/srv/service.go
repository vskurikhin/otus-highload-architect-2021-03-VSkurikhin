package srv

import (
	"database/sql"
	"github.com/Shopify/sarama"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/domain"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/domain/dto"
	"math/rand"
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

	rand.Seed(time.Now().Unix())
	for i := 0; i < MAX_COUNT; i++ {
		user, err := s.dao.User.ReadUserById(dm.To_user)
		if err != nil {
			if err != sql.ErrNoRows {
				logger.Errorf("upsertCounter: 1: %s", err)
				return
			}
		} else {
			if dm.Operation == "c" {
				err := s.dao.Counter.Upsert(user.Username)
				if err != nil {
					logger.Errorf("upsertCounter: 2: %s", err)
				} else {
					return
				}
			} else if dm.Operation == "u" {
				err := s.dao.Counter.Read(user.Username)
				if err != nil {
					logger.Errorf("upsertCounter: 3: %s", err)
				} else {
					return
				}
			}
		}
		time.Sleep(time.Duration(5*rand.Int31n(3)) * time.Second)
	}
}
