package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

	// for now...
	//dbrepo.TruncateTables() // first file in sequence only

	// use scanner to take each line at a time
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// get current max value of the project id inthe database
	pid := dbrepo.GetMaxPiD()

	for scanner.Scan() {
		projectLine := scanner.Text()
		if projectLine != "" {
			reader := strings.NewReader(string(projectLine))
			var fp jsonfile.Project
			if err := json.NewDecoder(reader).Decode(&fp); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			var title, summary string
			// see if object contains the words 'clinical' or 'health' or 'hospital' or 'care',
			// and 'trial' or 'study' in their title or summary...

			ptitle := dbrepo.GetStringValue(fp.Title)
			if ptitle != "null" {
				title = strings.ToLower(ptitle)
			} else {
				title = ""
			}

			psummary := dbrepo.GetStringValue(fp.Summary)
			if psummary != "null" {
				summary = strings.ToLower(psummary)
			} else {
				summary = ""
			}

			ProjectOfInterest := false
			if title != "" {
				if (strings.Contains(title, "clinical") || strings.Contains(title, "health") ||
					strings.Contains(title, "hospital")) &&
					(strings.Contains(title, "trial") || strings.Contains(title, "study")) {
					ProjectOfInterest = true
				}
			}

			if !ProjectOfInterest && summary != "" {
				if (strings.Contains(summary, "clinical") || strings.Contains(summary, "health") ||
					strings.Contains(summary, "hospital")) &&
					(strings.Contains(summary, "trial") || strings.Contains(summary, "study")) {
					ProjectOfInterest = true
				}
			}

			// if so get further details and store
			//ProjectOfInterest = true

			if ProjectOfInterest {
				pid++
				dp := dbrepo.DeriveDBProject(pid, fp)

				var dfs []dbrepo.Funding = nil
				var dhs []dbrepo.H2020 = nil
				var dss []dbrepo.PSubject = nil

				if len(fp.Funding) > 0 {
					dfs = dbrepo.DeriveDBFundings(pid, fp.Funding)
				}
				if len(fp.H2020Programme) > 0 {
					dhs = dbrepo.DeriveDBH2020(pid, fp.H2020Programme)
				}
				if len(fp.Subject) > 0 {
					dss = dbrepo.DeriveDBSubjects(pid, fp.Subject)
				}
				dbrepo.ProcessProjectData(dp, dfs, dhs, dss)
			}
		}

	}
}
