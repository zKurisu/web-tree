package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"
	"web-tree/conf"
)

// Every change? Delete a file? Already exist --> Add index/time
func Backup(name string) {
	if !IsTreeExist(name) {
		log.Fatal("Tree " + name + " does not exist")
	}
	backDir := conf.GetBackDir()
	storeDir := conf.GetStoreDir()
	_, err := os.Stat(backDir)

	if os.IsNotExist(err) {
		os.Mkdir(backDir, 0755)
	}

	dest := filepath.Join(backDir, AddFileExtention(name)+"."+GetFormatCurTime())
	src := filepath.Join(storeDir, AddFileExtention(name))
	BackFile(dest, src)
}

func BackFile(dest string, src string) error {
	src_hd, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer src_hd.Close()

	dest_hd, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer dest_hd.Close()

	_, err = io.Copy(dest_hd, src_hd)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func AddFileExtention(name string) string {
	pattern := regexp.QuoteMeta(`.yaml`)
	re := regexp.MustCompile(pattern + "$")
	if re.MatchString(name) {
		return name
	} else {
		return name + `.yaml`
	}
}

func RemoveFileExtention(name string) string {
	pattern := regexp.QuoteMeta(`.yaml`)
	re := regexp.MustCompile(pattern + "$")
	if re.MatchString(name) {
		return re.ReplaceAllString(name, "")
	} else {
		return name
	}
}

func GetFormatCurTime() string {
	curTime := time.Now()
	formatTime := curTime.Format("20060102-15-04-05")
	return formatTime
}
