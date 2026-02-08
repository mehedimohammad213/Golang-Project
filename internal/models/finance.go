package models

import (
	"time"
)

type PurchaseHistory struct {
	ID                  int64      `db:"id" json:"id"`
	CarID               int64      `db:"car_id" json:"car_id"`
	PurchaseDate        *time.Time `db:"purchase_date" json:"purchase_date"`
	PurchaseAmount      *float64   `db:"purchase_amount" json:"purchase_amount"`
	GovtDuty            *float64   `db:"govt_duty" json:"govt_duty"`
	CnfAmount           *float64   `db:"cnf_amount" json:"cnf_amount"`
	LcDate              *time.Time `db:"lc_date" json:"lc_date"`
	LcNumber            *string    `db:"lc_number" json:"lc_number"`
	LcBankName          *string    `db:"lc_bank_name" json:"lc_bank_name"`
	LcBankBranchName    *string    `db:"lc_bank_branch_name" json:"lc_bank_branch_name"`
	LcBankBranchAddress *string    `db:"lc_bank_branch_address" json:"lc_bank_branch_address"`
	TotalUnitsPerLc     *int       `db:"total_units_per_lc" json:"total_units_per_lc"`
	ForeignAmount       *float64   `db:"foreign_amount" json:"foreign_amount"`
	BDTAmount           *float64   `db:"bdt_amount" json:"bdt_amount"`
	CreatedAt           time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt           time.Time  `db:"updated_at" json:"updated_at"`
}

type PaymentHistory struct {
	ID                int64      `db:"id" json:"id"`
	CarID             *int64     `db:"car_id" json:"car_id"`
	ShowroomName      *string    `db:"showroom_name" json:"showroom_name"`
	WholesalerAddress *string    `db:"wholesaler_address" json:"wholesaler_address"`
	PurchaseAmount    *float64   `db:"purchase_amount" json:"purchase_amount"`
	PurchaseDate      *time.Time `db:"purchase_date" json:"purchase_date"`
	CustomerName      *string    `db:"customer_name" json:"customer_name"`
	NIDNumber         *string    `db:"nid_number" json:"nid_number"`
	TinCertificate    *string    `db:"tin_certificate" json:"tin_certificate"`
	CustomerAddress   *string    `db:"customer_address" json:"customer_address"`
	ContactNumber     *string    `db:"contact_number" json:"contact_number"`
	Email             *string    `db:"email" json:"email"`
	CreatedAt         time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time  `db:"updated_at" json:"updated_at"`
}

type Installment struct {
	ID               int64      `db:"id" json:"id"`
	PaymentHistoryID int64      `db:"payment_history_id" json:"payment_history_id"`
	InstallmentDate  *time.Time `db:"installment_date" json:"installment_date"`
	Description      *string    `db:"description" json:"description"`
	Amount           *float64   `db:"amount" json:"amount"`
	PaymentMethod    *string    `db:"payment_method" json:"payment_method"`
	BankName         *string    `db:"bank_name" json:"bank_name"`
	ChequeNumber     *string    `db:"cheque_number" json:"cheque_number"`
	Balance          *float64   `db:"balance" json:"balance"`
	Remarks          *string    `db:"remarks" json:"remarks"`
	CreatedAt        time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time  `db:"updated_at" json:"updated_at"`
}
