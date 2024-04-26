package cli

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/ini/v2/dotenv"
	. "github.com/kmou424/yukabot/global/logger"
)

var (
	EnvFile      string
	DataJsonFile string
)

func init() {
	flag.StringVar(&EnvFile, "env", ".env", "env file")
	flag.StringVar(&DataJsonFile, "data-json", "", "data file")
	flag.Parse()
	postOfParse()
}

func postOfParse() {
	if fsutil.IsFile(EnvFile) {
		err := dotenv.LoadFiles(EnvFile)
		if err != nil {
			Logger.Error("load .env failed", "err", err)
			os.Exit(1)
		}
	}

	if DataJsonFile == "" {
		Logger.Error("you must provide data file")
		os.Exit(1)
	}

	if dir := filepath.Dir(DataJsonFile); !fsutil.IsDir(dir) {
		if fsutil.IsFile(dir) {
			Logger.Error("data file dir is not a directory", "dir", dir)
			os.Exit(1)
		}
		err := fsutil.Mkdir(dir, os.ModeDir)
		if err != nil {
			Logger.Error("failed to create directory", "dir", dir, "err", err)
			os.Exit(1)
		}
	}
}
