class GroupFetch {
    static async getAll() {
        return getJSON('/api/groups')
    }

    static async create(group) {
        return postJSON('/api/groups', JSON.stringify(group))
    }

    static async delete(group) {
        return deleteJSON(`/api/groups/${group.id}`, '')
    }

    static async update(group) {
        return putJSON(`/api/groups/${group.id}`, JSON.stringify(group))
    }
}


class Group {
    static loaded = []

    static async createNewGroup() {
        let group = new Group('', 0)

        let newId = await GroupFetch.create(group)
        group.id = newId

        Group.loaded.push(group)

        return group
    }

    static async reloadAll() {
        const groupsData = await GroupFetch.getAll();
        Group.loaded = []

        if (groupsData === null) { return }

        groupsData.forEach(groupData => {
            Group.loaded.push(new Group(groupData))
        })

        Group.populateOptions()
    }

    static populateOptions() {
        const groupSelect = document.getElementById('groupSelect');

        Group.loaded.forEach(group => {
            const option = document.createElement('option');
            option.value = group.id;
            option.text = group.name;
            groupSelect.appendChild(option);
        });
    }

    constructor(data) {
        this.id = data.id
        this.name = data.name
        this.users = data.users
    }
}


document.getElementById('createNewGroup').addEventListener('click', function(event) {
    Group.createNewGroup()
})

Group.reloadAll()
