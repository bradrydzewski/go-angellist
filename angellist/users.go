// Copyright 2013 The go-angellist AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package angellist

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Followers int    `json:"follower_count"`
	Investor  bool   `json:"investor"`

	// Links to social websites where the
	// user has a professional presence.
	Blog      string `json:"blog_url"`
	OnlineBio string `json:"online_bio_url"`
	AngelList string `json:"angellist_url"`
	Twitter   string `json:"twitter_url"`
	Facebook  string `json:"facebook_url"`
	LinkedIn  string `json:"linkedin_url"`
	AboutMe   string `json:"aboutme_url"`
	GitHub    string `json:"github_url"`
	Dribbble  string `json:"dribbble_url"`
	Behance   string `json:"behance_url"`

	// Child collections
	Locations []*Location `json:"locations"`
	Roles     []*Role     `json:"roles"`
	Skills    []*Skill    `json:"skills"`

	// Optional Investor Detail
	InvestorDetail *InvestorDetail `json:"investor_details"`

}

type Location struct {
	Id      int    `json:"id"`
    Type    string `json:"tag_type"`
    Name    string `json:"name"`
    Display string `json:"display_name"`
    Url     string `json:"angellist_url"`
}

type Role struct {
	Id      int    `json:"id"`
    Type    string `json:"tag_type"`
    Name    string `json:"name"`
    Display string `json:"display_name"`
    Url     string `json:"angellist_url"`
}

type Skill struct {
	Id      int    `json:"id"`
    Type    string `json:"tag_type"`
    Name    string `json:"name"`
    Display string `json:"display_name"`
    Url     string `json:"angellist_url"`
}

type Market struct {
	Id      int    `json:"id"`
    Type    string `json:"tag_type"`
    Name    string `json:"name"`
    Display string `json:"display_name"`
    Url     string `json:"angellist_url"`
}

type Investment struct {
	Id      int    `json:"id"`
    Name    string `json:"name"`
	Quality int    `json:"quality"`
}

type InvestorDetail struct {
	Accreditation   string `json:"accreditation"`
    StartupsPerYear string `json:"startups_per_year"`
    AverageAmount   string `json:"average_amount"`

	// Child collections
	Locations   []*Location   `json:"locations"`
	Investments []*Investment `json:"investments"`
	Markets     []*Market     `json:"markets"`
}


type UserResource struct {
	client *Client
}

// Get a user's information given an id.
func (u *UserResource) Get(id int) (*User, error) {
	path := fmt.Sprintf("/1/users/%d?include_details=investor", id)
	user := User{}
	if err := u.client.do("GET", path, nil, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Get a user's information given an id.
func (u *UserResource) GetMulti(ids ...int) ([]*User, error) {
	// convert the int array to a string array
	idstr := []string{}
	for _, id := range ids {
		idstr = append(idstr, strconv.Itoa(id))
	}

	// convert the string array to a comma-separated string
	// and create the URL
	path := fmt.Sprintf("/1/users/batch?ids=%s&include_details=investor", strings.Join(idstr, ","))
	users := []*User{}
	if err := u.client.do("GET", path, nil, &users); err != nil {
		return nil, err
	}
	return users, nil
}

// Search for a user given a URL slug.
func (u *UserResource) GetSlug(slug string) (*User, error) {
	path := fmt.Sprintf("/1/users/search?include_details=investor&slug=%s", slug)
	user := User{}
	if err := u.client.do("GET", path, nil, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Search for a user given an MD5 email hash.
func (u *UserResource) GetEmail(email string) (*User, error) {
	// calculate emails md5 hash
	h := md5.New()
    io.WriteString(h, email)
    hash := h.Sum(nil)

	path := fmt.Sprintf("/1/users/search?include_details=investor&md5=%x", hash)
	user := User{}
	if err := u.client.do("GET", path, nil, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
