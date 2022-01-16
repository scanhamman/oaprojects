package dbrepo

type Projects []Project

type Project struct {
	Projectid         uint64
	OAID              string
	Title             string
	Acronym           string
	Summary           string
	Keywords          string
	Websiteurl        string
	Callidentifier    string
	Code              string
	Startdate         string
	Enddate           string
	GrantedCurr       string
	GrantedAmount     string
	GrantedTotCost    string
	Openaccessdataset string
	Openaccesspubs    string
}

type Funding struct {
	Projectid    uint64
	Description  string
	StreamID     string
	Jurisdiction string
	Name         string
	ShortName    string
}

type H2020 struct {
	Projectid   uint64
	Code        string
	Description string
}

type Subject struct {
	Projectid uint64
	Subject   string
}
