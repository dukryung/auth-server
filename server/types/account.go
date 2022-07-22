package types

type Account struct {
	Email        string `json:"email"`
	Mnemonic     string `json:"mnemonic"`
	NickName     string `json:"nick_name"`
	DeviceID  string `json:"device_token"`
	ProfileImage []byte `json:"profile_image"`
	PuzzleToken  string `json:"puzzle_token"`
}
