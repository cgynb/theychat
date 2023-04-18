package cache

import (
	"strconv"
	"theychat/utils"
)

func AddGroup(userId uint) (gid string) {
	gid = utils.GenGroupId()
	RDB.SAdd(gid, userId)
	return
}

func JoinGroup(cap int, gid string, userId uint) (ok bool) {
	if RDB.SCard(gid).Val() < int64(cap) {
		RDB.SAdd(gid, userId)
		ok = true
	} else {
		ok = false
	}
	return
}

func GetGroupMember(gid string) (members []uint) {
	ms := RDB.SMembers(gid).Val()
	members = make([]uint, cap(ms), len(ms))
	for i, m := range ms {
		val, _ := strconv.Atoi(m)
		members[i] = uint(val)
	}
	return
}

func IsMember(userId uint, gid string) (ok bool) {
	return RDB.SIsMember(gid, strconv.Itoa(int(userId))).Val()
}

func HasGroup(gid string) (ok bool) {
	return RDB.Exists(gid).Val() != 0
}
