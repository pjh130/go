package controllers

import (
	"github.com/pjh130/go/project/myfile/models"
)

func WatchFiles() {
	models.StartWatch()
}
