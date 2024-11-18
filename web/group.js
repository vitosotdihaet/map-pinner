class GroupFetch {
    static async getAll() {
        return getData('/api/groups')
    }

    static async create(group) {
        return postData('/api/groups', JSON.stringify(group))
    }

    static async delete(group) {
        return deleteData(`/api/groups/${group.id}`, '')
    }

    static async update(group) {
        return putData(`/api/groups/${group.id}`, JSON.stringify(group))
    }
}


class Group {
    constructor(name, id) {
        this.name = name
        this.id = id
    }

    static async createNewGroup() {
        let group = new Group('', 0)

        let newId = await GroupFetch.create(group)
        group.id = newId

        return group
    }
}


document.getElementById('createNewGroup').addEventListener('click', function(event) {
    Group.createNewGroup()
    loadGroups()
})


async function loadGroups() {
    const groups = await GroupFetch.getAll();
    const groupSelect = document.getElementById('groupSelect');

    groups.forEach(group => {
        const option = document.createElement('option');
        option.value = group.id;
        option.text = group.name;
        groupSelect.appendChild(option);
    });
}