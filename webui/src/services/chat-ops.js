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

export {
    API_get_conversation,
    API_create_conversation
}