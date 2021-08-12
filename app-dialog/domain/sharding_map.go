package domain

type ShardingMap map[string]uint8

var SHARDING_MAP = make(map[string]uint8)

type shardingMapTable struct {
	Id      int
	City    string
	ShardId int
}

const SELECT_ID_CITY_FROM_SHARDING_MAP = "SELECT id, city, shard_id FROM sharding_map"

func GetShardId(username string) uint8 {
	id := SHARDING_MAP[username]
	return id
}

func (s *shardingMap) ReadMap() ([]shardingMapTable, error) {

	stmtOut, err := s.dbRw.Prepare(SELECT_ID_CITY_FROM_SHARDING_MAP)
	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	defer func() { _ = stmtOut.Close() }()

	rows, err := stmtOut.Query()

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var maps []shardingMapTable
	for rows.Next() {

		var r shardingMapTable
		err = rows.Scan(&r.Id, &r.City, &r.ShardId)
		SHARDING_MAP[r.City] = uint8(r.ShardId)

		if err != nil {
			return nil, err
		}
		maps = append(maps, r)
	}
	return maps, nil
}
