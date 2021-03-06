// Package models contains the types for schema 'xodb'.
package models

// Code generated by xo. DO NOT EDIT.

// SSIdentity represents a row from '[custom s_s_identity]'.
type SSIdentity struct {
	TableName string // table_name
}

// SSIdentities runs a custom query, returning results as SSIdentity.
func SSIdentities(db XODB, schema string) ([]*SSIdentity, error) {
	var err error

	// sql query
	const sqlstr = `SELECT o.name as table_name ` +
		`FROM sys.objects o inner join sys.columns c on o.object_id = c.object_id ` +
		`WHERE c.is_identity = 1 ` +
		`AND schema_name(o.schema_id) = @p1 AND o.type = 'U'`

	// run query
	XOLog(sqlstr, schema)
	q, err := db.Query(sqlstr, schema)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*SSIdentity{}
	for q.Next() {
		ssi := SSIdentity{}

		// scan
		err = q.Scan(&ssi.TableName)
		if err != nil {
			return nil, err
		}

		res = append(res, &ssi)
	}

	return res, nil
}
