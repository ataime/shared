package enum

type Gender Gender

const (
	GenderNone   Gender = Gender(admin_grpc.Gender_UNKNOWN)
	GenderMale   Gender = Gender(admin_grpc.Gender_MALE)
	GenderFemale Gender = Gender(admin_grpc.Gender_FEMALE)
	GenderOther  Gender = Gender(admin_grpc.Gender_OTHER)
)
