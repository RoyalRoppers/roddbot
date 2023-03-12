package permissions

type Permission string

const (
	CTFCreate       Permission = "ctf.create"
	CTFUpdate       Permission = "ctf.update"
	CTFArchive      Permission = "ctf.archive"
	CTFPurge        Permission = "ctf.purge"
	CTFImportCTFD   Permission = "ctf.import-ctfd"
	ChallengeCreate Permission = "chall.create"
	ChallengeDelete Permission = "chall.delete"
	ChallengeSolve  Permission = "chall.solve"
	EventCreate     Permission = "event.create"
	EventList       Permission = "event.list"
)

type Role []Permission

type roles struct {
	Player Role
	Admin  Role
}

var Roles roles = roles{
	Admin:  Role{CTFCreate, CTFUpdate, CTFArchive, CTFPurge, CTFImportCTFD, ChallengeCreate, ChallengeDelete, ChallengeSolve, EventCreate, EventList},
	Player: Role{CTFCreate, CTFImportCTFD, ChallengeCreate, ChallengeDelete, ChallengeSolve},
}
