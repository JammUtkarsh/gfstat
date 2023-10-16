package main

type AccountType int

const (
	User AccountType = iota
	Organization
)

// MetaFollow stands for metadata of a follower/following.
type MetaFollow struct {
	Username string `json:"username"`
	HTMLURL  string `json:"html_url"`
}

// GoProfile is more abstracted version of the UserProfile struct.
// It holds the followers and following raw data, which could be processed immediately
// instead of having to make multiple API calls.
// This struct follows created for the same reason as UserProfile,
// to simplify the data structure and is subject to change based on needs.
type CurrentUser struct {
	Username       string       `json:"username"`
	Type           string       `json:"type"` // User or Organization
	HTMLURL        string       `json:"html_url"`
	Followers      []MetaFollow `json:"followers"`
	FollowersCount int
	Following      []MetaFollow `json:"following"`
	FollowingCount int
}

func (c *CurrentUser) setMetadata(username, htmlURL, accountType string) {
	c.Username = username
	c.HTMLURL = htmlURL
	c.Type = accountType
}

// Mutuals gives the list of mutuals between followers and following.
// TC: O(nLogn)
// SC: O(1)
func (c CurrentUser) Mutuals() []MetaFollow {
	var mutuals []MetaFollow
	for _, follower := range c.Followers {
		low, high := 0, len(c.Following)-1
		for low <= high {
			mid := low + (high-low)/2
			if follower.Username == c.Following[mid].Username {
				mutuals = append(mutuals, follower)
				break
			} else if follower.Username < c.Following[mid].Username {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return mutuals
}

// followers - following
// TC: O(n)
// SC: O(N)
func (c CurrentUser) FollowersYouDontFollow() []MetaFollow {
	m := make(map[string]MetaFollow)
	for _, following := range c.Following {
		m[following.Username] = following
	}

	var iDontFollow []MetaFollow
	for _, follower := range c.Followers {
		if _, ok := m[follower.Username]; !ok {
			iDontFollow = append(iDontFollow, follower)
		}
	}
	return iDontFollow
}

// following - followers
// TC: O(n)
// SC: O(N)
func (c CurrentUser) FollowingYouDontFollow() []MetaFollow {
	m := make(map[string]MetaFollow)
	for _, follower := range c.Followers {
		if _, ok := m[follower.Username]; !ok {
			m[follower.Username] = follower
		}
	}

	var iDontFollow []MetaFollow
	for _, following := range c.Following {
		if _, ok := m[following.Username]; !ok {
			iDontFollow = append(iDontFollow, following)
		}
	}

	return iDontFollow
}
