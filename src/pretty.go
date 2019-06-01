package main

import (
	"fmt"
	"github.com/kr/pretty"
)

func main() {

	type Project struct {
		Id    int64  `json:"project_id"`
		Title string `json:"title"`
		Name  string `json:"name"`
		//	Data Data `json:"data"`
		//	Commits Commits `json:"commits"`
	}

	var x = Project{1, "aaa", "{5, 6}"}
	fmt.Printf("%# v", pretty.Formatter(x)) //It will print all struct details

	fmt.Printf("%# v", pretty.Formatter(x.Id)) //It will print component one by one.

}
