package ante

import (
	"strings"

	disptypes "github.com/Sifchain/sifnode/x/dispensation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func NewAnteHandler(ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk ante.FeegrantKeeper, gasConsumer ante.SignatureVerificationGasConsumer, signModeHandler signing.SignModeHandler) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		NewAdjustGasPriceDecorator(),    // Custom decorator to adjust gas price for specific msg types
		ante.NewRejectExtensionOptionsDecorator(),
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.TxTimeoutHeightDecorator{},
		ante.NewValidateMemoDecorator(ak),
		ante.NewConsumeGasForTxSizeDecorator(ak),
		ante.NewSetPubKeyDecorator(ak), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(ak),
		ante.NewDeductFeeDecorator(ak, bk, fk),
		ante.NewSigGasConsumeDecorator(ak, gasConsumer),
		ante.NewSigVerificationDecorator(ak, signModeHandler),
		ante.NewIncrementSequenceDecorator(ak),
	)
}

// AdjustGasPriceDecorator is a custom decorator to reduce fee prices .
type AdjustGasPriceDecorator struct {
}

// NewAdjustGasPriceDecorator create a new instance of AdjustGasPriceDecorator
func NewAdjustGasPriceDecorator() AdjustGasPriceDecorator {
	return AdjustGasPriceDecorator{}
}

// AnteHandle adjusts the gas price based on the tx type.
func (r AdjustGasPriceDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	msgs := tx.GetMsgs()
	if len(msgs) == 1 && (strings.Contains(strings.ToLower(sdk.MsgTypeURL(msgs[0])), strings.ToLower(disptypes.MsgTypeCreateDistribution)) ||
		strings.Contains(strings.ToLower(sdk.MsgTypeURL(msgs[0])), strings.ToLower(disptypes.MsgTypeRunDistribution))) {

		minGasPrice := sdk.DecCoin{
			Denom:  "rowan",
			Amount: sdk.MustNewDecFromStr("0.00000005"),
		}
		if !minGasPrice.IsValid() {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrLogic, "invalid gas price")
		}
		ctx = ctx.WithMinGasPrices(sdk.NewDecCoins(minGasPrice))
		return next(ctx, tx, simulate)
	}
	minFee := sdk.ZeroInt()
	for i := range msgs {
		msgTypeURLLower := strings.ToLower(sdk.MsgTypeURL(msgs[i]))

		if strings.Contains(msgTypeURLLower, strings.ToLower(banktypes.TypeMsgSend)) ||
			strings.Contains(msgTypeURLLower, strings.ToLower(banktypes.TypeMsgMultiSend)) ||
			strings.Contains(msgTypeURLLower, "createuserclaim") ||
			strings.Contains(msgTypeURLLower, "swap") ||
			strings.Contains(msgTypeURLLower, "removeliquidity") ||
			strings.Contains(msgTypeURLLower, "addliquidity") {
			minFee = sdk.NewInt(100000000000000000) // 0.1
		} else if strings.Contains(msgTypeURLLower, "transfer") && minFee.LTE(sdk.NewInt(10000000000000000)) {
			minFee = sdk.NewInt(10000000000000000) // 0.01
		}
	}
	if minFee.Equal(sdk.ZeroInt()) {
		return next(ctx, tx, simulate)
	}
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "tx must be a FeeTx")
	}
	fees := feeTx.GetFee()
	rowanFee := sdk.ZeroInt()
	for j := range fees {
		if strings.EqualFold("rowan", fees[j].Denom) {
			rowanFee = fees[j].Amount
		}
	}
	if rowanFee.LTE(sdk.ZeroInt()) {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrLogic, "unsupported fee asset")
	}
	if rowanFee.LT(minFee) {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrLogic, "tx fee is too low")
	}
	return next(ctx, tx, simulate)
}
