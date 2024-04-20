# pbsql

[![Go Report Card](https://goreportcard.com/badge/github.com/rmilejcz/pbsql)](https://goreportcard.com/report/github.com/rmilejcz/pbsql)
[![codecov](https://codecov.io/gh/rmilejcz/pbsql/branch/master/graph/badge.svg)](https://codecov.io/gh/rmilejcz/pbsql)

A reflection-based, lightweight, efficient, and nullsafe query generator for protobufs

## Why?

Mostly because our database and current stack heavily utilizes null, and no existing sql library has
helped us solve the problem cleanly. By simply tagging protobuf fields as `nullable` and reading that value during
reflection we solve the problem by adding `ifnull()` (or `coalesce()`) and some runtime determined default value

## Dependencies

- [sqlx](https://github.com/jmoiron/sqlx)
- [protoc-go-inject-tags](https://github.com/favadi/protoc-go-inject-tag)

## Usage

This library assumes that:

- you are okay with sending and recieving the same message over the wire, e.g.:

  ```
  rpc Create(User) returns (User) {}
  rpc Get(User)    returns (User) {}
  rpc Update(User) returns (User) {}
  ```

- your proto structure is friendly to simple field masks (i.e. nested messages are not supported):

  ```
  map<string, int> field_mask = 1;
  ```

- all time values can be represented as `string` instead of `protobuf.Timestamp`:
  - this is especially convenient for SQL since a time value of `2019-09-12 08:30:00` can be queried with string literals
    such as `%2019%`, `%2019-09%`, etc

A protobuf message should utilize the tags `db:`, `nullable:`, and `primary_key:`

- `db:`
  - behaves exactly like sqlx, should be set to the database column name
- `nullable:`
  - set this to any non attempt string to prevent reading null values.
- `primary_key:`
  - make sure you denote the primary key to prevent it from being written into insert and update statements

### Usage example

Your proto should looke something like this:

```
service UserSvc{
  rpc Create(User) returns (User) {}
}

message User {
  // @inject-tag: db:"id" primary_key:"y"
  int32 id = 1;
  // @inject-tag: db:"first_name"
  string first_name = 2;
  // @inject-tag: db:"last_name"
  string last_name = 3;
  // @inject-tag: db:"user_phone" nullable:"y"
  string phone = 4;

  map<string, int> field_mask = 5;
}
```

Implementation:

```go
func (s *Service) Create(ctx context.Context, req *User) (*User, error) {

  // create query string, this also returns a []interface{} slice with each arg in order
  // this query string uses `?` for interpolation (not named vars)
  qry, args, err := pbsql.BuildCreateQuery("user", req)
  if err != nil {
    return nil, status.Errorf(codes.Internal, "failed to create prepared query string %v", err)
  }

  // now simply call the appropriate function and expand the args slice
  // sqlx.DB.Exec is great if you don't need a result set otherwise use
  // sqlx.DB.Queryx or sqlx.DB.QueryRowx
  res, err := s.DB.Exec(qry, args...)
  if err != nil {
    return nil, status.Errorf(codes.Internal, "failed to execute query %v", err)
  }
}
```

### Collation
To force collation of particular field, you can add `collation` tag with collation name as the value, e.g. `collation:"utf8mb4_0900_ai_ci"`.
The internal default collation is `utf8mb4_0900_ai_ci` so to use that collation we can simply define `collation:"default"` and the marked field will use `utf8mb4_0900_ai_ci`. 

Example:
```
type Task struct {
	..... 

	// @inject_tag: db:"date_performed" nullable:"y" collation:"default"
	DatePerformed string `protobuf:"bytes,30,opt,name=date_performed,json=datePerformed,proto3" json:"date_performed,omitempty" db:"date_performed" nullable:"y" collation:"default"`
	// @inject_tag: db:"spiff_tool_id" nullable:"y"
	SpiffToolId string `protobuf:"bytes,31,opt,name=spiff_tool_id,json=spiffToolId,proto3" json:"spiff_tool_id,omitempty" db:"spiff_tool_id" nullable:"y"`
	// @inject_tag: db:"spiff_tool_closeout_date" nullable:"y" collation:"default"
	OwnerName             string   `db:"owner_name" select_func:"name_of_user" func_arg_name:"external_id" collation:"default"`
	.....
}
```

## Caveats

The query builder doesn't handle any sort of limit or offset behavior, but since it returns a plain string this would be simple to implement:

```go
func (s *UserSvc) List(ctx context.Context, req *User) (*User, error) {
  qry, args, err := pbsql.BuildReadQuery("user", req)
  if err != nil {
    return nil, status.Errorf(codes.Internal, "failed to create prepared query string %v", err)
  }

  // add raw SQL to your query string
  limitedQry := qry + " OFFSET ?, LIMIT ?"

  // append the offset and limit to args before expanding
  rows, err := s.DB.Queryx(limitedQry, append(args, 0, 50)...)
  if err != nil {
    return status.Errorf(codes.Internal, "failed to execute query %v", err)
  }
  defer rows.Close()
}
```

## Roadmap

- [ ] Support a `default_value` tag in favor of guessing the default value at runtime
- [ ] Support including foreign key related entities in query results
- [ ] Support field masks in read queries
- [ ] Support other DBMS
