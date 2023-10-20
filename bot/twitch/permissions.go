package twitch

const (
	permissionNoBadge     uint = 1
	permissionSubscriber  uint = permissionNoBadge | (1 << 1)
	permissionVip         uint = permissionSubscriber | (1 << 2)
	permissionModerator   uint = permissionVip | (1 << 3)
	permissionBroadcaster uint = permissionModerator | (1 << 4)
)

func toPermissionsMap(permissions []string) map[string]int {
	permissionMap := make(map[string]int)

	for _, v := range permissions {
		permissionMap[v] = 1
	}

	return permissionMap
}

func getPermissionsMask(badges map[string]int) uint {
	permissions := permissionNoBadge

	if badge, ok := badges["broadcaster"]; ok && badge == 1 {
		permissions = permissions | permissionBroadcaster
	}

	if badge, ok := badges["moderator"]; ok && badge == 1 {
		permissions = permissions | permissionModerator
	}

	if badge, ok := badges["vip"]; ok && badge == 1 {
		permissions = permissions | permissionVip
	}

	if badge, ok := badges["subscriber"]; ok && badge == 1 {
		permissions = permissions | permissionSubscriber
	}

	return permissions
}

func hasPermissions(permissions uint, permissionLevel uint) bool {
	return permissions&permissionLevel != 0
}
