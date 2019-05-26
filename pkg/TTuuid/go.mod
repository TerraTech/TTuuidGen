module github.com/TerraTech/TTuuidGen/pkg/TTuuid

go 1.12

require (
	github.com/kr/pretty v0.1.0 // indirect
	github.com/satori/go.uuid v1.2.1-0.20180103161547-0ef6afb2f6cd
	github.com/stretchr/testify v1.3.0
	golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace github.com/TerraTech/TTuuidGen/pkg/TTuuid => ./
