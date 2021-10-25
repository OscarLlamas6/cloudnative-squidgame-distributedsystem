package helpers

// SingleGame structure
type SingleGame struct {
	Gamenumber  string `json:"gamenumber"`
	Gamename    string `json:"gamename"`
	Players     int64  `json:"players"`
	Rungames    int64  `json:"rungames"`
	Concurrence int64  `json:"concurrence"`
	Timeout     int64  `json:"timeout"`
	Request     int64  `json:"request"`
}

//NewSingleGame returned fuction
func NewSingleGame(n string, g string, p int64, r int64, c int64, t int64, req int64) *SingleGame {
	return &SingleGame{
		Gamenumber:  n,
		Gamename:    g,
		Players:     p,
		Rungames:    r,
		Concurrence: c,
		Timeout:     t,
		Request:     req,
	}
}

func (s *SingleGame) GetGameNumber() string {
	return s.Gamenumber
}

func (s *SingleGame) GetGameName() string {
	return s.Gamename
}

func (s *SingleGame) GetPlayers() int64 {
	return s.Players
}

func (s *SingleGame) GetRungames() int64 {
	return s.Rungames
}

func (s *SingleGame) GetConcurrence() int64 {
	return s.Concurrence
}

func (s *SingleGame) GetTimeout() int64 {
	return s.Timeout
}

func (s *SingleGame) GetRequest() int64 {
	return s.Request
}
