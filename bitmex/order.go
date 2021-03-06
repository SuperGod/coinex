package bitmex

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/SuperGod/coinex/bitmex/client/order"
	"github.com/SuperGod/coinex/bitmex/models"
	. "github.com/SuperGod/trademodel"
)

const (
	OrderBuy  = "Buy"
	OrderSell = "Sell"

	OrderTypeLimit     = "Limit"
	OrderTypeMarket    = "Market"
	OrderTypeStop      = "Stop"      // stop lose with market price, must set stopPx
	OrderTypeStopLimit = "StopLimit" // stop lose with limit price, must set stopPx
	PostOnly           = "ParticipateDoNotInitiate"
	ReduceOnly         = "ReduceOnly"
)

var (
	NoOrderFound = errors.New("no such order")
)

func transOrder(o *models.Order) (ret *Order) {
	ret = &Order{OrderID: *o.OrderID,
		Symbol:   o.Symbol,
		Currency: o.Currency,
		Amount:   float64(o.OrderQty),
		Price:    o.AvgPx,
		Status:   o.OrdStatus,
		Side:     o.Side,
		Time:     time.Time(o.Timestamp)}
	if ret.Price == 0 {
		ret.Price = o.Price
	}
	return
}

func (b *Bitmex) fixPrice(price float64) (ret float64, err error) {
	if price == 0 {
		return
	}
	err = b.preload()
	if err != nil {
		return
	}
	si, ok := b.contacts[b.symbol]
	if ok {
		ret = math.Floor((price / si.TickSize)) * si.TickSize
	} else {
		ret = price
	}
	return
}

// Buy open long with price
func (b *Bitmex) Buy(price float64, amount float64) (ret *Order, err error) {
	ret, err = b.OpenLong(price, amount)
	return
}

// Buy open long with price
func (b *Bitmex) Sell(price float64, amount float64) (ret *Order, err error) {
	ret, err = b.OpenShort(price, amount)
	return
}

// OpenLong open long with price
func (b *Bitmex) OpenLong(price float64, amount float64) (ret *Order, err error) {
	comment := "open long with bitmex api"
	side := "Buy"
	orderType := "Limit"
	nAmount := int32(amount)
	newOrder, err := b.createOrder(price, nAmount, side, orderType, comment, b.postOnlys...)
	if err != nil {
		return
	}
	fmt.Printf("%##v\n", *newOrder)
	ret = transOrder(newOrder)
	return
}

// CloseLong close long with price
func (b *Bitmex) CloseLong(price float64, amount float64) (ret *Order, err error) {
	comment := "close long with bitmex api"
	nAmount := 0 - int32(amount)
	newOrder, err := b.closeOrder(price, nAmount, OrderSell, OrderTypeLimit, comment, b.postOnlys...)
	if err != nil {
		return
	}
	ret = transOrder(newOrder)
	return
}

// OpenShort open short with price
func (b *Bitmex) OpenShort(price float64, amount float64) (ret *Order, err error) {
	comment := "open short with bitmex api"
	nAmount := 0 - int32(amount)
	newOrder, err := b.createOrder(price, nAmount, OrderSell, OrderTypeLimit, comment, b.postOnlys...)
	if err != nil {
		return
	}
	ret = transOrder(newOrder)
	return
}

// CloseShort close short with price
func (b *Bitmex) CloseShort(price float64, amount float64) (ret *Order, err error) {
	comment := "close short with bitmex api"
	nAmount := int32(amount)
	newOrder, err := b.closeOrder(price, nAmount, OrderBuy, OrderTypeLimit, comment)
	if err != nil {
		return
	}
	ret = transOrder(newOrder)
	return
}

// OpenLongMarket open long with market price
func (b *Bitmex) OpenLongMarket(amount float64) (ret *Order, err error) {
	comment := "open market long with bitmex api"
	nAmount := int32(amount)
	newOrder, err := b.createOrder(0, nAmount, OrderBuy, OrderTypeMarket, comment)
	if err != nil {
		return
	}
	ret = transOrder(newOrder)
	return
}

// CloseLongarket close long with market price
func (b *Bitmex) CloseLongMarket(amount float64) (ret *Order, err error) {
	comment := "close market long with bitmex api"
	nAmount := 0 - int32(amount)
	newOrder, err := b.closeOrder(0, nAmount, OrderSell, OrderTypeMarket, comment)
	if err != nil {
		return
	}
	ret = transOrder(newOrder)
	return
}

// OpenShortMarket open short with market price
func (b *Bitmex) OpenShortMarket(amount float64) (ret *Order, err error) {
	comment := "open market short with bitmex api"
	nAmount := 0 - int32(amount)
	newOrder, err := b.createOrder(0, nAmount, OrderSell, OrderTypeMarket, comment)
	if err != nil {
		return
	}
	ret = transOrder(newOrder)
	return
}

// CloseShortMarket close short with market price
func (b *Bitmex) CloseShortMarket(amount float64) (ret *Order, err error) {
	comment := "close market short with bitmex api"
	nAmount := int32(amount)
	newOrder, err := b.closeOrder(0, nAmount, OrderBuy, OrderTypeMarket, comment)
	if err != nil {
		return
	}
	ret = transOrder(newOrder)
	return
}

// StopLoseBuy when marketPrice>=stopPrice, create buy order with price and amount
func (b *Bitmex) StopLoseBuy(stopPrice, price, amount float64) (ret *Order, err error) {
	comment := "stop limit buy with bitmex api"
	nAmount := int32(amount)
	newOrder, err := b.createStopOrder(stopPrice, price, nAmount, OrderBuy, OrderTypeStopLimit, comment)
	if err != nil {
		return
	}
	ret = transOrder(newOrder)
	return
}

// StopLoseSell when marketPrice<=stopPrice, create sell order with price and amount
func (b *Bitmex) StopLoseSell(stopPrice, price, amount float64) (ret *Order, err error) {
	comment := "stop limit buy with bitmex api"
	nAmount := int32(amount)
	newOrder, err := b.createStopOrder(stopPrice, price, nAmount, OrderSell, OrderTypeStopLimit, comment)
	if err != nil {
		return
	}
	ret = transOrder(newOrder)
	return
}

// StopLoseSell when marketPrice>=stopPrice, create buy order with marketPrice and amount
func (b *Bitmex) StopLoseBuyMarket(price, amount float64) (ret *Order, err error) {
	comment := "stop market buy with bitmex api"
	nAmount := int32(amount)
	newOrder, err := b.createStopOrder(price, 0, nAmount, OrderBuy, OrderTypeStop, comment)
	if err != nil {
		return
	}
	ret = transOrder(newOrder)
	return
}

// StopLoseSell when marketPrice<=stopPrice, create buy order with marketPrice and amount
func (b *Bitmex) StopLoseSellMarket(price, amount float64) (ret *Order, err error) {
	comment := "stop market buy with bitmex api"
	nAmount := int32(amount)
	newOrder, err := b.createStopOrder(price, 0, nAmount, OrderSell, OrderTypeStop, comment)
	if err != nil {
		return
	}
	ret = transOrder(newOrder)
	return
}

// createStopOrder stop order
// if orderType is Buy, when marketPrice>= price, buy
// if orderType is Sell, when marketPrice<=price, sell
func (b *Bitmex) createStopOrder(stopPrice, price float64, amount int32, side, orderType, comment string) (newOrder *models.Order, err error) {
	price, err = b.fixPrice(price)
	if err != nil {
		return
	}
	stopPrice, err = b.fixPrice(stopPrice)
	if err != nil {
		return
	}
	execInst := "Close"

	params := order.OrderNewParams{
		StopPx:   &stopPrice,
		Side:     &side,
		Symbol:   b.symbol,
		Text:     &comment,
		OrderQty: &amount,
		OrdType:  &orderType,
		ExecInst: &execInst,
	}
	if stopPrice != 0 {
		params.StopPx = &stopPrice
	} else {
		execInst += ",LastPrice"
	}
	if price != 0 {
		params.Price = &price
	}
	buf, _ := json.Marshal(params)
	fmt.Println("param:", string(buf))
	orderInfo, err := b.api.Order.OrderNew(&params, nil)
	if err != nil {
		badReqErr, ok := err.(*order.OrderNewBadRequest)
		if ok && badReqErr != nil && badReqErr.Payload != nil && badReqErr.Payload.Error != nil {
			err = fmt.Errorf("bad request %s %s", badReqErr.Payload.Error.Name, badReqErr.Payload.Error.Message)
		}
		return
	}
	newOrder = orderInfo.Payload
	return
}

func (b *Bitmex) closeOrder(price float64, amount int32, side, orderType, comment string, execInsts ...string) (newOrder *models.Order, err error) {
	price, err = b.fixPrice(price)
	if err != nil {
		return
	}
	execInst := "Close"
	params := order.OrderNewParams{
		Side:     &side,
		Symbol:   b.symbol,
		Text:     &comment,
		OrderQty: &amount,
		OrdType:  &orderType,
		ExecInst: &execInst,
	}
	if len(execInsts) > 0 {
		execInsts = append(execInsts, execInst)
		execValue := strings.Join(execInsts, ",")
		params.ExecInst = &execValue
	}
	if price != 0 {
		params.Price = &price
	}
	orderInfo, err := b.api.Order.OrderNew(&params, nil)
	if err != nil {
		badReqErr, ok := err.(*order.OrderNewBadRequest)
		if ok && badReqErr != nil && badReqErr.Payload != nil && badReqErr.Payload.Error != nil {
			err = fmt.Errorf("bad request %s %s", badReqErr.Payload.Error.Name, badReqErr.Payload.Error.Message)
		}
		return
	}
	newOrder = orderInfo.Payload
	return
}

// CreteOrder create bitmex order,return bitmex model information
func (b *Bitmex) CreateOrder(price float64, amount int32, symbol, side, orderType, comment string, execInsts ...string) (newOrder *models.Order, err error) {
	price, err = b.fixPrice(price)
	if err != nil {
		return
	}
	params := order.OrderNewParams{
		Side:     &side,
		Symbol:   symbol,
		Text:     &comment,
		OrderQty: &amount,
		OrdType:  &orderType,
	}
	if len(execInsts) > 0 {
		execValue := strings.Join(execInsts, ",")
		params.ExecInst = &execValue
	}
	if price != 0 {
		params.Price = &price
	}
	orderInfo, err := b.api.Order.OrderNew(&params, nil)
	if err != nil {
		badReqErr, ok := err.(*order.OrderNewBadRequest)
		if ok && badReqErr != nil && badReqErr.Payload != nil && badReqErr.Payload.Error != nil {
			err = fmt.Errorf("bad request %s %s", badReqErr.Payload.Error.Name, badReqErr.Payload.Error.Message)
		}
		return
	}
	newOrder = orderInfo.Payload
	return
}

func (b *Bitmex) createOrder(price float64, amount int32, side, orderType, comment string, execInsts ...string) (newOrder *models.Order, err error) {
	newOrder, err = b.CreateOrder(price, amount, b.symbol, side, orderType, comment, execInsts...)
	return
}

// CancelOrder with oid
func (b *Bitmex) CancelOrder(oid string) (newOrder *Order, err error) {
	comment := "cancle order with bitmex api"
	params := order.OrderCancelParams{
		OrderID: &oid,
		Text:    &comment,
	}
	orderInfo, err := b.api.Order.OrderCancel(&params, nil)
	if err != nil {
		return
	}
	if len(orderInfo.Payload) == 0 {
		err = NoOrderFound
		return
	}
	newOrder = transOrder(orderInfo.Payload[0])
	return
}

// CancelAllOrders cancel all not filled orders
func (b *Bitmex) CancelAllOrders() (orders []*Order, err error) {
	comment := "cancle all order with bitmex api"
	params := order.OrderCancelAllParams{
		Symbol: &b.symbol,
		Text:   &comment,
	}
	orderInfo, err := b.api.Order.OrderCancelAll(&params, nil)
	if err != nil {
		return
	}
	for _, v := range orderInfo.Payload {
		orders = append(orders, transOrder(v))
	}
	return
}

// Orders get all active orders
func (b *Bitmex) Orders() (orders []*Order, err error) {
	filters := `{"ordStatus":"New"}`
	params := order.OrderGetOrdersParams{
		Symbol: &b.symbol,
		Filter: &filters,
	}
	orderInfo, err := b.api.Order.OrderGetOrders(&params, nil)
	if err != nil {
		return
	}
	for _, v := range orderInfo.Payload {
		orders = append(orders, transOrder(v))
	}
	return
}

// Orders get all active orders
func (b *Bitmex) Order(oid string) (newOrder *Order, err error) {
	filters := fmt.Sprintf(`{"orderID":"%s"}`, oid)
	params := order.OrderGetOrdersParams{
		Symbol: &b.symbol,
		Filter: &filters,
	}
	orderInfo, err := b.api.Order.OrderGetOrders(&params, nil)
	if err != nil {
		return
	}
	if len(orderInfo.Payload) == 0 {
		err = NoOrderFound
		return
	}
	newOrder = transOrder(orderInfo.Payload[0])
	return
}

func IsOrderFilled(o *Order) bool {
	return o.Status == "Filled"
}

func IsOrderCanceled(o *Order) bool {
	return o.Status == "Canceled"
}

func IsOrderLong(o *Order) bool {
	return strings.ToLower(o.Side) == "buy"
}
