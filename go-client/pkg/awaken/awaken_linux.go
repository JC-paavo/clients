package awaken

import (
	"fmt"
	"go-client/global"
	"go-client/pkg/config"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func getCommandFromArgs(connectInfo map[string]string, argFormat string) string {
	for key, value := range connectInfo {
		argFormat = strings.Replace(argFormat, "{"+key+"}", value, 1)
	}
	return argFormat
}

func awakenRDPCommand(filePath string) *exec.Cmd {
	global.LOG.Debug(filePath)
	cmd := exec.Command("remmina", filePath)
	return cmd
}

func awakenSSHCommand(r *Rouse, cfg *config.AppConfig) *exec.Cmd {
	var appItem *config.AppItem
	var appLst []config.AppItem
	switch r.Protocol {
	case "ssh":
		appLst = cfg.Linux.Terminal
	case "sftp":
		appLst = cfg.Linux.FileTransfer
	}

	for _, app := range appLst {
		if app.IsActive() && app.IsSupportProtocol(r.Protocol) {
			appItem = &app
			break
		}
	}
	if appItem == nil {
		return nil
	}
	var cmd *exec.Cmd
	connectMap := map[string]string{
		"name":     r.getName(),
		"protocol": r.Protocol,
		"username": r.getUserName(),
		"value":    r.Value,
		"host":     r.Host,
		"port":     strconv.Itoa(r.Port),
	}

	if appItem.IsInternal {
		currentPath := filepath.Dir(os.Args[0])
		commands := getCommandFromArgs(connectMap, appItem.ArgFormat)
		clientPath := filepath.Join(currentPath, "client")
		out, _ := exec.Command("bash", "-c", "echo $XDG_CURRENT_DESKTOP").CombinedOutput()
		currentDesktop := strings.ToLower(strings.Trim(string(out), "\n"))

		switch currentDesktop {
		case "gnome", "ubuntu:gnome":
			cmd = exec.Command(
				"gnome-terminal", "--", "bash", "-c",
				fmt.Sprintf("%s %s; exec bash -i", clientPath, commands),
			)
		case "deepin":
			cmd = exec.Command("deepin-terminal", "--keep-open", "-C", fmt.Sprintf("%s %s", clientPath, commands))
		default:
			msg := fmt.Sprintf("Not yet supported %s desktop system", currentDesktop)
			global.LOG.Info(msg)
		}
	} else {
		var appPath string
		appPath = appItem.Path
		commands := getCommandFromArgs(connectMap, appItem.ArgFormat)
		appPath = appItem.Path
		cmd = exec.Command(appPath, strings.Split(commands, " ")...)
	}
	return cmd
}

func awakenDBCommand(r *Rouse, cfg *config.AppConfig) *exec.Cmd {
	var appItem *config.AppItem
	appLst := cfg.Linux.Databases
	for _, app := range appLst {
		if app.IsSet && app.IsMatchProtocol(r.Protocol) {
			appItem = &app
			break
		}
	}
	if appItem == nil {
		return nil
	}
	var cmd *exec.Cmd
	appPath := appItem.Path
	connectMap := map[string]string{
		"name":     r.getName(),
		"protocol": r.Protocol,
		"username": r.getUserName(),
		"value":    r.Value,
		"host":     r.Host,
		"port":     strconv.Itoa(r.Port),
		"dbname":   r.DBName,
	}
	if appItem.IsInternal {
		var argFormat string
		switch r.Protocol {
		case "redis":
			argFormat = "redis-cli -h {host} -p {port} -a {username}@{value}"
		case "oracle":
			argFormat = "sqlplus {username}/{value}@{host}:{port}/{dbname}"
		case "postgresql":
			argFormat = "psql 'user={username} password={value} host={host} dbname={dbname} port={port}'"
		case "mysql", "mariadb":
			argFormat = "mysql -u {username} -p{value} -h {host} -P {port} {dbname}"
		}
		commands := getCommandFromArgs(connectMap, argFormat)

		out, _ := exec.Command("bash", "-c", "echo $XDG_CURRENT_DESKTOP").CombinedOutput()
		currentDesktop := strings.ToLower(strings.Trim(string(out), "\n"))

		switch currentDesktop {
		case "gnome", "ubuntu:gnome":
			cmd = exec.Command(
				"gnome-terminal", "--", "bash", "-c",
				fmt.Sprintf("%s; exec bash -i", commands),
			)
		case "deepin":
			cmd = exec.Command("deepin-terminal", "--keep-open", "-C", commands)
		default:
			msg := fmt.Sprintf("Not yet supported %s desktop system", currentDesktop)
			global.LOG.Info(msg)
		}
		return cmd
	} else {
		appPath = appItem.Path
		commands := getCommandFromArgs(connectMap, appItem.ArgFormat)
		return exec.Command(appPath, strings.Split(commands, " ")...)
	}
}

func awakenOtherCommand(r *Rouse, cfg *config.AppConfig) *exec.Cmd {
	cmd := new(exec.Cmd)
	return cmd
}
