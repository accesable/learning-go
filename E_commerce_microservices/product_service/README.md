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

## Pros & Cons of PUT request (Full Object Update)
- pros
    - Consistency: The client always sends the full representation of the resource, making it easy to ensure consistency.
    - Simpler Validation: You validate the entire object, which might be simpler than handling partial updates.
    - Clear Semantics: The PUT method is semantically defined to replace the resource at the given URI with the provided data.
- cons
    - Bandwidth: Sending the entire object can be inefficient, especially if the object is large and only a few fields are updated.
    - Concurrency: If multiple clients are updating the resource simultaneously, it might lead to overwriting changes unintentionally.
## Partial Update (PATCH)
- Pros:
    - Efficiency: Only the modified fields are sent, reducing bandwidth usage.
    - Granular Control: Allows more fine-grained updates, which can be useful in scenarios with frequent small changes.
    - Concurrency: Less likely to unintentionally overwrite other changes since only specific fields are updated.
- Cons:
    - Complexity: Handling partial updates can be more complex, especially in terms of validation and merging changes.
    - Semantic Ambiguity: The PATCH method can be semantically ambiguous, as it's not always clear how to apply partial updates.\
During Figure it how to insert with PATCH using `sqlc` . The `COALESCE()`
The `COALESCE` function is a powerful SQL feature that returns the first non-null value from a list of expressions. It is especially useful for handling situations where you might have optional or nullable columns and want to provide a default or fallback value.
```sql 
-- name: PartialUpdateItem :exec
UPDATE items
SET
    name = COALESCE(?, name),
    category_id = COALESCE(?, category_id),
    short_description = COALESCE(?, short_description),
    original_price = COALESCE(?, original_price),
    updated_at = ?
WHERE id = ?
```
