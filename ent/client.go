// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/ei-sugimoto/gortfolio/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ei-sugimoto/gortfolio/ent/article"
	"github.com/ei-sugimoto/gortfolio/ent/articlehistory"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Article is the client for interacting with the Article builders.
	Article *ArticleClient
	// ArticleHistory is the client for interacting with the ArticleHistory builders.
	ArticleHistory *ArticleHistoryClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Article = NewArticleClient(c.config)
	c.ArticleHistory = NewArticleHistoryClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Article:        NewArticleClient(cfg),
		ArticleHistory: NewArticleHistoryClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Article:        NewArticleClient(cfg),
		ArticleHistory: NewArticleHistoryClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Article.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Article.Use(hooks...)
	c.ArticleHistory.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Article.Intercept(interceptors...)
	c.ArticleHistory.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ArticleMutation:
		return c.Article.mutate(ctx, m)
	case *ArticleHistoryMutation:
		return c.ArticleHistory.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ArticleClient is a client for the Article schema.
type ArticleClient struct {
	config
}

// NewArticleClient returns a client for the Article from the given config.
func NewArticleClient(c config) *ArticleClient {
	return &ArticleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `article.Hooks(f(g(h())))`.
func (c *ArticleClient) Use(hooks ...Hook) {
	c.hooks.Article = append(c.hooks.Article, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `article.Intercept(f(g(h())))`.
func (c *ArticleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Article = append(c.inters.Article, interceptors...)
}

// Create returns a builder for creating a Article entity.
func (c *ArticleClient) Create() *ArticleCreate {
	mutation := newArticleMutation(c.config, OpCreate)
	return &ArticleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Article entities.
func (c *ArticleClient) CreateBulk(builders ...*ArticleCreate) *ArticleCreateBulk {
	return &ArticleCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ArticleClient) MapCreateBulk(slice any, setFunc func(*ArticleCreate, int)) *ArticleCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ArticleCreateBulk{err: fmt.Errorf("calling to ArticleClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ArticleCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ArticleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Article.
func (c *ArticleClient) Update() *ArticleUpdate {
	mutation := newArticleMutation(c.config, OpUpdate)
	return &ArticleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArticleClient) UpdateOne(a *Article) *ArticleUpdateOne {
	mutation := newArticleMutation(c.config, OpUpdateOne, withArticle(a))
	return &ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArticleClient) UpdateOneID(id int) *ArticleUpdateOne {
	mutation := newArticleMutation(c.config, OpUpdateOne, withArticleID(id))
	return &ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Article.
func (c *ArticleClient) Delete() *ArticleDelete {
	mutation := newArticleMutation(c.config, OpDelete)
	return &ArticleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ArticleClient) DeleteOne(a *Article) *ArticleDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ArticleClient) DeleteOneID(id int) *ArticleDeleteOne {
	builder := c.Delete().Where(article.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArticleDeleteOne{builder}
}

// Query returns a query builder for Article.
func (c *ArticleClient) Query() *ArticleQuery {
	return &ArticleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeArticle},
		inters: c.Interceptors(),
	}
}

// Get returns a Article entity by its id.
func (c *ArticleClient) Get(ctx context.Context, id int) (*Article, error) {
	return c.Query().Where(article.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArticleClient) GetX(ctx context.Context, id int) *Article {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Article.
func (c *ArticleClient) QueryOwner(a *Article) *ArticleHistoryQuery {
	query := (&ArticleHistoryClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(article.Table, article.FieldID, id),
			sqlgraph.To(articlehistory.Table, articlehistory.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, article.OwnerTable, article.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ArticleClient) Hooks() []Hook {
	return c.hooks.Article
}

// Interceptors returns the client interceptors.
func (c *ArticleClient) Interceptors() []Interceptor {
	return c.inters.Article
}

func (c *ArticleClient) mutate(ctx context.Context, m *ArticleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ArticleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ArticleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ArticleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Article mutation op: %q", m.Op())
	}
}

// ArticleHistoryClient is a client for the ArticleHistory schema.
type ArticleHistoryClient struct {
	config
}

// NewArticleHistoryClient returns a client for the ArticleHistory from the given config.
func NewArticleHistoryClient(c config) *ArticleHistoryClient {
	return &ArticleHistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `articlehistory.Hooks(f(g(h())))`.
func (c *ArticleHistoryClient) Use(hooks ...Hook) {
	c.hooks.ArticleHistory = append(c.hooks.ArticleHistory, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `articlehistory.Intercept(f(g(h())))`.
func (c *ArticleHistoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.ArticleHistory = append(c.inters.ArticleHistory, interceptors...)
}

// Create returns a builder for creating a ArticleHistory entity.
func (c *ArticleHistoryClient) Create() *ArticleHistoryCreate {
	mutation := newArticleHistoryMutation(c.config, OpCreate)
	return &ArticleHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ArticleHistory entities.
func (c *ArticleHistoryClient) CreateBulk(builders ...*ArticleHistoryCreate) *ArticleHistoryCreateBulk {
	return &ArticleHistoryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ArticleHistoryClient) MapCreateBulk(slice any, setFunc func(*ArticleHistoryCreate, int)) *ArticleHistoryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ArticleHistoryCreateBulk{err: fmt.Errorf("calling to ArticleHistoryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ArticleHistoryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ArticleHistoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ArticleHistory.
func (c *ArticleHistoryClient) Update() *ArticleHistoryUpdate {
	mutation := newArticleHistoryMutation(c.config, OpUpdate)
	return &ArticleHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArticleHistoryClient) UpdateOne(ah *ArticleHistory) *ArticleHistoryUpdateOne {
	mutation := newArticleHistoryMutation(c.config, OpUpdateOne, withArticleHistory(ah))
	return &ArticleHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArticleHistoryClient) UpdateOneID(id int) *ArticleHistoryUpdateOne {
	mutation := newArticleHistoryMutation(c.config, OpUpdateOne, withArticleHistoryID(id))
	return &ArticleHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ArticleHistory.
func (c *ArticleHistoryClient) Delete() *ArticleHistoryDelete {
	mutation := newArticleHistoryMutation(c.config, OpDelete)
	return &ArticleHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ArticleHistoryClient) DeleteOne(ah *ArticleHistory) *ArticleHistoryDeleteOne {
	return c.DeleteOneID(ah.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ArticleHistoryClient) DeleteOneID(id int) *ArticleHistoryDeleteOne {
	builder := c.Delete().Where(articlehistory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArticleHistoryDeleteOne{builder}
}

// Query returns a query builder for ArticleHistory.
func (c *ArticleHistoryClient) Query() *ArticleHistoryQuery {
	return &ArticleHistoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeArticleHistory},
		inters: c.Interceptors(),
	}
}

// Get returns a ArticleHistory entity by its id.
func (c *ArticleHistoryClient) Get(ctx context.Context, id int) (*ArticleHistory, error) {
	return c.Query().Where(articlehistory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArticleHistoryClient) GetX(ctx context.Context, id int) *ArticleHistory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryArticle queries the article edge of a ArticleHistory.
func (c *ArticleHistoryClient) QueryArticle(ah *ArticleHistory) *ArticleQuery {
	query := (&ArticleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ah.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(articlehistory.Table, articlehistory.FieldID, id),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, articlehistory.ArticleTable, articlehistory.ArticleColumn),
		)
		fromV = sqlgraph.Neighbors(ah.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ArticleHistoryClient) Hooks() []Hook {
	return c.hooks.ArticleHistory
}

// Interceptors returns the client interceptors.
func (c *ArticleHistoryClient) Interceptors() []Interceptor {
	return c.inters.ArticleHistory
}

func (c *ArticleHistoryClient) mutate(ctx context.Context, m *ArticleHistoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ArticleHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ArticleHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ArticleHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ArticleHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ArticleHistory mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Article, ArticleHistory []ent.Hook
	}
	inters struct {
		Article, ArticleHistory []ent.Interceptor
	}
)
