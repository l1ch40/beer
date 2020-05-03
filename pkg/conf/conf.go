package conf

import (
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"os/exec"
	user2 "os/user"
	"path/filepath"
	"strings"
)

var (
	configDir string = "%s/.config/beer/"
	configName string = "beer.ini"
	configPath string
	cfg *ini.File = nil
)

type Beer struct {
	Name string
	Command string
}

func init() {
	user, err := user2.Current()
	if err != nil {
		log.Fatal(err)
	}
	configDir = fmt.Sprintf(configDir, user.HomeDir)
	configPath = filepath.Join(configDir, configName)

	_, err = os.Stat(configDir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(configDir, os.ModePerm)
			if err != nil {
				panic(fmt.Sprintf("The beer create configDir failed %s", err))
			}
		}
	}

	_, err = os.Stat(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(configPath)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	cfg, err = ini.Load(configPath)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	cfg, _ = ini.InsensitiveLoad(configPath)
}

func Add(args []string) {
	softwareName, updateCommand := args[0], args[1]
	cfg.Section("").DeleteKey(softwareName)
	cfg.Section("").NewKey(softwareName, updateCommand)

	cfg.SaveTo(configPath)
}

func Remove(args []string) {
	for _, key := range args {
		cfg.Section("").DeleteKey(key)
	}

	cfg.SaveTo(configPath)
}

func Info(key string) (string, error) {
	yes := cfg.Section("").HasKey(key)
	if yes == false {
		return "", errors.New(fmt.Sprintf("The %s software doesn't exist", key))
	}
	value := cfg.Section("").Key(key).String()
	return value, nil
}

func List() ([]Beer, error) {
	names := cfg.Section("").KeyStrings()
	var beers []Beer
	for _, name := range names {
		beers = append(beers, Beer{
			Name: name,
			Command: cfg.Section("").Key(name).String(),
		})
	}
	return beers, nil
}

func Update() {
	names := cfg.Section("").KeyStrings()
	for _, name := range names {
		value := cfg.Section("").Key(name).String()
		commands := strings.Split(value, " ")
		cmd := exec.Command(commands[0], commands[1:]...)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}

func Upgrade(key string) {
	value := cfg.Section("").Key(key).String()
	commands := strings.Split(value, " ")
	cmd := exec.Command(commands[0], commands[1:]...)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}