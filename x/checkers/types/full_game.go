package types

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/alice/checkers/x/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (storedGame *StoredGame) GetCreatorAddress() (creator sdk.AccAddress, err error) {
	creator, errCreator := sdk.AccAddressFromBech32(storedGame.Creator)
	return creator, sdkerrors.Wrapf(errCreator, ErrInvalidCreator.Error(), storedGame.Creator)
}

func (storedGame *StoredGame) GetRedAddress() (red sdk.AccAddress, err error) {
	red, errRed := sdk.AccAddressFromBech32(storedGame.Red)
	return red, sdkerrors.Wrapf(errRed, ErrInvalidRed.Error(), storedGame.Red)
}

func (storedGame *StoredGame) GetBlackAddress() (black sdk.AccAddress, err error) {
	black, errBlack := sdk.AccAddressFromBech32(storedGame.Black)
	return black, sdkerrors.Wrapf(errBlack, ErrInvalidBlack.Error(), storedGame.Black)
}

func (storedGame *StoredGame) ParseGame() (game *rules.Game, err error) {
	game, errGame := rules.Parse(storedGame.Game)
	if errGame != nil {
		return nil, sdkerrors.Wrapf(errGame, ErrGameNotParseable.Error())
	}
	game.Turn = rules.StringPieces[storedGame.Turn].Player
	if game.Turn.Color == "" {
		return nil, sdkerrors.Wrapf(errors.New(fmt.Sprintf("Turn: %s", storedGame.Turn)), ErrGameNotParseable.Error())
	}
	return game, nil
}

func (storedGame *StoredGame) GetGameIndex() (index int, err error) {
	indexStr := storedGame.Index
	if indexStr == "" {
		return index, sdkerrors.Wrapf(errors.New("invalid index"), ErrInvalidCreator.Error(), storedGame.Creator)
	}
	index, err = strconv.Atoi(indexStr)
	if err != nil {
		return index, sdkerrors.Wrapf(err, ErrInvalidCreator.Error(), storedGame.Creator)
	}
	return index, nil
}

func (storedGame *StoredGame) Validate() (err error) {
	// in order for a stored game to be valid, all getter functions must be able to be
	// called without an error
	_, err = storedGame.GetCreatorAddress()
	if err != nil {
		return
	}
	_, err = storedGame.GetRedAddress()
	if err != nil {
		return
	}
	_, err = storedGame.GetBlackAddress()
	if err != nil {
		return
	}
	_, err = storedGame.ParseGame()
	if err != nil {
		return
	}
	_, err = storedGame.GetGameIndex()
	if err != nil {
		return
	}
	return
}
