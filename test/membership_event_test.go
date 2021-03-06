package github

import (
	"testing"
	"encoding/json"
)

func TestMembershipEvent(t *testing.T) {
	event := `{
	  "action": "added",
	  "scope": "team",
	  "member": {
	    "login": "kdaigle",
	    "id": 2501,
	    "avatar_url": "https://avatars.githubusercontent.com/u/2501?v=3",
	    "gravatar_id": "",
	    "url": "https://api.github.com/users/kdaigle",
	    "html_url": "https://github.com/kdaigle",
	    "followers_url": "https://api.github.com/users/kdaigle/followers",
	    "following_url": "https://api.github.com/users/kdaigle/following{/other_user}",
	    "gists_url": "https://api.github.com/users/kdaigle/gists{/gist_id}",
	    "starred_url": "https://api.github.com/users/kdaigle/starred{/owner}{/repo}",
	    "subscriptions_url": "https://api.github.com/users/kdaigle/subscriptions",
	    "organizations_url": "https://api.github.com/users/kdaigle/orgs",
	    "repos_url": "https://api.github.com/users/kdaigle/repos",
	    "events_url": "https://api.github.com/users/kdaigle/events{/privacy}",
	    "received_events_url": "https://api.github.com/users/kdaigle/received_events",
	    "type": "User",
	    "site_admin": true
	  },
	  "sender": {
	    "login": "baxterthehacker",
	    "id": 6752317,
	    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=2",
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
	  "team": {
	    "name": "Contractors",
	    "id": 123456,
	    "slug": "contractors",
	    "permission": "admin",
	    "url": "https://api.github.com/teams/123456",
	    "members_url": "https://api.github.com/teams/123456/members{/member}",
	    "repositories_url": "https://api.github.com/teams/123456/repos"
	  },
	  "organization": {
	    "login": "baxterandthehackers",
	    "id": 7649605,
	    "url": "https://api.github.com/orgs/baxterandthehackers",
	    "repos_url": "https://api.github.com/orgs/baxterandthehackers/repos",
	    "events_url": "https://api.github.com/orgs/baxterandthehackers/events",
	    "members_url": "https://api.github.com/orgs/baxterandthehackers/members{/member}",
	    "public_members_url": "https://api.github.com/orgs/baxterandthehackers/public_members{/member}",
	    "avatar_url": "https://avatars.githubusercontent.com/u/7649605?v=2"
	  }
	}`
	var membershipEvent MembershipEvent
	err := json.Unmarshal([]byte(event), &membershipEvent)
	if err != nil {
		t.Error(err)
	} else {
		if membershipEvent.Action != "added" {
			t.Error("membershipEvent.Action != added")
		}
		if membershipEvent.Member.Login != "kdaigle" {
			t.Error("membershipEvent.Member.Login != kdaigle")
		}
		if membershipEvent.Team.Name != "Contractors" {
			t.Error("membershipEvent.Team.Name != Contractors")
		}
		if membershipEvent.Organization.Id != 7649605 {
			t.Error("membershipEvent.Organization.Id != 7649605")
		}
		sender := membershipEvent.Sender
		if sender.Login != "baxterthehacker" {
			t.Error("membershipEvent.Sender.Login != baxterthehacker")
		}
	}
}
