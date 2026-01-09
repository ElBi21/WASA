import axios from '../services/axios.js';


async function API_get_conversation(chat_id, username) {
    let chat_object = null;

    await axios.get(`/chat/${chat_id}`,
        {
            headers: {
                Authorization: `Bearer ${username}`,
                "Content-Type": "application/json",
            },
        }
    ).then((result) => {
        chat_object = result.data;
    })

    return chat_object;
}

async function API_create_conversation(isPrivate, users, name, groupDescription, photo, username) {
    await axios.put(`/chat`, {
            is_private: isPrivate,
            users: users,
            name: name,
            group_description: groupDescription,
            photo: photo
        },
        {
            headers: {
                Authorization: `Bearer ${username}`,
                "Content-Type": "application/json",
            },
        }
    )
}

async function API_get_conversation_messages(chat_id, username) {
    let all_messages = null;

    await axios.get(`/chat/${chat_id}/messages`,
        {
            headers: {
                Authorization: `Bearer ${username}`,
                "Content-Type": "application/json",
            },
        }
    ).then((result) => {
        all_messages = result.data.messages;
    })

    return all_messages;
}

async function API_leave_group(chat_id, username) {
    await axios.delete(`/chat/${chat_id}/users/${username}`,
        {
            headers: {
                Authorization: `Bearer ${username}`,
                "Content-Type": "application/json",
            },
        }
    )
}

async function API_add_to_group(chat_id, user_id_to_add, username) {
    await axios.put(`/chat/${chat_id}/users`, {
            new_user: user_id_to_add
        },
        {
            headers: {
                Authorization: `Bearer ${username}`,
                "Content-Type": "application/json",
            },
        }
    )
}

async function API_set_group_name(chat_id, new_name, user_id) {
    await axios.post(`/chat/${chat_id}/name`, {
        new_name: new_name
    },
        {
            headers: {
                Authorization: `Bearer ${user_id}`,
                "Content-Type": "application/json",
            },
        })
}

async function API_set_group_description(chat_id, new_description, user_id) {
    await axios.post(`/chat/${chat_id}/description`, {
            new_description: new_description
        },
        {
            headers: {
                Authorization: `Bearer ${user_id}`,
                "Content-Type": "application/json",
            },
        })
}

async function API_set_group_photo(chat_id, new_photo, user_id) {
    await axios.post(`/chat/${chat_id}/photo`, {
            new_photo: new_photo
        },
        {
            headers: {
                Authorization: `Bearer ${user_id}`,
                "Content-Type": "application/json",
            },
        })
}

export {
    API_get_conversation,
    API_create_conversation,
    API_get_conversation_messages,
    API_leave_group,
    API_add_to_group,
    API_set_group_name,
    API_set_group_description,
    API_set_group_photo
}