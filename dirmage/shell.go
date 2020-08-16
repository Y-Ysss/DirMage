package dirmage

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	// "errors"
)

func Shell() {
	Select(RunShell)
}

func RunShell(info *dirInfo) {
	dirName := info.Name
	dirPath := info.Path
	rePttrn := regexp.MustCompile("%.*?%")
	homeDir, _ := os.UserHomeDir()
	dirPath = rePttrn.ReplaceAllString(string(dirPath), homeDir)
	chDirErr := os.Chdir(dirPath)
	if chDirErr != nil {
		fmt.Fprintln(os.Stderr, chDirErr)
	}
	promptStr := strings.Replace(conf.Prompt.Text, "{$DirName}", dirName, -1)

	rePttrn = regexp.MustCompile("{(0|3[0-9]|4[0-7])}")
	replaceFunc := func(s string) string {
		return fmt.Sprintf("\x1b[%sm", rePttrn.FindStringSubmatch(s)[1])
	}
	promptStr = rePttrn.ReplaceAllStringFunc(promptStr, replaceFunc)

	reader := bufio.NewReader(os.Stdin)
	for {
		prompt := strings.Replace(promptStr, "{$WorkingDir}", getWorkingDir(), -1)
		prompt = strings.Replace(prompt, "{$Git}", getGitBranch(), -1)
		fmt.Printf(prompt)
		cmdStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdStr)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func getGitBranch() string {
	_, err := os.Stat(".git")
	if err != nil {
		return ""
	}
	data, readErr := ioutil.ReadFile(".git/HEAD")
	if readErr != nil {
		log.Fatal(readErr)
		return ""
	}
	gitStRePttrn := regexp.MustCompile("ref: refs/heads/(.+)")
	s := gitStRePttrn.FindSubmatch(data)

	out, execErr := exec.Command("cmd", "/c", "git --no-optional-locks status --porcelain").Output()
	if execErr != nil {
		log.Fatal(execErr)
		return ""
	}
	color := "\x1b[37m" // White
	fmt.Printf("[[%d]]", len(out))
	if len(out) > 0 {
		color = "\x1b[31m" // Red
	}
	return fmt.Sprintf("%s(%s)\x1b[37m", color, s[1])
}

func getWorkingDir() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return wd
}

func runCommand(cmdStr string) error {
	cmdStr = strings.TrimSuffix(cmdStr, "\n")
	cmdStrArr := strings.Fields(cmdStr)
	if len(cmdStrArr) == 0 {
		return nil
	}
	switch cmdStrArr[0] {
	case "exit":
		os.Exit(0)
	case "cls", "clear":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		return cmd.Run()
	case "cd", "chdir":
		if len(cmdStrArr) == 1 {
			fmt.Printf("%s\n", getWorkingDir())
			return nil
		}
		return os.Chdir(cmdStrArr[1])
	case "start":
		if len(cmdStrArr) == 1 {
			return exec.Command("cmd", "/c", "start").Start()
		} else {
			return exec.Command("explorer", cmdStrArr[1:]...).Start()
		}
		// return cmd.Start()
	}
	cmd := exec.Command(cmdStrArr[0], cmdStrArr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
