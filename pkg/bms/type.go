package bms

type BMS struct {
	Header Header
}

type Header struct {
	Player    int
	Genre     string
	Artist    string
	BPM       int
	Playlevel int
	Rank      int
	Total     int
}
