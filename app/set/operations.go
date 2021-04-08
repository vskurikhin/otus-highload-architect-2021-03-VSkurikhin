package set

import "github.com/google/uuid"

// IntersectString находит пересечение массивов.
func IntersectString(args ...[]string) []string {
	// создать временную карту для хранения содержимого массивов
	arrLength := len(args)
	repetitionRate := make(map[string]int)
	for _, arg := range args {
		tempArr := DistinctString(arg)
		for idx := range tempArr {
			// сколько раз мы встречались с этим элементом?
			if _, ok := repetitionRate[tempArr[idx]]; ok {
				repetitionRate[tempArr[idx]]++
			} else {
				repetitionRate[tempArr[idx]] = 1
			}
		}
	}

	// найти ключи, равные длине входных аргументов
	result := make([]string, 0)
	for key, val := range repetitionRate {
		if val == arrLength {
			result = append(result, key)
		}
	}

	return result
}

// DifferenceString находит разницу двух массивов.
func DifferenceString(args ...[]string) []string {
	// создать временную карту для хранения содержимого массивов
	repetitionRate := make(map[string]int)
	for _, arg := range args {
		tempArr := DistinctString(arg)
		for idx := range tempArr {
			// сколько раз мы встречались с этим элементом?
			if _, ok := repetitionRate[tempArr[idx]]; ok {
				repetitionRate[tempArr[idx]]++
			} else {
				repetitionRate[tempArr[idx]] = 1
			}
		}
	}

	// записать окончательное значение diffMap в массив и вернуть
	result := make([]string, 0)
	for key, val := range repetitionRate {
		if val == 1 {
			result = append(result, key)
		}
	}

	return result
}

// DistinctString удаляет повторяющиеся значения из одного массива.
func DistinctString(arg []string) []string {
	tempMap := make(map[string]uint8)

	for idx := range arg {
		tempMap[arg[idx]] = 0
	}

	result := make([]string, 0)
	for key := range tempMap {
		result = append(result, key)
	}
	return result
}

// UnionString находит объединение двух массивов.
func UnionString(args ...[]string) []string {
	// создать временную карту для хранения содержимого массивов
	tempMap := make(map[string]uint8)

	// записать содержимое массивов как ключи к карте. Значения карты не имеют значения
	for _, arg := range args {
		for idx := range arg {
			tempMap[arg[idx]] = 0
		}
	}

	// ключи карты теперь являются уникальными экземплярами всего содержимого массива
	result := make([]string, 0)
	for key := range tempMap {
		result = append(result, key)
	}

	return result
}

// IntersectUUID находит пересечение массивов.
func IntersectUUID(args ...[]uuid.UUID) []uuid.UUID {
	// создать временную карту для хранения содержимого массивов
	arrLength := len(args)
	repetitionRate := make(map[uuid.UUID]int)
	for _, arg := range args {
		tempArr := DistinctUUID(arg)
		for idx := range tempArr {
			// сколько раз мы встречались с этим элементом?
			if _, ok := repetitionRate[tempArr[idx]]; ok {
				repetitionRate[tempArr[idx]]++
			} else {
				repetitionRate[tempArr[idx]] = 1
			}
		}
	}

	// найти ключи, равные длине входных аргументов
	result := make([]uuid.UUID, 0)
	for key, val := range repetitionRate {
		if val == arrLength {
			result = append(result, key)
		}
	}

	return result
}

// DifferenceUUID находит разницу двух массивов.
func DifferenceUUID(args ...[]uuid.UUID) []uuid.UUID {
	// создать временную карту для хранения содержимого массивов
	repetitionRate := make(map[uuid.UUID]int)
	for _, arg := range args {
		tempArr := DistinctUUID(arg)
		for idx := range tempArr {
			// сколько раз мы встречались с этим элементом?
			if _, ok := repetitionRate[tempArr[idx]]; ok {
				repetitionRate[tempArr[idx]]++
			} else {
				repetitionRate[tempArr[idx]] = 1
			}
		}
	}

	// записать окончательное значение diffMap в массив и вернуть
	result := make([]uuid.UUID, 0)
	for key, val := range repetitionRate {
		if val == 1 {
			result = append(result, key)
		}
	}

	return result
}

// DistinctUUID удаляет повторяющиеся значения из одного массива.
func DistinctUUID(arg []uuid.UUID) []uuid.UUID {
	tempMap := make(map[uuid.UUID]uint8)

	for idx := range arg {
		tempMap[arg[idx]] = 0
	}

	result := make([]uuid.UUID, 0)
	for key := range tempMap {
		result = append(result, key)
	}
	return result
}

// UnionUUID находит объединение двух массивов.
func UnionUUID(args ...[]uuid.UUID) []uuid.UUID {
	// создать временную карту для хранения содержимого массивов
	tempMap := make(map[uuid.UUID]uint8)

	// записать содержимое массивов как ключи к карте. Значения карты не имеют значения
	for _, arg := range args {
		for idx := range arg {
			tempMap[arg[idx]] = 0
		}
	}

	// ключи карты теперь являются уникальными экземплярами всего содержимого массива
	result := make([]uuid.UUID, 0)
	for key := range tempMap {
		result = append(result, key)
	}

	return result
}
