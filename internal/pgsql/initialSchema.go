package pgsql

/*

qs := []string{
	"CREATE TEMP TABLE regions (id integer NOT NULL, code character varying(4) NOT NULL, name text NOT NULL)",
	"CREATE TEMP TABLE departments (id integer NOT NULL, code character varying(4) NOT NULL, name text NOT NULL, region_code character varying(4) NOT NULL)",

	"INSERT INTO regions (id, code, name) VALUES (1, '01', 'Guadeloupe')",
	"INSERT INTO regions (id, code, name) VALUES (2, '02', 'Martinique')",

	"INSERT INTO departments (id, code, name, region_code) VALUES (1, '01', 'Ain', '82')",
	"INSERT INTO departments (id, code, name, region_code) VALUES (2, '02', 'Aisne', '22')",

	"ALTER TABLE ONLY regions ADD CONSTRAINT regions_id_key UNIQUE (id)",
	"ALTER TABLE ONLY regions ADD CONSTRAINT regions_code_key UNIQUE (code)",

	"ALTER TABLE ONLY departments ADD CONSTRAINT departments_id_key UNIQUE (id)",
	"ALTER TABLE ONLY departments ADD CONSTRAINT departments_code_key UNIQUE (code)",

	"ALTER TABLE ONLY departments ADD CONSTRAINT departments_region_fkey FOREIGN KEY (region_code) REFERENCES regions(code)",
}
for _, q := range qs {
	_, err := db.Exec(q)
	if err != nil {
		return err
	}
}

return nil
*/
