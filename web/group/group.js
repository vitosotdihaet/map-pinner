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
        groupsData = await GroupFetch.getAll();
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
    const option = document.createElement('option');
    option.value = group.id;
    option.text = group.name;
    groupSelect.appendChild(option);
}

Group.populateSelect = () => {
    groupSelect.options.length = 0
    Group.loaded.forEach(group => {
        Group.addOption(group) 
    });
    groupSelect.dispatchEvent(new Event('change'))
}



const newGroupNameInput = document.getElementById('groupName');
newGroupNameInput.addEventListener('input', () => {
    const inputLength = newGroupNameInput.value.length;
    if (inputLength >= 1 && inputLength <= 255) {
        newGroupButton.disabled = false;
    } else {
        newGroupButton.disabled = true;
    }
});

const newGroupButton = document.getElementById('createNewGroup');
newGroupButton.disabled = true;
newGroupButton.addEventListener('click', function(event) {
    event.preventDefault()
    const inputLength = newGroupNameInput.value.length;
    if (inputLength >= 1 && inputLength <= 255) {
        const newGroup = Group.createNewGroup(newGroupNameInput.value)
        groupSelect.value = newGroup.id;
        groupSelect.dispatchEvent(new Event('change'))
        newGroupNameInput.value = ''
    }
})


const groupSelect = document.getElementById('groupSelect');
groupSelect.addEventListener('change', async () => {
    const groupId = groupSelect.value

    if (groupId != '') {
        Group.currentGroup = new Group(await GroupFetch.getById(groupId))
        activateRegions()
        Region.reloadAll()
    } else {
        Group.currentGroup = null
        deactivateRegions()
        Region.reloadAll()
    }
});


const addUserToGroupInput = document.getElementById('addUserToGroupInput')
addUserToGroupInput.addEventListener('input', () => {
    const inputLength = addUserToGroupInput.value.length
    if (inputLength >= 1 && inputLength <= 255) {
        addUserToGroupButton.disabled = false
    } else {
        addUserToGroupButton.disabled = true
    }
});

const addUserToGroupRoleSelect = document.getElementById('addUserToGroupRoleSelect')
async function populateRoles() {
    roles = new Map(Object.entries(await RoleFetch.getAll()))

    for (let [id, name] of roles) {
        const option = document.createElement('option')
        option.value = id;
        option.text = name;
        addUserToGroupRoleSelect.appendChild(option)
    }
}
populateRoles()

const addUserToGroupButton = document.getElementById('addUserToGroupButton')
addUserToGroupButton.disabled = true;
addUserToGroupButton.addEventListener('click', async function(event) {
    event.preventDefault()
    const inputLength = addUserToGroupInput.value.length
    if (inputLength >= 1 && inputLength <= 255) {
        userName = addUserToGroupInput.value
        roleId = addUserToGroupRoleSelect.value

        add = await GroupFetch.addUserToGroup(groupSelect.value, userName, roleId)
        if (add.ok) {
            // TODO: add to a group user list
        } else {
            if (add.status == 400) {
                alert('No user with such name found! Check your spelling!')
            } else {
                alert('You don\'t have enough rights to add a user to this group!')
            }
        }
    }
})

Group.reloadAll()
