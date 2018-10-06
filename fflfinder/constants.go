package fflfinder

const (
	fflTest      = "select count(*) from ffl_lic "
	fflListQuery = " SELECT * FROM `ffl_lic` WHERE premise_zip_code like ? and " +
		" ((lic_type = '07') xor (lic_type = '01') or lic_type = '02') " +
		" order by license_name ASC"

	fflGetQueru = "SELECT * FROM `ffl_lic` WHERE id = ? "
)
