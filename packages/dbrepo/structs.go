package dbrepo

import (
	"strconv"

	"github.com/scanhamman/oaprojects/packages/jsonfile"
)

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

type PSubject struct {
	Projectid uint64
	Subject   string
}

func DeriveDBProject(pid uint64, t jsonfile.Project) Project {
	dp := Project{Projectid: pid}
	dp.Code = GetStringValue(t.Code)
	dp.OAID = GetStringValue(t.ID)
	dp.Title = GetStringValue(t.Title)
	dp.Acronym = GetStringValue(t.Acronym)
	dp.Summary = GetStringValue(t.Summary)
	dp.Keywords = GetStringValue(t.Keywords)
	dp.Websiteurl = GetStringValue(t.Websiteurl)
	dp.Callidentifier = GetStringValue(t.Callidentifier)
	dp.Code = GetStringValue(t.Code)
	dp.Startdate = GetStringValue(t.Startdate)
	dp.Enddate = GetStringValue(t.Enddate)
	dp.GrantedCurr = GetStringValue(t.Granted.Currency)
	dp.GrantedAmount = GetFloatValue(t.Granted.Fundedamount)
	dp.GrantedTotCost = GetFloatValue(t.Granted.Totalcost)
	dp.Openaccessdataset = GetBoolValue(t.Openaccessmandatefordataset)
	dp.Openaccesspubs = GetBoolValue(t.Openaccessmandateforpublications)
	return dp
}

func DeriveDBFundings(pid uint64, srcefundings []jsonfile.Fund) []Funding {
	var DBFundings []Funding
	for _, f := range srcefundings {
		df := Funding{Projectid: pid}
		df.Description = GetStringValue(f.FundingStream.Description)
		df.StreamID = GetStringValue(f.FundingStream.ID)
		df.Name = GetStringValue(f.Name)
		df.ShortName = GetStringValue(f.ShortName)
		df.Jurisdiction = GetStringValue(f.Jurisdiction)
		DBFundings = append(DBFundings, df)
	}
	return DBFundings
}

func DeriveDBH2020(pid uint64, srceh2020s []jsonfile.H2020) []H2020 {
	var DBH2020s []H2020
	for _, h := range srceh2020s {
		dh := H2020{Projectid: pid}
		dh.Code = GetStringValue(h.Code)
		dh.Description = GetStringValue(h.Description)
		DBH2020s = append(DBH2020s, dh)
	}
	return DBH2020s
}

func DeriveDBSubjects(pid uint64, srcesubjects []string) []PSubject {
	var DBSubjects []PSubject
	for _, s := range srcesubjects {
		ds := PSubject{Projectid: pid}
		ds.Subject = s
		DBSubjects = append(DBSubjects, ds)
	}
	return DBSubjects
}

func GetStringValue(fieldpointer *string) string {
	if fieldpointer == nil {
		return "null"
	} else {
		return *fieldpointer
	}
}

func GetFloatValue(fieldpointer *float64) string {
	if fieldpointer == nil {
		return "null"
	} else {
		return strconv.FormatFloat(*fieldpointer, 'f', -1, 64)
	}
}

func GetBoolValue(fieldpointer *bool) string {
	if fieldpointer == nil {
		return "null"
	} else {
		if *fieldpointer {
			return "true"
		} else {
			return "false"
		}
	}
}
