package data

/*
	Returns a map like:
	{ user: { Name: "Nick Landolfi"} }
	of form:
	{ <db.Kind>: <db.Model>}
*/
func Map(m Record) map[Kind]Record {
	return map[Kind]Record{
		m.Kind(): m,
	}
}
