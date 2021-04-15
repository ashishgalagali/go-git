package main

import (
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
	r := Checkout("https://github.com/ashishgalagali/SWEN610-project", "7368d5fcb7eec950161ed9d13b55caf5961326b6")

	h, err := r.ResolveRevision(plumbing.Revision("7368d5fcb7eec950161ed9d13b55caf5961326b6"))
	CheckIfError(err)
	commitObj, err := r.CommitObject(*h)
	CheckIfError(err)

	//hp, err := r.ResolveRevision(plumbing.Revision("79caa008ba1f9d06b34b4acc7c03d7fade185a63"))
	//CheckIfError(err)
	//
	//parentCommitObj, err := r.CommitObject(*hp)
	//CheckIfError(err)

	//TODO: Performance
	//TODO: Create detailed 2D matrix similar to the hercules override matrix

	_, _ = git.Blame(commitObj, "")
	//println(blame.Churns)

	//print(blame.)
	//if err == nil {
	//	for _, churn := range blame.Churns {
	//		// Convert structs to JSON.
	//		data, _ := json.Marshal(churn)
	//		fmt.Printf("%s\n", data)
	//		fmt.Println("\n")
	//	}
	//}

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
