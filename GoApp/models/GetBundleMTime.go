package models

import (
	"log"
	"os"
	"strconv"
)

func GetBundleMTime(bundleName string) (string, error) {
	stat, err := os.Stat("../public_html/" + bundleName)
	if err != nil {
		log.Fatal(err)
	}
	bundleModTime := stat.ModTime().Unix()
	return strconv.FormatInt(bundleModTime, 10), nil
}
