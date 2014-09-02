package gitwebhook

// Github push event payload as described in
// https://developer.github.com/v3/activity/events/types/#pushevent
type Push struct {
	Ref        string     `json:"ref"`
	After      string     `json:"after"`
	Before     string     `json:"before"`
	Created    bool       `json:"created"`
	Deleted    bool       `json:"deleted"`
	Forced     bool       `json:"forced"`
	Compare    string     `json:"compare"`
	Commits    []Commit   `json:"commits"`
	HeadCommit Commit     `json:"head_commit"`
	Repository Repository `json:"repository"`
	Pusher     Author     `json:"pusher"`
}

type Repository struct {
	Id               int64  `json:"id"`
	Name             string `json:"name"`
	FullName         string `json:"full_name"`
	Owner            Author `json:"owner"`
	Private          bool   `json:"private"`
	HtmlUrl          string `json:"html_url"`
	Description      string `json:"description"`
	Fork             bool   `json:"fork"`
	Url              string `json:"url"`
	ForksUrl         string `json:"forks_url"`
	KeysUrl          string `json:"keys_url"`
	CollaboratorsUrl string `json:"collaborators_url"`
	TeamsUrl         string `json:"teams_url"`
	HooksUrl         string `json:"hooks_url"`
	IssueEventsUrl   string `json:"issue_events_url"`
	EventsUrl        string `json:"events_url"`
	AssigneesUrl     string `json:"assignees_url"`
	BranchesUrl      string `json:"branches_url"`
	TagsUrl          string `json:"tags_url"`
	BlobsUrl         string `json:"blobs_url"`
	GitTagsUrl       string `json:"git_tags_url"`
	GitRefsUrl       string `json:"git_refs_url"`
	TreesUrl         string `json:"trees_url"`
	StatusesUrl      string `json:"statuses_url"`
	LanguagesUrl     string `json:"languages_url"`
	StargazersUrl    string `json:"stargazers_url"`
	ContributorsUrl  string `json:"contributors_url"`
	SubscribersUrl   string `json:"subscribers_url"`
	SubscriptionUrl  string `json:"subscription_url"`
	CommitsUrl       string `json:"commits_url"`
	GitCommitsUrl    string `json:"git_commits_url"`
	CommentsUrl      string `json:"comments_url"`
	IssueCommentUrl  string `json:"issue_comment_url"`
	ContentsUrl      string `json:"contents_url"`
	CompareUrl       string `json:"compare_url"`
	MergesUrl        string `json:"merges_url"`
	ArchiveUrl       string `json:"archive_url"`
	DownloadsUrl     string `json:"downloads_url"`
	IssuesUrl        string `json:"issues_url"`
	PullsUrl         string `json:"pulls_url"`
	MilstonesUrl     string `json:"milestones_url"`
	NotificationsUrl string `json:"notifications_url"`
	LabelsUrl        string `json:"labels_url"`
	ReleasesUrl      string `json:"releases_url"`
	Created          int64  `json:"created_at"`
	Updated          string `json:"updated_at"`
	Pushed           int64  `json:"pushed_at"`
	GitUrl           string `json:"git_url"`
	SshUrl           string `json:"ssh_url"`
	CloneUrl         string `json:"clone_url"`
	SvnUrl           string `json:"svn_url"`
	Homepage         string `json:"homepage"`
	Size             int64  `json:"size"`
	StargazersCount  int64  `json:"stargazers_count"`
	WatchersCount    int64  `json:"watchers_count"`
	Language         string `json:"language"`
	HasIssues        bool   `json:"has_issues"`
	HasDownloads     bool   `json:"has_downloads"`
	HasWiki          bool   `json:"has_wiki"`
	ForksCount       int64  `json:"forks_count"`
	MirrorUrl        string `json:"mirror_url"`
	OpenIssuesCount  int64  `json:"open_issues_count"`
	Forks            int64  `json:"forks"`
	OpenIssues       int64  `json:"open_issues"`
	Watchers         int64  `json:"watchers"`
	DefaultBranch    string `json:"default_branch"`
	Stargazers       int64  `json:"stargazers"`
	MasterBranch     string `json:"master_branch"`
}

type Commit struct {
	Id        string   `json:"id"`
	Distinct  bool     `json:"distinct"`
	Message   string   `json:"message"`
	Timestamp string   `json:"timestamp"`
	Url       string   `json:"url"`
	Author    Author   `json:"author"`
	Committer Author   `json:committer"`
	Added     []string `json:"added"`
	Removed   []string `json:"removed"`
	Modified  []string `json:"modified"`
}

type Author struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// Github ping event payload as describe in
// https://developer.github.com/webhooks/#ping-event
type Ping struct {
	Zen    string `json:"zen"`
	HookId int64  `json:"hook_id"`
	Hook   Hook   `json:"hook"`
}

type Hook struct {
	Url          string       `json:"url"`
	TestUrl      string       `json:"test_url"`
	Id           int          `json:"id"`
	Name         string       `json:"name"`
	Active       bool         `json:"active"`
	Events       []string     `json:"events"`
	Config       Config       `json:"config"`
	LastResopnse LastResponse `json:"last_response"`
	Updated      string       `json:"updated_at"`
	Created      string       `json:"created_at"`
}

type Config struct {
	Secret      string `json:"secret"`
	Url         string `json:"url"`
	ContentType string `json:"content_type"`
	InsecureSsl string `json:"insecure_ssl"`
}

type LastResponse struct {
	Code    string `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
