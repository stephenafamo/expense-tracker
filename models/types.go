// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Type is an object representing the database table.
type Type struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *typeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L typeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TypeColumns = struct {
	ID        string
	Name      string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// TypeRels is where relationship names are stored.
var TypeRels = struct {
	Transactions string
}{
	Transactions: "Transactions",
}

// typeR is where relationships are stored.
type typeR struct {
	Transactions TransactionSlice
}

// NewStruct creates a new relationship struct
func (*typeR) NewStruct() *typeR {
	return &typeR{}
}

// typeL is where Load methods for each relationship are stored.
type typeL struct{}

var (
	typeColumns               = []string{"id", "name", "created_at", "updated_at"}
	typeColumnsWithoutDefault = []string{"name", "updated_at"}
	typeColumnsWithDefault    = []string{"id", "created_at"}
	typePrimaryKeyColumns     = []string{"id"}
)

type (
	// TypeSlice is an alias for a slice of pointers to Type.
	// This should generally be used opposed to []Type.
	TypeSlice []*Type
	// TypeHook is the signature for custom Type hook methods
	TypeHook func(context.Context, boil.ContextExecutor, *Type) error

	typeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	typeType                 = reflect.TypeOf(&Type{})
	typeMapping              = queries.MakeStructMapping(typeType)
	typePrimaryKeyMapping, _ = queries.BindMapping(typeType, typeMapping, typePrimaryKeyColumns)
	typeInsertCacheMut       sync.RWMutex
	typeInsertCache          = make(map[string]insertCache)
	typeUpdateCacheMut       sync.RWMutex
	typeUpdateCache          = make(map[string]updateCache)
	typeUpsertCacheMut       sync.RWMutex
	typeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var typeBeforeInsertHooks []TypeHook
var typeBeforeUpdateHooks []TypeHook
var typeBeforeDeleteHooks []TypeHook
var typeBeforeUpsertHooks []TypeHook

var typeAfterInsertHooks []TypeHook
var typeAfterSelectHooks []TypeHook
var typeAfterUpdateHooks []TypeHook
var typeAfterDeleteHooks []TypeHook
var typeAfterUpsertHooks []TypeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Type) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range typeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Type) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range typeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Type) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range typeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Type) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range typeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Type) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range typeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Type) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range typeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Type) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range typeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Type) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range typeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Type) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range typeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTypeHook registers your hook function for all future operations.
func AddTypeHook(hookPoint boil.HookPoint, typeHook TypeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		typeBeforeInsertHooks = append(typeBeforeInsertHooks, typeHook)
	case boil.BeforeUpdateHook:
		typeBeforeUpdateHooks = append(typeBeforeUpdateHooks, typeHook)
	case boil.BeforeDeleteHook:
		typeBeforeDeleteHooks = append(typeBeforeDeleteHooks, typeHook)
	case boil.BeforeUpsertHook:
		typeBeforeUpsertHooks = append(typeBeforeUpsertHooks, typeHook)
	case boil.AfterInsertHook:
		typeAfterInsertHooks = append(typeAfterInsertHooks, typeHook)
	case boil.AfterSelectHook:
		typeAfterSelectHooks = append(typeAfterSelectHooks, typeHook)
	case boil.AfterUpdateHook:
		typeAfterUpdateHooks = append(typeAfterUpdateHooks, typeHook)
	case boil.AfterDeleteHook:
		typeAfterDeleteHooks = append(typeAfterDeleteHooks, typeHook)
	case boil.AfterUpsertHook:
		typeAfterUpsertHooks = append(typeAfterUpsertHooks, typeHook)
	}
}

// One returns a single type record from the query.
func (q typeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Type, error) {
	o := &Type{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for types")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Type records from the query.
func (q typeQuery) All(ctx context.Context, exec boil.ContextExecutor) (TypeSlice, error) {
	var o []*Type

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Type slice")
	}

	if len(typeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Type records in the query.
func (q typeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count types rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q typeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if types exists")
	}

	return count > 0, nil
}

// Transactions retrieves all the transaction's Transactions with an executor.
func (o *Type) Transactions(mods ...qm.QueryMod) transactionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"transactions\".\"type_id\"=?", o.ID),
	)

	query := Transactions(queryMods...)
	queries.SetFrom(query.Query, "\"transactions\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"transactions\".*"})
	}

	return query
}

// LoadTransactions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (typeL) LoadTransactions(ctx context.Context, e boil.ContextExecutor, singular bool, maybeType interface{}, mods queries.Applicator) error {
	var slice []*Type
	var object *Type

	if singular {
		object = maybeType.(*Type)
	} else {
		slice = *maybeType.(*[]*Type)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &typeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &typeR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`transactions`), qm.WhereIn(`type_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load transactions")
	}

	var resultSlice []*Transaction
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice transactions")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on transactions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for transactions")
	}

	if len(transactionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Transactions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &transactionR{}
			}
			foreign.R.Type = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.TypeID) {
				local.R.Transactions = append(local.R.Transactions, foreign)
				if foreign.R == nil {
					foreign.R = &transactionR{}
				}
				foreign.R.Type = local
				break
			}
		}
	}

	return nil
}

// AddTransactions adds the given related objects to the existing relationships
// of the type, optionally inserting them as new records.
// Appends related to o.R.Transactions.
// Sets related.R.Type appropriately.
func (o *Type) AddTransactions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Transaction) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.TypeID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"transactions\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"type_id"}),
				strmangle.WhereClause("\"", "\"", 2, transactionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.TypeID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &typeR{
			Transactions: related,
		}
	} else {
		o.R.Transactions = append(o.R.Transactions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &transactionR{
				Type: o,
			}
		} else {
			rel.R.Type = o
		}
	}
	return nil
}

// SetTransactions removes all previously related items of the
// type replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Type's Transactions accordingly.
// Replaces o.R.Transactions with related.
// Sets related.R.Type's Transactions accordingly.
func (o *Type) SetTransactions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Transaction) error {
	query := "update \"transactions\" set \"type_id\" = null where \"type_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Transactions {
			queries.SetScanner(&rel.TypeID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Type = nil
		}

		o.R.Transactions = nil
	}
	return o.AddTransactions(ctx, exec, insert, related...)
}

// RemoveTransactions relationships from objects passed in.
// Removes related items from R.Transactions (uses pointer comparison, removal does not keep order)
// Sets related.R.Type.
func (o *Type) RemoveTransactions(ctx context.Context, exec boil.ContextExecutor, related ...*Transaction) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.TypeID, nil)
		if rel.R != nil {
			rel.R.Type = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("type_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Transactions {
			if rel != ri {
				continue
			}

			ln := len(o.R.Transactions)
			if ln > 1 && i < ln-1 {
				o.R.Transactions[i] = o.R.Transactions[ln-1]
			}
			o.R.Transactions = o.R.Transactions[:ln-1]
			break
		}
	}

	return nil
}

// Types retrieves all the records using an executor.
func Types(mods ...qm.QueryMod) typeQuery {
	mods = append(mods, qm.From("\"types\""))
	return typeQuery{NewQuery(mods...)}
}

// FindType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindType(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Type, error) {
	typeObj := &Type{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"types\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, typeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from types")
	}

	return typeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Type) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no types provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(typeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	typeInsertCacheMut.RLock()
	cache, cached := typeInsertCache[key]
	typeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			typeColumns,
			typeColumnsWithDefault,
			typeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(typeType, typeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(typeType, typeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"types\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"types\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into types")
	}

	if !cached {
		typeInsertCacheMut.Lock()
		typeInsertCache[key] = cache
		typeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Type.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Type) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	typeUpdateCacheMut.RLock()
	cache, cached := typeUpdateCache[key]
	typeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			typeColumns,
			typePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update types, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"types\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, typePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(typeType, typeMapping, append(wl, typePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update types row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for types")
	}

	if !cached {
		typeUpdateCacheMut.Lock()
		typeUpdateCache[key] = cache
		typeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q typeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for types")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), typePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"types\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, typePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in type slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all type")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Type) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no types provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(typeColumnsWithDefault, o)

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

	typeUpsertCacheMut.RLock()
	cache, cached := typeUpsertCache[key]
	typeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			typeColumns,
			typeColumnsWithDefault,
			typeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			typeColumns,
			typePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert types, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(typePrimaryKeyColumns))
			copy(conflict, typePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryCockroachDB(dialect, "\"types\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(typeType, typeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(typeType, typeMapping, ret)
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

	if boil.DebugMode {
		_, _ = fmt.Fprintln(boil.DebugWriter, cache.query)
		_, _ = fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // CockcorachDB doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert types")
	}

	if !cached {
		typeUpsertCacheMut.Lock()
		typeUpsertCache[key] = cache
		typeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Type record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Type) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Type provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), typePrimaryKeyMapping)
	sql := "DELETE FROM \"types\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for types")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q typeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no typeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from types")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for types")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Type slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(typeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), typePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, typePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from type slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for types")
	}

	if len(typeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Type) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindType(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), typePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"types\".* FROM \"types\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, typePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TypeSlice")
	}

	*o = slice

	return nil
}

// TypeExists checks if the Type row exists.
func TypeExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"types\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if types exists")
	}

	return exists, nil
}
