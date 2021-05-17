package database

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/robfig/cron/v3"
	"os"
	"pdb/utils"
	"time"
)

func StartScheduler() {
	c := cron.New()

	_, _ = c.AddFunc("@midnight", func() {
		config := utils.GetConfig()

		if len(config.FtpUrl) > 0 {
			fmt.Println("Creating backup for database.db...")
			UploadFile(config)

			fmt.Println("Backup finished.")
		}
	})

	c.Start()
}

func UploadFile(config *utils.Config) {
	c, err := ftp.Dial(config.FtpUrl, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.Login(config.FtpUser, config.FtpPassword)
	if err != nil {
		fmt.Println(err)
		return
	}

	localFile, err := os.Open("data/database.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer localFile.Close()

	err = c.Stor("/pdb/database.db", localFile)
	if err != nil {
		fmt.Println(err)
	}
}
