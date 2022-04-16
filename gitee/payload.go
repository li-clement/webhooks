package gitee

import (
	"time"
)

type CommentEventPayload struct {
	Action        *string          `json:"action,omitempty"`
	Comment       *NoteHook        `json:"comment,omitempty"`
	Repository    *ProjectHook     `json:"repository,omitempty"`
	Project       *ProjectHook     `json:"project,omitempty"`
	Author        *UserHook        `json:"author,omitempty"`
	Sender        *UserHook        `json:"sender,omitempty"`
	URL           *string          `json:"url,omitempty"`
	Note          *string          `json:"note,omitempty"`
	NoteableType  *string          `json:"noteable_type,omitempty"`
	NoteableID    int64            `json:"noteable_id,omitempty"`
	Title         *string          `json:"title,omitempty"`
	PerIID        *string          `json:"per_iid,omitempty"`
	ShortCommitID *string          `json:"short_commit_id,omitempty"`
	Enterprise    *EnterpriseHook  `json:"enterprise,omitempty"`
	PullRequest   *PullRequestHook `json:"pull_request,omitempty"`
	Issue         *IssueHook       `json:"issue,omitempty"`
	HookName      *string          `json:"hook_name,omitempty"`
	Password      *string          `json:"password,omitempty"`
}

type PushEventPayload struct {
	Ref                *string         `json:"ref,omitempty"`
	Before             *string         `json:"before,omitempty"`
	After              *string         `json:"after,omitempty"`
	TotalCommitsCount  int64           `json:"total_commits_count,omitempty"`
	CommitsMoreThanTen *bool           `json:"commits_more_than_ten,omitempty"`
	Created            *bool           `json:"created,omitempty"`
	Deleted            *bool           `json:"deleted,omitempty"`
	Compare            *string         `json:"compare,omitempty"`
	Commits            []CommitHook    `json:"commits,omitempty"`
	HeadCommit         *CommitHook     `json:"head_commit,omitempty"`
	Repository         *ProjectHook    `json:"repository,omitempty"`
	Project            *ProjectHook    `json:"project,omitempty"`
	UserID             int64           `json:"user_id,omitempty"`
	UserName           *string         `json:"user_name,omitempty"`
	User               *UserHook       `json:"user,omitempty"`
	Pusher             *UserHook       `json:"pusher,omitempty"`
	Sender             *UserHook       `json:"sender,omitempty"`
	Enterprise         *EnterpriseHook `json:"enterprise,omitempty"`
	HookName           *string         `json:"hook_name,omitempty"`
	Password           *string         `json:"password,omitempty"`
}

type IssueEventPayload struct {
	Action      *string         `json:"action,omitempty"`
	Issue       *IssueHook      `json:"issue,omitempty"`
	Repository  *ProjectHook    `json:"repository,omitempty"`
	Project     *ProjectHook    `json:"project,omitempty"`
	Sender      *UserHook       `json:"sender,omitempty"`
	TargetUser  *UserHook       `json:"target_user,omitempty"`
	User        *UserHook       `json:"user,omitempty"`
	Assignee    *UserHook       `json:"assignee,omitempty"`
	UpdatedBy   *UserHook       `json:"updated_by,omitempty"`
	IID         string          `json:"iid,omitempty"`
	Title       *string         `json:"title,omitempty"`
	Description *string         `json:"description,omitempty"`
	State       *string         `json:"state,omitempty"`
	Milestone   *string         `json:"milestone,omitempty"`
	URL         *string         `json:"url,omitempty"`
	Enterprise  *EnterpriseHook `json:"enterprise,omitempty"`
	HookName    *string         `json:"hook_name,omitempty"`
	Password    *string         `json:"password,omitempty"`
}

type MergeRequestEventPayload struct {
	Action         *string          `json:"action,omitempty"`
	ActionDesc     *string          `json:"action_desc,omitempty"`
	PullRequest    *PullRequestHook `json:"pull_request,omitempty"`
	Number         int64            `json:"number,omitempty"`
	IID            int64            `json:"iid,omitempty"`
	Title          *string          `json:"title,omitempty"`
	Body           *string          `json:"body,omitempty"`
	State          *string          `json:"state,omitempty"`
	MergeStatus    *string          `json:"merge_status,omitempty"`
	MergeCommitSha *string          `json:"merge_commit_sha,omitempty"`
	URL            *string          `json:"url,omitempty"`
	SourceBranch   *string          `json:"source_branch,omitempty"`
	SourceRepo     *RepoInfo        `json:"source_repo,omitempty"`
	TargetBranch   *string          `json:"target_branch,omitempty"`
	TargetRepo     *RepoInfo        `json:"target_repo,omitempty"`
	Project        *ProjectHook     `json:"project,omitempty"`
	Repository     *ProjectHook     `json:"repository,omitempty"`
	Author         *UserHook        `json:"author,omitempty"`
	UpdatedBy      *UserHook        `json:"updated_by,omitempty"`
	Sender         *UserHook        `json:"sender,omitempty"`
	TargetUser     *UserHook        `json:"target_user,omitempty"`
	Enterprise     *EnterpriseHook  `json:"enterprise,omitempty"`
	HookName       *string          `json:"hook_name,omitempty"`
	Password       *string          `json:"password,omitempty"`
}

type TagEventPayload struct {
	Action *string `json:"action,omitempty"`
}

// RepoInfo : Repository information
type RepoInfo struct {
	Project    *ProjectHook `json:"project,omitempty"`
	Repository *ProjectHook `json:"repository,omitempty"`
}

// LabelHook : Label, issue and pull request labels
type LabelHook struct {
	Id    int32  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

// EnterpriseHook : Enterprise information
type EnterpriseHook struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

// NoteHook : comment information
type NoteHook struct {
	Id        int32     `json:"id,omitempty"`
	Body      string    `json:"body,omitempty"`
	User      *UserHook `json:"user,omitempty"`
	CreatedAt string    `json:"created_at,omitempty"`
	UpdatedAt string    `json:"updated_at,omitempty"`
	HtmlUrl   string    `json:"html_url,omitempty"`
	Position  string    `json:"position,omitempty"`
	CommitId  string    `json:"commit_id,omitempty"`
}

// UserHook : user information
type UserHook struct {
	Id        int32     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Username  string    `json:"username,omitempty"`
	UserName  string    `json:"user_name,omitempty"`
	Url       string    `json:"url,omitempty"`
	Login     string    `json:"login,omitempty"`
	AvatarUrl string    `json:"avatar_url,omitempty"`
	HtmlUrl   string    `json:"html_url,omitempty"`
	Type_     string    `json:"type,omitempty"`
	SiteAdmin bool      `json:"site_admin,omitempty"`
	Time      time.Time `json:"time,omitempty"`
	Remark    string    `json:"remark,omitempty"`
}

// CommitHook : git commit information
type CommitHook struct {
	Id        string    `json:"id,omitempty"`
	TreeId    string    `json:"tree_id,omitempty"`
	ParentIds []string  `json:"parent_ids,omitempty"`
	Message   string    `json:"message,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Url       string    `json:"url,omitempty"`
	Author    *UserHook `json:"author,omitempty"`
	Committer *UserHook `json:"committer,omitempty"`
	Distinct  bool      `json:"distinct,omitempty"`
	Added     []string  `json:"added,omitempty"`
	Removed   []string  `json:"removed,omitempty"`
	Modified  []string  `json:"modified,omitempty"`
}

// MilestoneHook : milestone information
type MilestoneHook struct {
	Id           int32     `json:"id,omitempty"`
	HtmlUrl      string    `json:"html_url,omitempty"`
	Number       int32     `json:"number,omitempty"`
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	OpenIssues   int32     `json:"open_issues,omitempty"`
	ClosedIssues int32     `json:"closed_issues,omitempty"`
	State        string    `json:"state,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	DueOn        string    `json:"due_on,omitempty"`
}

// IssueHook : issue information
type IssueHook struct {
	Id            int32          `json:"id,omitempty"`
	HtmlUrl       string         `json:"html_url,omitempty"`
	Number        string         `json:"number,omitempty"`
	Title         string         `json:"title,omitempty"`
	User          *UserHook      `json:"user,omitempty"`
	Labels        []LabelHook    `json:"labels,omitempty"`
	State         string         `json:"state,omitempty"`
	StateName     string         `json:"state_name,omitempty"`
	TypeName      string         `json:"type_name,omitempty"`
	Assignee      *UserHook      `json:"assignee,omitempty"`
	Collaborators []UserHook     `json:"collaborators,omitempty"`
	Milestone     *MilestoneHook `json:"milestone,omitempty"`
	Comments      int32          `json:"comments,omitempty"`
	CreatedAt     time.Time      `json:"created_at,omitempty"`
	UpdatedAt     time.Time      `json:"updated_at,omitempty"`
	Body          string         `json:"body,omitempty"`
}

// ProjectHook : project information
type ProjectHook struct {
	Id              int32     `json:"id,omitempty"`
	Name            string    `json:"name,omitempty"`
	Path            string    `json:"path,omitempty"`
	FullName        string    `json:"full_name,omitempty"`
	Owner           *UserHook `json:"owner,omitempty"`
	Private         bool      `json:"private,omitempty"`
	HtmlUrl         string    `json:"html_url,omitempty"`
	Url             string    `json:"url,omitempty"`
	Description     string    `json:"description,omitempty"`
	Fork            bool      `json:"fork,omitempty"`
	PushedAt        string    `json:"pushed_at,omitempty"`
	CreatedAt       string    `json:"created_at,omitempty"`
	UpdatedAt       string    `json:"updated_at,omitempty"`
	SshUrl          string    `json:"ssh_url,omitempty"`
	GitUrl          string    `json:"git_url,omitempty"`
	CloneUrl        string    `json:"clone_url,omitempty"`
	SvnUrl          string    `json:"svn_url,omitempty"`
	GitHttpUrl      string    `json:"git_http_url,omitempty"`
	GitSshUrl       string    `json:"git_ssh_url,omitempty"`
	GitSvnUrl       string    `json:"git_svn_url,omitempty"`
	Homepage        string    `json:"homepage,omitempty"`
	StargazersCount int32     `json:"stargazers_count,omitempty"`
	WatchersCount   int32     `json:"watchers_count,omitempty"`
	ForksCount      int32     `json:"forks_count,omitempty"`
	Language        string    `json:"language,omitempty"`

	HasIssues bool   `json:"has_issues,omitempty"`
	HasWiki   bool   `json:"has_wiki,omitempty"`
	HasPage   bool   `json:"has_pages,omitempty"`
	License   string `json:"license,omitempty"`

	OpenIssuesCount int32  `json:"open_issues_count,omitempty"`
	DefaultBranch   string `json:"default_branch,omitempty"`
	Namespace       string `json:"namespace,omitempty"`

	NameWithNamespace string `json:"name_with_namespace,omitempty"`
	PathWithNamespace string `json:"path_with_namespace,omitempty"`
}

// BranchHook : branch information
type BranchHook struct {
	Label string       `json:"label,omitempty"`
	Ref   string       `json:"ref,omitempty"`
	Sha   string       `json:"sha,omitempty"`
	User  *UserHook    `json:"user,omitempty"`
	Repo  *ProjectHook `json:"repo,omitempty"`
}

// PullRequestHook : PR information
type PullRequestHook struct {
	Id                 int32          `json:"id,omitempty"`
	Number             int32          `json:"number,omitempty"`
	State              string         `json:"state,omitempty"`
	HtmlUrl            string         `json:"html_url,omitempty"`
	DiffUrl            string         `json:"diff_url,omitempty"`
	PatchUrl           string         `json:"patch_url,omitempty"`
	Title              string         `json:"title,omitempty"`
	Body               string         `json:"body,omitempty"`
	StaleLabels        []LabelHook    `json:"stale_labels,omitempty"`
	Labels             []LabelHook    `json:"labels,omitempty"`
	CreatedAt          string         `json:"created_at,omitempty"`
	UpdatedAt          string         `json:"updated_at,omitempty"`
	ClosedAt           string         `json:"closed_at,omitempty"`
	MergedAt           string         `json:"merged_at,omitempty"`
	MergeCommitSha     string         `json:"merge_commit_sha,omitempty"`
	MergeReferenceName string         `json:"merge_reference_name,omitempty"`
	User               *UserHook      `json:"user,omitempty"`
	Assignee           *UserHook      `json:"assignee,omitempty"`
	Assignees          []UserHook     `json:"assignees,omitempty"`
	Tester             []UserHook     `json:"tester,omitempty"`
	Testers            []UserHook     `json:"testers,omitempty"`
	NeedTest           bool           `json:"need_test,omitempty"`
	NeedReview         bool           `json:"need_review,omitempty"`
	Milestone          *MilestoneHook `json:"milestone,omitempty"`
	Head               *BranchHook    `json:"head,omitempty"`
	Base               *BranchHook    `json:"base,omitempty"`
	Merged             bool           `json:"merged,omitempty"`
	Mergeable          bool           `json:"mergeable,omitempty"`
	MergeStatus        string         `json:"merge_status,omitempty"`
	UpdatedBy          *UserHook      `json:"updated_by,omitempty"`
	Comments           int32          `json:"comments,omitempty"`
	Commits            int32          `json:"commits,omitempty"`
	Additions          int32          `json:"additions,omitempty"`
	Deletions          int32          `json:"deletions,omitempty"`
	ChangedFiles       int32          `json:"changed_files,omitempty"`
}
