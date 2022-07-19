package types

type Account struct {
	Email        string `json:"email"`
	Mnemonic     string `json:"mnemonic"`
	NickName     string `json:"nick_name"`
	ProfileImage []byte `json:"profile_image"`
	PuzzleToken  string `json:"puzzle_token"`
}
