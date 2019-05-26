module github.com/TerraTech/TTuuidGen

go 1.12

require (
	github.com/TerraTech/FQversion v1.0.2
	github.com/TerraTech/FQversion/tools/genVersion v0.0.0-20190530061517-863bce5f8905
	github.com/TerraTech/TTuuidGen/pkg/TTuuid v0.0.0-00010101000000-000000000000
	github.com/TerraTech/pflag v1.0.4-0.20190523114457-3d0973303b16
	github.com/spf13/pflag v1.0.3 // indirect
	golang.org/x/sys v0.0.0-20190426135247-a129542de9ae // indirect
)

replace github.com/TerraTech/TTuuidGen/pkg/TTuuid => ./pkg/TTuuid
