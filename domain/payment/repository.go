package payment

// DBを用いたデータ取得
type IDBRepository interface {
	GetUserInfo()
}

// 外部APIを用いたデータ取得
type IAPIRepository interface {
	GetUserInfo()
}
