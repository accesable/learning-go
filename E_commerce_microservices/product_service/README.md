# Note During learning

The `for range` in Go
```go
func (s *Store) GetCategories(ctx context.Context) ([]types.Category, error) {
    db, err := s.queries.ListCategories(ctx)
    if err != nil {
        return nil, err
    }
    var categories []types.Category
    for i,v := range db {
        // first-case
        categories = append(categories, convertDBCategoryToPayloadCategory(&db[i]))
        // second-case 
        categories = append(categories, convertDBCategoryToPayloadCategory(&v))
    }
    return categories, nil
}

```
The `v` is in fact an copy of the struct of the items in the slices (or arrays)\
In every iteration the v is re-assign to the value of the current struct in that specific iteration if you want to use the reference to that items in array then when indexing like the `//first-case`
In Go or even in other programming language with GC(Garbage Collector like Java,C#) manual memory management is not recommended as it will interfere the Garbage Collector's algorithm and potentially lead to suboptimal performance or unexpected behavior. The garbage collector is designed to manage memory efficiently, and manual intervention can disrupt its operation.

## SQLC Notes 
- when using execresult for delete if an row is deleted then it show the affectedRows to be the number of deleted rows (e.g if one row deleted it showed one)
