package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RouterKey is they name of the bank module
const RouterKey = ModuleName

// NewMsgRequestData creates a new MsgRequestData instance.
func NewMsgRequestData(
	oracleScriptID OracleScriptID,
	calldata []byte,
	requestedValidatorCount int64,
	sufficientValidatorCount int64,
	clientID string,
	sender sdk.AccAddress,
) MsgRequestData {
	return MsgRequestData{
		OracleScriptID:           oracleScriptID,
		Calldata:                 calldata,
		RequestedValidatorCount:  requestedValidatorCount,
		SufficientValidatorCount: sufficientValidatorCount,
		ClientID:                 clientID,
		Sender:                   sender,
	}
}

// Route implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) Type() string { return "request" }

// ValidateBasic implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgRequestData: Sender address must not be empty.")
	}
	if msg.OracleScriptID <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgRequestData: Oracle script id (%d) must be positive.", msg.OracleScriptID)
	}
	if msg.SufficientValidatorCount <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg,
			"MsgRequestData: Sufficient validator count (%d) must be positive.",
			msg.SufficientValidatorCount,
		)
	}
	if msg.RequestedValidatorCount < msg.SufficientValidatorCount {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg,
			"MsgRequestData: Request validator count (%d) must not be less than sufficient validator count (%d).",
			msg.RequestedValidatorCount,
			msg.SufficientValidatorCount,
		)
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// GetSignBytes implements the sdk.Msg interface for MsgRequestData.
func (msg MsgRequestData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// NewMsgReportData creates a new MsgReportData instance.
func NewMsgReportData(
	requestID RequestID,
	dataSet []RawReport,
	validator sdk.ValAddress,
	reporter sdk.AccAddress,
) MsgReportData {
	return MsgReportData{
		RequestID: requestID,
		DataSet:   dataSet,
		Validator: validator,
		Reporter:  reporter,
	}
}

// Route implements the sdk.Msg interface for MsgReportData.
func (msg MsgReportData) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgReportData.
func (msg MsgReportData) Type() string { return "report" }

// ValidateBasic implements the sdk.Msg interface for MsgReportData.
func (msg MsgReportData) ValidateBasic() error {
	if msg.RequestID <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgReportData: Request id (%d) must be positive.", msg.RequestID)
	}
	if msg.DataSet == nil || len(msg.DataSet) == 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgReportData: Data set must not be empty.")
	}
	uniqueMap := make(map[ExternalID]bool)
	for _, rawReport := range msg.DataSet {
		if _, found := uniqueMap[rawReport.ExternalID]; found {
			return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgReportData: External IDs in dataset must be unique.")
		}
		uniqueMap[rawReport.ExternalID] = true
	}
	if msg.Validator.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgReportData: Validator address must not be empty.")
	}
	if msg.Reporter.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgReportData: Reporter address must not be empty.")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgReportData.
func (msg MsgReportData) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Reporter}
}

// GetSignBytes implements the sdk.Msg interface for MsgReportData.
func (msg MsgReportData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// NewMsgCreateDataSource creates a new MsgCreateDataSource instance.
func NewMsgCreateDataSource(
	owner sdk.AccAddress,
	name string,
	description string,
	fee sdk.Coins,
	executable []byte,
	sender sdk.AccAddress,
) MsgCreateDataSource {
	return MsgCreateDataSource{
		Owner:       owner,
		Name:        name,
		Description: description,
		Fee:         fee,
		Executable:  executable,
		Sender:      sender,
	}
}

// Route implements the sdk.Msg interface for MsgCreateDataSource.
func (msg MsgCreateDataSource) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgCreateDataSource.
func (msg MsgCreateDataSource) Type() string { return "create_data_source" }

// ValidateBasic implements the sdk.Msg interface for MsgCreateDataSource.
func (msg MsgCreateDataSource) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgCreateDataSource: Owner address must not be empty.")
	}
	if msg.Name == "" {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgCreateDataSource: Name must not be empty.")
	}
	if msg.Description == "" {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgCreateDataSource: Description must not be empty.")
	}
	if !msg.Fee.IsValid() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgCreateDataSource: Fee (%s) is not valid.", msg.Fee.String())
	}
	if msg.Executable == nil || len(msg.Executable) == 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgCreateDataSource: Executable must not be empty.")
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgCreateDataSource: Sender address must not be empty.")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgCreateDataSource.
func (msg MsgCreateDataSource) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// GetSignBytes implements the sdk.Msg interface for MsgCreateDataSource.
func (msg MsgCreateDataSource) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// NewMsgEditDataSource creates a new MsgEditDataSource instance.
func NewMsgEditDataSource(
	dataSourceID DataSourceID,
	owner sdk.AccAddress,
	name string,
	description string,
	fee sdk.Coins,
	executable []byte,
	sender sdk.AccAddress,
) MsgEditDataSource {
	return MsgEditDataSource{
		DataSourceID: dataSourceID,
		Owner:        owner,
		Name:         name,
		Description:  description,
		Fee:          fee,
		Executable:   executable,
		Sender:       sender,
	}
}

// Route implements the sdk.Msg interface for MsgEditDataSource.
func (msg MsgEditDataSource) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgEditDataSource.
func (msg MsgEditDataSource) Type() string { return "edit_data_source" }

// ValidateBasic implements the sdk.Msg interface for MsgEditDataSource.
func (msg MsgEditDataSource) ValidateBasic() error {
	if msg.DataSourceID <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgEditDataSource: Data source id (%d) must be positive.", msg.DataSourceID)
	}
	if msg.Owner.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgEditDataSource: Owner address must not be empty.")
	}
	if msg.Name == "" {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgEditDataSource: Name must not be empty.")
	}
	if msg.Description == "" {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgEditDataSource: Description must not be empty.")
	}
	if !msg.Fee.IsValid() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgEditDataSource: Fee (%s) is not valid.", msg.Fee.String())
	}
	if msg.Executable == nil || len(msg.Executable) == 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgEditDataSource: Executable must not be empty.")
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgEditDataSource: Sender address must not be empty.")
	}

	return nil
}

// GetSigners implements the sdk.Msg interface for MsgEditDataSource.
func (msg MsgEditDataSource) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// GetSignBytes implements the sdk.Msg interface for MsgEditDataSource.
func (msg MsgEditDataSource) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// NewMsgCreateOracleScript creates a new MsgCreateOracleScript instance.
func NewMsgCreateOracleScript(
	owner sdk.AccAddress,
	name string,
	description string,
	code []byte,
	schema string,
	sourceCodeURL string,
	sender sdk.AccAddress,
) MsgCreateOracleScript {
	return MsgCreateOracleScript{
		Owner:         owner,
		Name:          name,
		Description:   description,
		Code:          code,
		Schema:        schema,
		SourceCodeURL: sourceCodeURL,
		Sender:        sender,
	}
}

// Route implements the sdk.Msg interface for MsgCreateOracleScript.
func (msg MsgCreateOracleScript) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgCreateOracleScript.
func (msg MsgCreateOracleScript) Type() string { return "create_oracle_script" }

// ValidateBasic implements the sdk.Msg interface for MsgCreateOracleScript.
func (msg MsgCreateOracleScript) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgCreateOracleScript: Owner address must not be empty.")
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgCreateOracleScript: Sender address must not be empty.")
	}
	if msg.Name == "" {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgCreateOracleScript: Name must not be empty.")
	}
	if msg.Description == "" {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgCreateOracleScript: Description must not be empty.")
	}
	if msg.Code == nil || len(msg.Code) == 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgCreateOracleScript: Code must not be empty.")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgCreateOracleScript.
func (msg MsgCreateOracleScript) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// GetSignBytes implements the sdk.Msg interface for MsgCreateOracleScript.
func (msg MsgCreateOracleScript) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// NewMsgEditOracleScript creates a new MsgEditOracleScript instance.
func NewMsgEditOracleScript(
	oracleScriptID OracleScriptID,
	owner sdk.AccAddress,
	name string,
	description string,
	code []byte,
	schema string,
	sourceCodeURL string,
	sender sdk.AccAddress,
) MsgEditOracleScript {
	return MsgEditOracleScript{
		OracleScriptID: oracleScriptID,
		Owner:          owner,
		Name:           name,
		Description:    description,
		Code:           code,
		Schema:         schema,
		SourceCodeURL:  sourceCodeURL,
		Sender:         sender,
	}
}

// Route implements the sdk.Msg interface for MsgEditOracleScript.
func (msg MsgEditOracleScript) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgEditOracleScript.
func (msg MsgEditOracleScript) Type() string { return "edit_oracle_script" }

// ValidateBasic implements the sdk.Msg interface for MsgEditOracleScript.
func (msg MsgEditOracleScript) ValidateBasic() error {
	if msg.OracleScriptID <= 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgEditOracleScript: Oracle script id (%d) must be positive.", msg.OracleScriptID)
	}
	if msg.Owner.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgEditOracleScript: Owner address must not be empty.")
	}
	if msg.Sender.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgEditOracleScript: Sender address must not be empty.")
	}
	if msg.Name == "" {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgEditOracleScript: Name must not be empty.")
	}
	if msg.Code == nil || len(msg.Code) == 0 {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgEditOracleScript: Code must not be empty.")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgEditOracleScript.
func (msg MsgEditOracleScript) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// GetSignBytes implements the sdk.Msg interface for MsgEditOracleScript.
func (msg MsgEditOracleScript) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// NewMsgAddOracleAddress creates a new MsgAddOracleAddress instance.
func NewMsgAddOracleAddress(
	validator sdk.ValAddress,
	reporter sdk.AccAddress,
) MsgAddOracleAddress {
	return MsgAddOracleAddress{
		Validator: validator,
		Reporter:  reporter,
	}
}

// Route implements the sdk.Msg interface for MsgAddOracleAddress.
func (msg MsgAddOracleAddress) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgAddOracleAddress.
func (msg MsgAddOracleAddress) Type() string { return "add_oracle_address" }

// ValidateBasic implements the sdk.Msg interface for MsgAddOracleAddress.
func (msg MsgAddOracleAddress) ValidateBasic() error {
	if msg.Validator.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgAddOracleAddress: Validator address must not be empty.")
	}
	if msg.Reporter.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgAddOracleAddress: Reporter address must not be empty.")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgAddOracleAddress.
func (msg MsgAddOracleAddress) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Validator)}
}

// GetSignBytes implements the sdk.Msg interface for MsgAddOracleAddress.
func (msg MsgAddOracleAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// NewMsgRemoveOracleAddress creates a new MsgRemoveOracleAddress instance.
func NewMsgRemoveOracleAddress(
	validator sdk.ValAddress,
	reporter sdk.AccAddress,
) MsgRemoveOracleAddress {
	return MsgRemoveOracleAddress{
		Validator: validator,
		Reporter:  reporter,
	}
}

// Route implements the sdk.Msg interface for MsgRemoveOracleAddress.
func (msg MsgRemoveOracleAddress) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgRemoveOracleAddress.
func (msg MsgRemoveOracleAddress) Type() string { return "remove_oracle_address" }

// ValidateBasic implements the sdk.Msg interface for MsgRemoveOracleAddress.
func (msg MsgRemoveOracleAddress) ValidateBasic() error {
	if msg.Validator.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgRemoveOracleAddress: Validator address must not be empty.")
	}
	if msg.Reporter.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgRemoveOracleAddress: Reporter address must not be empty.")
	}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgRemoveOracleAddress.
func (msg MsgRemoveOracleAddress) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Validator)}
}

// GetSignBytes implements the sdk.Msg interface for MsgRemoveOracleAddress.
func (msg MsgRemoveOracleAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}