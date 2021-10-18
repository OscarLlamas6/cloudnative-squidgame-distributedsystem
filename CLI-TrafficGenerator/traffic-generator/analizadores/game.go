package analizadores

// SquidGameSet structure
type SquidGameSet struct {
	Gamename    string `json:"gamename"`
	Players     int64  `json:"players"`
	Rungames    int64  `json:"rungames"`
	Concurrence int64  `json:"concurrence"`
	Timeout     int64  `json:"timeout"`
}

//NewSquidGameSet returned fuction
func NewSquidGameSet(g string, p int64, r int64, c int64, t int64) *SquidGameSet {
	return &SquidGameSet{
		Gamename:    g,
		Players:     p,
		Rungames:    r,
		Concurrence: c,
		Timeout:     t,
	}
}

func (s *SquidGameSet) GetGameName() string {
	return s.Gamename
}

func (s *SquidGameSet) GetPlayers() int64 {
	return s.Players
}

func (s *SquidGameSet) GetRungames() int64 {
	return s.Rungames
}

func (s *SquidGameSet) GetConcurrence() int64 {
	return s.Concurrence
}

func (s *SquidGameSet) GetTimeout() int64 {
	return s.Timeout
}
