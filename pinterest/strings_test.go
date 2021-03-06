// Copyright 2013 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pinterest

import (
	"fmt"
	"testing"
)

func TestStringify(t *testing.T) {
	var nilPointer *string

	var tests = []struct {
		in  interface{}
		out string
	}{
		// basic types
		{"foo", `"foo"`},
		{123, `123`},
		{1.5, `1.5`},
		{false, `false`},
		{
			[]string{"a", "b"},
			`["a" "b"]`,
		},
		{
			struct {
				A []string
			}{nil},
			// nil slice is skipped
			`{}`,
		},
		{
			struct {
				A string
			}{"foo"},
			// structs not of a named type get no prefix
			`{A:"foo"}`,
		},

		// pointers
		{nilPointer, `<nil>`},
		{UserAccount{Username: String("abc")}, `pinterest.UserAccount{Username:"abc"}`},
	}

	for i, tt := range tests {
		s := Stringify(tt.in)
		if s != tt.out {
			t.Errorf("%d. Stringify(%q) => %q, want %q", i, tt.in, s, tt.out)
		}
	}
}

// Directly test the String() methods on various GitHub types. We don't do an
// exaustive test of all the various field types, since TestStringify() above
// takes care of that. Rather, we just make sure that Stringify() is being
// used to build the strings, which we do by verifying that pointers are
// stringified as their underlying value.
func TestString(t *testing.T) {
	var tests = []struct {
		in  interface{}
		out string
	}{
		{APIError{Code: 404, Message: "Pin not found."}, `pinterest.APIError{Code:404, Message:"Pin not found.", Status:"", Data:"", EndpointName:""}`},
		{UserAccount{Username: String("abc")}, `pinterest.UserAccount{Username:"abc"}`},
		{AuthorizationAPP{ClientID: "client id", ClientSecret: "client secret"}, `pinterest.AuthorizationAPP{ClientID:"client id", ClientSecret:"client secret", RedirectURI:"", Scope:""}`},
		{Metrics{Impression: Int64(3)}, `pinterest.Metrics{Impression:3}`},
		{DailyMetrics{DataStatus: String("READY"), Date: String("2022-02-10"), Metrics: &Metrics{Impression: Int64(3)}}, `pinterest.DailyMetrics{DataStatus:"READY", Date:"2022-02-10", Metrics:pinterest.Metrics{Impression:3}}`},
		{UserAccountAnalyticsMetrics{DailyMetrics: []*DailyMetrics{{DataStatus: String("READY"), Date: String("2022-02-10"), Metrics: &Metrics{Impression: Int64(3)}}}}, `pinterest.UserAccountAnalyticsMetrics{DailyMetrics:[pinterest.DailyMetrics{DataStatus:"READY", Date:"2022-02-10", Metrics:pinterest.Metrics{Impression:3}}]}`},
		{UserAccountAnalytics{All: &UserAccountAnalyticsMetrics{DailyMetrics: []*DailyMetrics{{DataStatus: String("READY"), Date: String("2022-02-10"), Metrics: &Metrics{Impression: Int64(3)}}}}}, `pinterest.UserAccountAnalytics{All:pinterest.UserAccountAnalyticsMetrics{DailyMetrics:[pinterest.DailyMetrics{DataStatus:"READY", Date:"2022-02-10", Metrics:pinterest.Metrics{Impression:3}}]}}`},
		{BoardOwner{Username: String("merleliukun")}, `pinterest.BoardOwner{Username:"merleliukun"}`},
		{Board{ID: String("1022106146619699845"), Name: String("City"), Description: String(""), Owner: &BoardOwner{Username: String("merleliukun")}, Privacy: String("PUBLIC")}, `pinterest.Board{ID:"1022106146619699845", Name:"City", Description:"", Owner:pinterest.BoardOwner{Username:"merleliukun"}, Privacy:"PUBLIC"}`},
		{BoardsResponse{Items: []*Board{{ID: String("1022106146619699845"), Name: String("City")}}}, `pinterest.BoardsResponse{Items:[pinterest.Board{ID:"1022106146619699845", Name:"City"}]}`},
		{BoardSection{ID: String("5215175925383086784"), Name: String("Day")}, `pinterest.BoardSection{ID:"5215175925383086784", Name:"Day"}`},
		{BoardSectionsResponse{Items: []*BoardSection{{ID: String("5215175925383086784"), Name: String("Day")}}}, `pinterest.BoardSectionsResponse{Items:[pinterest.BoardSection{ID:"5215175925383086784", Name:"Day"}]}`},
		{Pin{ID: String("1022106077902810180"), CreatedAt: String("2022-02-14T02:54:38")}, `pinterest.Pin{ID:"1022106077902810180", CreatedAt:"2022-02-14T02:54:38"}`},
		{PinsResponse{Items: []*Pin{{ID: String("1022106077902810180"), CreatedAt: String("2022-02-14T02:54:38")}}}, `pinterest.PinsResponse{Items:[pinterest.Pin{ID:"1022106077902810180", CreatedAt:"2022-02-14T02:54:38"}]}`},
		{Media{MediaType: String("image")}, `pinterest.Media{MediaType:"image"}`},
		{Image{Width: Int(10)}, `pinterest.Image{Width:10}`},
		{MediaUpload{MediaID: String("5216393791692388749"), MediaType: String("video")}, `pinterest.MediaUpload{MediaID:"5216393791692388749", MediaType:"video"}`},
		{MediaUploadsResponse{Items: []*MediaUpload{{MediaID: String("5216393791692388749")}}}, `pinterest.MediaUploadsResponse{Items:[pinterest.MediaUpload{MediaID:"5216393791692388749"}]}`},
		{RegisterMediaUploadResponse{MediaID: String("5216393791692388749"), UploadURL: String("https://p.com")}, `pinterest.RegisterMediaUploadResponse{MediaID:"5216393791692388749", UploadURL:"https://p.com"}`},
		{Ad{ID: String("687195134316"), Type: String("pinpromotion")}, `pinterest.Ad{ID:"687195134316", Type:"pinpromotion"}`},
		{AdsResponse{Items: []*Ad{{ID: String("687195134316"), IsRemovable: Bool(false)}}}, `pinterest.AdsResponse{Items:[pinterest.Ad{ID:"687195134316", IsRemovable:false}]}`},
		{AdAccount{ID: String("549763740754"), Name: String("SNS-SDKS")}, `pinterest.AdAccount{ID:"549763740754", Name:"SNS-SDKS"}`},
		{AdAccountsResponse{Items: []*AdAccount{{ID: String("549763740754")}}}, `pinterest.AdAccountsResponse{Items:[pinterest.AdAccount{ID:"549763740754"}]}`},
		{Campaign{ID: String("549755885175")}, `pinterest.Campaign{ID:"549755885175"}`},
		{CampaignsResponse{Items: []*Campaign{{ID: String("549755885175")}}}, `pinterest.CampaignsResponse{Items:[pinterest.Campaign{ID:"549755885175"}]}`},
		{TrackingURLs{Impression: []*string{String("URL1")}}, `pinterest.TrackingURLs{Impression:["URL1"]}`},
		{AdGroup{ID: String("2680060704746")}, `pinterest.AdGroup{ID:"2680060704746"}`},
		{AdGroupsResponse{Items: []*AdGroup{{ID: String("2680060704746")}}}, `pinterest.AdGroupsResponse{Items:[pinterest.AdGroup{ID:"2680060704746"}]}`},
	}

	for i, tt := range tests {
		s := tt.in.(fmt.Stringer).String()
		if s != tt.out {
			t.Errorf("%d. String() => %q, want %q", i, tt.in, tt.out)
		}
	}
}
