## brief

a piece of error handling codelet.

this code only demonstrates how to hide sql.ErrNoRows to the caller.

the Get() method defined in the pkg/project/dao/dao.go file will throw a sql.ErrNoRows error, if the Scan() detects no matching rows. after that I'll construct a new error without the sql.ErrNoRows error wrapped, so that the caller is informed with the error reason, but unware of the underlying sql implementation and no need to import the sql package. 

## notice

CAUTION: it will throw segmentation fault error since the required database connection is uninitialized.
