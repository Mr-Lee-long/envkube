package changArgs

import (
	"fmt"
	"os/exec"
)

var args []string

func renameFile(fileIndex []int)(map[int]string){

	newFileList := make(map[int]string)
	for _,k := range(fileIndex) {
		fmt.Println("args", args[k])

		newFileName := "ekube-" + args[k]
		cmd := fmt.Sprintf("envsubst < %s > %s",args[k], newFileName)
		c := exec.Command("bash", "-c", cmd)
		_, err := c.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		newFileList[k] = newFileName
	}	
	return newFileList
}

func ArgsHandle(args []string)([]string){
	// 如果没有参数，直接返回
	if len(args) == 0 {
		return args
	}

	// 如果有apply，  或者create -f | delete  -f  那么就将 -f 后面的文件名修改wei  ekube-filename.yaml
	var fileIndex [] int
	for index, arg := range(args) {
		switch {
		case arg == "-f":
			fileIndex = append(fileIndex, index + 1)
			fmt.Println(fileIndex)
		}
	}

	newFileList := renameFile(fileIndex)

	for index, fileName := range(newFileList) {
		args[index] = fileName
	}
	return args
}
