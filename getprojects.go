package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/scanhamman/oaprojects/packages/dbrepo"
	"github.com/scanhamman/oaprojects/packages/jsonfile"
)

func main() {
	fileName := os.Args[1]

	// Build the location of the source file.
	// filepath.Abs appends the file name to the default working directory.
	trialsFilePath, err := filepath.Abs(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Open the json file, with deferred closure.
	file, err := os.Open(trialsFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// Read in json file and decode to defined structures.
	var foundProjects jsonfile.Projects
	if err := json.NewDecoder(file).Decode(&foundProjects); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Obtain listed id for each trial
	// and add to a slice of strings.
	var pid uint64 = 0

	for _, t := range foundProjects {

		fmt.Printf("%+v\n\n", t)

		pid++
		dp := dbrepo.Project{Projectid: pid}
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
		fmt.Printf("%+v\n\n", dp)

		if len(t.Funding) > 0 {
			for _, f := range t.Funding {
				df := dbrepo.Funding{Projectid: pid}
				df.Description = GetStringValue(f.FundingStream.Description)
				df.StreamID = GetStringValue(f.FundingStream.ID)
				df.Name = GetStringValue(f.Name)
				df.ShortName = GetStringValue(f.ShortName)
				df.Jurisdiction = GetStringValue(f.Jurisdiction)
				fmt.Printf("%+v\n\n", df)
			}
		}

		if len(t.H2020Programme) > 0 {
			for _, h := range t.H2020Programme {
				dh := dbrepo.H2020{Projectid: pid}
				dh.Code = GetStringValue(h.Code)
				dh.Description = GetStringValue(h.Description)
				fmt.Printf("%+v\n\n", dh)
			}
		}

		if len(t.Subject) > 0 {
			for _, s := range t.Subject {
				ds := dbrepo.Subject{Projectid: pid}
				ds.Subject = s
				fmt.Printf("%+v\n\n", ds)
			}
		}
	}
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
