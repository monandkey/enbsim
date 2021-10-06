module github.com/monandkey/enbsim

go 1.16

replace local.packages/encoding/s1ap => ./pkg/encoding/s1ap

require (
	github.com/ishidawataru/sctp v0.0.0-20210707070123-9a39160e9062 // indirect
	local.packages/encoding/s1ap v0.0.0-00010101000000-000000000000 // indirect
)
