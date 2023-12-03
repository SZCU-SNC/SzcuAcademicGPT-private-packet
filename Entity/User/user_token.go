package User

import "encoding/json"

type Token struct {
	UserID int
	IP     string
	Role   []string
}

func (ut *Token) ToByte() ([]byte, error) {
	jsonBytes, err := json.Marshal(ut)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func (ut *Token) FromByte(jsonStr []byte) error {
	err := json.Unmarshal(jsonStr, ut)
	return err
}
