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

// PilotDatum is an object representing the database table.
type PilotDatum struct {
	// unique uuid for the pilot
	UUID string `boil:"uuid" json:"uuid" toml:"uuid" yaml:"uuid"`
	// Pilot First Name
	PilotFirstName string `boil:"pilot_first_name" json:"pilot_first_name" toml:"pilot_first_name" yaml:"pilot_first_name"`
	// Pilot Last Name
	PilotLastName string `boil:"pilot_last_name" json:"pilot_last_name" toml:"pilot_last_name" yaml:"pilot_last_name"`
	// Email for the Pilot
	PilotEmail string `boil:"pilot_email" json:"pilot_email" toml:"pilot_email" yaml:"pilot_email"`
	// The timestamp of the created Pilot
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	// The most recent time that the Pilot object has been updated.
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	// The timestamp for when a Pilot is removed from the Young Eagle Program
	DeletedTS null.Time `boil:"deleted_ts" json:"deleted_ts,omitempty" toml:"deleted_ts" yaml:"deleted_ts,omitempty"`
	// The associated EAA Chapter the Pilot is apart of.
	EaaChapterNumber int `boil:"eaa_chapter_number" json:"eaa_chapter_number" toml:"eaa_chapter_number" yaml:"eaa_chapter_number"`
	// The status of the Pilot for a given Young Eagle Event.
	Status null.String `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`

	R *pilotDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pilotDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PilotDatumColumns = struct {
	UUID             string
	PilotFirstName   string
	PilotLastName    string
	PilotEmail       string
	CreatedAt        string
	UpdatedAt        string
	DeletedTS        string
	EaaChapterNumber string
	Status           string
}{
	UUID:             "uuid",
	PilotFirstName:   "pilot_first_name",
	PilotLastName:    "pilot_last_name",
	PilotEmail:       "pilot_email",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedTS:        "deleted_ts",
	EaaChapterNumber: "eaa_chapter_number",
	Status:           "status",
}

var PilotDatumTableColumns = struct {
	UUID             string
	PilotFirstName   string
	PilotLastName    string
	PilotEmail       string
	CreatedAt        string
	UpdatedAt        string
	DeletedTS        string
	EaaChapterNumber string
	Status           string
}{
	UUID:             "pilot_data.uuid",
	PilotFirstName:   "pilot_data.pilot_first_name",
	PilotLastName:    "pilot_data.pilot_last_name",
	PilotEmail:       "pilot_data.pilot_email",
	CreatedAt:        "pilot_data.created_at",
	UpdatedAt:        "pilot_data.updated_at",
	DeletedTS:        "pilot_data.deleted_ts",
	EaaChapterNumber: "pilot_data.eaa_chapter_number",
	Status:           "pilot_data.status",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var PilotDatumWhere = struct {
	UUID             whereHelperstring
	PilotFirstName   whereHelperstring
	PilotLastName    whereHelperstring
	PilotEmail       whereHelperstring
	CreatedAt        whereHelpernull_Time
	UpdatedAt        whereHelpernull_Time
	DeletedTS        whereHelpernull_Time
	EaaChapterNumber whereHelperint
	Status           whereHelpernull_String
}{
	UUID:             whereHelperstring{field: "\"pilot_data\".\"uuid\""},
	PilotFirstName:   whereHelperstring{field: "\"pilot_data\".\"pilot_first_name\""},
	PilotLastName:    whereHelperstring{field: "\"pilot_data\".\"pilot_last_name\""},
	PilotEmail:       whereHelperstring{field: "\"pilot_data\".\"pilot_email\""},
	CreatedAt:        whereHelpernull_Time{field: "\"pilot_data\".\"created_at\""},
	UpdatedAt:        whereHelpernull_Time{field: "\"pilot_data\".\"updated_at\""},
	DeletedTS:        whereHelpernull_Time{field: "\"pilot_data\".\"deleted_ts\""},
	EaaChapterNumber: whereHelperint{field: "\"pilot_data\".\"eaa_chapter_number\""},
	Status:           whereHelpernull_String{field: "\"pilot_data\".\"status\""},
}

// PilotDatumRels is where relationship names are stored.
var PilotDatumRels = struct {
	PilotFlightInformations string
}{
	PilotFlightInformations: "PilotFlightInformations",
}

// pilotDatumR is where relationships are stored.
type pilotDatumR struct {
	PilotFlightInformations FlightInformationSlice `boil:"PilotFlightInformations" json:"PilotFlightInformations" toml:"PilotFlightInformations" yaml:"PilotFlightInformations"`
}

// NewStruct creates a new relationship struct
func (*pilotDatumR) NewStruct() *pilotDatumR {
	return &pilotDatumR{}
}

func (r *pilotDatumR) GetPilotFlightInformations() FlightInformationSlice {
	if r == nil {
		return nil
	}
	return r.PilotFlightInformations
}

// pilotDatumL is where Load methods for each relationship are stored.
type pilotDatumL struct{}

var (
	pilotDatumAllColumns            = []string{"uuid", "pilot_first_name", "pilot_last_name", "pilot_email", "created_at", "updated_at", "deleted_ts", "eaa_chapter_number", "status"}
	pilotDatumColumnsWithoutDefault = []string{"uuid", "pilot_first_name", "pilot_last_name", "pilot_email", "eaa_chapter_number"}
	pilotDatumColumnsWithDefault    = []string{"created_at", "updated_at", "deleted_ts", "status"}
	pilotDatumPrimaryKeyColumns     = []string{"uuid"}
	pilotDatumGeneratedColumns      = []string{}
)

type (
	// PilotDatumSlice is an alias for a slice of pointers to PilotDatum.
	// This should almost always be used instead of []PilotDatum.
	PilotDatumSlice []*PilotDatum

	pilotDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pilotDatumType                 = reflect.TypeOf(&PilotDatum{})
	pilotDatumMapping              = queries.MakeStructMapping(pilotDatumType)
	pilotDatumPrimaryKeyMapping, _ = queries.BindMapping(pilotDatumType, pilotDatumMapping, pilotDatumPrimaryKeyColumns)
	pilotDatumInsertCacheMut       sync.RWMutex
	pilotDatumInsertCache          = make(map[string]insertCache)
	pilotDatumUpdateCacheMut       sync.RWMutex
	pilotDatumUpdateCache          = make(map[string]updateCache)
	pilotDatumUpsertCacheMut       sync.RWMutex
	pilotDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single pilotDatum record from the query.
func (q pilotDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PilotDatum, error) {
	o := &PilotDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for pilot_data")
	}

	return o, nil
}

// All returns all PilotDatum records from the query.
func (q pilotDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (PilotDatumSlice, error) {
	var o []*PilotDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to PilotDatum slice")
	}

	return o, nil
}

// Count returns the count of all PilotDatum records in the query.
func (q pilotDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count pilot_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q pilotDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if pilot_data exists")
	}

	return count > 0, nil
}

// PilotFlightInformations retrieves all the flight_information's FlightInformations with an executor via pilot_uuid column.
func (o *PilotDatum) PilotFlightInformations(mods ...qm.QueryMod) flightInformationQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"flight_information\".\"pilot_uuid\"=?", o.UUID),
	)

	return FlightInformations(queryMods...)
}

// LoadPilotFlightInformations allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (pilotDatumL) LoadPilotFlightInformations(ctx context.Context, e boil.ContextExecutor, singular bool, maybePilotDatum interface{}, mods queries.Applicator) error {
	var slice []*PilotDatum
	var object *PilotDatum

	if singular {
		var ok bool
		object, ok = maybePilotDatum.(*PilotDatum)
		if !ok {
			object = new(PilotDatum)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePilotDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePilotDatum))
			}
		}
	} else {
		s, ok := maybePilotDatum.(*[]*PilotDatum)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePilotDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePilotDatum))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &pilotDatumR{}
		}
		args[object.UUID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pilotDatumR{}
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
		qm.From(`flight_information`),
		qm.WhereIn(`flight_information.pilot_uuid in ?`, argsSlice...),
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
		object.R.PilotFlightInformations = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &flightInformationR{}
			}
			foreign.R.Pilot = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.UUID == foreign.PilotUUID {
				local.R.PilotFlightInformations = append(local.R.PilotFlightInformations, foreign)
				if foreign.R == nil {
					foreign.R = &flightInformationR{}
				}
				foreign.R.Pilot = local
				break
			}
		}
	}

	return nil
}

// AddPilotFlightInformations adds the given related objects to the existing relationships
// of the pilot_datum, optionally inserting them as new records.
// Appends related to o.R.PilotFlightInformations.
// Sets related.R.Pilot appropriately.
func (o *PilotDatum) AddPilotFlightInformations(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*FlightInformation) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PilotUUID = o.UUID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"flight_information\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"pilot_uuid"}),
				strmangle.WhereClause("\"", "\"", 2, flightInformationPrimaryKeyColumns),
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

			rel.PilotUUID = o.UUID
		}
	}

	if o.R == nil {
		o.R = &pilotDatumR{
			PilotFlightInformations: related,
		}
	} else {
		o.R.PilotFlightInformations = append(o.R.PilotFlightInformations, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &flightInformationR{
				Pilot: o,
			}
		} else {
			rel.R.Pilot = o
		}
	}
	return nil
}

// PilotData retrieves all the records using an executor.
func PilotData(mods ...qm.QueryMod) pilotDatumQuery {
	mods = append(mods, qm.From("\"pilot_data\""), qmhelper.WhereIsNull("\"pilot_data\".\"deleted_ts\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"pilot_data\".*"})
	}

	return pilotDatumQuery{q}
}

// FindPilotDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPilotDatum(ctx context.Context, exec boil.ContextExecutor, uUID string, selectCols ...string) (*PilotDatum, error) {
	pilotDatumObj := &PilotDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"pilot_data\" where \"uuid\"=$1 and \"deleted_ts\" is null", sel,
	)

	q := queries.Raw(query, uUID)

	err := q.Bind(ctx, exec, pilotDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from pilot_data")
	}

	return pilotDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PilotDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no pilot_data provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(pilotDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pilotDatumInsertCacheMut.RLock()
	cache, cached := pilotDatumInsertCache[key]
	pilotDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pilotDatumAllColumns,
			pilotDatumColumnsWithDefault,
			pilotDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pilotDatumType, pilotDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pilotDatumType, pilotDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"pilot_data\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"pilot_data\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "dbmodels: unable to insert into pilot_data")
	}

	if !cached {
		pilotDatumInsertCacheMut.Lock()
		pilotDatumInsertCache[key] = cache
		pilotDatumInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the PilotDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PilotDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	pilotDatumUpdateCacheMut.RLock()
	cache, cached := pilotDatumUpdateCache[key]
	pilotDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pilotDatumAllColumns,
			pilotDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_ts"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update pilot_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"pilot_data\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, pilotDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pilotDatumType, pilotDatumMapping, append(wl, pilotDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodels: unable to update pilot_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for pilot_data")
	}

	if !cached {
		pilotDatumUpdateCacheMut.Lock()
		pilotDatumUpdateCache[key] = cache
		pilotDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q pilotDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for pilot_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for pilot_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PilotDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pilotDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"pilot_data\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, pilotDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in pilotDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all pilotDatum")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PilotDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("dbmodels: no pilot_data provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(pilotDatumColumnsWithDefault, o)

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

	pilotDatumUpsertCacheMut.RLock()
	cache, cached := pilotDatumUpsertCache[key]
	pilotDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			pilotDatumAllColumns,
			pilotDatumColumnsWithDefault,
			pilotDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			pilotDatumAllColumns,
			pilotDatumPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert pilot_data, could not build update column list")
		}

		ret := strmangle.SetComplement(pilotDatumAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(pilotDatumPrimaryKeyColumns) == 0 {
				return errors.New("dbmodels: unable to upsert pilot_data, could not build conflict column list")
			}

			conflict = make([]string, len(pilotDatumPrimaryKeyColumns))
			copy(conflict, pilotDatumPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"pilot_data\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(pilotDatumType, pilotDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pilotDatumType, pilotDatumMapping, ret)
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
		return errors.Wrap(err, "dbmodels: unable to upsert pilot_data")
	}

	if !cached {
		pilotDatumUpsertCacheMut.Lock()
		pilotDatumUpsertCache[key] = cache
		pilotDatumUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single PilotDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PilotDatum) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no PilotDatum provided for delete")
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pilotDatumPrimaryKeyMapping)
		sql = "DELETE FROM \"pilot_data\" WHERE \"uuid\"=$1"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedTS = null.TimeFrom(currTime)
		wl := []string{"deleted_ts"}
		sql = fmt.Sprintf("UPDATE \"pilot_data\" SET %s WHERE \"uuid\"=$2",
			strmangle.SetParamNames("\"", "\"", 1, wl),
		)
		valueMapping, err := queries.BindMapping(pilotDatumType, pilotDatumMapping, append(wl, pilotDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "dbmodels: unable to delete from pilot_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for pilot_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q pilotDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no pilotDatumQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_ts": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from pilot_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for pilot_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PilotDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pilotDatumPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM \"pilot_data\" WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pilotDatumPrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pilotDatumPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedTS = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_ts"}
		sql = fmt.Sprintf("UPDATE \"pilot_data\" SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 2, pilotDatumPrimaryKeyColumns, len(o)),
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
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from pilotDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for pilot_data")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PilotDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPilotDatum(ctx, exec, o.UUID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PilotDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PilotDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pilotDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"pilot_data\".* FROM \"pilot_data\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, pilotDatumPrimaryKeyColumns, len(*o)) +
		"and \"deleted_ts\" is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in PilotDatumSlice")
	}

	*o = slice

	return nil
}

// PilotDatumExists checks if the PilotDatum row exists.
func PilotDatumExists(ctx context.Context, exec boil.ContextExecutor, uUID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"pilot_data\" where \"uuid\"=$1 and \"deleted_ts\" is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, uUID)
	}
	row := exec.QueryRowContext(ctx, sql, uUID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if pilot_data exists")
	}

	return exists, nil
}

// Exists checks if the PilotDatum row exists.
func (o *PilotDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PilotDatumExists(ctx, exec, o.UUID)
}
