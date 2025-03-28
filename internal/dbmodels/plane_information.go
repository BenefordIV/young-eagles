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

// PlaneInformation is an object representing the database table.
type PlaneInformation struct {
	// The call_number is the unique identifier for a given airplane
	CallNumber string `boil:"call_number" json:"call_number" toml:"call_number" yaml:"call_number"`
	// This is the model of airplane
	Model null.String `boil:"model" json:"model,omitempty" toml:"model" yaml:"model,omitempty"`
	// the make of airplane
	Make      null.String `boil:"make" json:"make,omitempty" toml:"make" yaml:"make,omitempty"`
	CreatedTS null.Time   `boil:"created_ts" json:"created_ts,omitempty" toml:"created_ts" yaml:"created_ts,omitempty"`
	UpdatedTS null.Time   `boil:"updated_ts" json:"updated_ts,omitempty" toml:"updated_ts" yaml:"updated_ts,omitempty"`
	DeletedTS null.Time   `boil:"deleted_ts" json:"deleted_ts,omitempty" toml:"deleted_ts" yaml:"deleted_ts,omitempty"`

	R *planeInformationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L planeInformationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PlaneInformationColumns = struct {
	CallNumber string
	Model      string
	Make       string
	CreatedTS  string
	UpdatedTS  string
	DeletedTS  string
}{
	CallNumber: "call_number",
	Model:      "model",
	Make:       "make",
	CreatedTS:  "created_ts",
	UpdatedTS:  "updated_ts",
	DeletedTS:  "deleted_ts",
}

var PlaneInformationTableColumns = struct {
	CallNumber string
	Model      string
	Make       string
	CreatedTS  string
	UpdatedTS  string
	DeletedTS  string
}{
	CallNumber: "plane_information.call_number",
	Model:      "plane_information.model",
	Make:       "plane_information.make",
	CreatedTS:  "plane_information.created_ts",
	UpdatedTS:  "plane_information.updated_ts",
	DeletedTS:  "plane_information.deleted_ts",
}

// Generated where

var PlaneInformationWhere = struct {
	CallNumber whereHelperstring
	Model      whereHelpernull_String
	Make       whereHelpernull_String
	CreatedTS  whereHelpernull_Time
	UpdatedTS  whereHelpernull_Time
	DeletedTS  whereHelpernull_Time
}{
	CallNumber: whereHelperstring{field: "\"plane_information\".\"call_number\""},
	Model:      whereHelpernull_String{field: "\"plane_information\".\"model\""},
	Make:       whereHelpernull_String{field: "\"plane_information\".\"make\""},
	CreatedTS:  whereHelpernull_Time{field: "\"plane_information\".\"created_ts\""},
	UpdatedTS:  whereHelpernull_Time{field: "\"plane_information\".\"updated_ts\""},
	DeletedTS:  whereHelpernull_Time{field: "\"plane_information\".\"deleted_ts\""},
}

// PlaneInformationRels is where relationship names are stored.
var PlaneInformationRels = struct {
	PlaneCallNumberFlightInformations string
}{
	PlaneCallNumberFlightInformations: "PlaneCallNumberFlightInformations",
}

// planeInformationR is where relationships are stored.
type planeInformationR struct {
	PlaneCallNumberFlightInformations FlightInformationSlice `boil:"PlaneCallNumberFlightInformations" json:"PlaneCallNumberFlightInformations" toml:"PlaneCallNumberFlightInformations" yaml:"PlaneCallNumberFlightInformations"`
}

// NewStruct creates a new relationship struct
func (*planeInformationR) NewStruct() *planeInformationR {
	return &planeInformationR{}
}

func (r *planeInformationR) GetPlaneCallNumberFlightInformations() FlightInformationSlice {
	if r == nil {
		return nil
	}
	return r.PlaneCallNumberFlightInformations
}

// planeInformationL is where Load methods for each relationship are stored.
type planeInformationL struct{}

var (
	planeInformationAllColumns            = []string{"call_number", "model", "make", "created_ts", "updated_ts", "deleted_ts"}
	planeInformationColumnsWithoutDefault = []string{"call_number"}
	planeInformationColumnsWithDefault    = []string{"model", "make", "created_ts", "updated_ts", "deleted_ts"}
	planeInformationPrimaryKeyColumns     = []string{"call_number"}
	planeInformationGeneratedColumns      = []string{}
)

type (
	// PlaneInformationSlice is an alias for a slice of pointers to PlaneInformation.
	// This should almost always be used instead of []PlaneInformation.
	PlaneInformationSlice []*PlaneInformation

	planeInformationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	planeInformationType                 = reflect.TypeOf(&PlaneInformation{})
	planeInformationMapping              = queries.MakeStructMapping(planeInformationType)
	planeInformationPrimaryKeyMapping, _ = queries.BindMapping(planeInformationType, planeInformationMapping, planeInformationPrimaryKeyColumns)
	planeInformationInsertCacheMut       sync.RWMutex
	planeInformationInsertCache          = make(map[string]insertCache)
	planeInformationUpdateCacheMut       sync.RWMutex
	planeInformationUpdateCache          = make(map[string]updateCache)
	planeInformationUpsertCacheMut       sync.RWMutex
	planeInformationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single planeInformation record from the query.
func (q planeInformationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PlaneInformation, error) {
	o := &PlaneInformation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for plane_information")
	}

	return o, nil
}

// All returns all PlaneInformation records from the query.
func (q planeInformationQuery) All(ctx context.Context, exec boil.ContextExecutor) (PlaneInformationSlice, error) {
	var o []*PlaneInformation

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to PlaneInformation slice")
	}

	return o, nil
}

// Count returns the count of all PlaneInformation records in the query.
func (q planeInformationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count plane_information rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q planeInformationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if plane_information exists")
	}

	return count > 0, nil
}

// PlaneCallNumberFlightInformations retrieves all the flight_information's FlightInformations with an executor via plane_call_number column.
func (o *PlaneInformation) PlaneCallNumberFlightInformations(mods ...qm.QueryMod) flightInformationQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"flight_information\".\"plane_call_number\"=?", o.CallNumber),
	)

	return FlightInformations(queryMods...)
}

// LoadPlaneCallNumberFlightInformations allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (planeInformationL) LoadPlaneCallNumberFlightInformations(ctx context.Context, e boil.ContextExecutor, singular bool, maybePlaneInformation interface{}, mods queries.Applicator) error {
	var slice []*PlaneInformation
	var object *PlaneInformation

	if singular {
		var ok bool
		object, ok = maybePlaneInformation.(*PlaneInformation)
		if !ok {
			object = new(PlaneInformation)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePlaneInformation)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePlaneInformation))
			}
		}
	} else {
		s, ok := maybePlaneInformation.(*[]*PlaneInformation)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePlaneInformation)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePlaneInformation))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &planeInformationR{}
		}
		args[object.CallNumber] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &planeInformationR{}
			}
			args[obj.CallNumber] = struct{}{}
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
		qm.From(`flight_information`),
		qm.WhereIn(`flight_information.plane_call_number in ?`, argsSlice...),
		qmhelper.WhereIsNull(`flight_information.deleted_ts`),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load flight_information")
	}

	var resultSlice []*FlightInformation
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice flight_information")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on flight_information")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for flight_information")
	}

	if singular {
		object.R.PlaneCallNumberFlightInformations = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &flightInformationR{}
			}
			foreign.R.PlaneCallNumberPlaneInformation = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.CallNumber, foreign.PlaneCallNumber) {
				local.R.PlaneCallNumberFlightInformations = append(local.R.PlaneCallNumberFlightInformations, foreign)
				if foreign.R == nil {
					foreign.R = &flightInformationR{}
				}
				foreign.R.PlaneCallNumberPlaneInformation = local
				break
			}
		}
	}

	return nil
}

// AddPlaneCallNumberFlightInformations adds the given related objects to the existing relationships
// of the plane_information, optionally inserting them as new records.
// Appends related to o.R.PlaneCallNumberFlightInformations.
// Sets related.R.PlaneCallNumberPlaneInformation appropriately.
func (o *PlaneInformation) AddPlaneCallNumberFlightInformations(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*FlightInformation) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.PlaneCallNumber, o.CallNumber)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"flight_information\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"plane_call_number"}),
				strmangle.WhereClause("\"", "\"", 2, flightInformationPrimaryKeyColumns),
			)
			values := []interface{}{o.CallNumber, rel.UUID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.PlaneCallNumber, o.CallNumber)
		}
	}

	if o.R == nil {
		o.R = &planeInformationR{
			PlaneCallNumberFlightInformations: related,
		}
	} else {
		o.R.PlaneCallNumberFlightInformations = append(o.R.PlaneCallNumberFlightInformations, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &flightInformationR{
				PlaneCallNumberPlaneInformation: o,
			}
		} else {
			rel.R.PlaneCallNumberPlaneInformation = o
		}
	}
	return nil
}

// SetPlaneCallNumberFlightInformations removes all previously related items of the
// plane_information replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.PlaneCallNumberPlaneInformation's PlaneCallNumberFlightInformations accordingly.
// Replaces o.R.PlaneCallNumberFlightInformations with related.
// Sets related.R.PlaneCallNumberPlaneInformation's PlaneCallNumberFlightInformations accordingly.
func (o *PlaneInformation) SetPlaneCallNumberFlightInformations(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*FlightInformation) error {
	query := "update \"flight_information\" set \"plane_call_number\" = null where \"plane_call_number\" = $1"
	values := []interface{}{o.CallNumber}
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
		for _, rel := range o.R.PlaneCallNumberFlightInformations {
			queries.SetScanner(&rel.PlaneCallNumber, nil)
			if rel.R == nil {
				continue
			}

			rel.R.PlaneCallNumberPlaneInformation = nil
		}
		o.R.PlaneCallNumberFlightInformations = nil
	}

	return o.AddPlaneCallNumberFlightInformations(ctx, exec, insert, related...)
}

// RemovePlaneCallNumberFlightInformations relationships from objects passed in.
// Removes related items from R.PlaneCallNumberFlightInformations (uses pointer comparison, removal does not keep order)
// Sets related.R.PlaneCallNumberPlaneInformation.
func (o *PlaneInformation) RemovePlaneCallNumberFlightInformations(ctx context.Context, exec boil.ContextExecutor, related ...*FlightInformation) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.PlaneCallNumber, nil)
		if rel.R != nil {
			rel.R.PlaneCallNumberPlaneInformation = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("plane_call_number")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.PlaneCallNumberFlightInformations {
			if rel != ri {
				continue
			}

			ln := len(o.R.PlaneCallNumberFlightInformations)
			if ln > 1 && i < ln-1 {
				o.R.PlaneCallNumberFlightInformations[i] = o.R.PlaneCallNumberFlightInformations[ln-1]
			}
			o.R.PlaneCallNumberFlightInformations = o.R.PlaneCallNumberFlightInformations[:ln-1]
			break
		}
	}

	return nil
}

// PlaneInformations retrieves all the records using an executor.
func PlaneInformations(mods ...qm.QueryMod) planeInformationQuery {
	mods = append(mods, qm.From("\"plane_information\""), qmhelper.WhereIsNull("\"plane_information\".\"deleted_ts\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"plane_information\".*"})
	}

	return planeInformationQuery{q}
}

// FindPlaneInformation retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPlaneInformation(ctx context.Context, exec boil.ContextExecutor, callNumber string, selectCols ...string) (*PlaneInformation, error) {
	planeInformationObj := &PlaneInformation{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"plane_information\" where \"call_number\"=$1 and \"deleted_ts\" is null", sel,
	)

	q := queries.Raw(query, callNumber)

	err := q.Bind(ctx, exec, planeInformationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from plane_information")
	}

	return planeInformationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PlaneInformation) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no plane_information provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(planeInformationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	planeInformationInsertCacheMut.RLock()
	cache, cached := planeInformationInsertCache[key]
	planeInformationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			planeInformationAllColumns,
			planeInformationColumnsWithDefault,
			planeInformationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(planeInformationType, planeInformationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(planeInformationType, planeInformationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"plane_information\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"plane_information\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "dbmodels: unable to insert into plane_information")
	}

	if !cached {
		planeInformationInsertCacheMut.Lock()
		planeInformationInsertCache[key] = cache
		planeInformationInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the PlaneInformation.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PlaneInformation) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedTS, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	planeInformationUpdateCacheMut.RLock()
	cache, cached := planeInformationUpdateCache[key]
	planeInformationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			planeInformationAllColumns,
			planeInformationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_ts"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update plane_information, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"plane_information\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, planeInformationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(planeInformationType, planeInformationMapping, append(wl, planeInformationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodels: unable to update plane_information row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for plane_information")
	}

	if !cached {
		planeInformationUpdateCacheMut.Lock()
		planeInformationUpdateCache[key] = cache
		planeInformationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q planeInformationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for plane_information")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for plane_information")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PlaneInformationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), planeInformationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"plane_information\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, planeInformationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in planeInformation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all planeInformation")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PlaneInformation) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("dbmodels: no plane_information provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedTS).IsZero() {
			queries.SetScanner(&o.CreatedTS, currTime)
		}
		queries.SetScanner(&o.UpdatedTS, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(planeInformationColumnsWithDefault, o)

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

	planeInformationUpsertCacheMut.RLock()
	cache, cached := planeInformationUpsertCache[key]
	planeInformationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			planeInformationAllColumns,
			planeInformationColumnsWithDefault,
			planeInformationColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			planeInformationAllColumns,
			planeInformationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert plane_information, could not build update column list")
		}

		ret := strmangle.SetComplement(planeInformationAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(planeInformationPrimaryKeyColumns) == 0 {
				return errors.New("dbmodels: unable to upsert plane_information, could not build conflict column list")
			}

			conflict = make([]string, len(planeInformationPrimaryKeyColumns))
			copy(conflict, planeInformationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"plane_information\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(planeInformationType, planeInformationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(planeInformationType, planeInformationMapping, ret)
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
		return errors.Wrap(err, "dbmodels: unable to upsert plane_information")
	}

	if !cached {
		planeInformationUpsertCacheMut.Lock()
		planeInformationUpsertCache[key] = cache
		planeInformationUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single PlaneInformation record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PlaneInformation) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no PlaneInformation provided for delete")
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), planeInformationPrimaryKeyMapping)
		sql = "DELETE FROM \"plane_information\" WHERE \"call_number\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedTS = null.TimeFrom(currTime)
		wl := []string{"deleted_ts"}
		sql = fmt.Sprintf("UPDATE \"plane_information\" SET %s WHERE \"call_number\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(planeInformationType, planeInformationMapping, append(wl, planeInformationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodels: unable to delete from plane_information")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for plane_information")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q planeInformationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no planeInformationQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_ts": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from plane_information")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for plane_information")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PlaneInformationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), planeInformationPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"plane_information\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, planeInformationPrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), planeInformationPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedTS = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_ts"}
		sql = fmt.Sprintf("UPDATE \"plane_information\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, planeInformationPrimaryKeyColumns, len(o)),
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
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from planeInformation slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for plane_information")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PlaneInformation) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPlaneInformation(ctx, exec, o.CallNumber)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PlaneInformationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PlaneInformationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), planeInformationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"plane_information\".* FROM \"plane_information\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, planeInformationPrimaryKeyColumns, len(*o)) +
		"and \"deleted_ts\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in PlaneInformationSlice")
	}

	*o = slice

	return nil
}

// PlaneInformationExists checks if the PlaneInformation row exists.
func PlaneInformationExists(ctx context.Context, exec boil.ContextExecutor, callNumber string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"plane_information\" where \"call_number\"=$1 and \"deleted_ts\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, callNumber)
	}
	row := exec.QueryRowContext(ctx, sql, callNumber)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if plane_information exists")
	}

	return exists, nil
}

// Exists checks if the PlaneInformation row exists.
func (o *PlaneInformation) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PlaneInformationExists(ctx, exec, o.CallNumber)
}
