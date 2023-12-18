/*
main.go

Copyright (c) 2022 The OpenDataology Authors
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type Copyright struct {
	Id              int    `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	License_name    string `gorm:"type:TEXT" json:"license_name,omitempty"`
	Licensor        string `gorm:"type:TEXT" json:"licensor,omitempty"`
	License_content string `gorm:"type:TEXT" json:"license_content,omitempty"`
	License_url     string `gorm:"type:TEXT" json:"license_url,omitempty"`
}

func fetch(url string) (string, error) {
	c := colly.NewCollector()

	var content string
	c.OnHTML("p, h1, h2, h3, h4, h5, h6", func(e *colly.HTMLElement) {
		content += strings.TrimSpace(e.Text) + "\n"
	})

	err := c.Visit(url)
	if err != nil {
		return "", err
	}

	return content, nil
}

func main() {
	var copyright *Copyright
	content, err := fetch("https://www.baidu.com")
	if err != nil {
		return
	}

	// Possible copyright keywords
	keywords := []string{"Copyright", "©", "版权", "百度"}

	for _, keyword := range keywords {
		index := strings.LastIndex(content, keyword)
		if index != -1 {
			copyright = &Copyright{
				License_content: content[index:],
			}
			fmt.Println(copyright)
			fmt.Println(content)
			return
		}
	}
}
