package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	goPs "github.com/mitchellh/go-ps"
)

/*
	8. 	Minishell
		Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

		- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
		- pwd - показать путь до текущего каталога
		- echo <args> - вывод аргумента в STDOUT
		- kill <args> - "убить" процесс, переданный в качестве аргумента (пример: такой-то пример)
		- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		request := strings.Split(scanner.Text(), " ")
		switch request[0] {
		case "cd":
			cd(request)
		case "pwd":
			pwd(request)
		case "echo":
			echo(request)
		case "kill":
			kill(request)
		case "ps":
			ps(request)
		}
	}
}

func cd(request []string) {
	switch len(request) {
	case 1:
		fmt.Fprintln(os.Stderr, "where dir?")
	case 2:
		err := os.Chdir(request[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	default:
		fmt.Fprintln(os.Stderr, "too many arguments")
	}
}

func pwd(request []string) {
	if len(request) == 1 {
		path, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Println(path)
		}
	} else {
		fmt.Fprintln(os.Stderr, "too many arguments")
	}
}

func echo(request []string) {
	for i := 1; i < len(request); i++ {
		fmt.Printf("%s ", request[i])
	}
	fmt.Println()
}

func kill(request []string) {
	if len(request) == 1 {
		fmt.Fprintln(os.Stderr, "what?")
		return
	}
	if len(request) > 2 {
		fmt.Fprintln(os.Stderr, "too many arguments")
		return
	}
	pid, err := strconv.Atoi(request[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	err = process.Kill()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func ps(request []string) {
	if len(request) != 1 {
		fmt.Fprintln(os.Stderr, "too many arguments")
		return
	}
	sliceProc, _ := goPs.Processes()

	for _, proc := range sliceProc {
		fmt.Printf("Process name: %v process id: %v\n", proc.Executable(), proc.Pid())
	}
}
