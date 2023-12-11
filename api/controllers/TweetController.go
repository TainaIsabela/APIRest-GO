package controllers

import (
	entities "api/api/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *TweetController {
	tweetController := TweetController{}

	return &tweetController
}

func (tweetController *TweetController) FindAll (ctx *gin.Context) {
	ctx.JSON(http.StatusOK, tweetController.tweets)
}

func (tweetController *TweetController) Create (ctx *gin.Context) {
	tweet := entities.NewTweet()
	ctx.BindJSON(&tweet)
	tweetController.tweets = append(tweetController.tweets, *tweet)
	ctx.JSON(http.StatusOK, tweet)
}

func (tweetController *TweetController) Update (ctx *gin.Context) {
	id := ctx.Param("id")
	for i, tweet := range tweetController.tweets {
		if tweet.ID == id {
			ctx.BindJSON(&tweet)
			tweetController.tweets[i] = tweet
			ctx.JSON(http.StatusOK, tweet)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Tweet not found"})
}

func (tweetController *TweetController) Delete (ctx *gin.Context) {
	id := ctx.Param("id")
	for i, tweet := range tweetController.tweets {
		if tweet.ID == id {
			tweetController.tweets = append(tweetController.tweets[:i], tweetController.tweets[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Tweet deleted successfully"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Tweet not found"})
}

func (tweetController *TweetController) FindById (ctx *gin.Context) {
	id := ctx.Param("id")
	for _, tweet := range tweetController.tweets {
		if tweet.ID == id {
			ctx.JSON(http.StatusOK, tweet)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Tweet not found"})
}
