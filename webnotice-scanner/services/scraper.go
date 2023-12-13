/*
user.go

Copyright (c) 2022 The OpenDataology Authors
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package service

import (
	"net/http"

	"github.com/OpenDataology/webnotice-scanner/webnotice-scanner/models"
	"github.com/gin-gonic/gin"
)

func GetCopyrightService(c *gin.Context, url string) (h gin.H) {
	user, err := models.GetCopyright(url)
	if err != nil {
		if err.Error() == "Url is invalid!" {
			c.JSON(http.StatusOK, gin.H{"err": "Url is invalid!"})
		} else {
			c.JSON(http.StatusOK, gin.H{"err": "DB Error"})
		}
		return
	}
	res := gin.H{
		"data": &user,
	}
	return res
}
