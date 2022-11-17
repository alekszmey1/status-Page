package service

import (
	"fmt"
	"math"
	"strings"
)

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

func NewBillingData(b []bool) *BillingData {
	bd := BillingData{}
	bd.CreateCustomer = b[0]
	bd.Purchase = b[1]
	bd.Payout = b[2]
	bd.Recurring = b[3]
	bd.FraudControl = b[4]
	bd.CheckoutPage = b[5]
	return &bd
}

type StorageBD struct {
	storageBillingData map[int]*BillingData
}

func NewStorageBD() *StorageBD {
	return &StorageBD{storageBillingData: make(map[int]*BillingData)}
}
func (u *StorageBD) put(bd *BillingData, i int) {
	u.storageBillingData[i] = bd
}
func (u *StorageBD) getAll() []*BillingData {
	var billingDats []*BillingData
	for _, v := range u.storageBillingData {
		billingDats = append(billingDats, v)
	}
	return billingDats
}

func Billing() {
	BD := NewStorageBD()
	s := "00010011"
	rns := strings.Split(s, "")
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	var k float64
	var b []bool
	for i := 0; i < 6; i++ {
		if rns[i] == "1" {
			x := math.Pow(float64(2), float64(i))
			k += x
			b = append(b, true)
		} else {
			b = append(b, false)
		}

	}
	fmt.Println(k)

	l := NewBillingData(b)
	BD.put(l, 0)

	for _, v := range BD.getAll() {
		fmt.Println(v)
	}
}
