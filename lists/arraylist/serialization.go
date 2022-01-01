package arraylist

import "encoding/json"

// ToJSON outputs the JSON representation of list's elements
func (list *List) ToJSON() ([]byte, error) {
    return json.Marshal(list.elements[:list.size])
}

// FromJSON populates list's elements from the input JSON representation
func (list *List) FromJSON(data []byte) error {
    err := json.Unmarshal(data, &list.elements)
    if err == nil {
        list.size = len(list.elements)
    }
    return err
}