package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/desarso/goGraphql/graph/model"
)

func checkLastTimeSeenForDeletion(r *mutationResolver) {
	//print checking age
	// fmt.Println(1 * time.Minute)
	// //r.chessUsers[0].LastSeen into integer or number
	// thingy := r.chessUsers[0].LastSeen
	// thingys, err := strconv.Atoi(thingy)
	// fmt.Println(thingys)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(time.Now().UnixMilli())

	for {
		time.Sleep(10000)
		// fmt.Println("Checking age of users")

		for i, user := range r.chessUsers {
			//turn string into time
			lastSeen, err := strconv.Atoi(user.LastSeen)
			if err != nil {
				fmt.Println(err)
			}
			//if the user has not been seen in 1 minutes, delete them
			// fmt.Println(time.Now().UnixMilli() - int64(lastSeen))
			if time.Now().UnixMilli() > int64(lastSeen)+30000 {
				r.chessUsers = append(r.chessUsers[:i], r.chessUsers[i+1:]...)
				r.chessUsersChannel <- r.chessUsers
				//exit loop
				return
			}
		}
	}
}

// AddChessGame is the resolver for the addChessGame field.
func (r *mutationResolver) AddChessGame(ctx context.Context, fen string, gameID string, receiverID string, requesterID string, requesterColor string) (*model.ChessGame, error) {
	chessGame := &model.ChessGame{
		ID:            gameID,
		ReceiverID:    receiverID,
		ReceiverColor: requesterColor,
		RequesterID:   requesterID,
		Fen:           fen,
		Turn:          "white",
		Started:       false,
		Moves:         []*model.ChessMove{},
		Users:         []*model.ChessUser{},
	}
	r.chessGames = append(r.chessGames, chessGame)
	if r.chessGamesChannel == nil {
		r.chessGamesChannel = make(chan *model.ChessGame)
	}
	r.chessGamesChannel <- chessGame
	return chessGame, nil
}

// DeleteChessGame is the resolver for the deleteChessGame field.
func (r *mutationResolver) DeleteChessGame(ctx context.Context, id string) (*model.ChessGame, error) {
	//find the game
	for i, game := range r.chessGames {
		if game.ID == id {
			//delete the game
			r.chessGames = append(r.chessGames[:i], r.chessGames[i+1:]...)
			if r.chessGamesChannel == nil {
				r.chessGamesChannel = make(chan *model.ChessGame)
			}
			r.chessGamesChannel <- game
			return game, nil
		}
	}
	if r.chessGamesChannel == nil {
		r.chessGamesChannel = make(chan *model.ChessGame)
	}
	r.chessGamesChannel <- nil
	return nil, nil
}

// MutateChessGame is the resolver for the mutateChessGame field.
func (r *mutationResolver) MutateChessGame(ctx context.Context, id string, receiverID string, receiverColor string, requesterID string) (*model.ChessGame, error) {
	//find the game
	for _, game := range r.chessGames {
		if game.ID == id {
			//mutate the game
			game.ReceiverID = receiverID
			game.ReceiverColor = receiverColor
			game.RequesterID = requesterID
			if r.chessGamesChannel == nil {
				r.chessGamesChannel = make(chan *model.ChessGame)
			}
			r.chessGamesChannel <- game
			return game, nil
		}
	}
	if r.chessGamesChannel == nil {
		r.chessGamesChannel = make(chan *model.ChessGame)
	}
	r.chessGamesChannel <- nil
	return nil, nil
}

// MoveChessPiece is the resolver for the moveChessPiece field.
func (r *mutationResolver) MoveChessPiece(ctx context.Context, from string, to string, endFen string, gameID string, promotion *string) (*model.ChessGame, error) {
	//find the game
	for _, game := range r.chessGames {
		if game.ID == gameID {
			//add the move to the game
			chessMove := &model.ChessMove{
				From:      from,
				To:        to,
				Promotion: promotion,
				EndFen:    endFen,
			}
			game.Moves = append(game.Moves, chessMove)
			if r.chessGamesChannel == nil {
				r.chessGamesChannel = make(chan *model.ChessGame)
			}
			r.chessGamesChannel <- game
			return game, nil
		}
	}
	if r.chessGamesChannel == nil {
		r.chessGamesChannel = make(chan *model.ChessGame)
	}
	r.chessGamesChannel <- nil
	return nil, nil
}

// ChangeFen is the resolver for the changeFen field.
func (r *mutationResolver) ChangeFen(ctx context.Context, id string, fen string) (*model.ChessGame, error) {
	//find the game
	for _, game := range r.chessGames {
		if game.ID == id {
			//change the fen
			game.Fen = fen
			if r.chessGamesChannel == nil {
				r.chessGamesChannel = make(chan *model.ChessGame)
			}
			r.chessGamesChannel <- game
			return game, nil
		}
	}
	if r.chessGamesChannel == nil {
		r.chessGamesChannel = make(chan *model.ChessGame)
	}
	r.chessGamesChannel <- nil
	return nil, nil
}

// ChangeChessTurn is the resolver for the changeChessTurn field.
func (r *mutationResolver) ChangeChessTurn(ctx context.Context, id string, turn string) (*model.ChessGame, error) {
	//find the game
	for _, game := range r.chessGames {
		if game.ID == id {
			//change the turn
			game.Turn = turn
			if r.chessGamesChannel == nil {
				r.chessGamesChannel = make(chan *model.ChessGame)
			}
			r.chessGamesChannel <- game
			return game, nil
		}
	}
	if r.chessGamesChannel == nil {
		r.chessGamesChannel = make(chan *model.ChessGame)
	}
	r.chessGamesChannel <- nil
	return nil, nil
}

// StartChessGame is the resolver for the startChessGame field.
func (r *mutationResolver) StartChessGame(ctx context.Context, gameID string) (*model.ChessGame, error) {
	//find the game
	for _, game := range r.chessGames {
		if game.ID == gameID {
			//start the game
			game.Started = true
			if r.chessGamesChannel == nil {
				r.chessGamesChannel = make(chan *model.ChessGame)
			}
			r.chessGamesChannel <- game
			return game, nil
		}
	}
	if r.chessGamesChannel == nil {
		r.chessGamesChannel = make(chan *model.ChessGame)
	}
	r.chessGamesChannel <- nil
	return nil, nil
}

// AddChessUser is the resolver for the addChessUser field.
func (r *mutationResolver) AddChessUser(ctx context.Context, id string, username string, catURL string) (*model.ChessUser, error) {
	time := time.Now().UnixMilli()
	//tunr time to a string
	var timeString = fmt.Sprintf("%d", time)
	chessUser := &model.ChessUser{
		ID:            id,
		Username:      username,
		LastSeen:      timeString,
		CatURL:        catURL,
		ChessRequests: []*model.ChessRequest{},
	}
	go checkLastTimeSeenForDeletion(r)
	//make sure user does not exist already
	for _, user := range r.chessUsers {
		if user.ID == id {
			if r.chessUsersChannel == nil {
				r.chessUsersChannel = make(chan []*model.ChessUser)
			}
			r.chessUsersChannel <- r.chessUsers
			return nil, nil
		}
	}
	r.chessUsers = append(r.chessUsers, chessUser)
	//check thant channel has been initialized
	if r.chessUsersChannel == nil {
		r.chessUsersChannel = make(chan []*model.ChessUser)
	}
	r.chessUsersChannel <- r.chessUsers
	//return the user
	return chessUser, nil
}

// DeleteChessUser is the resolver for the deleteChessUser field.
func (r *mutationResolver) DeleteChessUser(ctx context.Context, id string) (*model.ChessUser, error) {
	//find the user
	for i, user := range r.chessUsers {
		if user.ID == id {
			//delete the user
			r.chessUsers = append(r.chessUsers[:i], r.chessUsers[i+1:]...)
			if r.chessUsersChannel == nil {
				r.chessUsersChannel = make(chan []*model.ChessUser)
			}
			r.chessUsersChannel <- r.chessUsers
			return user, nil
		}
	}
	if r.chessUsersChannel == nil {
		r.chessUsersChannel = make(chan []*model.ChessUser)
	}
	r.chessUsersChannel <- r.chessUsers
	return nil, nil
}

// MutateChessUser is the resolver for the mutateChessUser field.
func (r *mutationResolver) MutateChessUser(ctx context.Context, id string, username string, catURL string) (*model.ChessUser, error) {
	//find the user
	for _, user := range r.chessUsers {
		if user.ID == id {
			//mutate the user
			user.Username = username
			user.CatURL = catURL
			if r.chessUsersChannel == nil {
				r.chessUsersChannel = make(chan []*model.ChessUser)
			}
			r.chessUsersChannel <- r.chessUsers
			return user, nil
		}
	}
	if r.chessUsersChannel == nil {
		r.chessUsersChannel = make(chan []*model.ChessUser)
	}
	r.chessUsersChannel <- r.chessUsers
	return nil, nil
}

// UpdateLastSeenChess is the resolver for the updateLastSeenChess field.
func (r *mutationResolver) UpdateLastSeenChess(ctx context.Context, id string) (*model.ChessUser, error) {
	//find the user
	for _, user := range r.chessUsers {
		if user.ID == id {
			//update the last seen
			time := time.Now().UnixMilli()
			//tunr time to a string
			var timeString = fmt.Sprintf("%d", time)
			user.LastSeen = timeString
			if r.chessUsersChannel == nil {
				r.chessUsersChannel = make(chan []*model.ChessUser)
			}
			r.chessUsersChannel <- r.chessUsers
			return user, nil
		}
	}
	if r.chessUsersChannel == nil {
		r.chessUsersChannel = make(chan []*model.ChessUser)
	}
	r.chessUsersChannel <- r.chessUsers
	return nil, nil
}

// SendChessRequest is the resolver for the sendChessRequest field.
func (r *mutationResolver) SendChessRequest(ctx context.Context, gameID string, requesterID string, requesterColor string, receiverID string) (*model.ChessUser, error) {
	//find the user
	for _, user := range r.chessUsers {
		if user.ID == receiverID {
			//add the request to the user
			chessRequest := &model.ChessRequest{
				GameID:         gameID,
				RequesterID:    requesterID,
				RequesterColor: requesterColor,
				ReceiverID:     receiverID,
			}
			user.ChessRequests = append(user.ChessRequests, chessRequest)
			if r.chessRequestsChannel == nil {
				r.chessRequestsChannel = make(chan *model.ChessRequest)
			}
			r.chessRequestsChannel <- chessRequest
			return user, nil
		}
	}
	if r.chessRequestsChannel == nil {
		r.chessRequestsChannel = make(chan *model.ChessRequest)
	}
	r.chessRequestsChannel <- nil
	return nil, nil
}

// GetChessGames is the resolver for the getChessGames field.
func (r *queryResolver) GetChessGames(ctx context.Context) ([]*model.ChessGame, error) {
	return r.chessGames, nil
}

// GetChessGame is the resolver for the getChessGame field.
func (r *queryResolver) GetChessGame(ctx context.Context, id string) (*model.ChessGame, error) {
	//find the game
	for _, game := range r.chessGames {
		if game.ID == id {
			return game, nil
		}
	}
	return nil, nil
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.ChessUser, error) {
	return r.chessUsers, nil
}

// ChessGamesSub is the resolver for the chessGamesSub field.
func (r *subscriptionResolver) ChessGamesSub(ctx context.Context, id string) (<-chan *model.ChessGame, error) {

	if r.chessGamesChannel == nil {
		r.chessGamesChannel = make(chan *model.ChessGame)
	}

	go func() {
		for {
			//find the game
			var foundGame *model.ChessGame
			for _, game := range r.chessGames {
				if game.ID == id {

					foundGame = game
				}
			}
			select {
			case r.chessGamesChannel <- foundGame:

			default:
				fmt.Println("Channel closed")

				return
			}
		}
	}()
	return r.chessGamesChannel, nil
}

// ChessUsersSub is the resolver for the chessUsersSub field.
func (r *subscriptionResolver) ChessUsersSub(ctx context.Context) (<-chan []*model.ChessUser, error) {
	if r.chessUsersChannel == nil {
		r.chessUsersChannel = make(chan []*model.ChessUser)
	}
	go func() {
		for {
			select {
			case r.chessUsersChannel <- r.chessUsers:

			default:
				fmt.Println("Channel closed")

				return
			}
		}
	}()
	return r.chessUsersChannel, nil
}

// ChessRequestsSub is the resolver for the chessRequestsSub field.
func (r *subscriptionResolver) ChessRequestsSub(ctx context.Context) (<-chan *model.ChessRequest, error) {
	if r.chessRequestsChannel == nil {
		r.chessRequestsChannel = make(chan *model.ChessRequest)
	}
	go func() {
		for {
			select {
			case chessRequest := <-r.chessRequestsChannel:
				fmt.Println("Chess Request: ", chessRequest)

			default:
				fmt.Println("Channel closed")

				return
			}
		}
	}()
	return r.chessRequestsChannel, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
