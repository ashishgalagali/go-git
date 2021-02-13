package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"os"
	"time"
)

func main() {

	defer fmt.Println(time.Since(time.Now()))
	r := Checkout("https://github.com/apache/struts", "4b1262ef0fdf72aa3bbdf601e0fa3a255561e3c8")

	h, err := r.ResolveRevision(plumbing.Revision("4b1262ef0fdf72aa3bbdf601e0fa3a255561e3c8"))
	CheckIfError(err)
	commitObj, err := r.CommitObject(*h)
	CheckIfError(err)

	//TODO: blame is now specific to a file
	//TODO: Aggregation over all the files and for a range of commits
	//TODO: Create detailed 2D matrix similar to the hercules override matrix
	//TODO: Performance
	// TODO: check diff 16/24




	blame, err := git.Blame(commitObj, "README.md")
	//println(blame.Churns)

	print(blame.Churns)
	if err == nil {
		for _, churn := range blame.Churns {
			// Convert structs to JSON.
			data, _ := json.Marshal(churn)
			fmt.Printf("%s\n", data)
			fmt.Println("\n")
		}
	}

}

func Checkout(repoUrl, hash string) *git.Repository {
	//PrintInBlue("git clone " + repoUrl)

	r := GetRepo(repoUrl)
	w, err := r.Worktree()
	CheckIfError(err)

	// ... checking out to commit
	//PrintInBlue("git checkout %s", hash)
	err = w.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(hash),
	})
	CheckIfError(err)
	return r
}

func GetRepo(repoUrl string) *git.Repository {
	//defer helper.Duration(helper.Track("GetRepo"))

	//PrintInBlue("git clone " + repoUrl)

	var r *git.Repository
	var err error
	//if strings.HasPrefix(repoUrl, "https://github.com") {
	r, err = git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{
		URL: repoUrl,
	})
	CheckIfError(err)
	//} else {
	//	r, err = git.PlainOpen(repoUrl)
	//	CheckIfError(err)
	//}
	return r
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
