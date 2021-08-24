package srv

import (
	"github.com/Shopify/sarama"
	"github.com/savsgio/go-logger/v2"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/domain"
)

type Service struct {
	// dao *domain.DAO
}

func NewService(dao *domain.DAO) *Service {
	return nil
}

func (s *Service) Do(msg *sarama.ConsumerMessage) error {

	if logger.DebugEnabled() {
		logger.Debugf(
			"Received message Topic(%s) | Key(%s) | Message(%s) \n", msg.Topic, string(msg.Key), string(msg.Value),
		)
	}
	// value, err := s.unmarshal(msg.Value)

	/*	if err != nil {
			return err
		}
		// Определяем значение бизнес-ключа из входящего сообщения(1)
		bk, err := domain.NewBusinessKey(value)

		if err != nil {
			return err
		}
		// Ищем в БД запись с таким значением бизнес-ключа (2)
		exists, err := s.dao.ABResultCumulative.IsExistByBK(bk)

		if err != nil {
			return err
		}
		if logger.DebugEnabled() {
			logger.Debugf("IsBusinessKeyExists(%s) -> %v", bk.JsonString(), exists)
		}
		// Запись в БД найдена?
		if !exists {
			// Пытаемся вставить новую запись (5)
			err = s.insert(bk, value)
			// Вставка прошла успешно? (6)
			if err == nil {
				// Записываем INFO в лог
				logger.Infof(" Record with BusinessKey(%s) INSERTED. ", bk.JsonString())
			}
			return err
		}
		// Записываем WARNING в лог
		logger.Warningf(" Record with BusinessKey(%s) already exists! ", bk.JsonString())
	*/
	return nil
}
