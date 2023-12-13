/*
user.go

Copyright (c) 2022 The OpenDataology Authors
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package models

type Copyright struct {
	Id              int    `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	License_name    string `gorm:"type:TEXT" json:"license_name,omitempty"`
	Licensor        string `gorm:"type:TEXT" json:"licensor,omitempty"`
	License_content string `gorm:"type:TEXT" json:"license_content,omitempty"`
	License_url     string `gorm:"type:TEXT" json:"license_url,omitempty"`
}

func GetCopyright(url string) (copyright *Copyright, err error) {

	return
}
