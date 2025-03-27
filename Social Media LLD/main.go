package main

import (
	"fmt"
	"strings"
)

type User struct {
	Username  string
	Privacy   string
	Tweets    []Tweet
	Followers map[string]bool
	Following map[string]bool
}

type Tweet struct {
	User    string
	Message string
}

var users = make(map[string]*User)

func createUser(username, privacy string) *User {

	if privacy != "public" && privacy != "private" {
		fmt.Println("Invaild Privacy Setting. ")
		return nil
	}

	user := &User{
		Username:  username,
		Privacy:   privacy,
		Tweets:    []Tweet{},
		Followers: make(map[string]bool),
		Following: make(map[string]bool),
	}

	users[username] = user
	return user
}

func postTweet(username, message string) {
	user, exists := users[username]
	if !exists {
		fmt.Println("User is not found!")
		return
	}

	if len(message) > 280 {
		fmt.Println("Tweet is more than 280 characters!")
		return
	}

	tweet := Tweet{
		User:    username,
		Message: message,
	}

	user.Tweets = append(user.Tweets, tweet)
	fmt.Println(username + " posted a tweet: " + message)
}

func followUser(follower, followee string) {
	followeeUser, exists := users[followee]
	if !exists {
		fmt.Println("Followee is not found!")
		return
	}

	followerUser, exists := users[follower]
	if !exists {
		fmt.Println("Follower is not found!")
		return
	}

	if followeeUser.Privacy == "public" {
		followeeUser.Followers[follower] = true
		followerUser.Following[followee] = true
		fmt.Println(follower + " followed " + followee)
	} else {
		followerUser.Following[followee] = false
		fmt.Println(follower + " sent a request to follow " + followee)
	}
}

func approveFollowRequest(follower, followee string) {
	followeeUser, exists := users[followee]
	if !exists {
		fmt.Println("Followee is not found!")
		return
	}

	followerUser, exists := users[follower]
	if !exists {
		fmt.Println("Followee is not found!")
		return
	}

	if followeeUser.Privacy == "private" && followeeUser.Following[follower] == false {
		followeeUser.Followers[follower] = true
		followerUser.Followers[followee] = true
		fmt.Println(followee + " approves the follow request " + follower)
	} else {
		fmt.Println("There is no request from " + follower + " to " + followee)
	}
}

func displayTweets(username string) {
	user, exists := users[username]
	if !exists {
		fmt.Println("User is not found!")
		return
	}

	fmt.Println(username + "'s Tweets:")
	for _, tweet := range user.Tweets {
		if user.Privacy == "public" {
			fmt.Println("- " + tweet.Message)
		} else {
			fmt.Println("- (Private Tweet)")
		}
	}

}

func searchTweets(username, searchWord string) {
	user, exists := users[username]
	if !exists {
		fmt.Println("User is not found!")
		return
	}

	count := 0
	fmt.Println("Sreach for Tweet with: " + searchWord)
	for _, tweet := range user.Tweets {
		if strings.Contains(tweet.Message, searchWord) {
			fmt.Println("- " + tweet.Message)
			count++
		}

		if count == 10 {
			break
		}
		if count == 0 {
			fmt.Println("No Tweets found!")
		}
	}

}

func main() {

	createUser("Bob", "private")
	createUser("Bill", "public")

	postTweet("Bob", "Hi, I am Bob.")
	postTweet("Bob", "Hi Again.")

	postTweet("Bill", "Hi, I am Bill.")
	postTweet("Bill", "Who is Bob? and Hi")

	followUser("Bill", "Bob")
	approveFollowRequest("Bill", "Bob")

	displayTweets("Bob")
	displayTweets("Bill")

	searchTweets("Bill", "Hi")
}
