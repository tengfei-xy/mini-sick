module mini-sick

go 1.15

replace print => ./print

replace env => ./env

require (
	env v0.0.0-00010101000000-000000000000
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.2
	github.com/go-sql-driver/mysql v1.5.0
	print v0.0.0-00010101000000-000000000000
)
