package database

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/robfig/cron/v3"
	"os"
	"pdb/utils"
	"strings"
	"time"
)

func StartScheduler() {
	c := cron.New()

	_, _ = c.AddFunc("@midnight", func() {
		config := utils.GetConfig()

		if len(config.FtpUrl) > 0 {
			fmt.Println("Creating backup for database.db...")
			CreateBackup(config)

			fmt.Println("Backup finished.")
		}
	})

	c.Start()
}

func ConnectFTP(config *utils.Config) (*ftp.ServerConn, error) {
	c, err := ftp.Dial(config.FtpUrl, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		fmt.Printf("Error whilst connecting to ftp: %v\n", err)
		return nil, err
	}

	if err = c.Login(config.FtpUser, config.FtpPassword); err != nil {
		fmt.Printf("Error whilst login in to ftp: %v\n", err)
		return nil, err
	}

	return c, nil
}

func CreateBackup(config *utils.Config) {
	c, _ := ConnectFTP(config)
	if c == nil {
		return
	}

	defer c.Quit()

	UploadFile(c)
	DeleteOld(c)
}

func DeleteOld(ftp *ftp.ServerConn) {
	files, err := ftp.List("/pdb/")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if !strings.HasPrefix(file.Name, "database") {
			continue
		}

		now := time.Now()
		diff := now.Sub(file.Time)
		fmt.Printf("Files: %s,  %v, diff: %f\n", file.Name, file.Time, diff.Hours())

		if (diff.Hours() / 24.0) >= 3 {
			fileName := fmt.Sprintf("/pdb/%s", file.Name)
			if err = ftp.Delete(fileName); err != nil {
				fmt.Println(err)
			}
		}

	}
}

func UploadFile(ftp *ftp.ServerConn) {
	localFile, err := os.Open("data/database.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer localFile.Close()

	filePath := fmt.Sprintf("/pdb/database-%d.db", time.Now().Unix())

	err = ftp.Stor(filePath, localFile)
	if err != nil {
		fmt.Println(err)
	}
}
