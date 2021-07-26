package main

import (
	"fmt"
	"time"
	"whosbugAssigns"
)

func main() {
	t := time.Now()
	fmt.Println("Start!")
	whosbugAssigns.GetInputConfig()
	projectId := whosbugAssigns.Config.ProjectId
	branchName := whosbugAssigns.Config.BranchName
	repoPath := whosbugAssigns.Config.ProjectUrl
	resCommits := whosbugAssigns.Analysis(repoPath, branchName, projectId)
	whosbugAssigns.Result(resCommits, projectId, "1.0.0")
	//for _, resCommit := range resCommits {
	//	fmt.Println(resCommit.Commit, " ", resCommit.CommitDiffs[0].DiffContent[0]["Name"])
	//}

	fmt.Println("Whosbug analysis done: ", time.Since(t))
}
