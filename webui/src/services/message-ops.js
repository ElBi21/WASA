import axios from "./axios";

async function API_send_message(username, chatID, content, photo, replying) {
    await axios.put(`/message`,
        {
            sender: username,
            chat: chatID,
            message_content: content,
            photo: photo,
            replying: replying === true ? 1 : 0
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

export {
    API_send_message,
    API_delete_message
}