/*
user.go

Copyright (c) 2022 The OpenDataology Authors
All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package controllers

import (
	"net/http"

	service "github.com/OpenDataology/webnotice-scanner/webnotice-scanner/services"

	"github.com/gin-gonic/gin"
)

func (a *BasicInfo) GetCopyright(c *gin.Context) {
	url := c.GetString("Url")

	res := service.GetCopyrightService(c, url)
	a.JsonSuccess(c, http.StatusOK, res)
}
