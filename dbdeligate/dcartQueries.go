package dbdeligate

const (
	dcartTest        = "select count(*) from dcart_user "
	dcartGetByStore  = "select * from dcart_user where store_url = ? "
	dcartInsertStore = "insert into dcart_user (store_url, public_key, token, action, entered, enabled) values(?, ?, ?, ?, ?, ?)"
	dcartUpdateStore = "update dcart_user set public_key = ?, token = ?, action = ?, modified = ?, enabled = true where id = ? "
	dcartRemoveStore = "update dcart_user set action = ?, modified = ?, enabled = false where id = ? "
)
