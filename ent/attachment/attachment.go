// Code generated by ent, DO NOT EDIT.

package attachment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the attachment type in the database.
	Label = "attachment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFilePath holds the string denoting the file_path field in the database.
	FieldFilePath = "file_path"
	// FieldFileType holds the string denoting the file_type field in the database.
	FieldFileType = "file_type"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeVisit holds the string denoting the visit edge name in mutations.
	EdgeVisit = "visit"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"
	// Table holds the table name of the attachment in the database.
	Table = "attachments"
	// VisitTable is the table that holds the visit relation/edge.
	VisitTable = "attachments"
	// VisitInverseTable is the table name for the Visit entity.
	// It exists in this package in order to avoid circular dependency with the "visit" package.
	VisitInverseTable = "visits"
	// VisitColumn is the table column denoting the visit relation/edge.
	VisitColumn = "visit_attachment"
	// PatientTable is the table that holds the patient relation/edge.
	PatientTable = "attachments"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "patient_attachment"
)

// Columns holds all SQL columns for attachment fields.
var Columns = []string{
	FieldID,
	FieldFilePath,
	FieldFileType,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "attachments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"patient_attachment",
	"visit_attachment",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)

// OrderOption defines the ordering options for the Attachment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFilePath orders the results by the file_path field.
func ByFilePath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFilePath, opts...).ToFunc()
}

// ByFileType orders the results by the file_type field.
func ByFileType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileType, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByVisitField orders the results by visit field.
func ByVisitField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVisitStep(), sql.OrderByField(field, opts...))
	}
}

// ByPatientField orders the results by patient field.
func ByPatientField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPatientStep(), sql.OrderByField(field, opts...))
	}
}
func newVisitStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VisitInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, VisitTable, VisitColumn),
	)
}
func newPatientStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PatientInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
	)
}
