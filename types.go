package iblclient

type Channel struct {
	HasSchedule   bool
	ID            string
	MasterBrandID string
	Title         string
	Type          string
}

type Supplier struct {
	ID    string
	Title string
}

type Synopses struct {
	Large  string
	Medium string
	Small  string
}

type MasterBrand struct {
	Attribution string
	ID          string
	IdentID     string
	Titles      struct {
		Large  string
		Medium string
		Small  string
	}
}

type Duration struct {
	Text  string
	Value string
}

type Availability struct {
	End       string
	Remaining string
	Start     string
}

type Version struct {
	Availability   Availability
	Download       bool
	Duration       Duration
	FirstBroadcast string
	Hd             bool
	ID             string
	Kind           string
	Type           string
}

type Labels struct {
	Editorial string
}

type Image struct {
	Standard string
}

/* This can be a programme or an episode */
type Item struct {
	AudioDescribed    bool
	Categories        []string
	Count             int
	Guidance          bool
	ID                string
	Image             Image `json:"images"`
	InitialChildren   []Item
	LexicalSortLetter string
	MasterBrand       MasterBrand
	RelatedLinks      []interface{}
	ReleaseDate       string
	ReleaseDateTime   string
	Signed            bool
	Labels            Labels
	Status            string
	Suppliers         []Supplier
	Synopses          Synopses
	Title             string
	Subtitle          string
	TleoID            string
	TleoType          string
	Type              string
	Versions          []Version
}

type ChannelProgrammes struct {
	Channel  Channel
	Count    int
	Elements []Item
	Page     int
	PerPage  int
}

type ChannelProgrammesResponse struct {
	ChannelProgrammes `json:"channel_programmes"`
}

type ChannelsResponse struct {
	Channels []Channel `json:"channels"`
}

type Search struct {
	Page    int
	PerPage int
	Count   int
	Query   string
	Items   []Item `json:"results"`
}

type SearchResponse struct {
	Search Search
}
