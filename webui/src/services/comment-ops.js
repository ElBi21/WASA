import axios from '../services/axios.js';


async function API_comment_message(messageID, commenterID, commentContent) {
    await axios.put(`/comment`, {
        message: messageID,
        sender: commenterID,
        reaction_content: commentContent
    }, {
        headers: {
            Authorization: `Bearer ${commenterID}`,
            "Content-Type": "application/json",
        },
    })
}

async function API_uncomment_message(commentID, userRemover) {
    await axios.delete(`/comment/${commentID}`, {
        headers: {
            Authorization: `Bearer ${userRemover}`,
            "Content-Type": "application/json",
        },
    })
}


export {
    API_comment_message,
    API_uncomment_message
}