package models

import (
	"reflect"
	"strings"
	"testing"
)

const csvheaders = "zip,type,decommissioned,primary_city,acceptable_cities,unacceptable_cities,state,county,timezone,area_codes,world_region,country,latitude,longitude,irs_estimated_population_2014"
const validdata = csvheaders + `
98101,STANDARD,0,Seattle,,"Times Square",WA,"King County",America/Los_Angeles,"206,253,360,425",NA,US,47.61,-122.33,10270
98102,STANDARD,0,Seattle,,"Broadway, Capitol Hill",WA,"King County",America/Los_Angeles,"206,360,425",NA,US,47.63,-122.32,20490`
const invaliddata = `h1,h2
too,short`

func TestLoadZips(t *testing.T) {
	cases := []struct {
		name           string
		csvdata        string
		expectedOutput ZipSlice
		expectError    bool
	}{
		{
			"empty input",
			"",
			nil,
			true,
		},
		{
			"only headers",
			csvheaders,
			ZipSlice{},
			false,
		},
		{
			"valid rows",
			validdata,
			ZipSlice{
				&Zip{"98101", "Seattle", "WA"},
				&Zip{"98102", "Seattle", "WA"},
			},
			false,
		},
		{
			"invalid rows",
			invaliddata,
			nil,
			true,
		},
	}

	for _, c := range cases {
		output, err := LoadZips(strings.NewReader(c.csvdata), 2)
		if err != nil && !c.expectError {
			t.Errorf("case %s: unexpected error: %v", c.name, err)
		}
		if c.expectError && err == nil {
			t.Errorf("case %s: expected error not returned", c.name)
		}
		if !reflect.DeepEqual(output, c.expectedOutput) {
			t.Errorf("case %s: incorrect output: expected\n%v but got\n%v", c.name, c.expectedOutput, output)
		}
	}
}
