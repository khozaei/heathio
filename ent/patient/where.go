// Code generated by ent, DO NOT EDIT.

package patient

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/khozaei/healthio/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldID, id))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldFirstName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldLastName, v))
}

// FatherName applies equality check predicate on the "father_name" field. It's identical to FatherNameEQ.
func FatherName(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldFatherName, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldPhone, v))
}

// NationalCode applies equality check predicate on the "national_code" field. It's identical to NationalCodeEQ.
func NationalCode(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldNationalCode, v))
}

// IdentityCode applies equality check predicate on the "identity_code" field. It's identical to IdentityCodeEQ.
func IdentityCode(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldIdentityCode, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldDeletedAt, v))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameIsNil applies the IsNil predicate on the "first_name" field.
func FirstNameIsNil() predicate.Patient {
	return predicate.Patient(sql.FieldIsNull(FieldFirstName))
}

// FirstNameNotNil applies the NotNil predicate on the "first_name" field.
func FirstNameNotNil() predicate.Patient {
	return predicate.Patient(sql.FieldNotNull(FieldFirstName))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContainsFold(FieldFirstName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameIsNil applies the IsNil predicate on the "last_name" field.
func LastNameIsNil() predicate.Patient {
	return predicate.Patient(sql.FieldIsNull(FieldLastName))
}

// LastNameNotNil applies the NotNil predicate on the "last_name" field.
func LastNameNotNil() predicate.Patient {
	return predicate.Patient(sql.FieldNotNull(FieldLastName))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContainsFold(FieldLastName, v))
}

// FatherNameEQ applies the EQ predicate on the "father_name" field.
func FatherNameEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldFatherName, v))
}

// FatherNameNEQ applies the NEQ predicate on the "father_name" field.
func FatherNameNEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldFatherName, v))
}

// FatherNameIn applies the In predicate on the "father_name" field.
func FatherNameIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldFatherName, vs...))
}

// FatherNameNotIn applies the NotIn predicate on the "father_name" field.
func FatherNameNotIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldFatherName, vs...))
}

// FatherNameGT applies the GT predicate on the "father_name" field.
func FatherNameGT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldFatherName, v))
}

// FatherNameGTE applies the GTE predicate on the "father_name" field.
func FatherNameGTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldFatherName, v))
}

// FatherNameLT applies the LT predicate on the "father_name" field.
func FatherNameLT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldFatherName, v))
}

// FatherNameLTE applies the LTE predicate on the "father_name" field.
func FatherNameLTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldFatherName, v))
}

// FatherNameContains applies the Contains predicate on the "father_name" field.
func FatherNameContains(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContains(FieldFatherName, v))
}

// FatherNameHasPrefix applies the HasPrefix predicate on the "father_name" field.
func FatherNameHasPrefix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasPrefix(FieldFatherName, v))
}

// FatherNameHasSuffix applies the HasSuffix predicate on the "father_name" field.
func FatherNameHasSuffix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasSuffix(FieldFatherName, v))
}

// FatherNameIsNil applies the IsNil predicate on the "father_name" field.
func FatherNameIsNil() predicate.Patient {
	return predicate.Patient(sql.FieldIsNull(FieldFatherName))
}

// FatherNameNotNil applies the NotNil predicate on the "father_name" field.
func FatherNameNotNil() predicate.Patient {
	return predicate.Patient(sql.FieldNotNull(FieldFatherName))
}

// FatherNameEqualFold applies the EqualFold predicate on the "father_name" field.
func FatherNameEqualFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEqualFold(FieldFatherName, v))
}

// FatherNameContainsFold applies the ContainsFold predicate on the "father_name" field.
func FatherNameContainsFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContainsFold(FieldFatherName, v))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContains(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.Patient {
	return predicate.Patient(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.Patient {
	return predicate.Patient(sql.FieldNotNull(FieldPhone))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContainsFold(FieldPhone, v))
}

// NationalCodeEQ applies the EQ predicate on the "national_code" field.
func NationalCodeEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldNationalCode, v))
}

// NationalCodeNEQ applies the NEQ predicate on the "national_code" field.
func NationalCodeNEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldNationalCode, v))
}

// NationalCodeIn applies the In predicate on the "national_code" field.
func NationalCodeIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldNationalCode, vs...))
}

// NationalCodeNotIn applies the NotIn predicate on the "national_code" field.
func NationalCodeNotIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldNationalCode, vs...))
}

// NationalCodeGT applies the GT predicate on the "national_code" field.
func NationalCodeGT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldNationalCode, v))
}

// NationalCodeGTE applies the GTE predicate on the "national_code" field.
func NationalCodeGTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldNationalCode, v))
}

// NationalCodeLT applies the LT predicate on the "national_code" field.
func NationalCodeLT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldNationalCode, v))
}

// NationalCodeLTE applies the LTE predicate on the "national_code" field.
func NationalCodeLTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldNationalCode, v))
}

// NationalCodeContains applies the Contains predicate on the "national_code" field.
func NationalCodeContains(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContains(FieldNationalCode, v))
}

// NationalCodeHasPrefix applies the HasPrefix predicate on the "national_code" field.
func NationalCodeHasPrefix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasPrefix(FieldNationalCode, v))
}

// NationalCodeHasSuffix applies the HasSuffix predicate on the "national_code" field.
func NationalCodeHasSuffix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasSuffix(FieldNationalCode, v))
}

// NationalCodeIsNil applies the IsNil predicate on the "national_code" field.
func NationalCodeIsNil() predicate.Patient {
	return predicate.Patient(sql.FieldIsNull(FieldNationalCode))
}

// NationalCodeNotNil applies the NotNil predicate on the "national_code" field.
func NationalCodeNotNil() predicate.Patient {
	return predicate.Patient(sql.FieldNotNull(FieldNationalCode))
}

// NationalCodeEqualFold applies the EqualFold predicate on the "national_code" field.
func NationalCodeEqualFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEqualFold(FieldNationalCode, v))
}

// NationalCodeContainsFold applies the ContainsFold predicate on the "national_code" field.
func NationalCodeContainsFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContainsFold(FieldNationalCode, v))
}

// IdentityCodeEQ applies the EQ predicate on the "identity_code" field.
func IdentityCodeEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldIdentityCode, v))
}

// IdentityCodeNEQ applies the NEQ predicate on the "identity_code" field.
func IdentityCodeNEQ(v string) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldIdentityCode, v))
}

// IdentityCodeIn applies the In predicate on the "identity_code" field.
func IdentityCodeIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldIdentityCode, vs...))
}

// IdentityCodeNotIn applies the NotIn predicate on the "identity_code" field.
func IdentityCodeNotIn(vs ...string) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldIdentityCode, vs...))
}

// IdentityCodeGT applies the GT predicate on the "identity_code" field.
func IdentityCodeGT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldIdentityCode, v))
}

// IdentityCodeGTE applies the GTE predicate on the "identity_code" field.
func IdentityCodeGTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldIdentityCode, v))
}

// IdentityCodeLT applies the LT predicate on the "identity_code" field.
func IdentityCodeLT(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldIdentityCode, v))
}

// IdentityCodeLTE applies the LTE predicate on the "identity_code" field.
func IdentityCodeLTE(v string) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldIdentityCode, v))
}

// IdentityCodeContains applies the Contains predicate on the "identity_code" field.
func IdentityCodeContains(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContains(FieldIdentityCode, v))
}

// IdentityCodeHasPrefix applies the HasPrefix predicate on the "identity_code" field.
func IdentityCodeHasPrefix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasPrefix(FieldIdentityCode, v))
}

// IdentityCodeHasSuffix applies the HasSuffix predicate on the "identity_code" field.
func IdentityCodeHasSuffix(v string) predicate.Patient {
	return predicate.Patient(sql.FieldHasSuffix(FieldIdentityCode, v))
}

// IdentityCodeIsNil applies the IsNil predicate on the "identity_code" field.
func IdentityCodeIsNil() predicate.Patient {
	return predicate.Patient(sql.FieldIsNull(FieldIdentityCode))
}

// IdentityCodeNotNil applies the NotNil predicate on the "identity_code" field.
func IdentityCodeNotNil() predicate.Patient {
	return predicate.Patient(sql.FieldNotNull(FieldIdentityCode))
}

// IdentityCodeEqualFold applies the EqualFold predicate on the "identity_code" field.
func IdentityCodeEqualFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldEqualFold(FieldIdentityCode, v))
}

// IdentityCodeContainsFold applies the ContainsFold predicate on the "identity_code" field.
func IdentityCodeContainsFold(v string) predicate.Patient {
	return predicate.Patient(sql.FieldContainsFold(FieldIdentityCode, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Patient {
	return predicate.Patient(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Patient {
	return predicate.Patient(sql.FieldNotNull(FieldUpdatedAt))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Patient {
	return predicate.Patient(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Patient {
	return predicate.Patient(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Patient {
	return predicate.Patient(sql.FieldNotNull(FieldDeletedAt))
}

// HasAttachment applies the HasEdge predicate on the "attachment" edge.
func HasAttachment() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttachmentTable, AttachmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttachmentWith applies the HasEdge predicate on the "attachment" edge with a given conditions (other predicates).
func HasAttachmentWith(preds ...predicate.Attachment) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := newAttachmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReception applies the HasEdge predicate on the "reception" edge.
func HasReception() predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReceptionTable, ReceptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReceptionWith applies the HasEdge predicate on the "reception" edge with a given conditions (other predicates).
func HasReceptionWith(preds ...predicate.Reception) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		step := newReceptionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Patient) predicate.Patient {
	return predicate.Patient(func(s *sql.Selector) {
		p(s.Not())
	})
}
