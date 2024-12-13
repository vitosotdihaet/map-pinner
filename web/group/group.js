Group.createNewGroup = async (name) => {
    let group = new Group({ name:name })

    let newId = await GroupFetch.create(group)
    group.id = newId.id

    Group.loaded.push(group)
    Group.addOption(group)

    groupSelect.value = group.id
    groupSelect.dispatchEvent(new Event('change'))

    return group
}

Group.reloadAll = async () => {
    Group.loaded = []

    groupsData = null
    try {
        groupsData = await GroupFetch.getAll()
    } catch (error) {
        console.error(`${error}`)
    }

    if (groupsData !== null) {
        groupsData.forEach(groupData => {
            Group.loaded.push(new Group(groupData))
        })
    }

    Group.populateSelect()
}

Group.addOption = (group) => {
    const option = document.createElement('option')
    option.value = group.id
    option.text = group.name
    groupSelect.appendChild(option)
}

Group.populateSelect = () => {
    groupSelect.options.length = 0
    Group.loaded.forEach(group => {
        Group.addOption(group) 
    })
    groupSelect.dispatchEvent(new Event('change'))
}



const newGroupNameInput = document.getElementById('groupName')
newGroupNameInput.addEventListener('input', () => {
    const inputLength = newGroupNameInput.value.length
    if (inputLength >= 1 && inputLength <= 255) {
        newGroupButton.disabled = false
    } else {
        newGroupButton.disabled = true
    }
})

const newGroupButton = document.getElementById('createNewGroup')
newGroupButton.disabled = true
newGroupButton.addEventListener('click', function(event) {
    event.preventDefault()
    const inputLength = newGroupNameInput.value.length
    if (inputLength >= 1 && inputLength <= 255) {
        const newGroup = Group.createNewGroup(newGroupNameInput.value)
        groupSelect.value = newGroup.id
        groupSelect.dispatchEvent(new Event('change'))
        newGroupNameInput.value = ''
    }
})


function hideSelect() {
    groupSelectLabel.hidden = true
    groupSelect.hidden = true
}

function unhideSelect() {
    groupSelectLabel.hidden = false
    groupSelect.hidden = false
}


const groupSelectLabel = document.getElementById('groupSelectLabel')
const groupSelect = document.getElementById('groupSelect')
groupSelect.addEventListener('change', async () => {
    const groupId = groupSelect.value

    addUserToGroupInput.value = ''

    if (groupId != '') {
        Group.currentGroup = new Group(await GroupFetch.getById(groupId))
        Role.currentRoleID = (await RoleFetch.getRole(Group.currentGroup.id)).role_id
        unhideSelect()
        unhideRegions()
        unhideAddUserToGroup()
    } else {
        Group.currentGroup = null
        Role.currentRoleID = null
        hideSelect()
        hideRegions()
        hideAddUserToGroup()
    }

    Group.populateUsers()
    Region.reloadAll()
})


addUserToGroupLabel = document.getElementById('addUserToGroupLabel')
addUserToGroupWithRoleLabel = document.getElementById('addUserToGroupWithRoleLabel')
function hideAddUserToGroup() {
    addUserToGroupInput.hidden = true
    addUserToGroupLabel.hidden = true
    addUserToGroupInput.hidden = true
    addUserToGroupRoleSelect.hidden = true
    addUserToGroupButton.hidden = true
    addUserToGroupWithRoleLabel.hidden = true
}

function unhideAddUserToGroup() {
    addUserToGroupInput.hidden = false
    addUserToGroupLabel.hidden = false
    addUserToGroupInput.hidden = false
    addUserToGroupRoleSelect.hidden = false
    addUserToGroupButton.hidden = false
    addUserToGroupWithRoleLabel.hidden = false
}

const addUserToGroupInput = document.getElementById('addUserToGroupInput')
addUserToGroupInput.addEventListener('input', () => {
    if (!Role.hasAtLeastRole("admin")) { return }
    const inputLength = addUserToGroupInput.value.length
    if (inputLength >= 8 && inputLength <= 32 && Group.currentGroup != null) {
        addUserToGroupButton.disabled = false
    } else {
        addUserToGroupButton.disabled = true
    }
})

const addUserToGroupRoleSelect = document.getElementById('addUserToGroupRoleSelect')
async function populateRoles() {
    while (Role.allRoles == null) {
        await new Promise(resolve => setTimeout(resolve, 50))
    }
    for (let [id, name] of Role.allRoles) {
        const option = document.createElement('option')
        option.value = id
        option.text = name
        addUserToGroupRoleSelect.appendChild(option)
    }
}
populateRoles()

const addUserToGroupButton = document.getElementById('addUserToGroupButton')
addUserToGroupButton.disabled = true
addUserToGroupButton.addEventListener('click', async function(event) {
    if (!Role.hasAtLeastRole("admin")) { return }
    const inputLength = addUserToGroupInput.value.length
    if (inputLength >= 8 && inputLength <= 32 && Group.currentGroup != null) {
        userName = addUserToGroupInput.value
        roleId = addUserToGroupRoleSelect.value

        add = await GroupFetch.addUserToGroup(groupSelect.value, userName, roleId)
        if (add.ok) {
            Group.populateUsers()
        } else {
            if (add.status == 400) {
                alert('No user with such name found! Check your spelling!')
            } else {
                alert('You don\'t have enough rights to add a user to this group!')
            }
        }
    }
})


Group.addUserTableEntry = (userRole) => {
    const row = usersInGroup.insertRow()
    let role = row.insertCell(0)
    role.innerHTML = userRole.role
    let name = row.insertCell(1)
    name.innerHTML = userRole.name
}

function hideUsersInGroup() {
    usersInGroup.hidden = true
    usersInGroupLabel.hidden = true
}

function unhideUsersInGroup() {
    usersInGroup.hidden = false
    usersInGroupLabel.hidden = false
}

const usersInGroupLabel = document.getElementById('usersInGroupLabel')
const usersInGroup = document.getElementById('usersInGroup')
Group.populateUsers = async () => {
    if (Group.currentGroup == null) {
        hideUsersInGroup()
        return
    } else {
        unhideUsersInGroup()
    }
    const rows = usersInGroup.querySelectorAll("tr")
    usersRoles = await GroupFetch.getAllUsers(Group.currentGroup.id)

    // remove previously loaded rows
    rows.forEach((row, index) => {
        if (index > 0) {
            row.remove()
        }
    })

    for (let i = 0; i < usersRoles.users.length; ++i) {
        Group.addUserTableEntry({name: usersRoles.users[i].name, role: usersRoles.roles[i]})
    }
}

Group.reloadAll()
