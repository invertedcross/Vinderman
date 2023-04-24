package vinderman

import (
	"fmt"
	"net/http"
	"time"

	"github.com/0xDistrust/Vinderman/consts"
	"github.com/0xDistrust/Vinderman/request"
)

type Friend struct {
	AccountId string        `json:"accountId"`
	Groups    []interface{} `json:"groups"`
	Mutual    int           `json:"mutual"`
	Alias     string        `json:"alias"`
	Note      string        `json:"note"`
	Favorite  bool          `json:"favorite"`
	Created   time.Time     `json:"created"`
}

type PendingFriend struct {
	AccountId string    `json:"accountId"`
	Mutual    int       `json:"mutual"`
	Favorite  bool      `json:"favorite"`
	Created   time.Time `json:"created"`
}

type SuggestedFriend struct {
	AccountId   string `json:"accountId"`
	Connections struct {
		Epic struct {
			Id          string `json:"id"`
			SortFactors struct {
				X int       `json:"x"`
				Y int       `json:"y"`
				K time.Time `json:"k"`
				L time.Time `json:"l"`
			} `json:"sortFactors"`
		} `json:"epic"`
	} `json:"connections"`
	Mutual int `json:"mutual"`
}

type FriendsSummary struct {
	Friends []Friend `json:"friends"`
	Incoming []PendingFriend `json:"incoming"`
	Outgoing []PendingFriend `json:"outgoing"`
	Suggested []SuggestedFriend `json:"suggested"`
	Blocklist []interface{} `json:"blocklist"`
	Settings  struct {
		AcceptInvites string `json:"acceptInvites"`
		MutualPrivacy string `json:"mutualPrivacy"`
	} `json:"settings"`
	LimitsReached struct {
		Incoming bool `json:"incoming"`
		Outgoing bool `json:"outgoing"`
		Accepted bool `json:"accepted"`
	} `json:"limitsReached"`
}

func (c Client) AddFriend(credentials UserCredentials, friendID string) (err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("POST", fmt.Sprintf("%s/friends/api/v1/%s/friends/%s", consts.FRIENDS_SERVICE, credentials.AccountID, friendID), headers, "")
	if err != nil {
		return
	}

	if resp.StatusCode >= 300 {
		return fmt.Errorf("failed to add friend: %s", resp.Status)
	}

	return
}

func (c Client) FetchFriend(credentials UserCredentials, friendID string) (friendship Friend, err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("GET", fmt.Sprintf("%s/friends/api/v1/%s/friends/%s", consts.FRIENDS_SERVICE, credentials.AccountID, friendID), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[Friend](resp)

	return res.Body, err
}

func (c Client) FetchFriends(credentials UserCredentials) (friendship []Friend, err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("GET", fmt.Sprintf("%s/friends/api/v1/%s/friends", consts.FRIENDS_SERVICE, credentials.AccountID), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[[]Friend](resp)

	return res.Body, err
}

func (c Client) RemoveFriend(credentials UserCredentials, friendID string) (err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("DELETE", fmt.Sprintf("%s/friends/api/v1/%s/friends/%s", consts.FRIENDS_SERVICE, credentials.AccountID, friendID), headers, "")
	if err != nil {
		return
	}

	if resp.StatusCode >= 300 {
		return fmt.Errorf("failed to remove friend: %s", resp.Status)
	}

	return
}

func (c Client) RemoveFriendsBulk(credentials UserCredentials) (err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("DELETE", fmt.Sprintf("%s/friends/api/v1/%s/friends", consts.FRIENDS_SERVICE, credentials.AccountID), headers, "")
	if err != nil {
		return
	}

	if resp.StatusCode >= 300 {
		return fmt.Errorf("failed to remove friends: %s", resp.Status)
	}

	return
}

func (c Client) SetFriendNickname(credentials UserCredentials, friendID string, nickname string) (err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("PUT", fmt.Sprintf("%s/friends/api/v1/%s/friends/%s/alias", consts.FRIENDS_SERVICE, credentials.AccountID, friendID), headers, nickname)
	if err != nil {
		return
	}

	if resp.StatusCode >= 300 {
		return fmt.Errorf("failed to set friend nickname: %s", resp.Status)
	}

	return
}

func (c Client) RemoveFriendNickname(credentials UserCredentials, friendID string) (err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("DELETE", fmt.Sprintf("%s/friends/api/v1/%s/friends/%s/alias", consts.FRIENDS_SERVICE, credentials.AccountID, friendID), headers, "")
	if err != nil {
		return
	}

	if resp.StatusCode >= 300 {
		return fmt.Errorf("failed to remove friend nickname: %s", resp.Status)
	}

	return
}

func (c Client) FetchFriendsSummary(credentials UserCredentials) (friendship FriendsSummary, err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("GET", fmt.Sprintf("%s/friends/api/v1/%s/summary", consts.FRIENDS_SERVICE, credentials.AccountID), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[FriendsSummary](resp)

	return res.Body, err
}

func (c Client) FetchFriendsIncoming(credentials UserCredentials) (friendship []Friend, err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("GET", fmt.Sprintf("%s/friends/api/v1/%s/incoming", consts.FRIENDS_SERVICE, credentials.AccountID), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[[]Friend](resp)

	return res.Body, err
}

func (c Client) FetchFriendsOutgoing(credentials UserCredentials) (friendship []Friend, err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("GET", fmt.Sprintf("%s/friends/api/v1/%s/outgoing", consts.FRIENDS_SERVICE, credentials.AccountID), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[[]Friend](resp)

	return res.Body, err
}

func (c Client) FetchFriendsSuggested(credentials UserCredentials) (friendship []SuggestedFriend, err error) {
	headers := http.Header{}
	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))

	resp, err := c.Request("GET", fmt.Sprintf("%s/friends/api/v1/%s/suggested", consts.FRIENDS_SERVICE, credentials.AccountID), headers, "")
	if err != nil {
		return
	}

	res, err := request.ResponseParser[[]SuggestedFriend](resp)

	return res.Body, err
}