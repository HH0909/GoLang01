package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

// declare struct SelPg
type SelpgArgs struct {
	startP int    //start page number
	endP   int    //end page number
	len    int    //how many lines per page
	typeP  int    //type of page decide whether to yse '/f'
	desP   string //destination to print
	name   string //file name
}

func main() {
	sa := SelpgArgs{
		startP: -1,
		endP:   -1,
		len:    72,
		typeP:  1,
		name:   "",
		desP:   "",
	}
	flag.IntVar(&sa.startP, "s", -1, "specify start page")
	flag.IntVar(&sa.endP, "e", -1, "specify end page")
	flag.IntVar(&sa.len, "l", 72, "specify how many lines per page,default 72")
	typeP := flag.Bool("f", false, "-f means page(type2), and you can't set -f and page length at the same time")
	desP := flag.String("d", "", "specify print dest.")
	flag.Usage = usage
	flag.Parse()

	if sa.startP == -1 || sa.endP == -1 || sa.startP > sa.endP || sa.startP < 1 || sa.endP < 1 {
		flag.Usage()
		return
	}

	if sa.len != 72 && *typeP == true {
		flag.Usage()
		return
	}

	if *typeP == true {
		sa.typeP = 2
	}

	if *desP != "" {
		sa.desP = *desP
	}

	if len(flag.Args()) > 1 {
		flag.Usage()
		return
	}

	if len(flag.Args()) == 1 {
		sa.name = flag.Args()[0]
	}

	if sa.typeP == 1 {
		Processor1(sa, sa.name != "", sa.desP != "")
	} else {
		Processor2(sa, sa.name != "", sa.desP != "")
	}
}

func Processor1(sa SelpgArgs, f bool, p bool) {
	cmd := exec.Command("cat", "-n")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	curP := 1
	curL := 0
	if f {
		fileIn, err := os.OpenFile(sa.name, os.O_RDWR, os.ModeType)
		defer fileIn.Close()
		if err != nil {
			panic(err)
			return
		}
		line := bufio.NewScanner(fileIn)
		for line.Scan() {
			if curP >= sa.startP && curP <= sa.endP {
				os.Stdout.Write([]byte(line.Text() + "\n"))
				stdin.Write([]byte(line.Text() + "\n"))
			}
			curL++
			if curL %= sa.len; curL == 0 {
				curP++
			}
		}
	} else {
		line := bufio.NewScanner(os.Stdin)
		for line.Scan() {
			if curP >= sa.startP && curP <= sa.endP {
				os.Stdout.Write([]byte(line.Text() + "\n"))
				stdin.Write([]byte(line.Text() + "\n"))
			}
			curL++
			if curL %= sa.len; curL == 0 {
				curP++
			}
		}
	}
	if curP < sa.endP {
		fmt.Fprint(os.Stderr, "We don't have this page\n")
	}
	if p {
		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Start()
	}

}

func Processor2(sa SelpgArgs, f bool, p bool) {
	cmd := exec.Command("cat", "-n")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	curP := 1
	if f {
		fileIn, err := os.OpenFile(sa.name, os.O_RDWR, os.ModeType)
		defer fileIn.Close()
		if err != nil {
			panic(err)
			return
		}
		line := bufio.NewScanner(fileIn)
		for line.Scan() {
			flag := false
			for _, c := range line.Text() {
				if c == '\f' {
					if curP >= sa.startP && curP <= sa.endP {
						flag = true
						os.Stdout.Write([]byte("\n"))
						stdin.Write([]byte("\n"))
					}
					curP++
				} else {
					if curP >= sa.startP && curP <= sa.endP {
						os.Stdout.Write([]byte(string(c)))
						stdin.Write([]byte(string(c)))
					}
				}
			}
			if flag != true && curP >= sa.startP && curP <= sa.endP {
				os.Stdout.Write([]byte("\n"))
				stdin.Write([]byte("\n"))
			}
			flag = false
		}
	} else {
		line := bufio.NewScanner(os.Stdin)
		for line.Scan() {
			flag := false
			for _, c := range line.Text() {
				if c == '\f' {
					if curP >= sa.startP && curP <= sa.endP {
						flag = true
						os.Stdout.Write([]byte("\n"))
						stdin.Write([]byte("\n"))
					}
					curP++
				} else {
					if curP >= sa.startP && curP <= sa.endP {
						os.Stdout.Write([]byte(string(c)))
						stdin.Write([]byte(string(c)))
					}
				}
			}
			if flag != true && curP >= sa.startP && curP <= sa.endP {
				os.Stdout.Write([]byte("\n"))
				stdin.Write([]byte("\n"))
			}
			flag = false
		}
	}
	if curP < sa.endP {
		fmt.Fprint(os.Stderr, "We don't have this page\n")
	}
	if p {

		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Start()
	}
}

func usage() {
	fmt.Fprint(os.Stderr, `Usage: [-s start page(>=1)] [e end page(>=s)]
    [-l page length(dafault 72)] [-f file type(default true)] [-d destination]
    [filename specify input file]`)
}
