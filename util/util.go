package util

type ListItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func MapConverter(data []*ListItem) map[string]string {
	dataMap := make(map[string]string)
	for _, item := range data {
		dataMap[item.Key] = item.Value
	}
	return dataMap
}

func ListConverter(data map[string]string) []*ListItem {
	l := make([]*ListItem, 0)
	for k, v := range data {
		l = append(l, &ListItem{Key: k, Value: v})
	}
	return l
}

func ListConverterFromByte(data map[string][]byte) []*ListItem {
	l := make([]*ListItem, 0)
	for k, v := range data {
		l = append(l, &ListItem{Key: k, Value: string(v)})
	}
	return l
}
