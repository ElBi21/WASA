import axios from "./axios";

async function API_send_message(username, chatID, content, photo, replying) {
    await axios.put(`/message`,
        {
            sender: username,
            chat: chatID,
            message_content: content,
            photo: photo,
            replying: replying
        },
        {
            headers: {
                Authorization: `Bearer ${username}`,
                "Content-Type": "application/json",
            },
        }
    )
}

async function API_delete_message(message_id, username) {
    await axios.delete(`/message/${message_id}`,
        {
            headers: {
                Authorization: `Bearer ${username}`
            },
        }
    )
}

async function API_retrieve_message(message_id, username) {
    let message = null;

    await axios.get(`/message/${message_id}`,
        {
            headers: {
                Authorization: `Bearer ${username}`
            },
        }
    ).then((result) => {
        message = result.data;
    });

    return message;
}

async function API_forward_message(message_obj, chat_to_forward, username) {
    await axios.post(`/message/${message_obj.message_id}`,
        {
            sender: username,
            chat_id: chat_to_forward
        },
        {
            headers: {
                Authorization: `Bearer ${username}`,
                "Content-Type": "application/json",
            },
    })
}

export {
    API_send_message,
    API_delete_message,
    API_retrieve_message,
    API_forward_message
}