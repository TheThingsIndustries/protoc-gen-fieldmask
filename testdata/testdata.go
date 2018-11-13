package testdata

import "errors"

type CustomType []byte

func (t CustomType) Marshal() ([]byte, error) {
	return t, nil
}
func (t *CustomType) MarshalTo(data []byte) (n int, err error) {
	data = append(data, *t...)
	return len(*t), nil
}
func (t *CustomType) Unmarshal(data []byte) error {
	*t = make([]byte, len(data))
	copy(*t, data)
	return nil
}
func (t *CustomType) Size() int {
	return 42
}

func (t CustomType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + string(t) + `"`), nil
}
func (t *CustomType) UnmarshalJSON(data []byte) error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("Invalid data")
	}
	*t = make([]byte, len(data)-2)
	copy(*t, data[1:len(data)-1])
	return nil
}
