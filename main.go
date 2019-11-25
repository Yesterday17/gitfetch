package main

import (
	"fmt"
	"github.com/Yesterday17/gitfetch/git"
	"github.com/Yesterday17/gitfetch/git/object"
)

func main() {
	f := git.NewFetcher("http://localhost:8884/.git/")
	head, _ := f.FetchHead()

	latestCommit := object.NewObject(f.FetchObject(f.FetchRef(head))).(object.CommitObject)
	fmt.Println(latestCommit.String())

	latestTree := object.NewObject(f.FetchObject(latestCommit.Tree())).(object.TreeObject)
	fmt.Println(latestTree.String())

	latestItems := latestTree.Items()
	for _, item := range latestItems {
		fmt.Println(item.String())

		//obj := object.NewObject(f.FetchObject(item.Hash()))
		//fmt.Println(obj.String())
	}
}
