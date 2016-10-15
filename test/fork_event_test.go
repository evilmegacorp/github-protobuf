package github

import (
	"testing"
	"encoding/json"
)

func TestForkEvent(t *testing.T) {
	event := `{
	  "forkee": {
	    "id": 35129393,
	    "name": "public-repo",
	    "full_name": "baxterandthehackers/public-repo",
	    "owner": {
	      "login": "baxterandthehackers",
	      "id": 7649605,
	      "avatar_url": "https://avatars.githubusercontent.com/u/7649605?v=3",
	      "gravatar_id": "",
	      "url": "https://api.github.com/users/baxterandthehackers",
	      "html_url": "https://github.com/baxterandthehackers",
	      "followers_url": "https://api.github.com/users/baxterandthehackers/followers",
	      "following_url": "https://api.github.com/users/baxterandthehackers/following{/other_user}",
	      "gists_url": "https://api.github.com/users/baxterandthehackers/gists{/gist_id}",
	      "starred_url": "https://api.github.com/users/baxterandthehackers/starred{/owner}{/repo}",
	      "subscriptions_url": "https://api.github.com/users/baxterandthehackers/subscriptions",
	      "organizations_url": "https://api.github.com/users/baxterandthehackers/orgs",
	      "repos_url": "https://api.github.com/users/baxterandthehackers/repos",
	      "events_url": "https://api.github.com/users/baxterandthehackers/events{/privacy}",
	      "received_events_url": "https://api.github.com/users/baxterandthehackers/received_events",
	      "type": "Organization",
	      "site_admin": false
	    },
	    "private": false,
	    "html_url": "https://github.com/baxterandthehackers/public-repo",
	    "description": "",
	    "fork": true,
	    "url": "https://api.github.com/repos/baxterandthehackers/public-repo",
	    "forks_url": "https://api.github.com/repos/baxterandthehackers/public-repo/forks",
	    "keys_url": "https://api.github.com/repos/baxterandthehackers/public-repo/keys{/key_id}",
	    "collaborators_url": "https://api.github.com/repos/baxterandthehackers/public-repo/collaborators{/collaborator}",
	    "teams_url": "https://api.github.com/repos/baxterandthehackers/public-repo/teams",
	    "hooks_url": "https://api.github.com/repos/baxterandthehackers/public-repo/hooks",
	    "issue_events_url": "https://api.github.com/repos/baxterandthehackers/public-repo/issues/events{/number}",
	    "events_url": "https://api.github.com/repos/baxterandthehackers/public-repo/events",
	    "assignees_url": "https://api.github.com/repos/baxterandthehackers/public-repo/assignees{/user}",
	    "branches_url": "https://api.github.com/repos/baxterandthehackers/public-repo/branches{/branch}",
	    "tags_url": "https://api.github.com/repos/baxterandthehackers/public-repo/tags",
	    "blobs_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/blobs{/sha}",
	    "git_tags_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/tags{/sha}",
	    "git_refs_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/refs{/sha}",
	    "trees_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/trees{/sha}",
	    "statuses_url": "https://api.github.com/repos/baxterandthehackers/public-repo/statuses/{sha}",
	    "languages_url": "https://api.github.com/repos/baxterandthehackers/public-repo/languages",
	    "stargazers_url": "https://api.github.com/repos/baxterandthehackers/public-repo/stargazers",
	    "contributors_url": "https://api.github.com/repos/baxterandthehackers/public-repo/contributors",
	    "subscribers_url": "https://api.github.com/repos/baxterandthehackers/public-repo/subscribers",
	    "subscription_url": "https://api.github.com/repos/baxterandthehackers/public-repo/subscription",
	    "commits_url": "https://api.github.com/repos/baxterandthehackers/public-repo/commits{/sha}",
	    "git_commits_url": "https://api.github.com/repos/baxterandthehackers/public-repo/git/commits{/sha}",
	    "comments_url": "https://api.github.com/repos/baxterandthehackers/public-repo/comments{/number}",
	    "issue_comment_url": "https://api.github.com/repos/baxterandthehackers/public-repo/issues/comments{/number}",
	    "contents_url": "https://api.github.com/repos/baxterandthehackers/public-repo/contents/{+path}",
	    "compare_url": "https://api.github.com/repos/baxterandthehackers/public-repo/compare/{base}...{head}",
	    "merges_url": "https://api.github.com/repos/baxterandthehackers/public-repo/merges",
	    "archive_url": "https://api.github.com/repos/baxterandthehackers/public-repo/{archive_format}{/ref}",
	    "downloads_url": "https://api.github.com/repos/baxterandthehackers/public-repo/downloads",
	    "issues_url": "https://api.github.com/repos/baxterandthehackers/public-repo/issues{/number}",
	    "pulls_url": "https://api.github.com/repos/baxterandthehackers/public-repo/pulls{/number}",
	    "milestones_url": "https://api.github.com/repos/baxterandthehackers/public-repo/milestones{/number}",
	    "notifications_url": "https://api.github.com/repos/baxterandthehackers/public-repo/notifications{?since,all,participating}",
	    "labels_url": "https://api.github.com/repos/baxterandthehackers/public-repo/labels{/name}",
	    "releases_url": "https://api.github.com/repos/baxterandthehackers/public-repo/releases{/id}",
	    "created_at": "2015-05-05T23:40:30Z",
	    "updated_at": "2015-05-05T23:40:30Z",
	    "pushed_at": "2015-05-05T23:40:27Z",
	    "git_url": "git://github.com/baxterandthehackers/public-repo.git",
	    "ssh_url": "git@github.com:baxterandthehackers/public-repo.git",
	    "clone_url": "https://github.com/baxterandthehackers/public-repo.git",
	    "svn_url": "https://github.com/baxterandthehackers/public-repo",
	    "homepage": null,
	    "size": 0,
	    "stargazers_count": 0,
	    "watchers_count": 0,
	    "language": null,
	    "has_issues": false,
	    "has_downloads": true,
	    "has_wiki": true,
	    "has_pages": true,
	    "forks_count": 0,
	    "mirror_url": null,
	    "open_issues_count": 0,
	    "forks": 0,
	    "open_issues": 0,
	    "watchers": 0,
	    "default_branch": "master",
	    "public": true
	  },
	  "repository": {
	    "id": 35129377,
	    "name": "public-repo",
	    "full_name": "baxterthehacker/public-repo",
	    "owner": {
	      "login": "baxterthehacker",
	      "id": 6752317,
	      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
	      "gravatar_id": "",
	      "url": "https://api.github.com/users/baxterthehacker",
	      "html_url": "https://github.com/baxterthehacker",
	      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
	      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
	      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
	      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
	      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
	      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
	      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
	      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
	      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
	      "type": "User",
	      "site_admin": false
	    },
	    "private": false,
	    "html_url": "https://github.com/baxterthehacker/public-repo",
	    "description": "",
	    "fork": false,
	    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
	    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
	    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
	    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
	    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
	    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
	    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
	    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
	    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
	    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
	    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
	    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
	    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
	    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
	    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
	    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
	    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
	    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
	    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
	    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
	    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
	    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
	    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
	    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
	    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
	    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
	    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
	    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
	    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
	    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
	    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
	    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
	    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
	    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
	    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
	    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
	    "created_at": "2015-05-05T23:40:12Z",
	    "updated_at": "2015-05-05T23:40:30Z",
	    "pushed_at": "2015-05-05T23:40:27Z",
	    "git_url": "git://github.com/baxterthehacker/public-repo.git",
	    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
	    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
	    "svn_url": "https://github.com/baxterthehacker/public-repo",
	    "homepage": null,
	    "size": 0,
	    "stargazers_count": 0,
	    "watchers_count": 0,
	    "language": null,
	    "has_issues": true,
	    "has_downloads": true,
	    "has_wiki": true,
	    "has_pages": true,
	    "forks_count": 1,
	    "mirror_url": null,
	    "open_issues_count": 2,
	    "forks": 1,
	    "open_issues": 2,
	    "watchers": 0,
	    "default_branch": "master"
	  },
	  "sender": {
	    "login": "baxterandthehackers",
	    "id": 7649605,
	    "avatar_url": "https://avatars.githubusercontent.com/u/7649605?v=3",
	    "gravatar_id": "",
	    "url": "https://api.github.com/users/baxterandthehackers",
	    "html_url": "https://github.com/baxterandthehackers",
	    "followers_url": "https://api.github.com/users/baxterandthehackers/followers",
	    "following_url": "https://api.github.com/users/baxterandthehackers/following{/other_user}",
	    "gists_url": "https://api.github.com/users/baxterandthehackers/gists{/gist_id}",
	    "starred_url": "https://api.github.com/users/baxterandthehackers/starred{/owner}{/repo}",
	    "subscriptions_url": "https://api.github.com/users/baxterandthehackers/subscriptions",
	    "organizations_url": "https://api.github.com/users/baxterandthehackers/orgs",
	    "repos_url": "https://api.github.com/users/baxterandthehackers/repos",
	    "events_url": "https://api.github.com/users/baxterandthehackers/events{/privacy}",
	    "received_events_url": "https://api.github.com/users/baxterandthehackers/received_events",
	    "type": "Organization",
	    "site_admin": false
	  }
	}`
	var forkEvent ForkEvent
	err := json.Unmarshal([]byte(event), &forkEvent)
	if err != nil {
		t.Error(err)
	} else {
		if forkEvent.Forkee.Id != 35129393 {
			t.Error("forkEvent.Forkee.Id != 35129393")
		}
		if forkEvent.Forkee.FullName != "baxterandthehackers/public-repo" {
			t.Error("forkEvent.Forkee.FullName != baxterandthehackers/public-repo")
		}
		repo := forkEvent.Repository
		if repo.Id != 35129377 {
			t.Error("forkEvent.Repository.Id != 35129377")
		}
		if repo.Owner.Login != "baxterthehacker" {
			t.Error("forkEvent.Repository.Owner.Login != baxterthehacker")
		}
		sender := forkEvent.Sender
		if sender.Login != "baxterandthehackers" {
			t.Error("forkEvent.Sender.Login != baxterandthehackers")
		}
	}
}
