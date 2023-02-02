// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// GuildRoleMapping is an object representing the database table.
type GuildRoleMapping struct {
	GuildID      string      `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	AdminRoleID  null.String `boil:"admin_role_id" json:"admin_role_id,omitempty" toml:"admin_role_id" yaml:"admin_role_id,omitempty"`
	PlayerRoleID null.String `boil:"player_role_id" json:"player_role_id,omitempty" toml:"player_role_id" yaml:"player_role_id,omitempty"`

	R *guildRoleMappingR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L guildRoleMappingL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GuildRoleMappingColumns = struct {
	GuildID      string
	AdminRoleID  string
	PlayerRoleID string
}{
	GuildID:      "guild_id",
	AdminRoleID:  "admin_role_id",
	PlayerRoleID: "player_role_id",
}

var GuildRoleMappingTableColumns = struct {
	GuildID      string
	AdminRoleID  string
	PlayerRoleID string
}{
	GuildID:      "guild_role_mappings.guild_id",
	AdminRoleID:  "guild_role_mappings.admin_role_id",
	PlayerRoleID: "guild_role_mappings.player_role_id",
}

// Generated where

var GuildRoleMappingWhere = struct {
	GuildID      whereHelperstring
	AdminRoleID  whereHelpernull_String
	PlayerRoleID whereHelpernull_String
}{
	GuildID:      whereHelperstring{field: "\"guild_role_mappings\".\"guild_id\""},
	AdminRoleID:  whereHelpernull_String{field: "\"guild_role_mappings\".\"admin_role_id\""},
	PlayerRoleID: whereHelpernull_String{field: "\"guild_role_mappings\".\"player_role_id\""},
}

// GuildRoleMappingRels is where relationship names are stored.
var GuildRoleMappingRels = struct {
	Guild string
}{
	Guild: "Guild",
}

// guildRoleMappingR is where relationships are stored.
type guildRoleMappingR struct {
	Guild *Guild `boil:"Guild" json:"Guild" toml:"Guild" yaml:"Guild"`
}

// NewStruct creates a new relationship struct
func (*guildRoleMappingR) NewStruct() *guildRoleMappingR {
	return &guildRoleMappingR{}
}

func (r *guildRoleMappingR) GetGuild() *Guild {
	if r == nil {
		return nil
	}
	return r.Guild
}

// guildRoleMappingL is where Load methods for each relationship are stored.
type guildRoleMappingL struct{}

var (
	guildRoleMappingAllColumns            = []string{"guild_id", "admin_role_id", "player_role_id"}
	guildRoleMappingColumnsWithoutDefault = []string{"guild_id"}
	guildRoleMappingColumnsWithDefault    = []string{"admin_role_id", "player_role_id"}
	guildRoleMappingPrimaryKeyColumns     = []string{"guild_id"}
	guildRoleMappingGeneratedColumns      = []string{}
)

type (
	// GuildRoleMappingSlice is an alias for a slice of pointers to GuildRoleMapping.
	// This should almost always be used instead of []GuildRoleMapping.
	GuildRoleMappingSlice []*GuildRoleMapping
	// GuildRoleMappingHook is the signature for custom GuildRoleMapping hook methods
	GuildRoleMappingHook func(context.Context, boil.ContextExecutor, *GuildRoleMapping) error

	guildRoleMappingQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	guildRoleMappingType                 = reflect.TypeOf(&GuildRoleMapping{})
	guildRoleMappingMapping              = queries.MakeStructMapping(guildRoleMappingType)
	guildRoleMappingPrimaryKeyMapping, _ = queries.BindMapping(guildRoleMappingType, guildRoleMappingMapping, guildRoleMappingPrimaryKeyColumns)
	guildRoleMappingInsertCacheMut       sync.RWMutex
	guildRoleMappingInsertCache          = make(map[string]insertCache)
	guildRoleMappingUpdateCacheMut       sync.RWMutex
	guildRoleMappingUpdateCache          = make(map[string]updateCache)
	guildRoleMappingUpsertCacheMut       sync.RWMutex
	guildRoleMappingUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var guildRoleMappingAfterSelectHooks []GuildRoleMappingHook

var guildRoleMappingBeforeInsertHooks []GuildRoleMappingHook
var guildRoleMappingAfterInsertHooks []GuildRoleMappingHook

var guildRoleMappingBeforeUpdateHooks []GuildRoleMappingHook
var guildRoleMappingAfterUpdateHooks []GuildRoleMappingHook

var guildRoleMappingBeforeDeleteHooks []GuildRoleMappingHook
var guildRoleMappingAfterDeleteHooks []GuildRoleMappingHook

var guildRoleMappingBeforeUpsertHooks []GuildRoleMappingHook
var guildRoleMappingAfterUpsertHooks []GuildRoleMappingHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *GuildRoleMapping) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildRoleMappingAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *GuildRoleMapping) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildRoleMappingBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *GuildRoleMapping) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildRoleMappingAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *GuildRoleMapping) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildRoleMappingBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *GuildRoleMapping) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildRoleMappingAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *GuildRoleMapping) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildRoleMappingBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *GuildRoleMapping) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildRoleMappingAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *GuildRoleMapping) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildRoleMappingBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *GuildRoleMapping) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range guildRoleMappingAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGuildRoleMappingHook registers your hook function for all future operations.
func AddGuildRoleMappingHook(hookPoint boil.HookPoint, guildRoleMappingHook GuildRoleMappingHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		guildRoleMappingAfterSelectHooks = append(guildRoleMappingAfterSelectHooks, guildRoleMappingHook)
	case boil.BeforeInsertHook:
		guildRoleMappingBeforeInsertHooks = append(guildRoleMappingBeforeInsertHooks, guildRoleMappingHook)
	case boil.AfterInsertHook:
		guildRoleMappingAfterInsertHooks = append(guildRoleMappingAfterInsertHooks, guildRoleMappingHook)
	case boil.BeforeUpdateHook:
		guildRoleMappingBeforeUpdateHooks = append(guildRoleMappingBeforeUpdateHooks, guildRoleMappingHook)
	case boil.AfterUpdateHook:
		guildRoleMappingAfterUpdateHooks = append(guildRoleMappingAfterUpdateHooks, guildRoleMappingHook)
	case boil.BeforeDeleteHook:
		guildRoleMappingBeforeDeleteHooks = append(guildRoleMappingBeforeDeleteHooks, guildRoleMappingHook)
	case boil.AfterDeleteHook:
		guildRoleMappingAfterDeleteHooks = append(guildRoleMappingAfterDeleteHooks, guildRoleMappingHook)
	case boil.BeforeUpsertHook:
		guildRoleMappingBeforeUpsertHooks = append(guildRoleMappingBeforeUpsertHooks, guildRoleMappingHook)
	case boil.AfterUpsertHook:
		guildRoleMappingAfterUpsertHooks = append(guildRoleMappingAfterUpsertHooks, guildRoleMappingHook)
	}
}

// One returns a single guildRoleMapping record from the query.
func (q guildRoleMappingQuery) One(ctx context.Context, exec boil.ContextExecutor) (*GuildRoleMapping, error) {
	o := &GuildRoleMapping{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for guild_role_mappings")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all GuildRoleMapping records from the query.
func (q guildRoleMappingQuery) All(ctx context.Context, exec boil.ContextExecutor) (GuildRoleMappingSlice, error) {
	var o []*GuildRoleMapping

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to GuildRoleMapping slice")
	}

	if len(guildRoleMappingAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all GuildRoleMapping records in the query.
func (q guildRoleMappingQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count guild_role_mappings rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q guildRoleMappingQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if guild_role_mappings exists")
	}

	return count > 0, nil
}

// Guild pointed to by the foreign key.
func (o *GuildRoleMapping) Guild(mods ...qm.QueryMod) guildQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.GuildID),
	}

	queryMods = append(queryMods, mods...)

	return Guilds(queryMods...)
}

// LoadGuild allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (guildRoleMappingL) LoadGuild(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGuildRoleMapping interface{}, mods queries.Applicator) error {
	var slice []*GuildRoleMapping
	var object *GuildRoleMapping

	if singular {
		var ok bool
		object, ok = maybeGuildRoleMapping.(*GuildRoleMapping)
		if !ok {
			object = new(GuildRoleMapping)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeGuildRoleMapping)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeGuildRoleMapping))
			}
		}
	} else {
		s, ok := maybeGuildRoleMapping.(*[]*GuildRoleMapping)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeGuildRoleMapping)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeGuildRoleMapping))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &guildRoleMappingR{}
		}
		args = append(args, object.GuildID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &guildRoleMappingR{}
			}

			for _, a := range args {
				if a == obj.GuildID {
					continue Outer
				}
			}

			args = append(args, obj.GuildID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`guilds`),
		qm.WhereIn(`guilds.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Guild")
	}

	var resultSlice []*Guild
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Guild")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for guilds")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for guilds")
	}

	if len(guildAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Guild = foreign
		if foreign.R == nil {
			foreign.R = &guildR{}
		}
		foreign.R.GuildRoleMapping = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GuildID == foreign.ID {
				local.R.Guild = foreign
				if foreign.R == nil {
					foreign.R = &guildR{}
				}
				foreign.R.GuildRoleMapping = local
				break
			}
		}
	}

	return nil
}

// SetGuild of the guildRoleMapping to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildRoleMapping.
func (o *GuildRoleMapping) SetGuild(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Guild) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"guild_role_mappings\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"guild_id"}),
		strmangle.WhereClause("\"", "\"", 2, guildRoleMappingPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.GuildID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GuildID = related.ID
	if o.R == nil {
		o.R = &guildRoleMappingR{
			Guild: related,
		}
	} else {
		o.R.Guild = related
	}

	if related.R == nil {
		related.R = &guildR{
			GuildRoleMapping: o,
		}
	} else {
		related.R.GuildRoleMapping = o
	}

	return nil
}

// GuildRoleMappings retrieves all the records using an executor.
func GuildRoleMappings(mods ...qm.QueryMod) guildRoleMappingQuery {
	mods = append(mods, qm.From("\"guild_role_mappings\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"guild_role_mappings\".*"})
	}

	return guildRoleMappingQuery{q}
}

// FindGuildRoleMapping retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGuildRoleMapping(ctx context.Context, exec boil.ContextExecutor, guildID string, selectCols ...string) (*GuildRoleMapping, error) {
	guildRoleMappingObj := &GuildRoleMapping{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"guild_role_mappings\" where \"guild_id\"=$1", sel,
	)

	q := queries.Raw(query, guildID)

	err := q.Bind(ctx, exec, guildRoleMappingObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from guild_role_mappings")
	}

	if err = guildRoleMappingObj.doAfterSelectHooks(ctx, exec); err != nil {
		return guildRoleMappingObj, err
	}

	return guildRoleMappingObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *GuildRoleMapping) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no guild_role_mappings provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(guildRoleMappingColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	guildRoleMappingInsertCacheMut.RLock()
	cache, cached := guildRoleMappingInsertCache[key]
	guildRoleMappingInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			guildRoleMappingAllColumns,
			guildRoleMappingColumnsWithDefault,
			guildRoleMappingColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(guildRoleMappingType, guildRoleMappingMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(guildRoleMappingType, guildRoleMappingMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"guild_role_mappings\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"guild_role_mappings\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into guild_role_mappings")
	}

	if !cached {
		guildRoleMappingInsertCacheMut.Lock()
		guildRoleMappingInsertCache[key] = cache
		guildRoleMappingInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the GuildRoleMapping.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *GuildRoleMapping) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	guildRoleMappingUpdateCacheMut.RLock()
	cache, cached := guildRoleMappingUpdateCache[key]
	guildRoleMappingUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			guildRoleMappingAllColumns,
			guildRoleMappingPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update guild_role_mappings, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"guild_role_mappings\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, guildRoleMappingPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(guildRoleMappingType, guildRoleMappingMapping, append(wl, guildRoleMappingPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update guild_role_mappings row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for guild_role_mappings")
	}

	if !cached {
		guildRoleMappingUpdateCacheMut.Lock()
		guildRoleMappingUpdateCache[key] = cache
		guildRoleMappingUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q guildRoleMappingQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for guild_role_mappings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for guild_role_mappings")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GuildRoleMappingSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), guildRoleMappingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"guild_role_mappings\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, guildRoleMappingPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in guildRoleMapping slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all guildRoleMapping")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *GuildRoleMapping) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no guild_role_mappings provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(guildRoleMappingColumnsWithDefault, o)

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

	guildRoleMappingUpsertCacheMut.RLock()
	cache, cached := guildRoleMappingUpsertCache[key]
	guildRoleMappingUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			guildRoleMappingAllColumns,
			guildRoleMappingColumnsWithDefault,
			guildRoleMappingColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			guildRoleMappingAllColumns,
			guildRoleMappingPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert guild_role_mappings, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(guildRoleMappingPrimaryKeyColumns))
			copy(conflict, guildRoleMappingPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"guild_role_mappings\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(guildRoleMappingType, guildRoleMappingMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(guildRoleMappingType, guildRoleMappingMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert guild_role_mappings")
	}

	if !cached {
		guildRoleMappingUpsertCacheMut.Lock()
		guildRoleMappingUpsertCache[key] = cache
		guildRoleMappingUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single GuildRoleMapping record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *GuildRoleMapping) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no GuildRoleMapping provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), guildRoleMappingPrimaryKeyMapping)
	sql := "DELETE FROM \"guild_role_mappings\" WHERE \"guild_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from guild_role_mappings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for guild_role_mappings")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q guildRoleMappingQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no guildRoleMappingQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from guild_role_mappings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for guild_role_mappings")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GuildRoleMappingSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(guildRoleMappingBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), guildRoleMappingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"guild_role_mappings\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, guildRoleMappingPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from guildRoleMapping slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for guild_role_mappings")
	}

	if len(guildRoleMappingAfterDeleteHooks) != 0 {
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
func (o *GuildRoleMapping) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGuildRoleMapping(ctx, exec, o.GuildID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GuildRoleMappingSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GuildRoleMappingSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), guildRoleMappingPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"guild_role_mappings\".* FROM \"guild_role_mappings\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, guildRoleMappingPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GuildRoleMappingSlice")
	}

	*o = slice

	return nil
}

// GuildRoleMappingExists checks if the GuildRoleMapping row exists.
func GuildRoleMappingExists(ctx context.Context, exec boil.ContextExecutor, guildID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"guild_role_mappings\" where \"guild_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, guildID)
	}
	row := exec.QueryRowContext(ctx, sql, guildID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if guild_role_mappings exists")
	}

	return exists, nil
}

// Exists checks if the GuildRoleMapping row exists.
func (o *GuildRoleMapping) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return GuildRoleMappingExists(ctx, exec, o.GuildID)
}