package keeper

import (
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) RemoveFromFifo(ctx sdk.Context, game *types.StoredGame, nextGame *types.NextGame) {
	// Does it have a predecessor?
	if game.BeforeId != types.NoFifoIdKey {
		predecessorGame, found := k.GetStoredGame(ctx, game.BeforeId)
		if !found {
			panic("Game before in FIFO was not found")
		}
		predecessorGame.AfterId = game.AfterId
		k.SetStoredGame(ctx, predecessorGame)
		if game.AfterId == types.NoFifoIdKey {
			nextGame.FifoTail = predecessorGame.Index
		}
		// Is it at the FIFO head?
	} else if nextGame.FifoHead == game.Index {
		nextGame.FifoHead = game.AfterId
	} // Else it is the first game in the fifo
	// Does it have a successor?
	if game.AfterId != types.NoFifoIdKey {
		successorGame, found := k.GetStoredGame(ctx, game.AfterId)
		if !found {
			panic("Game after in FIFO was not found")
		}
		successorGame.BeforeId = game.BeforeId
		k.SetStoredGame(ctx, successorGame)
		if game.BeforeId == types.NoFifoIdKey {
			nextGame.FifoHead = successorGame.Index
		}
		// Is it at the FIFO tail?
	} else if nextGame.FifoTail == game.Index {
		nextGame.FifoTail = game.BeforeId
	} // Else it is the first game in the fifo
	// these changes to game, as well as previous changes to nextGame, are not saved to
	// memory, that should be done after this function call
	game.BeforeId = types.NoFifoIdKey
	game.AfterId = types.NoFifoIdKey
}

func (k Keeper) SendToFifoTail(ctx sdk.Context, game *types.StoredGame, nextGame *types.NextGame) {
	if nextGame.FifoHead == types.NoFifoIdKey && nextGame.FifoTail == types.NoFifoIdKey {
		// If this is the first game in the FIFO list
		game.BeforeId = types.NoFifoIdKey
		game.AfterId = types.NoFifoIdKey
		nextGame.FifoHead = game.Index
		nextGame.FifoTail = game.Index
	} else if nextGame.FifoHead == types.NoFifoIdKey || nextGame.FifoTail == types.NoFifoIdKey {
		panic("Next Game should have both FIFO head and FIFO tail or none")
	} else if nextGame.FifoTail == game.Index {
		// Nothing to do, already at tail
	} else {
		// Snip game out
		k.RemoveFromFifo(ctx, game, nextGame)

		// Now add to tail
		currentTail, found := k.GetStoredGame(ctx, nextGame.FifoTail)
		if !found {
			panic("Current FIFO tail was not found")
		}
		currentTail.AfterId = game.Index
		k.SetStoredGame(ctx, currentTail)

		game.BeforeId = currentTail.Index
		nextGame.FifoTail = game.Index
	}
	// these changes to game and nextGame, are not saved to memory, that should be done after this function call
}
