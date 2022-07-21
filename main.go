package main

import (
	"log"
	"strings"

	"github.com/clubo-app/aggregator-service/config"
	authhandler "github.com/clubo-app/aggregator-service/handler/auth_handler"
	commenthandler "github.com/clubo-app/aggregator-service/handler/comment_handler"
	partyhandler "github.com/clubo-app/aggregator-service/handler/party_handler"
	profilehandler "github.com/clubo-app/aggregator-service/handler/profile_handler"
	relationhandler "github.com/clubo-app/aggregator-service/handler/relation_handler"
	storyhandler "github.com/clubo-app/aggregator-service/handler/story_handler"
	"github.com/clubo-app/packages/utils/middleware"
	ag "github.com/clubo-app/protobuf/auth"
	cg "github.com/clubo-app/protobuf/comment"
	pg "github.com/clubo-app/protobuf/party"
	prf "github.com/clubo-app/protobuf/profile"
	rg "github.com/clubo-app/protobuf/relation"
	sg "github.com/clubo-app/protobuf/story"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	prf, err := prf.NewClient(c.PROFILE_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to profile service: %v", err)
	}
	ac, err := ag.NewClient(c.AUTH_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to auth service: %v", err)
	}
	pc, err := pg.NewClient(c.PARTY_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to party service: %v", err)
	}
	sc, err := sg.NewClient(c.STORY_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to story service: %v", err)
	}
	rc, err := rg.NewClient(c.RELATION_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to relation service: %v", err)
	}
	cc, err := cg.NewClient(c.COMMENT_SERVICE_ADDRESS)
	if err != nil {
		log.Fatalf("did not connect to comment service: %v", err)
	}

	authHandler := authhandler.NewAuthGatewayHandler(ac, prf)
	profileHandler := profilehandler.NewUserGatewayHandler(prf, rc, ac)
	partyHandler := partyhandler.NewPartyGatewayHandler(pc, prf, sc, rc)
	storyHandler := storyhandler.NewStoryGatewayHandler(sc, prf)
	relationHandler := relationhandler.NewRelationGatewayHandler(rc, prf)
	commentHandler := commenthandler.NewCommentGatewayHandler(cc, prf)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Send custom error in json format
			return ctx.Status(code).JSON(err)
		},
	})
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())

	auth := app.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Post("/register", authHandler.Register)
	auth.Post("/google-login", authHandler.GoogleLogin)
	auth.Get("/refresh/:rt", authHandler.RefreshAccessToken)

	profile := app.Group("/profile")
	profile.Patch("/", middleware.AuthRequired(c.TOKEN_SECRET), profileHandler.UpdateUser)
	profile.Get("/me", middleware.AuthRequired(c.TOKEN_SECRET), profileHandler.GetMe)
	profile.Get("/:id", middleware.AuthOptional(c.TOKEN_SECRET), profileHandler.GetProfile)
	profile.Get("/username-taken/:username", profileHandler.UsernameTaken)

	party := app.Group("/party")
	party.Post("/", middleware.AuthRequired(c.TOKEN_SECRET), partyHandler.CreateParty)
	party.Delete("/:id", middleware.AuthRequired(c.TOKEN_SECRET), partyHandler.DeleteParty)
	party.Get("/:id", partyHandler.GetParty)
	party.Get("/user/:id", partyHandler.GetPartyByUser)
	party.Patch("/:id", middleware.AuthRequired(c.TOKEN_SECRET), partyHandler.UpdateParty)
	party.Get("/:id/favorite/user", partyHandler.GetFavorisingUsersByParty)
	party.Get("/search/:lat/:long", partyHandler.GeoSearch)

	party.Put("/favorite/:id", middleware.AuthRequired(c.TOKEN_SECRET), relationHandler.FavorParty)
	party.Delete("/favorite/:id", middleware.AuthRequired(c.TOKEN_SECRET), relationHandler.DefavorParty)
	party.Get("/favorite/user/:id", partyHandler.GetFavoritePartiesByUser)

	story := app.Group("/story")
	story.Post("/", middleware.AuthRequired(c.TOKEN_SECRET), storyHandler.CreateStory)
	story.Delete("/:id", storyHandler.DeleteStory)
	story.Get("/:id", storyHandler.GetStory)
	story.Get("/party/:id", storyHandler.GetStoryByParty)
	story.Get("/user/:id", storyHandler.GetStoryByUser)
	story.Get("/presign/:key", storyHandler.PresignURL)

	friend := app.Group("/friend")
	friend.Get("/:id", middleware.AuthOptional(c.TOKEN_SECRET), relationHandler.GetFriends)
	friend.Put("/request/:id", middleware.AuthRequired(c.TOKEN_SECRET), relationHandler.CreateFriendRequest)
	friend.Put("/accept/:id", middleware.AuthRequired(c.TOKEN_SECRET), relationHandler.AcceptFriendRequest)
	friend.Delete("/request/:id", middleware.AuthRequired(c.TOKEN_SECRET), relationHandler.DeclineFriendRequest)
	friend.Delete("/:id", middleware.AuthRequired(c.TOKEN_SECRET), relationHandler.RemoveFriend)

	comment := app.Group("/comment")
	comment.Post("/party/:id", middleware.AuthRequired(c.TOKEN_SECRET), commentHandler.CreateComment)
	comment.Get("/party/:id", commentHandler.GetCommentByParty)
	comment.Delete("/:id/party/:pId", middleware.AuthRequired(c.TOKEN_SECRET), commentHandler.DeleteComment)
	comment.Post("/:id/reply", middleware.AuthRequired(c.TOKEN_SECRET), commentHandler.CreateReply)
	comment.Get("/:id/reply", commentHandler.GetReplyByComment)
	comment.Delete("/:id/reply/:rId", middleware.AuthRequired(c.TOKEN_SECRET), commentHandler.DeleteReply)

	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(c.PORT)

	if err := app.Listen(sb.String()); err != nil {
		log.Fatal(err)
	}
}
