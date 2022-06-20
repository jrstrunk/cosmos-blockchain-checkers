package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/checkers module sentinel errors
var (
	ErrInvalidCreator     = sdkerrors.Register(ModuleName, 1100, "creator address is invalid: %s")
	ErrInvalidRed         = sdkerrors.Register(ModuleName, 1101, "red address is invalid: %s")
	ErrInvalidBlack       = sdkerrors.Register(ModuleName, 1102, "black address is invalid: %s")
	ErrGameNotParseable   = sdkerrors.Register(ModuleName, 1103, "game cannot be parsed")
	ErrInvalidIndex       = sdkerrors.Register(ModuleName, 1104, "game index is invalid: %s")
	ErrGameNotFound       = sdkerrors.Register(ModuleName, 1105, "game by id not found: %s")
	ErrCreatorNotPlayer   = sdkerrors.Register(ModuleName, 1106, "message creator is not a player: %s")
	ErrNotPlayerTurn      = sdkerrors.Register(ModuleName, 1107, "player tried to play out of turn: %s")
	ErrWrongMove          = sdkerrors.Register(ModuleName, 1108, "wrong move")
	ErrRedAlreadyPlayed   = sdkerrors.Register(ModuleName, 1109, "red player has already played")
	ErrBlackAlreadyPlayed = sdkerrors.Register(ModuleName, 1110, "black player has already played")
	ErrInvalidDeadline    = sdkerrors.Register(ModuleName, 1111, "deadline cannot be parsed: %s")
)
