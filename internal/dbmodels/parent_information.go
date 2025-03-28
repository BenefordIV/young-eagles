// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// ParentInformation is an object representing the database table.
type ParentInformation struct {
	// uuid of the parent information. used in tables such as child_information
	UUID        string      `boil:"uuid" json:"uuid" toml:"uuid" yaml:"uuid"`
	FirstName   null.String `boil:"first_name" json:"first_name,omitempty" toml:"first_name" yaml:"first_name,omitempty"`
	LastName    null.String `boil:"last_name" json:"last_name,omitempty" toml:"last_name" yaml:"last_name,omitempty"`
	Email       null.String `boil:"email" json:"email,omitempty" toml:"email" yaml:"email,omitempty"`
	PhoneNumber null.String `boil:"phone_number" json:"phone_number,omitempty" toml:"phone_number" yaml:"phone_number,omitempty"`
	CreatedTS   null.Time   `boil:"created_ts" json:"created_ts,omitempty" toml:"created_ts" yaml:"created_ts,omitempty"`
	UpdatedTS   null.Time   `boil:"updated_ts" json:"updated_ts,omitempty" toml:"updated_ts" yaml:"updated_ts,omitempty"`
	DeletedTS   null.Time   `boil:"deleted_ts" json:"deleted_ts,omitempty" toml:"deleted_ts" yaml:"deleted_ts,omitempty"`

	R *parentInformationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L parentInformationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ParentInformationColumns = struct {
	UUID        string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	CreatedTS   string
	UpdatedTS   string
	DeletedTS   string
}{
	UUID:        "uuid",
	FirstName:   "first_name",
	LastName:    "last_name",
	Email:       "email",
	PhoneNumber: "phone_number",
	CreatedTS:   "created_ts",
	UpdatedTS:   "updated_ts",
	DeletedTS:   "deleted_ts",
}

var ParentInformationTableColumns = struct {
	UUID        string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	CreatedTS   string
	UpdatedTS   string
	DeletedTS   string
}{
	UUID:        "parent_information.uuid",
	FirstName:   "parent_information.first_name",
	LastName:    "parent_information.last_name",
	Email:       "parent_information.email",
	PhoneNumber: "parent_information.phone_number",
	CreatedTS:   "parent_information.created_ts",
	UpdatedTS:   "parent_information.updated_ts",
	DeletedTS:   "parent_information.deleted_ts",
}

// Generated where

var ParentInformationWhere = struct {
	UUID        whereHelperstring
	FirstName   whereHelpernull_String
	LastName    whereHelpernull_String
	Email       whereHelpernull_String
	PhoneNumber whereHelpernull_String
	CreatedTS   whereHelpernull_Time
	UpdatedTS   whereHelpernull_Time
	DeletedTS   whereHelpernull_Time
}{
	UUID:        whereHelperstring{field: "\"parent_information\".\"uuid\""},
	FirstName:   whereHelpernull_String{field: "\"parent_information\".\"first_name\""},
	LastName:    whereHelpernull_String{field: "\"parent_information\".\"last_name\""},
	Email:       whereHelpernull_String{field: "\"parent_information\".\"email\""},
	PhoneNumber: whereHelpernull_String{field: "\"parent_information\".\"phone_number\""},
	CreatedTS:   whereHelpernull_Time{field: "\"parent_information\".\"created_ts\""},
	UpdatedTS:   whereHelpernull_Time{field: "\"parent_information\".\"updated_ts\""},
	DeletedTS:   whereHelpernull_Time{field: "\"parent_information\".\"deleted_ts\""},
}

// ParentInformationRels is where relationship names are stored.
var ParentInformationRels = struct {
	ParentChildInformations string
}{
	ParentChildInformations: "ParentChildInformations",
}

// parentInformationR is where relationships are stored.
type parentInformationR struct {
	ParentChildInformations ChildInformationSlice `boil:"ParentChildInformations" json:"ParentChildInformations" toml:"ParentChildInformations" yaml:"ParentChildInformations"`
}

// NewStruct creates a new relationship struct
func (*parentInformationR) NewStruct() *parentInformationR {
	return &parentInformationR{}
}

func (r *parentInformationR) GetParentChildInformations() ChildInformationSlice {
	if r == nil {
		return nil
	}
	return r.ParentChildInformations
}

// parentInformationL is where Load methods for each relationship are stored.
type parentInformationL struct{}

var (
	parentInformationAllColumns            = []string{"uuid", "first_name", "last_name", "email", "phone_number", "created_ts", "updated_ts", "deleted_ts"}
	parentInformationColumnsWithoutDefault = []string{"uuid"}
	parentInformationColumnsWithDefault    = []string{"first_name", "last_name", "email", "phone_number", "created_ts", "updated_ts", "deleted_ts"}
	parentInformationPrimaryKeyColumns     = []string{"uuid"}
	parentInformationGeneratedColumns      = []string{}
)

type (
	// ParentInformationSlice is an alias for a slice of pointers to ParentInformation.
	// This should almost always be used instead of []ParentInformation.
	ParentInformationSlice []*ParentInformation

	parentInformationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	parentInformationType                 = reflect.TypeOf(&ParentInformation{})
	parentInformationMapping              = queries.MakeStructMapping(parentInformationType)
	parentInformationPrimaryKeyMapping, _ = queries.BindMapping(parentInformationType, parentInformationMapping, parentInformationPrimaryKeyColumns)
	parentInformationInsertCacheMut       sync.RWMutex
	parentInformationInsertCache          = make(map[string]insertCache)
	parentInformationUpdateCacheMut       sync.RWMutex
	parentInformationUpdateCache          = make(map[string]updateCache)
	parentInformationUpsertCacheMut       sync.RWMutex
	parentInformationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single parentInformation record from the query.
func (q parentInformationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ParentInformation, error) {
	o := &ParentInformation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for parent_information")
	}

	return o, nil
}

// All returns all ParentInformation records from the query.
func (q parentInformationQuery) All(ctx context.Context, exec boil.ContextExecutor) (ParentInformationSlice, error) {
	var o []*ParentInformation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to ParentInformation slice")
	}

	return o, nil
}

// Count returns the count of all ParentInformation records in the query.
func (q parentInformationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count parent_information rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q parentInformationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if parent_information exists")
	}

	return count > 0, nil
}

// ParentChildInformations retrieves all the child_information's ChildInformations with an executor via parent_id column.
func (o *ParentInformation) ParentChildInformations(mods ...qm.QueryMod) childInformationQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"child_information\".\"parent_id\"=?", o.UUID),
	)

	return ChildInformations(queryMods...)
}

// LoadParentChildInformations allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (parentInformationL) LoadParentChildInformations(ctx context.Context, e boil.ContextExecutor, singular bool, maybeParentInformation interface{}, mods queries.Applicator) error {
	var slice []*ParentInformation
	var object *ParentInformation

	if singular {
		var ok bool
		object, ok = maybeParentInformation.(*ParentInformation)
		if !ok {
			object = new(ParentInformation)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeParentInformation)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeParentInformation))
			}
		}
	} else {
		s, ok := maybeParentInformation.(*[]*ParentInformation)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeParentInformation)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeParentInformation))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &parentInformationR{}
		}
		args[object.UUID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &parentInformationR{}
			}
			args[obj.UUID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`child_information`),
		qm.WhereIn(`child_information.parent_id in ?`, argsSlice...),
		qmhelper.WhereIsNull(`child_information.deleted_ts`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load child_information")
	}

	var resultSlice []*ChildInformation
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice child_information")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on child_information")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for child_information")
	}

	if singular {
		object.R.ParentChildInformations = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &childInformationR{}
			}
			foreign.R.Parent = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.UUID, foreign.ParentID) {
				local.R.ParentChildInformations = append(local.R.ParentChildInformations, foreign)
				if foreign.R == nil {
					foreign.R = &childInformationR{}
				}
				foreign.R.Parent = local
				break
			}
		}
	}

	return nil
}

// AddParentChildInformations adds the given related objects to the existing relationships
// of the parent_information, optionally inserting them as new records.
// Appends related to o.R.ParentChildInformations.
// Sets related.R.Parent appropriately.
func (o *ParentInformation) AddParentChildInformations(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ChildInformation) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ParentID, o.UUID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"child_information\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"parent_id"}),
				strmangle.WhereClause("\"", "\"", 2, childInformationPrimaryKeyColumns),
			)
			values := []interface{}{o.UUID, rel.UUID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.ParentID, o.UUID)
		}
	}

	if o.R == nil {
		o.R = &parentInformationR{
			ParentChildInformations: related,
		}
	} else {
		o.R.ParentChildInformations = append(o.R.ParentChildInformations, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &childInformationR{
				Parent: o,
			}
		} else {
			rel.R.Parent = o
		}
	}
	return nil
}

// SetParentChildInformations removes all previously related items of the
// parent_information replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Parent's ParentChildInformations accordingly.
// Replaces o.R.ParentChildInformations with related.
// Sets related.R.Parent's ParentChildInformations accordingly.
func (o *ParentInformation) SetParentChildInformations(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ChildInformation) error {
	query := "update \"child_information\" set \"parent_id\" = null where \"parent_id\" = $1"
	values := []interface{}{o.UUID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.ParentChildInformations {
			queries.SetScanner(&rel.ParentID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Parent = nil
		}
		o.R.ParentChildInformations = nil
	}

	return o.AddParentChildInformations(ctx, exec, insert, related...)
}

// RemoveParentChildInformations relationships from objects passed in.
// Removes related items from R.ParentChildInformations (uses pointer comparison, removal does not keep order)
// Sets related.R.Parent.
func (o *ParentInformation) RemoveParentChildInformations(ctx context.Context, exec boil.ContextExecutor, related ...*ChildInformation) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.ParentID, nil)
		if rel.R != nil {
			rel.R.Parent = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("parent_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.ParentChildInformations {
			if rel != ri {
				continue
			}

			ln := len(o.R.ParentChildInformations)
			if ln > 1 && i < ln-1 {
				o.R.ParentChildInformations[i] = o.R.ParentChildInformations[ln-1]
			}
			o.R.ParentChildInformations = o.R.ParentChildInformations[:ln-1]
			break
		}
	}

	return nil
}

// ParentInformations retrieves all the records using an executor.
func ParentInformations(mods ...qm.QueryMod) parentInformationQuery {
	mods = append(mods, qm.From("\"parent_information\""), qmhelper.WhereIsNull("\"parent_information\".\"deleted_ts\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"parent_information\".*"})
	}

	return parentInformationQuery{q}
}

// FindParentInformation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindParentInformation(ctx context.Context, exec boil.ContextExecutor, uUID string, selectCols ...string) (*ParentInformation, error) {
	parentInformationObj := &ParentInformation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"parent_information\" where \"uuid\"=$1 and \"deleted_ts\" is null", sel,
	)

	q := queries.Raw(query, uUID)

	err := q.Bind(ctx, exec, parentInformationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from parent_information")
	}

	return parentInformationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ParentInformation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no parent_information provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedTS).IsZero() {
			queries.SetScanner(&o.CreatedTS, currTime)
		}
		if queries.MustTime(o.UpdatedTS).IsZero() {
			queries.SetScanner(&o.UpdatedTS, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(parentInformationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	parentInformationInsertCacheMut.RLock()
	cache, cached := parentInformationInsertCache[key]
	parentInformationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			parentInformationAllColumns,
			parentInformationColumnsWithDefault,
			parentInformationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(parentInformationType, parentInformationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(parentInformationType, parentInformationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"parent_information\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"parent_information\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to insert into parent_information")
	}

	if !cached {
		parentInformationInsertCacheMut.Lock()
		parentInformationInsertCache[key] = cache
		parentInformationInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the ParentInformation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ParentInformation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedTS, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	parentInformationUpdateCacheMut.RLock()
	cache, cached := parentInformationUpdateCache[key]
	parentInformationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			parentInformationAllColumns,
			parentInformationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_ts"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update parent_information, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"parent_information\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, parentInformationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(parentInformationType, parentInformationMapping, append(wl, parentInformationPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update parent_information row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for parent_information")
	}

	if !cached {
		parentInformationUpdateCacheMut.Lock()
		parentInformationUpdateCache[key] = cache
		parentInformationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q parentInformationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for parent_information")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for parent_information")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ParentInformationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dbmodels: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), parentInformationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"parent_information\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, parentInformationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in parentInformation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all parentInformation")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ParentInformation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("dbmodels: no parent_information provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedTS).IsZero() {
			queries.SetScanner(&o.CreatedTS, currTime)
		}
		queries.SetScanner(&o.UpdatedTS, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(parentInformationColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	parentInformationUpsertCacheMut.RLock()
	cache, cached := parentInformationUpsertCache[key]
	parentInformationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			parentInformationAllColumns,
			parentInformationColumnsWithDefault,
			parentInformationColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			parentInformationAllColumns,
			parentInformationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert parent_information, could not build update column list")
		}

		ret := strmangle.SetComplement(parentInformationAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(parentInformationPrimaryKeyColumns) == 0 {
				return errors.New("dbmodels: unable to upsert parent_information, could not build conflict column list")
			}

			conflict = make([]string, len(parentInformationPrimaryKeyColumns))
			copy(conflict, parentInformationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"parent_information\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(parentInformationType, parentInformationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(parentInformationType, parentInformationMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to upsert parent_information")
	}

	if !cached {
		parentInformationUpsertCacheMut.Lock()
		parentInformationUpsertCache[key] = cache
		parentInformationUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single ParentInformation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ParentInformation) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no ParentInformation provided for delete")
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), parentInformationPrimaryKeyMapping)
		sql = "DELETE FROM \"parent_information\" WHERE \"uuid\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedTS = null.TimeFrom(currTime)
		wl := []string{"deleted_ts"}
		sql = fmt.Sprintf("UPDATE \"parent_information\" SET %s WHERE \"uuid\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(parentInformationType, parentInformationMapping, append(wl, parentInformationPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), valueMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete from parent_information")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for parent_information")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q parentInformationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no parentInformationQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_ts": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from parent_information")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for parent_information")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ParentInformationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), parentInformationPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"parent_information\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, parentInformationPrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), parentInformationPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedTS = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_ts"}
		sql = fmt.Sprintf("UPDATE \"parent_information\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, parentInformationPrimaryKeyColumns, len(o)),
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		args = append([]interface{}{currTime}, args...)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from parentInformation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for parent_information")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ParentInformation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindParentInformation(ctx, exec, o.UUID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ParentInformationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ParentInformationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), parentInformationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"parent_information\".* FROM \"parent_information\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, parentInformationPrimaryKeyColumns, len(*o)) +
		"and \"deleted_ts\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in ParentInformationSlice")
	}

	*o = slice

	return nil
}

// ParentInformationExists checks if the ParentInformation row exists.
func ParentInformationExists(ctx context.Context, exec boil.ContextExecutor, uUID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"parent_information\" where \"uuid\"=$1 and \"deleted_ts\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, uUID)
	}
	row := exec.QueryRowContext(ctx, sql, uUID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if parent_information exists")
	}

	return exists, nil
}

// Exists checks if the ParentInformation row exists.
func (o *ParentInformation) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ParentInformationExists(ctx, exec, o.UUID)
}
