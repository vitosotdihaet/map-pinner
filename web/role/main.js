class Role {
    static allRoles = null
    static currentRoleID = null

    static hasAtLeastRole(roleName) {
        for (let [id, name] of Role.allRoles) {
            if (name == roleName && Role.currentRoleID <= id) {
                return true
            }
        }
        return false;
    }
}

(async function _() {
    Role.allRoles = new Map(Object.entries(await RoleFetch.getAll()))
})();

(async function _() {
    Role.currentRoleID = (await RoleFetch.getRole(Group.currentGroup.id)).role_id
})();
