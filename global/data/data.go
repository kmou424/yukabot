package data

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/gookit/goutil/fsutil"
	"github.com/kmou424/dingtalkbot"
	"github.com/kmou424/yukabot/cmd/cli"
	. "github.com/kmou424/yukabot/global/logger"
)

type SharedData struct {
	Channels *dingtalkbot.RWMap[string, string]
}

var Data = &SharedData{
	Channels: dingtalkbot.NewRWMap[string, string](),
}

func init() {
	if fsutil.IsFile(cli.DataJsonFile) {
		content := fsutil.ReadFile(cli.DataJsonFile)
		err := json.Unmarshal(content, Data)
		if err != nil {
			Logger.Warn("failed to read data file", "file", cli.DataJsonFile, "err", err)
		}
	}
}

func SyncDataToFS() {
	content, _ := json.MarshalIndent(Data, "", "    ")
	if dir := filepath.Dir(cli.DataJsonFile); !fsutil.IsDir(dir) {
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
	err := fsutil.WriteFile(cli.DataJsonFile, content, 0o644)
	if err != nil {
		Logger.Error("failed to write data file", "file", cli.DataJsonFile, "err", err)
	}
}
