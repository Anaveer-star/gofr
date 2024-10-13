package main

import (
	"fmt"
	"strconv"
	"strings"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/datasource/file"
	"gofr.dev/pkg/gofr/datasource/file/ftp"
	"gofr.dev/pkg/gofr/logging"
)

type FileServerType int

var logger logging.Logger

const (
	FTP FileServerType = iota
	SFTP
)

// This can be a common function to configure both FTP and SFTP server.
func configureFileServer(app *gofr.App) file.FileSystemProvider {
	port, _ := strconv.Atoi(app.Config.Get("PORT"))

	return ftp.New(&ftp.Config{
		Host:      app.Config.Get("HOST"),
		User:      app.Config.Get("USER_NAME"),
		Password:  app.Config.Get("PASSWORD"),
		Port:      port,
		RemoteDir: app.Config.Get("REMOTE_DIR_PATH"),
	})
}

func printFiles(files []file.FileInfo, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		for _, f := range files {
			fmt.Println(f.Name())
		}
	}
}

func grepFiles(files []file.FileInfo, keyword string, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		for _, f := range files {
			if strings.HasPrefix(f.Name(), keyword) {
				fmt.Println(f.Name())
			}
		}
	}
}

func registerPwdCommand(app *gofr.App, fs file.FileSystemProvider) {
	app.SubCommand("pwd", func(_ *gofr.Context) (interface{}, error) {
		workingDirectory, err := fs.Getwd()

		return workingDirectory, err
	})
}

func registerLsCommand(app *gofr.App, fs file.FileSystemProvider) {
	app.SubCommand("ls", func(c *gofr.Context) (interface{}, error) {
		path := c.Param("path")
		files, err := fs.ReadDir(path)
		printFiles(files, err)

		return "", err
	})
}

func registerGrepCommand(app *gofr.App, fs file.FileSystemProvider) {
	app.SubCommand("grep", func(c *gofr.Context) (interface{}, error) {
		keyword := c.Param("keyword")
		path := c.Param("path")
		files, err := fs.ReadDir(path)
		grepFiles(files, keyword, err)

		return "", err
	})
}

func registerCreateFileCommand(app *gofr.App, fs file.FileSystemProvider) {
	app.SubCommand("createfile", func(c *gofr.Context) (interface{}, error) {
		fileName := c.Param("filename")
		logger.Log("Creating file : ", fileName)
		_, err := fs.Create(fileName)

		if err == nil {
			logger.Log("Successfully created file: ", fileName)
		}

		return "", err
	})
}

func registerRmCommand(app *gofr.App, fs file.FileSystemProvider) {
	app.SubCommand("rm", func(c *gofr.Context) (interface{}, error) {
		fileName := c.Param("filename")
		logger.Log("Removing file : ", fileName)
		err := fs.Remove(fileName)

		if err == nil {
			logger.Log("Successfully removed file: ", fileName)
		}

		return "", err
	})
}

func main() {
	app := gofr.NewCMD()

	logger = gofr.New().Logger()

	fileSystemProvider := configureFileServer(app)

	app.AddFileStore(fileSystemProvider)

	registerPwdCommand(app, fileSystemProvider)

	registerLsCommand(app, fileSystemProvider)

	registerGrepCommand(app, fileSystemProvider)

	registerCreateFileCommand(app, fileSystemProvider)

	registerRmCommand(app, fileSystemProvider)

	app.Run()
}
