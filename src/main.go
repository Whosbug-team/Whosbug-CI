package main

import (
	"fmt"
	"os"
	"time"
	"whosbugAssigns"
)

func main12() {
	t := time.Now()
	fmt.Println("Start!")
	s, _ := os.Getwd()
	fmt.Println(s)
	whosbugAssigns.GetInputConfig()
	projectId := whosbugAssigns.Config.ProjectId
	branchName := whosbugAssigns.Config.BranchName
	repoPath := whosbugAssigns.Config.ProjectUrl
	resCommits := whosbugAssigns.Analysis(repoPath, branchName, projectId)
	whosbugAssigns.Result(resCommits, projectId, "1.0.0")
	fmt.Println("All time cost: ", time.Since(t))
}
