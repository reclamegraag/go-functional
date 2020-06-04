package functions

type StringSlice []string
type IntSlice []int
type StringStringMap map[string]string
type StringIntMap map[string]int
type IntStringMap map[int]string

/*
	MAP FUNCTIONS
*/

/* Perform a function on every value of a StringSlice */
func (stringSlice StringSlice) Map(mapFunction func(string) string) StringSlice {
	newStringSlice := make([]string, len(stringSlice))
	for i, value := range stringSlice {
		newStringSlice[i] = mapFunction(value)
	}
	return newStringSlice
}

/* Perform a function on every value of an IntSlice */
func (intSlice IntSlice) Map(mapFunction func(int) int) IntSlice {
	newIntSlice := make(IntSlice, len(intSlice))
	for i, value := range intSlice {
		newIntSlice[i] = mapFunction(value)
	}
	return newIntSlice
}

/* Perform a function on every value of a Map with the key and value as a string */
func (stringStringMap StringStringMap) Map(mapFunction func(string) string) StringStringMap {
	newStringStringMap := make(StringStringMap, len(stringStringMap))
	for key, value := range stringStringMap {
		newStringStringMap[key] = mapFunction(value)
	}
	return newStringStringMap
}

/* Perform a function on every value of a Map with the key as a string and the value as an int */
func (stringIntMap StringIntMap) Map(mapFunction func(int) int) StringIntMap {
	newStringIntMap := make(StringIntMap, len(stringIntMap))
	for key, value := range stringIntMap {
		newStringIntMap[key] = mapFunction(value)
	}
	return newStringIntMap
}

/* Perform a function on every value of a Map with the key as an int and the value as a string */
func (intStringMap IntStringMap) Map(mapFunction func(string) string) IntStringMap {
	newIntStringMap := make(IntStringMap, len(intStringMap))
	for key, value := range intStringMap {
		newIntStringMap[key] = mapFunction(value)
	}
	return newIntStringMap
}

/*
	UNIQUE FUNCTIONS
*/

/* Create a new StringSlice with only unique strings */
func (stringSlice StringSlice) Unique() StringSlice {
	newStringSlice := make([]string, 0, len(stringSlice))
	temporaryMap := make(map[string]bool)
	for _, value := range stringSlice {
		if _, ok := temporaryMap[value]; !ok {
			temporaryMap[value] = true
			newStringSlice = append(newStringSlice, value)
		}
	}
	return newStringSlice
}

/* Create a new IntSlice with only unique ints */
func (intSlice IntSlice) Unique() IntSlice {
	newIntSlice := make([]int, 0, len(intSlice))
	temporaryMap := make(map[int]bool)
	for _, value := range intSlice {
		if _, ok := temporaryMap[value]; !ok {
			temporaryMap[value] = true
			newIntSlice = append(newIntSlice, value)
		}
	}
	return newIntSlice
}
