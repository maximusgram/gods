package doublylinkedlist

import "encoding/json"

// toJSON outputs the JSON representation of list's elements
func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list.Values())
}

func (list *List) FromJSON(data []byte) error {
	elements := []interface{}{}
	err := json.Unmarshal(data, &elements)

	if err == nil {
		list.Clear()
		list.Add(elements...)
	}

	return err
} 