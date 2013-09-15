// Copyright 2013 The go-angellist AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package angellist

import (
	"fmt"
	"time"
)

type Startup struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Desc      string `json:"product_desc"`
	Concept   string `json:"high_concept"`
	Website   string `json:"company_url"`
	Followers int    `json:"follower_count"`
	Hidden    bool   `json:"hidden"`

	// associated media URLs
	Logo  string `json:"logo_url"`
	Thumb string `json:"thumb_url"`
	Video string `json:"video_url"`

	// community_profile is true if the company's
	// profile was automatically generated and has not
	// been 'claimed' by anyone at the company.
	CommunityProfile bool `json:"community_profile"`

	// quality is an integer between 0 and 10, calculated
	// every 48 hours, and reflects the company's rank on
	// AngelList. Higher numbers mean better quality.
	Quality int `json:"quality"`

	Created *time.Time `json:"created_at"` // format "2011-03-18T00:24:29Z"
	Updated *time.Time `json:"updated_at"`

	// Links to social websites where the
	// user has a professional presence.
	Blog       string `json:"blog_url"`
	Twitter    string `json:"twitter_url"`
	AngelList  string `json:"angellist_url"`
	Crunchbase string `json:"crunchbase_url"`

	Screenshots []*Screenshot  `json:"screenshots"`
	CompanyType []*CompanyType `json:"company_type"`
	Locations   []*Location    `json:"locations"`
	Markets     []*Market      `json:"markets"`
}

type Screenshot struct {
    Thumbnail string `json:"thumb"`
    Original  string `json:"original"`
}

type CompanyType struct {
	Id      int    `json:"id"`
    Type    string `json:"tag_type"`
    Name    string `json:"name"`
    Display string `json:"display_name"`
    Url     string `json:"angellist_url"`
}

type Comment struct {
	Id      int        `json:"id"`
	Comment string     `json:"comment"`
	Created *time.Time `json:"created_at"`
	User    *User      `json:"user"`
}


type StartupResource struct {
	client *Client
}

// Get a company's information given an id. 
// see http://goo.gl/ktfUI5
func (s *StartupResource) Get(id int) (*Startup, error) {
	path := fmt.Sprintf("/1/startups/%d", id)
	startup := Startup{}
	if err := s.client.do("GET", path, nil, &startup); err != nil {
		return nil, err
	}

	return &startup, nil
}

// Returns the comments on the given company.
// see http://goo.gl/NLTiRX
func (s *StartupResource) GetComments(id int) ([]*Comment, error) {
	path := fmt.Sprintf("/1/startups/%d/comments", id)
	comments := []*Comment{}
	if err := s.client.do("GET", path, nil, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}
