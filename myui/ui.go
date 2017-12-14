package main

import (
	"fmt"
	"os/exec"
)

type ShellCommand struct {
	Sh  string //shell命令
	arg string //命令参数
}

func (sh *ShellCommand) Initsh(src string) {
	sh.Sh, sh.arg = GetJob(src)

}

func (sh *ShellCommand) Clear() {
	var bi ShellCommand
	*sh = bi
}

type Jobs struct {
	Src       string //原字符串，字符串以 ; 分隔
	Command   string //解析出来的命令
	date      string
	ShCommand ShellCommand //shell命令
}

func (job *Jobs) InitJob() {
	job.Command, job.date = GetJob(job.Src)
	if job.Command == "sh" {
		job.ShCommand.Sh, job.ShCommand.arg = GetJob(job.date)
	}

}

func (job *Jobs) Clear() {
	var bi Jobs
	*job = bi
}

func main() {

	var job Jobs
	fmt.Printf("请输入:\n")
	for {
		job.Clear()
		//fmt.Println("bi", job)
		fmt.Printf("\033[32m Note$ \033[0m")
		fmt.Scanln(&job.Src)
		job.InitJob()
		//fmt.Println(job)
		switch job.Command {
		case "conn":
			fmt.Printf("conn: %s\n", job.date)
		case "get":
			fmt.Printf("get: %s\n", job.date)
		case "send":
			fmt.Printf("send: %s\n", job.date)
		case "sh":
			fmt.Printf("shell: %s\n", job.date)
			Shell(job.ShCommand.Sh, job.ShCommand.arg)
		case "exit":
			fmt.Printf("nice day~\n")
			return
		default:
			fmt.Printf("Try again~\n")

		}
	}

	fmt.Printf("nice day~\n")
}

func Shell(sh string, arg string) {
	if sh == "a" {
		ShellAlways()
	} else {
		ShellShow(sh, arg)
	}

}

func ShellAlways() {
	var sh ShellCommand
	var s string
	for {
		sh.Clear()
		fmt.Printf("\033[32m Shell$ \033[0m")
		fmt.Scanln(&s)
		sh.Initsh(s)
		if sh.Sh == "exit" {
			return
		} else {
			ShellShow(sh.Sh, sh.arg)
		}

	}
}

func ShellShow(sh string, arg string) {
	num := len(arg)
	if num == 0 {
		f, err := exec.Command(sh).Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(f))
	} else {
		f, err := exec.Command(sh, arg).Output()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(f))
	}

}

func GetJob(s string) (job string, date string) {
	sbuf := []byte(s)
	num := len(sbuf)
	//fmt.Println(sbuf, num)
	bi := num
	for i := 0; i < num; i++ {
		if string(sbuf[i]) == ";" {
			bi = i
			break
		}
	}
	job = string(sbuf[:bi])
	if bi != num {
		date = string(sbuf[bi+1 : num])
	}
	//fmt.Println(job)
	//fmt.Println(date)
	return
}
