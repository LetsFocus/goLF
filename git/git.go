package git

import (
	"fmt"
	"os"
	"os/exec"
)

func CloneRepository(repoURL, destinationPath string) {
	cmd := exec.Command("git", "clone", repoURL, destinationPath)
	runCommand(cmd)
}

func SetUser(username string) {
	cmd := exec.Command("git", "config", "--global", "user.name", username)
	runCommand(cmd)
}

func SetEmail(email string) {
	cmd := exec.Command("git", "config", "--global", "user.email", email)
	runCommand(cmd)
}

func AddFiles() {
	cmd := exec.Command("git", "add", ".")
	runCommand(cmd)
}

func CommitChanges(message string) {
	cmd := exec.Command("git", "commit", "-m", message)
	runCommand(cmd)
}

func PushChanges(remote, branch string) {
	cmd := exec.Command("git", "push", remote, branch)
	runCommand(cmd)
}

func CreateBranch(branchName string) {
	cmd := exec.Command("git", "branch", branchName)
	runCommand(cmd)
}

func SwitchBranch(branchName string) {
	cmd := exec.Command("git", "checkout", branchName)
	runCommand(cmd)
}

func MergeBranch(sourceBranch string) {
	cmd := exec.Command("git", "merge", sourceBranch)
	runCommand(cmd)
}

func FetchChanges(remote string) {
	cmd := exec.Command("git", "fetch", remote)
	runCommand(cmd)
}

func ListBranches() {
	cmd := exec.Command("git", "branch", "--list")
	runCommand(cmd)
}

func DeleteBranch(branchName string) {
	cmd := exec.Command("git", "branch", "-d", branchName)
	runCommand(cmd)
}

func RenameBranch(oldName, newName string) {
	cmd := exec.Command("git", "branch", "-m", oldName, newName)
	runCommand(cmd)
}

func ShowCommitHistory() {
	cmd := exec.Command("git", "log")
	runCommand(cmd)
}

func ShowDiff(commitHash1, commitHash2 string) {
	cmd := exec.Command("git", "diff", commitHash1, commitHash2)
	runCommand(cmd)
}

func ShowCommitDiff(commitHash string) {
	cmd := exec.Command("git", "show", commitHash)
	runCommand(cmd)
}

func CheckoutFile(fileName string) {
	cmd := exec.Command("git", "checkout", fileName)
	runCommand(cmd)
}

func ResetStaged(fileName string) {
	cmd := exec.Command("git", "reset", fileName)
	runCommand(cmd)
}

func ResetHard(commitHash string) {
	cmd := exec.Command("git", "reset", "--hard", commitHash)
	runCommand(cmd)
}

func TagCommit(tagName, message, commitHash string) {
	cmd := exec.Command("git", "tag", "-a", tagName, "-m", message, commitHash)
	runCommand(cmd)
}

func ListTags() {
	cmd := exec.Command("git", "tag")
	runCommand(cmd)
}

func DeleteTag(tagName string) {
	cmd := exec.Command("git", "tag", "-d", tagName)
	runCommand(cmd)
}

func ShowTagDiff(tagName string) {
	cmd := exec.Command("git", "show", tagName)
	runCommand(cmd)
}

func FetchTag(remote, tagName string) {
	cmd := exec.Command("git", "fetch", remote, tagName)
	runCommand(cmd)
}

func CherryPickCommit(commitHash string) {
	cmd := exec.Command("git", "cherry-pick", commitHash)
	runCommand(cmd)
}

func StashChanges() {
	cmd := exec.Command("git", "stash")
	runCommand(cmd)
}

func ApplyStash() {
	cmd := exec.Command("git", "stash", "apply")
	runCommand(cmd)
}

func DropStash() {
	cmd := exec.Command("git", "stash", "drop")
	runCommand(cmd)
}

func CleanFiles() {
	cmd := exec.Command("git", "clean", "-f")
	runCommand(cmd)
}

func InitializeEmptyRepository() {
	cmd := exec.Command("git", "init")
	runCommand(cmd)
}

func GetStatus() {
	cmd := exec.Command("git", "status")
	runCommand(cmd)
}

func GetRemoteURL() {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	runCommand(cmd)
}

func runCommand(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing command:", err)
	}
}
