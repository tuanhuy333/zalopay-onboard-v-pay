package proto

//go:generate protoc ./disbursement.proto --go_out=../../pkg/disbursement/pb --go_opt=Mdisbursement.proto=/.;pb --go_opt=paths=source_relative --go-grpc_out=../../pkg/disbursement/pb --go-grpc_opt=paths=source_relative --go-grpc_opt=Mdisbursement.proto=/.;pb
//go:generate protoc-go-inject-tag -input=../../pkg/disbursement/pb/*.pb.go
