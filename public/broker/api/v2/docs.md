# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [public/broker/api/v2/api.proto](#public_broker_api_v2_api-proto)
    - [AddExchangeAPIKeyRequest](#cryptowatch-broker-api-v2-AddExchangeAPIKeyRequest)
    - [AddExchangeAPIKeyRequest.MetaFieldsEntry](#cryptowatch-broker-api-v2-AddExchangeAPIKeyRequest-MetaFieldsEntry)
    - [AddExchangeAPIKeyResponse](#cryptowatch-broker-api-v2-AddExchangeAPIKeyResponse)
    - [Asset](#cryptowatch-broker-api-v2-Asset)
    - [Balance](#cryptowatch-broker-api-v2-Balance)
    - [Balances](#cryptowatch-broker-api-v2-Balances)
    - [Balances.BalancesEntry](#cryptowatch-broker-api-v2-Balances-BalancesEntry)
    - [CancelAllOrdersRequest](#cryptowatch-broker-api-v2-CancelAllOrdersRequest)
    - [CancelAllOrdersResponse](#cryptowatch-broker-api-v2-CancelAllOrdersResponse)
    - [CancelAllOrdersResponse.FailedIdsEntry](#cryptowatch-broker-api-v2-CancelAllOrdersResponse-FailedIdsEntry)
    - [CancelOrderRequest](#cryptowatch-broker-api-v2-CancelOrderRequest)
    - [CancelOrderResponse](#cryptowatch-broker-api-v2-CancelOrderResponse)
    - [ClosePositionRequest](#cryptowatch-broker-api-v2-ClosePositionRequest)
    - [ClosePositionResponse](#cryptowatch-broker-api-v2-ClosePositionResponse)
    - [Column](#cryptowatch-broker-api-v2-Column)
    - [CreateOrderRequest](#cryptowatch-broker-api-v2-CreateOrderRequest)
    - [CreateOrderRequest.MetadataEntry](#cryptowatch-broker-api-v2-CreateOrderRequest-MetadataEntry)
    - [CreateOrderResponse](#cryptowatch-broker-api-v2-CreateOrderResponse)
    - [CreateOrdersRequest](#cryptowatch-broker-api-v2-CreateOrdersRequest)
    - [CreateOrdersResponse](#cryptowatch-broker-api-v2-CreateOrdersResponse)
    - [ErrorResponse](#cryptowatch-broker-api-v2-ErrorResponse)
    - [Exchange](#cryptowatch-broker-api-v2-Exchange)
    - [ExchangeAPIKey](#cryptowatch-broker-api-v2-ExchangeAPIKey)
    - [ExchangeAPIKeysInformation](#cryptowatch-broker-api-v2-ExchangeAPIKeysInformation)
    - [ExchangeAccount](#cryptowatch-broker-api-v2-ExchangeAccount)
    - [ExchangeAccountWallets](#cryptowatch-broker-api-v2-ExchangeAccountWallets)
    - [ExchangeAccountWallets.WalletsEntry](#cryptowatch-broker-api-v2-ExchangeAccountWallets-WalletsEntry)
    - [ExchangeAccounts](#cryptowatch-broker-api-v2-ExchangeAccounts)
    - [ExchangeAccounts.ExchangeAccountsEntry](#cryptowatch-broker-api-v2-ExchangeAccounts-ExchangeAccountsEntry)
    - [ExchangeConnectionUpdate](#cryptowatch-broker-api-v2-ExchangeConnectionUpdate)
    - [ExchangeConnectionUpdate.UserConnectionsEntry](#cryptowatch-broker-api-v2-ExchangeConnectionUpdate-UserConnectionsEntry)
    - [ExchangeInformation](#cryptowatch-broker-api-v2-ExchangeInformation)
    - [Filters](#cryptowatch-broker-api-v2-Filters)
    - [GetAllMarketsInformationRequest](#cryptowatch-broker-api-v2-GetAllMarketsInformationRequest)
    - [GetAllMarketsInformationResponse](#cryptowatch-broker-api-v2-GetAllMarketsInformationResponse)
    - [GetExchangeInformationRequest](#cryptowatch-broker-api-v2-GetExchangeInformationRequest)
    - [GetExchangeInformationResponse](#cryptowatch-broker-api-v2-GetExchangeInformationResponse)
    - [GetExchangeInformationResponse.ExchangeInfoEntry](#cryptowatch-broker-api-v2-GetExchangeInformationResponse-ExchangeInfoEntry)
    - [GetFeeScheduleRequest](#cryptowatch-broker-api-v2-GetFeeScheduleRequest)
    - [GetFeeScheduleResponse](#cryptowatch-broker-api-v2-GetFeeScheduleResponse)
    - [GetMarketInformationRequest](#cryptowatch-broker-api-v2-GetMarketInformationRequest)
    - [GetMarketInformationResponse](#cryptowatch-broker-api-v2-GetMarketInformationResponse)
    - [GetOrderRequest](#cryptowatch-broker-api-v2-GetOrderRequest)
    - [GetOrderResponse](#cryptowatch-broker-api-v2-GetOrderResponse)
    - [GetServerTimeRequest](#cryptowatch-broker-api-v2-GetServerTimeRequest)
    - [GetServerTimeResponse](#cryptowatch-broker-api-v2-GetServerTimeResponse)
    - [GetUserSettingsRequest](#cryptowatch-broker-api-v2-GetUserSettingsRequest)
    - [GetUserSettingsResponse](#cryptowatch-broker-api-v2-GetUserSettingsResponse)
    - [LeverageLevels](#cryptowatch-broker-api-v2-LeverageLevels)
    - [LeverageLevels.LeverageBound](#cryptowatch-broker-api-v2-LeverageLevels-LeverageBound)
    - [LeverageLevels.LeverageList](#cryptowatch-broker-api-v2-LeverageLevels-LeverageList)
    - [ListAvailableFiltersRequest](#cryptowatch-broker-api-v2-ListAvailableFiltersRequest)
    - [ListAvailableFiltersResponse](#cryptowatch-broker-api-v2-ListAvailableFiltersResponse)
    - [ListExchangeAccountsRequest](#cryptowatch-broker-api-v2-ListExchangeAccountsRequest)
    - [ListExchangeAccountsResponse](#cryptowatch-broker-api-v2-ListExchangeAccountsResponse)
    - [ListExchangeAccountsResponse.ExchangeAccountsByExchangeEntry](#cryptowatch-broker-api-v2-ListExchangeAccountsResponse-ExchangeAccountsByExchangeEntry)
    - [ListExchangeWalletsRequest](#cryptowatch-broker-api-v2-ListExchangeWalletsRequest)
    - [ListExchangeWalletsResponse](#cryptowatch-broker-api-v2-ListExchangeWalletsResponse)
    - [ListExchangeWalletsResponse.ExchangeAccountWalletsEntry](#cryptowatch-broker-api-v2-ListExchangeWalletsResponse-ExchangeAccountWalletsEntry)
    - [ListOrdersRequest](#cryptowatch-broker-api-v2-ListOrdersRequest)
    - [ListOrdersResponse](#cryptowatch-broker-api-v2-ListOrdersResponse)
    - [ListPositionsRequest](#cryptowatch-broker-api-v2-ListPositionsRequest)
    - [ListPositionsResponse](#cryptowatch-broker-api-v2-ListPositionsResponse)
    - [ListTradesRequest](#cryptowatch-broker-api-v2-ListTradesRequest)
    - [ListTradesResponse](#cryptowatch-broker-api-v2-ListTradesResponse)
    - [MFAAuthenticationResult](#cryptowatch-broker-api-v2-MFAAuthenticationResult)
    - [Market](#cryptowatch-broker-api-v2-Market)
    - [MarketInformation](#cryptowatch-broker-api-v2-MarketInformation)
    - [MessageAcknowledgementResponse](#cryptowatch-broker-api-v2-MessageAcknowledgementResponse)
    - [OptionalOrderType](#cryptowatch-broker-api-v2-OptionalOrderType)
    - [OptionalSide](#cryptowatch-broker-api-v2-OptionalSide)
    - [Order](#cryptowatch-broker-api-v2-Order)
    - [PauseOrdersRequest](#cryptowatch-broker-api-v2-PauseOrdersRequest)
    - [PauseOrdersResponse](#cryptowatch-broker-api-v2-PauseOrdersResponse)
    - [PauseOrdersResponse.FailedIdsEntry](#cryptowatch-broker-api-v2-PauseOrdersResponse-FailedIdsEntry)
    - [PlayOrdersRequest](#cryptowatch-broker-api-v2-PlayOrdersRequest)
    - [PlayOrdersResponse](#cryptowatch-broker-api-v2-PlayOrdersResponse)
    - [PlayOrdersResponse.FailedIdsEntry](#cryptowatch-broker-api-v2-PlayOrdersResponse-FailedIdsEntry)
    - [Position](#cryptowatch-broker-api-v2-Position)
    - [ReplaceOrderRequest](#cryptowatch-broker-api-v2-ReplaceOrderRequest)
    - [ReplaceOrderResponse](#cryptowatch-broker-api-v2-ReplaceOrderResponse)
    - [RevokeAllExchangeAPIKeysRequest](#cryptowatch-broker-api-v2-RevokeAllExchangeAPIKeysRequest)
    - [RevokeAllExchangeAPIKeysResponse](#cryptowatch-broker-api-v2-RevokeAllExchangeAPIKeysResponse)
    - [RevokeExchangeAPIKeyRequest](#cryptowatch-broker-api-v2-RevokeExchangeAPIKeyRequest)
    - [RevokeExchangeAPIKeyResponse](#cryptowatch-broker-api-v2-RevokeExchangeAPIKeyResponse)
    - [SettlePositionRequest](#cryptowatch-broker-api-v2-SettlePositionRequest)
    - [SettlePositionResponse](#cryptowatch-broker-api-v2-SettlePositionResponse)
    - [SubscribeToExchangeAccountsRequest](#cryptowatch-broker-api-v2-SubscribeToExchangeAccountsRequest)
    - [SubscribeToExchangeAccountsResponse](#cryptowatch-broker-api-v2-SubscribeToExchangeAccountsResponse)
    - [Tab](#cryptowatch-broker-api-v2-Tab)
    - [TickSizes](#cryptowatch-broker-api-v2-TickSizes)
    - [Trade](#cryptowatch-broker-api-v2-Trade)
    - [TradingWebsocketRequest](#cryptowatch-broker-api-v2-TradingWebsocketRequest)
    - [TradingWebsocketResponse](#cryptowatch-broker-api-v2-TradingWebsocketResponse)
    - [TrailingOffsetPrice](#cryptowatch-broker-api-v2-TrailingOffsetPrice)
    - [UpdateUserSettingsRequest](#cryptowatch-broker-api-v2-UpdateUserSettingsRequest)
    - [UpdateUserSettingsResponse](#cryptowatch-broker-api-v2-UpdateUserSettingsResponse)
    - [UserSettings](#cryptowatch-broker-api-v2-UserSettings)
    - [Wallets](#cryptowatch-broker-api-v2-Wallets)
    - [Wallets.BalancesEntry](#cryptowatch-broker-api-v2-Wallets-BalancesEntry)
  
    - [APIKeyStatus](#cryptowatch-broker-api-v2-APIKeyStatus)
    - [ContractSchema](#cryptowatch-broker-api-v2-ContractSchema)
    - [ExchangeStatus](#cryptowatch-broker-api-v2-ExchangeStatus)
    - [FeeCurrency](#cryptowatch-broker-api-v2-FeeCurrency)
    - [MFAAuthenticationStatus](#cryptowatch-broker-api-v2-MFAAuthenticationStatus)
    - [MarginAction](#cryptowatch-broker-api-v2-MarginAction)
    - [MarketType](#cryptowatch-broker-api-v2-MarketType)
    - [OrderType](#cryptowatch-broker-api-v2-OrderType)
    - [OrdersSortKey](#cryptowatch-broker-api-v2-OrdersSortKey)
    - [ReferencePrice](#cryptowatch-broker-api-v2-ReferencePrice)
    - [Side](#cryptowatch-broker-api-v2-Side)
    - [Sorting](#cryptowatch-broker-api-v2-Sorting)
    - [Status](#cryptowatch-broker-api-v2-Status)
    - [TimeInForce](#cryptowatch-broker-api-v2-TimeInForce)
    - [TradesSortKey](#cryptowatch-broker-api-v2-TradesSortKey)
    - [TrailingOffsetType](#cryptowatch-broker-api-v2-TrailingOffsetType)
    - [UserConnectionStatus](#cryptowatch-broker-api-v2-UserConnectionStatus)
  
    - [TradingService](#cryptowatch-broker-api-v2-TradingService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="public_broker_api_v2_api-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## public/broker/api/v2/api.proto



<a name="cryptowatch-broker-api-v2-AddExchangeAPIKeyRequest"></a>

### AddExchangeAPIKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_id | [int32](#int32) |  | The ID of the exchange we want to add an API key pair for. Could be found with the REST API under https://api.cryptowat.ch/exchanges. |
| exchange_account_id | [string](#string) |  | [optional] The exchange account to attach this API key too. By not specifying this, the keys will be attached to the default exchange account. |
| key | [string](#string) |  | The exchange key to register. |
| secret | [string](#string) |  | The secret associated with the key. |
| sub_account | [string](#string) |  | [optional] The name of the sub-account to use with this key |
| meta_fields | [AddExchangeAPIKeyRequest.MetaFieldsEntry](#cryptowatch-broker-api-v2-AddExchangeAPIKeyRequest-MetaFieldsEntry) | repeated | meta_fields is a flexible list of parameters to pass on when registering an API key. For example with Coinbase, a passphrase is necessary in order to be able to use the API. For Bitstamp you want to add a `customer_id`. A full list of fields to pass on could be found with the `GetExchangeInformation()` call which lists all the required fields for exchanges. |






<a name="cryptowatch-broker-api-v2-AddExchangeAPIKeyRequest-MetaFieldsEntry"></a>

### AddExchangeAPIKeyRequest.MetaFieldsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="cryptowatch-broker-api-v2-AddExchangeAPIKeyResponse"></a>

### AddExchangeAPIKeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_account_name | [string](#string) |  | The account name of the exchange account that the API key was attached too. |
| shortened_key | [string](#string) |  | The first 4 characters of the key that was just added. This value is shortened for security purposes since it is sensitive data. |






<a name="cryptowatch-broker-api-v2-Asset"></a>

### Asset
Represents an asset, for example BTC or ETH.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| asset_id | [int32](#int32) |  | The ID of the asset. (mapping could be found with the REST API under https://api.cryptowat.ch/assets) |
| asset_symbol | [string](#string) |  | The symbol of the asset: ie. btc, eth, usd, etc. |
| asset_name | [string](#string) |  | The asset name, for example: bitcoin, ethereum, euro, dollar, etc. |






<a name="cryptowatch-broker-api-v2-Balance"></a>

### Balance
Balance represents the total and available amount of asset for a given
exchange, asset and wallet type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [string](#string) |  | The total balance owned for that asset. |
| available | [string](#string) |  | The available balance for this asset to use for trading. Ensure to checkout the `virtual_available` balance as well |
| asset_name | [string](#string) |  | The name of the asset |
| exchange_name | [string](#string) |  | The name of the exchange |
| exchange_account_name | [string](#string) |  | The name of the exchange account for which this balance belongs to |
| virtual_available | [string](#string) |  | virtual_available is the available balance including synthetic orders. If a user has any open synethetic order, this value will denote how much would they have left if the entire synthetic order was filled. This will be useful for the user to know how much to keep in their account so their synthetic orders can completely be placed. |
| asset_symbol | [string](#string) |  | The symbol of the asset |






<a name="cryptowatch-broker-api-v2-Balances"></a>

### Balances
Balances represent all of the balances the user have accross 
all of their assets.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balances | [Balances.BalancesEntry](#cryptowatch-broker-api-v2-Balances-BalancesEntry) | repeated | balances is a map of all of the user balances where the key of the map is the asset ID, the value being the Balance for that asset. |






<a name="cryptowatch-broker-api-v2-Balances-BalancesEntry"></a>

### Balances.BalancesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [Balance](#cryptowatch-broker-api-v2-Balance) |  |  |






<a name="cryptowatch-broker-api-v2-CancelAllOrdersRequest"></a>

### CancelAllOrdersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_id | [int32](#int32) |  | [optional] If specified, deletes orders for this exchange only. (Could be found with the REST API under https://api.cryptowat.ch/exchanges) |
| market_id | [int32](#int32) |  | [optional] If specified, deletes orders belonging to this market only. (could be found with the REST API under https://api.cryptowat.ch/markets) |
| exchange_account_id | [string](#string) |  | [optional] If specified, deletes orders belonging to this exchange account only. |






<a name="cryptowatch-broker-api-v2-CancelAllOrdersResponse"></a>

### CancelAllOrdersResponse
CancelAllOrdersResponse returns a list of failed order ids (key) with the error/reason for the failure (value).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| failed_ids | [CancelAllOrdersResponse.FailedIdsEntry](#cryptowatch-broker-api-v2-CancelAllOrdersResponse-FailedIdsEntry) | repeated |  |






<a name="cryptowatch-broker-api-v2-CancelAllOrdersResponse-FailedIdsEntry"></a>

### CancelAllOrdersResponse.FailedIdsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [ErrorResponse](#cryptowatch-broker-api-v2-ErrorResponse) |  |  |






<a name="cryptowatch-broker-api-v2-CancelOrderRequest"></a>

### CancelOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [string](#string) |  | order_id is the exchange&#39;s identifier for the order to be cancelled |
| exchange_id | [int32](#int32) |  | [optional] The exchange ID that the order is on. Because order IDs are not guaranteed to be unique across exchanges, there is a very small change that a user could have two orders with the same order ID, and end up cancelling the wrong order. Providing the exchange_id removes that possibility, since order IDs are unique within an exchange. (Could be found with the REST API under https://api.cryptowat.ch/exchange) |






<a name="cryptowatch-broker-api-v2-CancelOrderResponse"></a>

### CancelOrderResponse
CancelOrderResponse is an empty response






<a name="cryptowatch-broker-api-v2-ClosePositionRequest"></a>

### ClosePositionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| position_id | [string](#string) |  | position_id is the exchange&#39;s identifier for the position to close. |
| exchange_id | [int32](#int32) |  | [optional] The exchange ID that the position is on. Because position IDs are not guaranteed to be unique across exchanges, there is a very small change that a user could have two positions with the same position ID, and end up closing the wrong position. Providing the exchange_id removes that possibility, since position IDs are unique within an exchange. (Could be found with the REST API under https://api.cryptowat.ch/exchange) |






<a name="cryptowatch-broker-api-v2-ClosePositionResponse"></a>

### ClosePositionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders | [Order](#cryptowatch-broker-api-v2-Order) | repeated | The orders created as a result of closing the position. Not guaranteed to be filled, some exchanges do not include this in the response. |






<a name="cryptowatch-broker-api-v2-Column"></a>

### Column
column setting inside each tab


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | name of column |
| fixed | [bool](#bool) |  | is column fixed |
| visibility | [bool](#bool) |  | is column visible |






<a name="cryptowatch-broker-api-v2-CreateOrderRequest"></a>

### CreateOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_account_id | [string](#string) |  | [optional] id of the exchange account to place the order through. If not provided, the default account will be used. |
| side | [Side](#cryptowatch-broker-api-v2-Side) |  | Could be either &#39;Buy&#39; or &#39;Sell&#39;. |
| type | [OrderType](#cryptowatch-broker-api-v2-OrderType) |  | Could be any of the values defined in the &#39;OrderType&#39; enum except `SETTLE_POSITION`. If you need to settle a position make a request using `SettlePosition()`. Type will define how we should place this order and usually which parameters we are going to use (if a stop loss, we&#39;ll use the stop price, etc.). Default value is MARKET. |
| market_id | [int32](#int32) |  | The ID of the desired market to place an order upon (could be found with the REST API under https://api.cryptowat.ch/markets) |
| leverage | [string](#string) |  | [optional] Leverage value for a margin order (default: 0 / no leverage). |
| amount | [string](#string) |  | Amount of base asset to buy or sell. |
| price | [string](#string) |  | Quote asset price to buy or sell the base asset. Used for `LIMIT` orders. |
| stop_price | [string](#string) |  | The stop price for a `Stop Market` or `Stop Limit` order. |
| stop_limit_price | [string](#string) |  | The limit order execution price for `STOP_LIMIT` order. |
| trigger_price | [string](#string) |  | The trigger price for a `MARKET_IF_TOUCHED` order. |
| execution_price | [string](#string) |  | The limit order execution price for `LIMIT_IF_TOUCHED` order. |
| trailing_stop_market_offset | [TrailingOffsetPrice](#cryptowatch-broker-api-v2-TrailingOffsetPrice) |  | The trailing offset price for a `TRAILING_STOP_MARKET` order. |
| trailing_stop_limit_offset | [TrailingOffsetPrice](#cryptowatch-broker-api-v2-TrailingOffsetPrice) |  | The limit price offset for a `TRAILING_STOP_LIMIT` order. |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | [optional] Timestamp defining when the order should be placed. |
| expire_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | [optional] Timestamp for the time the order should be canceled. |
| post_only | [bool](#bool) |  | [optional] Defines if the order should de deleted if it won&#39;t appear on the orderbook and fill immediately (default: false). |
| immediate_or_cancel | [bool](#bool) |  | [optional] Defines if the order should be immediate or cancel (default: false) |
| fill_or_kill | [bool](#bool) |  | [optional] Defines if the order should be a fill or kill order (default: false) |
| hidden | [bool](#bool) |  | [optional] Determines if the order should be hidden from the orderbook or not (default: false). |
| fee_currency | [FeeCurrency](#cryptowatch-broker-api-v2-FeeCurrency) |  | [optional] Currency to use in order to pay for the trading fees. (default: base currency) |
| closing_order | [CreateOrderRequest](#cryptowatch-broker-api-v2-CreateOrderRequest) |  | [optional] If the order fills in its entirety, place the following closing order. |
| reduce_only | [bool](#bool) |  | [optional] Defines if the order should be immediately canceled, if it would increase the size of an open position. (default: false) |
| iceberg_quantity | [string](#string) |  | [optional] An order type to trade in smaller predetermined quantities, called iceberg_quantity, in order to hide the total order quantity. |
| reference_price | [ReferencePrice](#cryptowatch-broker-api-v2-ReferencePrice) |  | [optional] Trigger price references are used in futures trading to denote the kind of price at which a trade action is to be taken, if supported for a market. (default: Default) |
| synthetic | [bool](#bool) |  | [optional] Specify if an order is synthetic or not. (default: false) |
| metadata | [CreateOrderRequest.MetadataEntry](#cryptowatch-broker-api-v2-CreateOrderRequest-MetadataEntry) | repeated | We expect following be present in this field Key : &#34;platform&#34;, Values = [&#34;desktop&#34;, &#34;web&#34;, &#34;mobile&#34;, &#34;bcli&#34;, &#34;other&#34;] This represents platform where trade was placed

 Key: &#34;tradeform&#34;, Values = [&#34;legacy&#34;, &#34;standard&#34;, &#34;active&#34;, &#34;ladder&#34;, &#34;tui&#34;, &#34;other&#34;] This represents which trade form was used to place order |






<a name="cryptowatch-broker-api-v2-CreateOrderRequest-MetadataEntry"></a>

### CreateOrderRequest.MetadataEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="cryptowatch-broker-api-v2-CreateOrderResponse"></a>

### CreateOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order | [Order](#cryptowatch-broker-api-v2-Order) | repeated | CreateOrder may result in multiple orders returned since some order types can cause multiple orders to be returned from the exchange. Example: Order type `ONE_CANCELS_OTHER` will return multiple orders from the exchange. |






<a name="cryptowatch-broker-api-v2-CreateOrdersRequest"></a>

### CreateOrdersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders | [CreateOrderRequest](#cryptowatch-broker-api-v2-CreateOrderRequest) | repeated | list of orders to be created |






<a name="cryptowatch-broker-api-v2-CreateOrdersResponse"></a>

### CreateOrdersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders | [Order](#cryptowatch-broker-api-v2-Order) | repeated | multiple orders can be returned even if only one order was sent in `CreateOrdersRequest` since some order types can result in multiple orders returned from the exchange. |






<a name="cryptowatch-broker-api-v2-ErrorResponse"></a>

### ErrorResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [int32](#int32) |  |  |
| message | [string](#string) |  |  |






<a name="cryptowatch-broker-api-v2-Exchange"></a>

### Exchange



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_id | [int32](#int32) |  | Cryptowatch identifier for the exchange. Could be found with the REST API under https://api.cryptowat.ch/exchanges. |
| exchange_name | [string](#string) |  | Name of the exchange |






<a name="cryptowatch-broker-api-v2-ExchangeAPIKey"></a>

### ExchangeAPIKey



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shortened_key | [string](#string) |  | The shortened version of the key to identify it, it is shortened because this is a sensitive information. |
| status | [APIKeyStatus](#cryptowatch-broker-api-v2-APIKeyStatus) |  | The status of the API key, whether it is valid or not. See the `APIKeyStatus` enumeration for more information about the statuses. |






<a name="cryptowatch-broker-api-v2-ExchangeAPIKeysInformation"></a>

### ExchangeAPIKeysInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| api_keys | [ExchangeAPIKey](#cryptowatch-broker-api-v2-ExchangeAPIKey) | repeated | The list of api keys for a given exchange. |






<a name="cryptowatch-broker-api-v2-ExchangeAccount"></a>

### ExchangeAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | cryptowatch identifier of this exchange account |
| exchange | [Exchange](#cryptowatch-broker-api-v2-Exchange) |  | the exchange that this exchange account belongs to |
| read_only | [bool](#bool) |  | whether this account is read only or not. Enforced by cryptowatch, not necessarily by the exchange API key permissions. |
| account_name | [string](#string) |  | human readable account name assigned by the user |
| is_default | [bool](#bool) |  | whether this is the default exchange account for this exchange or not. Only one exchange account per exchange will be the default. |
| api_key | [ExchangeAPIKey](#cryptowatch-broker-api-v2-ExchangeAPIKey) |  | the API keys attached to this account |






<a name="cryptowatch-broker-api-v2-ExchangeAccountWallets"></a>

### ExchangeAccountWallets



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| wallets | [ExchangeAccountWallets.WalletsEntry](#cryptowatch-broker-api-v2-ExchangeAccountWallets-WalletsEntry) | repeated | A map of exchange account IDs to the wallets for that exchange account. |






<a name="cryptowatch-broker-api-v2-ExchangeAccountWallets-WalletsEntry"></a>

### ExchangeAccountWallets.WalletsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Wallets](#cryptowatch-broker-api-v2-Wallets) |  |  |






<a name="cryptowatch-broker-api-v2-ExchangeAccounts"></a>

### ExchangeAccounts



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_accounts | [ExchangeAccounts.ExchangeAccountsEntry](#cryptowatch-broker-api-v2-ExchangeAccounts-ExchangeAccountsEntry) | repeated | exchange_accounts is a map of exchange account ID to the corresponding exchange account. |






<a name="cryptowatch-broker-api-v2-ExchangeAccounts-ExchangeAccountsEntry"></a>

### ExchangeAccounts.ExchangeAccountsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [ExchangeAccount](#cryptowatch-broker-api-v2-ExchangeAccount) |  |  |






<a name="cryptowatch-broker-api-v2-ExchangeConnectionUpdate"></a>

### ExchangeConnectionUpdate
ExchangeConnectionUpdate represents Cryptowatch connection to the exchange.
Users connected to high performance trading platforms need to know the realtime 
status of an exchange prior to placing orders at that exchange. If an exchange is down 
or degraded for any reason, a trader may not want to risk their money by sending orders there.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [ExchangeStatus](#cryptowatch-broker-api-v2-ExchangeStatus) |  | `status` represent the general status of the exchange. |
| exchange_id | [int32](#int32) |  | exchange unique identifier |
| user_connection_status | [UserConnectionStatus](#cryptowatch-broker-api-v2-UserConnectionStatus) |  | **Deprecated.** DEPRECATED, please use `user_connections` field instead. To maintain backwards compatibility: - When sent as a response to a request, this field will contain the connection status of the default exchange account. - When pushed as an update, this field will contain the connection status of whichever exchange account the update is for. If you have multiple subscriptions and need to disambiguate, please use `user_connections`. |
| user_connections | [ExchangeConnectionUpdate.UserConnectionsEntry](#cryptowatch-broker-api-v2-ExchangeConnectionUpdate-UserConnectionsEntry) | repeated | `user_connection_status` represent user specific connection status with the exchange for each exchange account, where the key is the exchange account ID and the value is the connection status for that exchange account. This is different than `status` since `status` represents the exchange&#39;s overall status. An exchange status may be `AVAILABLE` but the user&#39;s connection to the exchange isn&#39;t. Examples where user may not have a `CONNECTION_AVAILABLE` status even when the exchange is `AVAILABLE`: - user being rate limited - user key is revoked - user doesn&#39;t have a websocket connection to the exchange in the backend |






<a name="cryptowatch-broker-api-v2-ExchangeConnectionUpdate-UserConnectionsEntry"></a>

### ExchangeConnectionUpdate.UserConnectionsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [UserConnectionStatus](#cryptowatch-broker-api-v2-UserConnectionStatus) |  |  |






<a name="cryptowatch-broker-api-v2-ExchangeInformation"></a>

### ExchangeInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_id | [int32](#int32) |  | Unique ID for the exchange. Could be found with the REST API under https://api.cryptowat.ch/exchanges. |
| exchange_name | [string](#string) |  | Name of the exchange. |
| fields | [string](#string) | repeated | The fields required in order to enter a new API key successfully. For example: - Kraken requires: [api_key, api_secret] - Coinbase Pro requires: [api_key, api_secret, passphrase] |
| exchange_accounts | [ExchangeAccounts](#cryptowatch-broker-api-v2-ExchangeAccounts) |  | information about any exchange accounts the user has registered. if there are accounts, it is guaranteed one of them is marked as the default account for the exchange. Default account is used for operations where user does not specify a specific exchange account, so Cryptowatch perform the action using the exchange&#39;s default account. |
| connection | [ExchangeConnectionUpdate](#cryptowatch-broker-api-v2-ExchangeConnectionUpdate) |  | information about the current exchange status |






<a name="cryptowatch-broker-api-v2-Filters"></a>

### Filters
Filters are user independent data. Each user may see different 
filters based on the orders/trades/positions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchanges | [Exchange](#cryptowatch-broker-api-v2-Exchange) | repeated | All exchanges available to be filtered |
| markets | [Market](#cryptowatch-broker-api-v2-Market) | repeated | All markets availables to be filtered |
| types | [OrderType](#cryptowatch-broker-api-v2-OrderType) | repeated | All order types availables to be filtered Only for orders |
| base_assets | [Asset](#cryptowatch-broker-api-v2-Asset) | repeated | All base assets availables to be filtered |
| quote_assets | [Asset](#cryptowatch-broker-api-v2-Asset) | repeated | All quote assets availables to be filtered |
| statuses | [Status](#cryptowatch-broker-api-v2-Status) | repeated | All order statuses availables to be filtered Only for orders |






<a name="cryptowatch-broker-api-v2-GetAllMarketsInformationRequest"></a>

### GetAllMarketsInformationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_id | [int32](#int32) |  | The exchange to get all of its markets information. Exchange Ids can be found with the REST API under https://api.cryptowat.ch/exchanges. |






<a name="cryptowatch-broker-api-v2-GetAllMarketsInformationResponse"></a>

### GetAllMarketsInformationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| markets_info | [MarketInformation](#cryptowatch-broker-api-v2-MarketInformation) | repeated | list of market information |






<a name="cryptowatch-broker-api-v2-GetExchangeInformationRequest"></a>

### GetExchangeInformationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| only_registered_exchanges | [bool](#bool) |  | if true, will only respond with information about exchanges the user is registered with. (exchanges that the user has keys setup for in Cryptowatch) |






<a name="cryptowatch-broker-api-v2-GetExchangeInformationResponse"></a>

### GetExchangeInformationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_info | [GetExchangeInformationResponse.ExchangeInfoEntry](#cryptowatch-broker-api-v2-GetExchangeInformationResponse-ExchangeInfoEntry) | repeated | exchange_info is a map where the key is the exchange id and the value is the corresponding ExchangeInformation |






<a name="cryptowatch-broker-api-v2-GetExchangeInformationResponse-ExchangeInfoEntry"></a>

### GetExchangeInformationResponse.ExchangeInfoEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [ExchangeInformation](#cryptowatch-broker-api-v2-ExchangeInformation) |  |  |






<a name="cryptowatch-broker-api-v2-GetFeeScheduleRequest"></a>

### GetFeeScheduleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| market_id | [int32](#int32) |  | The market id to get the fee schedule for. A cryptowatch market identifier (could be found with the REST API under https://api.cryptowat.ch/markets) |






<a name="cryptowatch-broker-api-v2-GetFeeScheduleResponse"></a>

### GetFeeScheduleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| makerFee | [string](#string) |  | makerFee is a fee for orders that provide liquidity. This fee is unique per user. Ex: &#34;0.1&#34; means 0.1% fee |
| takerFee | [string](#string) |  | takerFee is a fee for orders that take away liquidity. This fee is unique per user. Ex: &#34;0.1&#34; means 0.1% fee |
| last_updated | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | represent when was the last time the user fee schedules were synced with the exchange |






<a name="cryptowatch-broker-api-v2-GetMarketInformationRequest"></a>

### GetMarketInformationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| market_id | [int32](#int32) |  | the market id to get the information about. A cryptowatch market identifier (could be found with the REST API under https://api.cryptowat.ch/markets) |






<a name="cryptowatch-broker-api-v2-GetMarketInformationResponse"></a>

### GetMarketInformationResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| market_info | [MarketInformation](#cryptowatch-broker-api-v2-MarketInformation) |  |  |






<a name="cryptowatch-broker-api-v2-GetOrderRequest"></a>

### GetOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [string](#string) |  | order_id is the exchange&#39;s identifier for the order. |
| exchange_id | [int32](#int32) |  | [optional] The exchange ID that the order is on. Because order IDs are not guaranteed to be unique across exchanges, there is a very small change that a user could have two orders with the same order ID, and end up retrieving the wrong order. Providing the exchange_id removes that possibility, since order IDs are unique within an exchange. (Could be found with the REST API under https://api.cryptowat.ch/exchanges.) |






<a name="cryptowatch-broker-api-v2-GetOrderResponse"></a>

### GetOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order | [Order](#cryptowatch-broker-api-v2-Order) |  |  |






<a name="cryptowatch-broker-api-v2-GetServerTimeRequest"></a>

### GetServerTimeRequest
GetServerTimeRequest is an empty request






<a name="cryptowatch-broker-api-v2-GetServerTimeResponse"></a>

### GetServerTimeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| server_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | server_time is the current server time of the instance that served the user request |






<a name="cryptowatch-broker-api-v2-GetUserSettingsRequest"></a>

### GetUserSettingsRequest
GetUserSettingsRequest is an empty request






<a name="cryptowatch-broker-api-v2-GetUserSettingsResponse"></a>

### GetUserSettingsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| settings | [UserSettings](#cryptowatch-broker-api-v2-UserSettings) |  | cryptowatch settings |






<a name="cryptowatch-broker-api-v2-LeverageLevels"></a>

### LeverageLevels



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| leverage_list | [LeverageLevels.LeverageList](#cryptowatch-broker-api-v2-LeverageLevels-LeverageList) |  |  |
| leverage_bound | [LeverageLevels.LeverageBound](#cryptowatch-broker-api-v2-LeverageLevels-LeverageBound) |  |  |






<a name="cryptowatch-broker-api-v2-LeverageLevels-LeverageBound"></a>

### LeverageLevels.LeverageBound



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| min_leverage | [string](#string) |  | Used for arbitrary values (up to `precision` decimals) within `minLeverage`/`maxLeverage` range |
| max_leverage | [string](#string) |  |  |
| precision | [string](#string) |  | The maximum precision that could be used for arbitrary values within the leverage bounds |






<a name="cryptowatch-broker-api-v2-LeverageLevels-LeverageList"></a>

### LeverageLevels.LeverageList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| list | [string](#string) | repeated | Used for fixed set of leverage values, e.g: [1,10,20,30,40,50] |






<a name="cryptowatch-broker-api-v2-ListAvailableFiltersRequest"></a>

### ListAvailableFiltersRequest
ListAvailableFiltersRequest is an empty request






<a name="cryptowatch-broker-api-v2-ListAvailableFiltersResponse"></a>

### ListAvailableFiltersResponse
ListAvailableFiltersResponse is unique per user
since the response is depended upon the orders, trades 
and positions a user have.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders_filters | [Filters](#cryptowatch-broker-api-v2-Filters) |  | orders filters a user have available |
| trades_filters | [Filters](#cryptowatch-broker-api-v2-Filters) |  | trades filters a user have available |






<a name="cryptowatch-broker-api-v2-ListExchangeAccountsRequest"></a>

### ListExchangeAccountsRequest
ListExchangeAccountsRequest is an empty request






<a name="cryptowatch-broker-api-v2-ListExchangeAccountsResponse"></a>

### ListExchangeAccountsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_accounts_by_exchange | [ListExchangeAccountsResponse.ExchangeAccountsByExchangeEntry](#cryptowatch-broker-api-v2-ListExchangeAccountsResponse-ExchangeAccountsByExchangeEntry) | repeated | exchange_accounts_by_exchange is a map of exchange ID to a the users exchange accounts for that particular exchange. If there are exchange accounts, it is guaranteed that one account is set as the default account for the exchange. |






<a name="cryptowatch-broker-api-v2-ListExchangeAccountsResponse-ExchangeAccountsByExchangeEntry"></a>

### ListExchangeAccountsResponse.ExchangeAccountsByExchangeEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [ExchangeAccounts](#cryptowatch-broker-api-v2-ExchangeAccounts) |  |  |






<a name="cryptowatch-broker-api-v2-ListExchangeWalletsRequest"></a>

### ListExchangeWalletsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_id | [int32](#int32) |  | [optional] If specified, only returns the wallets and balances for that exchange. (Could be found with the REST API under https://api.cryptowat.ch/exchanges) |
| exchange_account_id | [string](#string) |  | [optional] If specified, only returns the wallets and balances for that exchange account. |






<a name="cryptowatch-broker-api-v2-ListExchangeWalletsResponse"></a>

### ListExchangeWalletsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_account_wallets | [ListExchangeWalletsResponse.ExchangeAccountWalletsEntry](#cryptowatch-broker-api-v2-ListExchangeWalletsResponse-ExchangeAccountWalletsEntry) | repeated | A map of exchange IDs associated with their complete wallet set and balances information. |






<a name="cryptowatch-broker-api-v2-ListExchangeWalletsResponse-ExchangeAccountWalletsEntry"></a>

### ListExchangeWalletsResponse.ExchangeAccountWalletsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int32](#int32) |  |  |
| value | [ExchangeAccountWallets](#cryptowatch-broker-api-v2-ExchangeAccountWallets) |  |  |






<a name="cryptowatch-broker-api-v2-ListOrdersRequest"></a>

### ListOrdersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int32](#int32) |  | [optional] Limits the number of orders shown, if unset, the default is set to 200. |
| exchange_account_id | [string](#string) |  | [optional] If specified, only retrieve only orders belonging to the exchange account |
| exchange_id | [int32](#int32) |  | [optional] If specified, only retrieves orders for this exchange. (Could be found with the REST API under https://api.cryptowat.ch/exchanges) |
| market_id | [int32](#int32) |  | [optional] If specified, only retrieves orders for this market. (could be found with the REST API under https://api.cryptowat.ch/markets) |
| group_id | [string](#string) |  | [optional] If specified, only lists orders belonging to a given order group, if it exists. |
| start_timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | [optional] If specified, only retrieve orders that were executed after (strictly &gt;) |
| end_timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | [optional] If specified, only retrieve orders that were executed before (strictly &lt;) |
| offset | [int32](#int32) |  | [optional] Used for pagination |
| side | [OptionalSide](#cryptowatch-broker-api-v2-OptionalSide) |  | [optional] Could be either &#39;Buy&#39; or &#39;Sell&#39;. |
| type | [OptionalOrderType](#cryptowatch-broker-api-v2-OptionalOrderType) |  | [optional] Type of order to filter on |
| quote_asset_id | [int32](#int32) |  | [optional] The quote asset for this market. (mapping could be found with the REST API under https://api.cryptowat.ch/assets) |
| base_asset_id | [int32](#int32) |  | [optional] The base asset for this market. (mapping could be found with the REST API under https://api.cryptowat.ch/assets) |
| statuses | [Status](#cryptowatch-broker-api-v2-Status) | repeated | [optional] Filter by statuses |
| sorting | [Sorting](#cryptowatch-broker-api-v2-Sorting) |  | [optional] defaults to DESCENDING NOTE: Anything other than DESCENDING is disabled. |
| sort_on | [OrdersSortKey](#cryptowatch-broker-api-v2-OrdersSortKey) |  | [optional] defaults to CREATION_TIME NOTE: Anything other than CREATION_TIME is disabled. |






<a name="cryptowatch-broker-api-v2-ListOrdersResponse"></a>

### ListOrdersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders | [Order](#cryptowatch-broker-api-v2-Order) | repeated | list of orders that belong to the user |
| total | [int32](#int32) |  | Total field is not the number of all orders ever, unless there were no filter criteria in the request. It is also not the length of the orders in the response. It is the total number of possible results as defined by the filter criteria without pagination criteria. For example, a user has 10 orders, 7 of them belongs to BTC/USD market. If the request is: `ListOrdersRequest { limit=2 market_id=87 (BTC/USD) offset=3 }` `total` will be equal 7 and the `orders` field will contains 2 objects. NB: ListOrdersResponse can be used as a websocket response in a update of orders, in this specific case `total` will represent the amount of orders that have been updated. |






<a name="cryptowatch-broker-api-v2-ListPositionsRequest"></a>

### ListPositionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int32](#int32) |  | [optional] Limits the number of positions shown, if unset, the default is set to 200. |
| exchange_account_id | [string](#string) |  | [optional] If specified, only retrieves positions for this exchange account. |
| exchange_id | [int32](#int32) |  | [optional] If specified, only retrieves positions for this exchange. (Could be found with the REST API under https://api.cryptowat.ch/exchanges) |
| market_id | [int32](#int32) |  | [optional] If specified, only retrieves positions for this market. (could be found with the REST API under https://api.cryptowat.ch/markets) |
| timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | [optional] If specified, only retrieve positions that were opened after (strictly &gt;) to this Timestamp. |
| offset | [int32](#int32) |  | [optional] Used for pagination |






<a name="cryptowatch-broker-api-v2-ListPositionsResponse"></a>

### ListPositionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| positions | [Position](#cryptowatch-broker-api-v2-Position) | repeated | list of positions that belong to the user |
| total | [int32](#int32) |  | Total field is not the number of all positions ever, unless there were no filter criteria in the request. It is also not the length of the positions in the response. It is the total number of possible results as defined by the filter criteria without pagination criteria. For example, a user has 10 postions, 7 of them belongs to BTC/USD market. If the request is: ListPositionRequest { limit=2 market_id=87 (BTC/USD) offset=3 } `total` will be equal 7 and the `positions` field will contains 2 objects. |






<a name="cryptowatch-broker-api-v2-ListTradesRequest"></a>

### ListTradesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int32](#int32) |  | [optional] Limits the number of trades shown, if unset, the default is set to 200. |
| exchange_account_id | [string](#string) |  | [optional] If sepcified, only retrieves trades for this exchange account. |
| exchange_id | [int32](#int32) |  | [optional] If specified, only retrieves trades for this exchange. (Could be found with the REST API under https://api.cryptowat.ch/exchanges) |
| market_id | [int32](#int32) |  | [optional] If specified, only retrieves trades for this market. (could be found with the REST API under https://api.cryptowat.ch/markets) |
| start_timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | [optional] If specified, only retrieve trades that were executed after (strictly &gt;) |
| end_timestamp | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | [optional] If specified, only retrieve trades that were executed before (strictly &lt;) |
| offset | [int32](#int32) |  | [optional] Used for pagination |
| side | [OptionalSide](#cryptowatch-broker-api-v2-OptionalSide) |  | [optional] Could be either &#39;Buy&#39; or &#39;Sell&#39;. |
| quote_asset_id | [int32](#int32) |  | [optional] The quote asset for this market. |
| base_asset_id | [int32](#int32) |  | [optional] The base asset for this market. |
| consolidate | [bool](#bool) |  | [optional] Whether or not to consolidate trades per order |
| sorting | [Sorting](#cryptowatch-broker-api-v2-Sorting) |  | [optional] defaults to DESCENDING NOTE: Anything other than DESCENDING is disabled. |
| sort_on | [TradesSortKey](#cryptowatch-broker-api-v2-TradesSortKey) |  | [optional] defaults to TRADE_EXECUTION_TIME NOTE: Anything other than TRADE_EXECUTION_TIME is disabled. |






<a name="cryptowatch-broker-api-v2-ListTradesResponse"></a>

### ListTradesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trades | [Trade](#cryptowatch-broker-api-v2-Trade) | repeated | list of trades that belong to the user |
| total | [int32](#int32) |  | Total field is not the number of all trades ever, unless there were no filter criteria in the request. It is also not the length of the trades in the response. It is the total number of possible results as defined by the filter criteria without pagination criteria. For example, a user has 10 trades, 7 of them belongs to BTC/USD market. If the request is: `ListTradesRequest { limit=2g market_id=87 (BTC/USD) offset=3 }` `total` will be equal 7 and the `trades` field will contains 2 objects. NB: ListTradesResponse can be used as a websocket response in a update of trades, in this specific case `total` will represent the amount of trades that have been updated. |






<a name="cryptowatch-broker-api-v2-MFAAuthenticationResult"></a>

### MFAAuthenticationResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [MFAAuthenticationStatus](#cryptowatch-broker-api-v2-MFAAuthenticationStatus) |  | user&#39;s mfa authentication status with cryptowatch |






<a name="cryptowatch-broker-api-v2-Market"></a>

### Market



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| market_id | [int32](#int32) |  | The ID of the market the trade executed on (could be found with the REST API under https://api.cryptowat.ch/markets) |
| market_symbol | [string](#string) |  | The symbol for that market, ie. btcusd, etheur, ltcusd, etc. |
| base | [Asset](#cryptowatch-broker-api-v2-Asset) |  | The base asset for this market. |
| quote | [Asset](#cryptowatch-broker-api-v2-Asset) |  | The quote asset for this market. |






<a name="cryptowatch-broker-api-v2-MarketInformation"></a>

### MarketInformation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| market_id | [int32](#int32) |  | A cryptowatch market identifier (could be found with the REST API under https://api.cryptowat.ch/markets) |
| exchange | [Exchange](#cryptowatch-broker-api-v2-Exchange) |  | The exchange this market belongs to. |
| symbol | [string](#string) |  | Symbol is a unique upper-case string identifier for the underlying instrument. |
| tick_sizes | [TickSizes](#cryptowatch-broker-api-v2-TickSizes) |  | Specifies precision info for order size &amp; order price. |
| market_type | [MarketType](#cryptowatch-broker-api-v2-MarketType) |  | Differentiates between regular pair markets and futures markets. |
| contract_schema | [ContractSchema](#cryptowatch-broker-api-v2-ContractSchema) |  | Specifies the futures contract schema. Only applies to futures markets. |
| supports_margin | [bool](#bool) |  | Determines if the exchange supports margin and leverage options. |
| margin_actions | [MarginAction](#cryptowatch-broker-api-v2-MarginAction) | repeated | List of supported margin actions available atomically for this market. |
| leverage_levels | [LeverageLevels](#cryptowatch-broker-api-v2-LeverageLevels) |  | Specifies the set of allowed leverage levels. Only applies if the market supports margin trading. |
| lot_size | [string](#string) |  | Lot size is the size of a traded unit on a market. When a trade is done on an exchange, a lot is either bought or sold. The lot size is how many units of an instrument composes a lot. |
| order_types | [OrderType](#cryptowatch-broker-api-v2-OrderType) | repeated | Order types that this specific market supports. |
| supports_post_only | [bool](#bool) |  | Determines if this market supports post-only option to limit orders. |
| supports_fee_currency | [bool](#bool) |  | Determines if this market supports fee currency on order creation. |
| supports_conditional_close | [bool](#bool) |  | Determines if this market supports conditional closing orders. |
| time_in_force | [TimeInForce](#cryptowatch-broker-api-v2-TimeInForce) | repeated | Time in force options supported by this market. |
| supports_hidden | [bool](#bool) |  | Determines if this market supports hidden option to limit orders. |
| supports_reduce_only | [bool](#bool) |  | Determines if this market supports reduce-only option to margin orders. |
| supports_iceberg | [bool](#bool) |  | Determines if this market supports iceberg option to limit orders. |
| synthetic_order_types | [OrderType](#cryptowatch-broker-api-v2-OrderType) | repeated | Synthetic order types that this specific market supports. |






<a name="cryptowatch-broker-api-v2-MessageAcknowledgementResponse"></a>

### MessageAcknowledgementResponse
MessageAcknowledgementResponse is sent to acknowledge that broker has received and is currently processing a
request. It is only sent for specific requests, and it is only sent if the request has provided a request_id.
The specific requests for which broker will send an acknowledgement response are `CreateOrderRequest`, `CreateOrdersRequest`,
`ReplaceOrderRequest`, `CancelOrderRequest`, `CancelAllOrdersRequest`, `ClosePositionRequest`, `SettlePositionRequest`, 
`PlayOrdersRequest` and `PauseOrdersRequest`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| server_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | server_time is the current server time of the instance that served the user request |






<a name="cryptowatch-broker-api-v2-OptionalOrderType"></a>

### OptionalOrderType
OptionalOrderType helps cryotwatch determine if a user specified an order type.
Since the order types offered are enums, not sending an order type will default
to the default value 0. In this case, the request will be interpreted as the user
specified the order type that belongs to enum 0. To avoid the scenario above, 
OptionalOrderType is a wrapper around OrderType which will be evaluated if and only if valid flag is set.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| valid | [bool](#bool) |  | Valid needs to be set to true otherwise type will be ignored. |
| type | [OrderType](#cryptowatch-broker-api-v2-OrderType) |  | OrderType value |






<a name="cryptowatch-broker-api-v2-OptionalSide"></a>

### OptionalSide
OptionalSide helps cryotwatch determine if a user specified an order side.
Since the sides offered are enums, not sending an order type will default
to the default value 0. In this case, the request will be interpreted as the user
specified the order side that belongs to enum 0. To avoid the scenario above, 
OptionalSide is a wrapper around Side which will be evaluated if and only if valid flag is set.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| valid | [bool](#bool) |  | Valid needs to be set to true otherwise side will be ignored. |
| side | [Side](#cryptowatch-broker-api-v2-Side) |  | Side value |






<a name="cryptowatch-broker-api-v2-Order"></a>

### Order
Order defines the information for an order after it has been placed on the
exchange side. It is generally returned when we list open orders on an
exchange or after we place the order using the &#39;CreateOrder&#39; endpoint.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Cryptowatch identifier for the order, uniqueness is guaranteed by cryptowatch. This ID could be used to interact with the orders through the API. |
| external_id | [string](#string) |  | Exchange identifier for the order, uniqueness is ensured by the exchange itself. This ID could be used to interact with the orders through the API. |
| group_id | [string](#string) |  | Group Identifier for the order. Appears if the order is part of an OCO. |
| exchange_account_id | [string](#string) |  | The exchange account this order belongs to. |
| exchange | [Exchange](#cryptowatch-broker-api-v2-Exchange) |  | The exchange this order was placed on. |
| creation_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | Creation timestamp of the order. |
| side | [Side](#cryptowatch-broker-api-v2-Side) |  | Could be either &#39;Buy&#39; or &#39;Sell&#39;. |
| type | [OrderType](#cryptowatch-broker-api-v2-OrderType) |  | Could be any of the values defined in the &#39;OrderType&#39; enum. It defines how we this order was placed and usually which parameters were used (if a stop loss, we&#39;ll see the stop price information filled, etc.). |
| market | [Market](#cryptowatch-broker-api-v2-Market) |  | Denotes the market this order has been opened on. |
| status | [Status](#cryptowatch-broker-api-v2-Status) |  | The abstracted status of the order, reduced to the minimum subset supported by most exchanges. |
| leverage | [string](#string) |  | Leverage value for a margin order. |
| amount | [string](#string) |  | Amount of base asset to buy or sell. |
| amount_filled | [string](#string) |  | Amount filled for the order, appears if the order is partially filled. |
| current_stop | [string](#string) |  | If this is a trailing stop, defines the current stop price. |
| initial_stop | [string](#string) |  | Initial stop price as defined when the order was placed. |
| price | [string](#string) |  | Quote asset price to buy or sell the base asset. |
| stop_price | [string](#string) |  | The stop price for a Stop Market or Stop Limit order. |
| stop_limit_price | [string](#string) |  | The limit order execution price for `STOP_LIMIT` order. |
| trigger_price | [string](#string) |  | The trigger price for a `MARKET_IF_TOUCHED` order. |
| execution_price | [string](#string) |  | The limit order execution price for `LIMIT_IF_TOUCHED` order. |
| trailing_stop_market_offset | [TrailingOffsetPrice](#cryptowatch-broker-api-v2-TrailingOffsetPrice) |  | The trailing offset price for a `TRAILING_STOP_MARKET` order. |
| trailing_stop_limit_offset | [TrailingOffsetPrice](#cryptowatch-broker-api-v2-TrailingOffsetPrice) |  | The limit price offset for a `TRAILING_STOP_LIMIT` order. |
| oco_stop_price | [string](#string) |  | The stop price for an OCO order. |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | Timestamp defining when the order should be placed. |
| expire_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | Timestamp for the time the order should be canceled. |
| post_only | [bool](#bool) |  | Defines if the order should de deleted if it won&#39;t appear in the orderbook and fill immediately (default: false) |
| immediate_or_cancel | [bool](#bool) |  | Defines if the order should be immediate or cancel (default: false) |
| fill_or_kill | [bool](#bool) |  | Defines if the order should be a fill or kill order (default: false) |
| hidden | [bool](#bool) |  | Determines if the order should be hidden from the orderbook or not. (default: false) |
| fee_currency | [FeeCurrency](#cryptowatch-broker-api-v2-FeeCurrency) |  | The currency used to pay for the fees (could be base or quote). |
| closing_order | [Order](#cryptowatch-broker-api-v2-Order) |  | Contains the details for the closing order if the parent order fills fully. |
| reduce_only | [bool](#bool) |  | [optional] Defines if the order should be immediately canceled, if it would increase the size of an open position. (default: false) |
| iceberg_quantity | [string](#string) |  | An order type to trade in smaller predetermined quantities, called iceberg_quantity, in order to hide the total order quantity. |
| reference_price | [ReferencePrice](#cryptowatch-broker-api-v2-ReferencePrice) |  | [optional] Trigger price references are used in futures trading to denote the kind of price at which a trade action is to be taken, if supported for a market. (default: Default) |
| synthetic | [bool](#bool) |  | If true, order is a synthetic order |






<a name="cryptowatch-broker-api-v2-PauseOrdersRequest"></a>

### PauseOrdersRequest
PauseOrdersRequest expects either of the filters to be specified. If none of them are set, we pause all orders.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_ids | [string](#string) | repeated | providing an array of specific order IDs will only pause those specific orders |
| market_id | [int32](#int32) |  | providing the market_id will only pause orders on that market Could be found with the REST API under https://api.cryptowat.ch/exchanges. |
| exchange_id | [int32](#int32) |  | providing the exchange_id will only pause orders on that exchange Could be found with the REST API under https://api.cryptowat.ch/markets. |






<a name="cryptowatch-broker-api-v2-PauseOrdersResponse"></a>

### PauseOrdersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders | [Order](#cryptowatch-broker-api-v2-Order) | repeated | array of modified orders |
| all_successful | [bool](#bool) |  | boolean indicating whether all of the orders have been successfully paused |
| failed_ids | [PauseOrdersResponse.FailedIdsEntry](#cryptowatch-broker-api-v2-PauseOrdersResponse-FailedIdsEntry) | repeated | list of failed order ids (key) with the error/reason for the failure (value). |






<a name="cryptowatch-broker-api-v2-PauseOrdersResponse-FailedIdsEntry"></a>

### PauseOrdersResponse.FailedIdsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [ErrorResponse](#cryptowatch-broker-api-v2-ErrorResponse) |  |  |






<a name="cryptowatch-broker-api-v2-PlayOrdersRequest"></a>

### PlayOrdersRequest
PlayOrdersRequest expects either of the filters to be specified. If none of them are set, we play all orders.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_ids | [string](#string) | repeated | providing an array of specific order IDs will only play those specific orders |
| market_id | [int32](#int32) |  | providing the market_id will only play orders on that market Could be found with the REST API under https://api.cryptowat.ch/exchanges. |
| exchange_id | [int32](#int32) |  | providing the exchange_id will only play orders on that exchange Could be found with the REST API under https://api.cryptowat.ch/markets. |






<a name="cryptowatch-broker-api-v2-PlayOrdersResponse"></a>

### PlayOrdersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders | [Order](#cryptowatch-broker-api-v2-Order) | repeated | array of modified orders |
| all_successful | [bool](#bool) |  | boolean indicating whether all of the orders have been successfully played |
| failed_ids | [PlayOrdersResponse.FailedIdsEntry](#cryptowatch-broker-api-v2-PlayOrdersResponse-FailedIdsEntry) | repeated | list of failed order ids (key) with the error/reason for the failure (value). |






<a name="cryptowatch-broker-api-v2-PlayOrdersResponse-FailedIdsEntry"></a>

### PlayOrdersResponse.FailedIdsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [ErrorResponse](#cryptowatch-broker-api-v2-ErrorResponse) |  |  |






<a name="cryptowatch-broker-api-v2-Position"></a>

### Position



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Cryptowatch identifier for the position, uniqueness is guaranteed by cryptowatch. This ID could be used to interact with positions through the API. |
| external_id | [string](#string) |  | The ID of the position coming from the exchange. |
| exchange_account_id | [string](#string) |  | The exchange account for which this position belongs to. |
| exchange | [Exchange](#cryptowatch-broker-api-v2-Exchange) |  | The exchange for which this position is opened on. |
| market | [Market](#cryptowatch-broker-api-v2-Market) |  | The market for which this position is valid and open. |
| creation_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | The time at which the given position was created. |
| side | [Side](#cryptowatch-broker-api-v2-Side) |  | Could be either &#39;Buy&#39; or &#39;Sell&#39;. |
| average_price | [string](#string) |  | The average price for the position, made from all the orders contributing to this position. |
| amount_open | [string](#string) |  | Amount of that position still open after closing part of it with buy or sell orders. |
| amount_closed | [string](#string) |  | Amount that was closed by selling or buying assets, reducing the position size. |
| leverage | [string](#string) |  | The current leverage used to open this position. Can be any decimal value supported by the final exchange. For better compatibility, use the values returned by the GetExchangeSupportInformation endpoint. |
| profit_loss | [string](#string) |  | The profit loss for this position, determining the gain or loss incurred. |
| order_ids | [string](#string) | repeated | The list of order IDs (if this could be determined), contributing to this position. |
| trade_ids | [string](#string) | repeated | The list of trade IDs (if this could be determined), contributing to this position. |
| liquidation_price | [string](#string) |  | The price at which the position will get liquidated by the exchange. |
| lending_fee | [string](#string) |  | The fee paid to lend assets from other traders or the exchange itself when opening a margin position. |






<a name="cryptowatch-broker-api-v2-ReplaceOrderRequest"></a>

### ReplaceOrderRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| order_id | [string](#string) |  | Order ID of the original order |
| order | [CreateOrderRequest](#cryptowatch-broker-api-v2-CreateOrderRequest) |  | New order to replace the original with |
| exchange_id | [int32](#int32) |  | [optional] The exchange ID that the order is on. Because order IDs are not guaranteed to be unique across exchanges, there is a very small change that a user could have two orders with the same order ID, and end up replacing the wrong order. Providing the exchange_id removes that possibility, since order IDs are unique within an exchange. Could be found with the REST API under https://api.cryptowat.ch/exchanges. |






<a name="cryptowatch-broker-api-v2-ReplaceOrderResponse"></a>

### ReplaceOrderResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| canceled | [bool](#bool) |  | Some exchanges don&#39;t have an atomic replace function, so broker manually cancels the original order and then places the new one. Because of this, it&#39;s possible the original order was canceled but the new order failed to place for some reason. The canceled field indicates whether the original order was cancelled or not. |
| orders | [Order](#cryptowatch-broker-api-v2-Order) | repeated | affected orders |






<a name="cryptowatch-broker-api-v2-RevokeAllExchangeAPIKeysRequest"></a>

### RevokeAllExchangeAPIKeysRequest
RevokeAllExchangeAPIKeysRequest is an empty request






<a name="cryptowatch-broker-api-v2-RevokeAllExchangeAPIKeysResponse"></a>

### RevokeAllExchangeAPIKeysResponse
RevokeAllExchangeAPIKeysResponse is an empty response






<a name="cryptowatch-broker-api-v2-RevokeExchangeAPIKeyRequest"></a>

### RevokeExchangeAPIKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_account_id | [string](#string) |  | the id of the exchange account from which to remove API keys |






<a name="cryptowatch-broker-api-v2-RevokeExchangeAPIKeyResponse"></a>

### RevokeExchangeAPIKeyResponse
RevokeExchangeAPIKeyResponse is an empty response






<a name="cryptowatch-broker-api-v2-SettlePositionRequest"></a>

### SettlePositionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| position_id | [string](#string) |  | the exchange&#39;s identifier for the position to settle |
| exchange_id | [int32](#int32) |  | [optional] The exchange ID that the position is on. Because position IDs are not guaranteed to be unique across exchanges, there is a very small change that a user could have two positions with the same position ID, and end up closing the wrong position. Providing the exchange_id removes that possibility, since position IDs are unique within an exchange. (Could be found with the REST API under https://api.cryptowat.ch/exchanges) |






<a name="cryptowatch-broker-api-v2-SettlePositionResponse"></a>

### SettlePositionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| orders | [Order](#cryptowatch-broker-api-v2-Order) | repeated | The orders created as a result of settling the position. Not guaranteed to be filled, some exchanges do not include this in the response. |






<a name="cryptowatch-broker-api-v2-SubscribeToExchangeAccountsRequest"></a>

### SubscribeToExchangeAccountsRequest
Subscribe to receive all balance/order/trade/position changes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_ids | [int32](#int32) | repeated | by specifying an exchange_id, a subscription for each exchange account the user has on the given exchange will started. Could be found with the REST API under https://api.cryptowat.ch/exchanges. |
| exchange_account_ids | [string](#string) | repeated | by specifying an exchange_account_id, a subscription for that particular exchange account will be started. |






<a name="cryptowatch-broker-api-v2-SubscribeToExchangeAccountsResponse"></a>

### SubscribeToExchangeAccountsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| exchange_account_ids | [string](#string) | repeated | which exchange accounts where subscribed to. |






<a name="cryptowatch-broker-api-v2-Tab"></a>

### Tab



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | valid names are &#34;open-orders&#34;, &#34;holdings&#34;, &#34;positions&#34;, &#34;historical-orders&#34;, &#34;trades&#34; |
| visibility | [bool](#bool) |  |  |
| columns | [Column](#cryptowatch-broker-api-v2-Column) | repeated |  |






<a name="cryptowatch-broker-api-v2-TickSizes"></a>

### TickSizes
The min, max, and tick_size specify the set of valid inputs for each of the
base and quote (aka order size and order price) when creating an order.
min specifies the minimum value, max specifies the max value, and tick_size
specifies the size of the increment.

As an example, if a market specifies a base_min of 10, a base_max of 12,
and a base_tick_size of 0.5, then only the values [10, 10.5, 11, 11.5, 12]
may be used to specify the order amount.

The min, max and tick_size are unique per market. A market belongs to an exchange so 
the values for a market will differ between exchanges. Some exchanges do not offer
this data and you may get -1 for the value indicating the exchange does not have this data.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_min | [string](#string) |  | base asset minimum order size |
| base_max | [string](#string) |  | base asset maximum order size |
| base_tick_size | [string](#string) |  | base asset tick size |
| quote_min | [string](#string) |  | quote asset minimum price |
| quote_max | [string](#string) |  | quote asset maximum price |
| quote_tick_size | [string](#string) |  | quote asset tick size |






<a name="cryptowatch-broker-api-v2-Trade"></a>

### Trade
Trade denotes the information after an order has been filled and becomes a
trade.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Cryptowatch identifier for the trade, uniqueness is guaranteed by cryptowatch. This ID could be used to interact with the trades through the API. empty (zero value) if is_consolidated is true |
| external_id | [string](#string) |  | The ID of the trade, coming from the exchange itself or assigned by Cryptowatch. empty (zero value) if is_consolidated is true |
| order_external_id | [string](#string) |  | The ID of the order, coming from the exchange itself or assigned by Cryptowatch. |
| exchange_account_id | [string](#string) |  | The exchange account this trade belongs to. |
| exchange | [Exchange](#cryptowatch-broker-api-v2-Exchange) |  | The exchange this trade was executed on. |
| market | [Market](#cryptowatch-broker-api-v2-Market) |  | The market for which this trade was executed. |
| execution_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | The time at which the trade was executed, in millisecond precision. empty (zero value) if is_consolidated is true |
| price | [string](#string) |  | Quote asset price that was used to buy or sell the base asset. |
| amount | [string](#string) |  | Amount of base asset that was bought/sold. |
| side | [Side](#cryptowatch-broker-api-v2-Side) |  | Could be either &#39;Buy&#39; or &#39;Sell&#39;. |
| fee | [string](#string) |  | The fee paid for this trade. |
| fee_currency | [FeeCurrency](#cryptowatch-broker-api-v2-FeeCurrency) |  | The currency used to pay for the fees (could be base or quote). |
| is_consolidated | [bool](#bool) |  | A consolidated trade aggregates all trades under the same order ID. It acts as if the related order created a single trade: - fee(s) are added up - price(s) are weighted averaged, weighted by the amount - amount(s) are added up - execution_time is set to its most recent constituent trades execution_time. - id and external_id are ignored (zero&#39;d value) - other attributes (market, exchange, exchange_account_id, order_external_id, side, fee_currency) remain untouched |
| total | [string](#string) |  | total = price * amount |






<a name="cryptowatch-broker-api-v2-TradingWebsocketRequest"></a>

### TradingWebsocketRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  | Client specified opaque string that we echo back on the response. This is so that the client can match responses to requests when making multiple in parallel. |
| get_exchange_information_request | [GetExchangeInformationRequest](#cryptowatch-broker-api-v2-GetExchangeInformationRequest) |  |  |
| get_market_information_request | [GetMarketInformationRequest](#cryptowatch-broker-api-v2-GetMarketInformationRequest) |  |  |
| get_all_markets_information_request | [GetAllMarketsInformationRequest](#cryptowatch-broker-api-v2-GetAllMarketsInformationRequest) |  |  |
| create_order_request | [CreateOrderRequest](#cryptowatch-broker-api-v2-CreateOrderRequest) |  |  |
| create_orders_request | [CreateOrdersRequest](#cryptowatch-broker-api-v2-CreateOrdersRequest) |  |  |
| replace_order_request | [ReplaceOrderRequest](#cryptowatch-broker-api-v2-ReplaceOrderRequest) |  |  |
| cancel_order_request | [CancelOrderRequest](#cryptowatch-broker-api-v2-CancelOrderRequest) |  |  |
| cancel_all_orders_request | [CancelAllOrdersRequest](#cryptowatch-broker-api-v2-CancelAllOrdersRequest) |  |  |
| get_order_request | [GetOrderRequest](#cryptowatch-broker-api-v2-GetOrderRequest) |  |  |
| list_orders_request | [ListOrdersRequest](#cryptowatch-broker-api-v2-ListOrdersRequest) |  |  |
| list_trades_request | [ListTradesRequest](#cryptowatch-broker-api-v2-ListTradesRequest) |  |  |
| list_positions_request | [ListPositionsRequest](#cryptowatch-broker-api-v2-ListPositionsRequest) |  |  |
| list_exchange_wallets_request | [ListExchangeWalletsRequest](#cryptowatch-broker-api-v2-ListExchangeWalletsRequest) |  |  |
| close_position_request | [ClosePositionRequest](#cryptowatch-broker-api-v2-ClosePositionRequest) |  |  |
| settle_position_request | [SettlePositionRequest](#cryptowatch-broker-api-v2-SettlePositionRequest) |  |  |
| subscribe_to_exchange_accounts_request | [SubscribeToExchangeAccountsRequest](#cryptowatch-broker-api-v2-SubscribeToExchangeAccountsRequest) |  |  |
| add_exchange_api_key_request | [AddExchangeAPIKeyRequest](#cryptowatch-broker-api-v2-AddExchangeAPIKeyRequest) |  |  |
| revoke_exchange_api_key_request | [RevokeExchangeAPIKeyRequest](#cryptowatch-broker-api-v2-RevokeExchangeAPIKeyRequest) |  |  |
| revoke_all_exchange_api_keys_request | [RevokeAllExchangeAPIKeysRequest](#cryptowatch-broker-api-v2-RevokeAllExchangeAPIKeysRequest) |  |  |
| list_exchange_accounts_request | [ListExchangeAccountsRequest](#cryptowatch-broker-api-v2-ListExchangeAccountsRequest) |  |  |
| get_user_settings_request | [GetUserSettingsRequest](#cryptowatch-broker-api-v2-GetUserSettingsRequest) |  |  |
| update_user_settings_request | [UpdateUserSettingsRequest](#cryptowatch-broker-api-v2-UpdateUserSettingsRequest) |  |  |
| get_server_time_request | [GetServerTimeRequest](#cryptowatch-broker-api-v2-GetServerTimeRequest) |  |  |
| get_fee_schedule_request | [GetFeeScheduleRequest](#cryptowatch-broker-api-v2-GetFeeScheduleRequest) |  |  |
| list_available_filters_request | [ListAvailableFiltersRequest](#cryptowatch-broker-api-v2-ListAvailableFiltersRequest) |  |  |
| play_orders_request | [PlayOrdersRequest](#cryptowatch-broker-api-v2-PlayOrdersRequest) |  |  |
| pause_orders_request | [PauseOrdersRequest](#cryptowatch-broker-api-v2-PauseOrdersRequest) |  |  |






<a name="cryptowatch-broker-api-v2-TradingWebsocketResponse"></a>

### TradingWebsocketResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request_id | [string](#string) |  |  |
| get_exchange_information_response | [GetExchangeInformationResponse](#cryptowatch-broker-api-v2-GetExchangeInformationResponse) |  |  |
| get_market_information_response | [GetMarketInformationResponse](#cryptowatch-broker-api-v2-GetMarketInformationResponse) |  |  |
| get_all_markets_information_response | [GetAllMarketsInformationResponse](#cryptowatch-broker-api-v2-GetAllMarketsInformationResponse) |  |  |
| create_order_response | [CreateOrderResponse](#cryptowatch-broker-api-v2-CreateOrderResponse) |  |  |
| create_orders_response | [CreateOrdersResponse](#cryptowatch-broker-api-v2-CreateOrdersResponse) |  |  |
| cancel_order_response | [CancelOrderResponse](#cryptowatch-broker-api-v2-CancelOrderResponse) |  |  |
| cancel_all_orders_response | [CancelAllOrdersResponse](#cryptowatch-broker-api-v2-CancelAllOrdersResponse) |  |  |
| get_order_response | [GetOrderResponse](#cryptowatch-broker-api-v2-GetOrderResponse) |  |  |
| list_orders_response | [ListOrdersResponse](#cryptowatch-broker-api-v2-ListOrdersResponse) |  |  |
| list_trades_response | [ListTradesResponse](#cryptowatch-broker-api-v2-ListTradesResponse) |  |  |
| list_positions_response | [ListPositionsResponse](#cryptowatch-broker-api-v2-ListPositionsResponse) |  |  |
| list_exchange_wallets_response | [ListExchangeWalletsResponse](#cryptowatch-broker-api-v2-ListExchangeWalletsResponse) |  |  |
| close_position_response | [ClosePositionResponse](#cryptowatch-broker-api-v2-ClosePositionResponse) |  |  |
| subscribe_to_exchange_accounts_response | [SubscribeToExchangeAccountsResponse](#cryptowatch-broker-api-v2-SubscribeToExchangeAccountsResponse) |  |  |
| add_exchange_api_key_response | [AddExchangeAPIKeyResponse](#cryptowatch-broker-api-v2-AddExchangeAPIKeyResponse) |  |  |
| revoke_exchange_api_key_response | [RevokeExchangeAPIKeyResponse](#cryptowatch-broker-api-v2-RevokeExchangeAPIKeyResponse) |  |  |
| revoke_all_exchange_api_keys_response | [RevokeAllExchangeAPIKeysResponse](#cryptowatch-broker-api-v2-RevokeAllExchangeAPIKeysResponse) |  |  |
| list_exchange_accounts_response | [ListExchangeAccountsResponse](#cryptowatch-broker-api-v2-ListExchangeAccountsResponse) |  |  |
| error_response | [ErrorResponse](#cryptowatch-broker-api-v2-ErrorResponse) |  |  |
| replace_order_response | [ReplaceOrderResponse](#cryptowatch-broker-api-v2-ReplaceOrderResponse) |  |  |
| get_user_settings_response | [GetUserSettingsResponse](#cryptowatch-broker-api-v2-GetUserSettingsResponse) |  |  |
| update_user_settings_response | [UpdateUserSettingsResponse](#cryptowatch-broker-api-v2-UpdateUserSettingsResponse) |  |  |
| settle_position_response | [SettlePositionResponse](#cryptowatch-broker-api-v2-SettlePositionResponse) |  |  |
| get_server_time_response | [GetServerTimeResponse](#cryptowatch-broker-api-v2-GetServerTimeResponse) |  |  |
| exchange_connection_update | [ExchangeConnectionUpdate](#cryptowatch-broker-api-v2-ExchangeConnectionUpdate) |  |  |
| acknowledgement | [MessageAcknowledgementResponse](#cryptowatch-broker-api-v2-MessageAcknowledgementResponse) |  |  |
| mfa_authentication_result | [MFAAuthenticationResult](#cryptowatch-broker-api-v2-MFAAuthenticationResult) |  |  |
| get_fee_schedule_response | [GetFeeScheduleResponse](#cryptowatch-broker-api-v2-GetFeeScheduleResponse) |  |  |
| list_available_filters_response | [ListAvailableFiltersResponse](#cryptowatch-broker-api-v2-ListAvailableFiltersResponse) |  |  |
| play_orders_response | [PlayOrdersResponse](#cryptowatch-broker-api-v2-PlayOrdersResponse) |  |  |
| pause_orders_response | [PauseOrdersResponse](#cryptowatch-broker-api-v2-PauseOrdersResponse) |  |  |






<a name="cryptowatch-broker-api-v2-TrailingOffsetPrice"></a>

### TrailingOffsetPrice
TrailingOffsetPrice defines price offset to be used for Trailing Stop Limit and Trailing
Stop Market orders.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [TrailingOffsetType](#cryptowatch-broker-api-v2-TrailingOffsetType) |  |  |
| value | [string](#string) |  |  |






<a name="cryptowatch-broker-api-v2-UpdateUserSettingsRequest"></a>

### UpdateUserSettingsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| settings | [UserSettings](#cryptowatch-broker-api-v2-UserSettings) |  | the settings to update in cryptowatch |






<a name="cryptowatch-broker-api-v2-UpdateUserSettingsResponse"></a>

### UpdateUserSettingsResponse







<a name="cryptowatch-broker-api-v2-UserSettings"></a>

### UserSettings



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_timeout_seconds | [int32](#int32) |  | Session timeout in seconds define how frequently user will need to reauthenticate current session |
| tabs | [Tab](#cryptowatch-broker-api-v2-Tab) | repeated | Allows you to change the order and related properties of the tab |
| privacy_mode | [bool](#bool) |  | Privacy mode toggle allows a user to hide the value of their holdings and their PnL. This is recommended when working in public so that your account details are not visible to people glancing at your screen. |
| consolidating_fills | [bool](#bool) |  | Consolidating fills toggle allows you to toggle between seeing order fills with all the constituent partial fills, or as one whole fill. For example, an order to buy 10 ETH/USD might have gotten filled over 20 orders of many sizes at one price, but a user might just want to see that they bought 10 ETH/USD in one order and the average price. |
| show_all_pairs_positions | [bool](#bool) |  | This is used to persist setting if we want to show positions for all pairs in cockpit |
| show_all_pairs_open_order | [bool](#bool) |  | This is used to persist setting if we want to show open orders for all pairs in cockpit |






<a name="cryptowatch-broker-api-v2-Wallets"></a>

### Wallets



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| balances | [Wallets.BalancesEntry](#cryptowatch-broker-api-v2-Wallets-BalancesEntry) | repeated | A map of arbitrary wallet strings (ie. Spot, Margin, etc.) linked to their list of account balances per assets. |






<a name="cryptowatch-broker-api-v2-Wallets-BalancesEntry"></a>

### Wallets.BalancesEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [Balances](#cryptowatch-broker-api-v2-Balances) |  |  |





 


<a name="cryptowatch-broker-api-v2-APIKeyStatus"></a>

### APIKeyStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 | The key is valid and we can contact the exchange successfully. |
| INVALID_KEY | 1 | The key pair is invalid, we cannot communicate with the exchange. |
| UNPRIVILEGED | 2 | The key pair is valid but has insufficient privileges set from the exchange website. |
| NOT_EXCLUSIVE | 3 | The key pair is potentially used elsewhere because we&#39;re getting an elevated amount |



<a name="cryptowatch-broker-api-v2-ContractSchema"></a>

### ContractSchema


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN_CONTRACT_SCHEMA | 0 | contract schema isn&#39;t known |
| VANILLA | 1 | contract that is settled in its quote asset. |
| INVERSE | 2 | contract that is settled in its base asset. |
| QUANTO | 3 | contract that is settled in an asset that is neither its base or quote. |



<a name="cryptowatch-broker-api-v2-ExchangeStatus"></a>

### ExchangeStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | exchange status is not known |
| AVAILABLE | 1 | exchange is up and running - user should be able to make requests, provided API keys are valid |
| DEGRADED | 2 | exchange is seeing degraded performance - user should be able to make requests, but they might take longer than usual or fail |
| UNAVAILABLE | 3 | exchange is currently unavailable - user should expect calls to fail |



<a name="cryptowatch-broker-api-v2-FeeCurrency"></a>

### FeeCurrency


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 |  |
| BASE | 1 |  |
| QUOTE | 2 |  |



<a name="cryptowatch-broker-api-v2-MFAAuthenticationStatus"></a>

### MFAAuthenticationStatus


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN_MFA_AUTH_STATUS | 0 |  |
| AUTHENTICATED | 1 | AUTHENTICATED is returned when the user is authenticated with MFA. |
| MFA_NOT_CONFIGURED | 2 | MFA_NOT_CONFIGURED is returned when the user does not have MFA configured whatsoever. |
| MFA_EXPIRED | 3 | MFA_EXPIRED is returned when the user has MFA, but the authentication has expired. In this case, we require them to authenticate again. |



<a name="cryptowatch-broker-api-v2-MarginAction"></a>

### MarginAction


| Name | Number | Description |
| ---- | ------ | ----------- |
| POSITION_CLOSE | 0 |  |
| POSITION_SETTLE | 1 |  |



<a name="cryptowatch-broker-api-v2-MarketType"></a>

### MarketType


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN_MARKET_TYPE | 0 | market type isn&#39;t known |
| PAIR | 1 | regular pair market |
| FUTURES | 2 | futures market |



<a name="cryptowatch-broker-api-v2-OrderType"></a>

### OrderType


| Name | Number | Description |
| ---- | ------ | ----------- |
| MARKET | 0 |  |
| LIMIT | 1 |  |
| STOP_MARKET | 2 |  |
| STOP_LIMIT | 3 |  |
| MARKET_IF_TOUCHED | 4 | previously take_profit |
| LIMIT_IF_TOUCHED | 5 | previously take_profit_limit |
| TRAILING_STOP_MARKET | 6 |  |
| TRAILING_STOP_LIMIT | 7 |  |
| SETTLE_POSITION | 8 |  |
| ONE_CANCELS_OTHER | 9 |  |
| ICEBERG | 10 |  |
| TRIGGER_ENTRY | 11 |  |



<a name="cryptowatch-broker-api-v2-OrdersSortKey"></a>

### OrdersSortKey


| Name | Number | Description |
| ---- | ------ | ----------- |
| ORDER_CREATION_TIME | 0 |  |
| ORDER_STATUS | 1 |  |
| ORDER_AMOUNT | 2 |  |
| ORDER_PRICE | 3 |  |
| ORDER_TRIGGER_PRICE | 4 |  |
| ORDER_MARKET | 5 |  |
| ORDER_TYPE | 6 |  |
| ORDER_SIDE | 7 |  |



<a name="cryptowatch-broker-api-v2-ReferencePrice"></a>

### ReferencePrice


| Name | Number | Description |
| ---- | ------ | ----------- |
| DEFAULT | 0 | Default type for an exchange |
| MARK | 1 | Mark Price is Index price &#43; 30 seconds EMA (protection against freak trades) Note that mark is first option because it is also the default value if nothing is passed |
| INDEX | 2 | Index Price is the underlying benchmark which is based on collective prices from different exchange. |
| LTP | 3 | Last traded price for contract |



<a name="cryptowatch-broker-api-v2-Side"></a>

### Side


| Name | Number | Description |
| ---- | ------ | ----------- |
| SELL | 0 |  |
| BUY | 1 |  |



<a name="cryptowatch-broker-api-v2-Sorting"></a>

### Sorting


| Name | Number | Description |
| ---- | ------ | ----------- |
| DESCENDING | 0 |  |
| ASCENDING | 1 |  |



<a name="cryptowatch-broker-api-v2-Status"></a>

### Status


| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN_STATUS | 0 | Unknown status. |
| NEW | 1 | The order has been accepted by the exchange engine. |
| OPEN | 2 | The order has been accepted by the engine and posted to the orderbook. |
| CANCELED | 3 | The order was canceled by the user. |
| FILLED | 4 | The order has been completed and filled entirely. |
| PARTIALLY_FILLED | 5 | The order has been partially filled. |
| EXPIRED | 6 | The order has expired (if it was set with an expiry date). |
| REJECTED | 7 | The order was not accepted by the engine and got rejected. |
| PAUSED | 8 | The order has been paused and ready to be restarted. |



<a name="cryptowatch-broker-api-v2-TimeInForce"></a>

### TimeInForce


| Name | Number | Description |
| ---- | ------ | ----------- |
| GOOD_TIL_CANCELED | 0 | order must be either filled or manually cancelled |
| IMMEDIATE_OR_CANCEL | 1 | order is partially filled as much as possible immediately, all remaining is cancelled |
| FILL_OR_KILL | 2 | entire order must fill entirely immediately or it&#39;s cancelled |
| GOOD_TILL_DATE | 3 | order is good until a specified datetime |
| PLACE_AT | 4 | order is good after a specified datetime |



<a name="cryptowatch-broker-api-v2-TradesSortKey"></a>

### TradesSortKey


| Name | Number | Description |
| ---- | ------ | ----------- |
| TRADE_EXECUTION_TIME | 0 |  |
| TRADE_EXTERNAL_ID | 1 |  |
| TRADE_ORDER_ID | 2 |  |
| TRADE_AMOUNT | 3 |  |
| TRADE_PRICE | 4 |  |
| TRADE_MARKET | 5 |  |
| TRADE_FEE | 6 |  |
| TRADE_SIDE | 7 |  |



<a name="cryptowatch-broker-api-v2-TrailingOffsetType"></a>

### TrailingOffsetType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ABSOLUTE | 0 |  |
| RELATIVE | 1 |  |
| RELATIVE_PERCENT | 2 |  |



<a name="cryptowatch-broker-api-v2-UserConnectionStatus"></a>

### UserConnectionStatus
UserConnectionStatus is user specific connection status to an exchange

| Name | Number | Description |
| ---- | ------ | ----------- |
| USER_CONNECTION_UNKNOWN | 0 | user connection to the exchange is unknown. Cryptowatch can not predetermine if the user calls to the exchange will fail or succeed. This uncertainty occurs if the user does not have a session with cryptowatch so cryptowatch can&#39;t keep track of their interactions to determine if their connection to the exchange is well. |
| USER_CONNECTION_AVAILABLE | 1 | user connection to the exchange is working well as expected. No interuption or errors occured with the user flow and the exchange is responding properly for all of the user&#39;s specific trading requests. |
| USER_CONNECTION_DEGRADED | 2 | user connection to the exchange is unstable and have a higher probability of resulting in errors for operations performed on the exchange for the specific user. This status can occur for any inconsistent errors occurring while the user perform their trading operations on the exchange. Example: if the user is rate limited by the exchange then their connection is degraded as some calls may error out depending on how the exchange implements their rate limit. |
| USER_CONNECTION_UNAVAILABLE | 3 | the user connection to the exchange is unavailable and all calls made to the exchange is expected to fail for the user. If this is the status then the user should avoid making trading operations to the exchange until the status is updated to a better value. |


 

 


<a name="cryptowatch-broker-api-v2-TradingService"></a>

### TradingService
TradingService is a GRPC Service that handles all trading related operations
accross multiple exchanges in a uniformed format.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetServerTime | [GetServerTimeRequest](#cryptowatch-broker-api-v2-GetServerTimeRequest) | [GetServerTimeResponse](#cryptowatch-broker-api-v2-GetServerTimeResponse) | GetServerTime returns the current server time for the trading server. This can be used to test connectivity to the trading service or this can be used as an application level ping to keep a websocket connection alive. |
| GetExchangeInformation | [GetExchangeInformationRequest](#cryptowatch-broker-api-v2-GetExchangeInformationRequest) | [GetExchangeInformationResponse](#cryptowatch-broker-api-v2-GetExchangeInformationResponse) | GetExchangeInformation returns all exchange information for each exchange. example use cases: - Determine what fields are needed to authenticate with the exchange. - Find the exchanges the user is registred with. - List the exchanges available with their ids for the user to use in the API. |
| GetMarketInformation | [GetMarketInformationRequest](#cryptowatch-broker-api-v2-GetMarketInformationRequest) | [GetMarketInformationResponse](#cryptowatch-broker-api-v2-GetMarketInformationResponse) | GetMarketInformation allows user to get all trading information for a specific market. Two Assets trade on a given market; a market belongs to an exchange. Example of an asset (Bitcoin). Example of a market (Bitcoin/Ethereum on kraken exchange). example use cases: - determine trading ticks for a market - determine if a market supports margin |
| GetAllMarketsInformation | [GetAllMarketsInformationRequest](#cryptowatch-broker-api-v2-GetAllMarketsInformationRequest) | [GetAllMarketsInformationResponse](#cryptowatch-broker-api-v2-GetAllMarketsInformationResponse) | GetAllMarketsInformation allows user to get all trading information for all markets supported on an exchange. example use cases: - find supported markets for an exchange - trading info for all markets on an exchange |
| CreateOrder | [CreateOrderRequest](#cryptowatch-broker-api-v2-CreateOrderRequest) | [CreateOrderResponse](#cryptowatch-broker-api-v2-CreateOrderResponse) | CreateOrder allows a user to place an order for a market; Creating an order does not mean a trade will occur, an order is simply the user&#39;s intent to make the trade. When an order is filled, a formal trade is made. Since an order can be partially filled, this can result in multiple formal trades created to fulfill a single order. Order can have many different types: limit, market, etc. Synthetic order types are order types that cryptowatch manages for the user for order types that aren&#39;t supported natively on the exchange. Example, Iceberg order type may not be available on a specific exchange but we may still support it and allow the user to create that order on the exchange that does not support it; Cryptowatch makes this feature possible by handling the logistics of the iceberg order to allow the user to place their order on their desired exchange. But it is possible to have a specific order type that is supported by this exchange and can also be synthetically created by Cryptowatch. Example, Iceberg order are supported natively by Binance and also by Cryptowach synthetic orders feature. The following are the current synthetic order types and the exchanges cryptowatch allows the synthetic order type on: __Iceberg__ __Order__ __Type__: [kraken, coinbase, ftx, kucoin, bitfinex, binance, poloniex] |
| CreateOrders | [CreateOrdersRequest](#cryptowatch-broker-api-v2-CreateOrdersRequest) | [CreateOrdersResponse](#cryptowatch-broker-api-v2-CreateOrdersResponse) | CreateOrders allows a user to create a batch of orders. Each order can have different a market on different exchanges and each order will be created exactly as does CreateOrder method permits the user. CreateOrders failures are handled as follow: if an order creation failed in the middle of the process, the rest of the orders are aborted and an error is returned to the user. Example: the user sent 3 orders to be created and the failure occurred on the second order creation, the third order will not be placed and only the first order is placed successfully. |
| ReplaceOrder | [ReplaceOrderRequest](#cryptowatch-broker-api-v2-ReplaceOrderRequest) | [ReplaceOrderResponse](#cryptowatch-broker-api-v2-ReplaceOrderResponse) | ReplaceOrder allows a user to replace an order for a different one. This method can be used in a variety of use cases: - replace an entire order to a different order - updating some values for an order such as amount, price, etc. Some exchanges support replacing an order natively but some do not. For the exchanges that do not support replacing an order, we cancel the original order and replace it by creating a new one with the updated values provided in the ReplaceOrderRequest. This method may result in multiple orders being returned since some orders will cause multiple orders to be returned from the exchange. Example: Order type `ONE_CANCELS_OTHER` will return multiple orders from the exchange. The following exchanges __do not__ support Replaceorder __natively__: [kraken, binance, bitmex, bitstamp, bittrex, coinbase, ftx hitbtc, huboi, kucoin, okex] ReplaceOrder possible failures: for exchanges that do not support replace order natively, we may cancel the old order and fail to create the new one. The user will know if the original order was cancelled part of the response. |
| CancelOrder | [CancelOrderRequest](#cryptowatch-broker-api-v2-CancelOrderRequest) | [CancelOrderResponse](#cryptowatch-broker-api-v2-CancelOrderResponse) | CancelOrder allows a user to cancel an existing open order. If the order requested to be cancelled is one of the synthetic orders supported by cryptowatch, cancelling that order will cause cancelling the rest of the orders for that synthetic order. |
| CancelAllOrders | [CancelAllOrdersRequest](#cryptowatch-broker-api-v2-CancelAllOrdersRequest) | [CancelAllOrdersResponse](#cryptowatch-broker-api-v2-CancelAllOrdersResponse) | CancelAllOrders allows a user to batch cancel their open orders. A user may choose to provide the optional filters listed in CancelAllOrdersRequest to narrow down the orders they would like to cancel. Example, cancelling all orders on a specific exchange instead of all their orders on all exchanges. CancelAllOrders failures are handled the same way CreateOrders failures are handled. If a failure occurred in the middle of the process, the remaining orders are not cancelled. However, if a failure occurred because an order had been cancelled previously on the exchange, we ignore that failure and continue the process. |
| GetOrder | [GetOrderRequest](#cryptowatch-broker-api-v2-GetOrderRequest) | [GetOrderResponse](#cryptowatch-broker-api-v2-GetOrderResponse) | GetOrder allows the user to fetch an existing order&#39;s information. |
| ListOrders | [ListOrdersRequest](#cryptowatch-broker-api-v2-ListOrdersRequest) | [ListOrdersResponse](#cryptowatch-broker-api-v2-ListOrdersResponse) | ListOrders allows the user to list their existing orders. The user may choose any of the filters provided in ListOrdersRequest to narrow down the list returned. |
| ListTrades | [ListTradesRequest](#cryptowatch-broker-api-v2-ListTradesRequest) | [ListTradesResponse](#cryptowatch-broker-api-v2-ListTradesResponse) | ListTrades allows the user to list all of their trades. The user may choose to provide any of the filters provided in ListTradesRequest to narrow down the list returned. Temporary workflow: Currently, the total value is not the user&#39;s total number of trades but rather the number of trades returned. possible scenario: A user may not see the most up-to-date trades that they have with the exchange and that is due to possible failures occurring in the process of syncing the user trades with the exchange. Failures can be the exchange not returning trades, exchange has issues when cryptowatch requests to sync user trades, exchange has downtime etc. |
| ListPositions | [ListPositionsRequest](#cryptowatch-broker-api-v2-ListPositionsRequest) | [ListPositionsResponse](#cryptowatch-broker-api-v2-ListPositionsResponse) | ListPositions allows the user to list their positions. The user may choose to provide any of the filters provided in ListPositionsRequest to narrow down the list returned. possible scenario: A user may not see the most up-to-date positions that they have with the exchange and that is due to possible failures occurring in the process of syncing the user positions with the exchange. Failures can be the exchange not returning positions, exchange has issues when cryptowatch requests to sync user positions, exchange has downtime etc. |
| ListAvailableFilters | [ListAvailableFiltersRequest](#cryptowatch-broker-api-v2-ListAvailableFiltersRequest) | [ListAvailableFiltersResponse](#cryptowatch-broker-api-v2-ListAvailableFiltersResponse) | ListAvailableFilters allows user to get their current available filters based on their orders/trades/positions. The filters are user specific since the data returned is based on the user&#39;s data. For example, if user has only Ethereum and Bictoin orders, they should only get the set of (ETH, BTC) for orders filters. |
| ListExchangeWallets | [ListExchangeWalletsRequest](#cryptowatch-broker-api-v2-ListExchangeWalletsRequest) | [ListExchangeWalletsResponse](#cryptowatch-broker-api-v2-ListExchangeWalletsResponse) | ListExchangeWallets allows user to get their current balances for all of their registered exchanges or the balances narrowed to the filters provided in ListExchangeWalletsRequest. possible scenario: A user may not see the most up-to-date balance that they have with the exchange and that is due to possible failures occurring in the process of syncing the user balances with the exchange. Failures can be the exchange not returning balances, exchange has issues when cryptowatch requests to sync user balances, exchange has downtime etc. |
| ClosePosition | [ClosePositionRequest](#cryptowatch-broker-api-v2-ClosePositionRequest) | [ClosePositionResponse](#cryptowatch-broker-api-v2-ClosePositionResponse) | ClosePosition allows the user to close an open position on an exchange. Closing a position means taking the position back to zero. possible failures: A user may get a &#34;not found&#34; error for a position that exists on the exchange. Since Cryptowatch syncs with the exchange for the user positions it is possible for numerous kind of errors to occur during the syncing process that would restrict cryptowatch from getting a user&#39;s position. Failures can be the exchange not returning position, exchange has issues when cryptowatch requests to sync user positions, exchange has downtime etc. |
| SettlePosition | [SettlePositionRequest](#cryptowatch-broker-api-v2-SettlePositionRequest) | [SettlePositionResponse](#cryptowatch-broker-api-v2-SettlePositionResponse) | SettlePosition allows a user to settle their open position on an exchange. A valid exchange position id must be provided. If an exchange id is provided as well, it must be an exchange that supports margin. Exchanges give traders the option to ‘settle’ the position to reduce their liquidation risk by paying back the exchange the amount borrowed to originally open the position. possible failures: A user may get a &#34;not found&#34; error for a position that exists on the exchange. Since Cryptowatch syncs with the exchange for the user positions it is possible for numerous kind of errors to occur during the syncing process that would restrict cryptowatch from getting a user&#39;s position. Failures can be the exchange not returning position, exchange has issues when cryptowatch requests to sync user positions, exchange has downtime etc. |
| GetUserSettings | [GetUserSettingsRequest](#cryptowatch-broker-api-v2-GetUserSettingsRequest) | [GetUserSettingsResponse](#cryptowatch-broker-api-v2-GetUserSettingsResponse) | GetUserSettings allows the user to see their set settings on cryptowatch. |
| UpdateUserSettings | [UpdateUserSettingsRequest](#cryptowatch-broker-api-v2-UpdateUserSettingsRequest) | [UpdateUserSettingsResponse](#cryptowatch-broker-api-v2-UpdateUserSettingsResponse) | UpdateUserSettings allows the user to update their cryptowatch settings. This is mainly used to manipulate cryptowatch&#39;s trading user interface to the user&#39;s preference |
| ListExchangeAccounts | [ListExchangeAccountsRequest](#cryptowatch-broker-api-v2-ListExchangeAccountsRequest) | [ListExchangeAccountsResponse](#cryptowatch-broker-api-v2-ListExchangeAccountsResponse) | ListExchangeAccounts allows the user to list their registered exchange accounts with cryptowatch. If the user does not have an account registered with cryptowatch for an exchange, the map will be empty for that exchange. If a user has at least 1 account, it is guaranteed that one account is marked as the default account for that exchange. |
| GetFeeSchedule | [GetFeeScheduleRequest](#cryptowatch-broker-api-v2-GetFeeScheduleRequest) | [GetFeeScheduleResponse](#cryptowatch-broker-api-v2-GetFeeScheduleResponse) | GetFeeSchedule allows a user to get their fee schedule for a market. The market&#39;s exchange must be registered with cryptowatch with valid keys to be able to fetch the fee schedule for that market. Also, fee schedules are synced with the exchange when a user requests their fee schedule, or if a user has an active session. The sync process only occur every 24 hours, if the user requests their fee schedule and it has been synced within the past 24 hours of the request, the user will get what cryptowatch has for the fees schedule at that time; else, cryptowatch requests to sync the user fee schedule for that market and then return the new data. |
| PlayOrders | [PlayOrdersRequest](#cryptowatch-broker-api-v2-PlayOrdersRequest) | [PlayOrdersResponse](#cryptowatch-broker-api-v2-PlayOrdersResponse) | PlayOrders allows a user to play their paused orders. Not providing any argument will cause all of the user&#39;s paused orders to be played. Using the arguments provided helps narrowing down the orders the user would like to be played. PlayOrders failures are handled as CreateOrders failures: if playing an order failed in the middle of the process, the rest of the orders are aborted and an error is returned to the user. Example, the user sent 3 orders to be played and the failure occurred on the second order, the third order will not be played and only the first order is played successfully. |
| PauseOrders | [PauseOrdersRequest](#cryptowatch-broker-api-v2-PauseOrdersRequest) | [PauseOrdersResponse](#cryptowatch-broker-api-v2-PauseOrdersResponse) | PauseOrders allows a user to pause their orders. Not providing any argument will cause all of the user&#39;s __open__ orders to be paused. Using the arguments provided helps narrowing down the orders the user would like to be paused. PauseOrders failures are handled as CreateOrders failures: if pausing an order failed in the middle of the process, the rest of the orders are not paused and an error is returned to the user. Example, the user sent 3 orders to be paused and the failure occurred on the second order, the third order will not be paused and only the first order is paused successfully. |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

