# pbsql

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
  // create query string
  qryString := pbsql.BuildCreateQuery("user", req)
  // since the query string uses named args we can use sqlx.Named to convert it
  // to a query string using `?` instead, sqlx.Named also returns a slice of all
  // the target args
  qry, args, err := sqlx.Named(qryStr, req)
  if err != nil {
    return nil, status.Errorf(codes.Internal, "failed to create prepared query string %v", err)
  }

  // now simple call the appropriate function and expand the args slice
  // sqlx.DB.Exec is great if you don't need a result set otherwise use
  // sqlx.DB.Queryx or sqlx.DB.QueryRowx
  res, err := s.DB.Exec(qry, args...)
  if err != nil {
    return nil, status.Errorf(codes.Internal, "failed to execute query %v", err)
  }
}
```

## Caveats

The query builder doesn't handle any sort of limit or offset behavior, but since it returns a plain string this would be simple to implement:

```go
func (s *UserSvc) List(ctx context.Context, req *User) (*User, error) {
  qryString := pbsql.BuildReadQuery("user", req)
  qry, args, err := sqlx.Named(qryStr, req)
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
