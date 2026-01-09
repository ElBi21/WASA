import * as regist_ops from "./registration"
import axios from '../services/axios.js';

// The login function performs login logic
async function API_login(username) {
    let userData;
    let shouldRegister;

    // POST request to the backend
    await axios.post("/session", {
        headers: {
            "Content-Type": "application/json",
        },

        name: username
    }).then((result) => {
        // Load user data and change flag
        userData = result.data;
        shouldRegister = result.status === 201;
    })

    return {userData, shouldRegister}
}

// The set_new_display_name function sets a new display name for the user
async function API_set_new_display_name(username, newDisplayName) {
    await axios.post(`/user/${username}/name`,
        {
            new_display_name: newDisplayName,
        },
        {
            headers: {
                Authorization: `Bearer ${username}`,
                "Content-Type": "application/json",
            }
        }
    )
}

// The set_new_biography function performs a change of biography
async function API_set_new_biography(username, newBiography) {
    await axios.post(`/user/${username}/bio`,
        {
            new_bio: newBiography
        },
        {
            headers: {
                Authorization: `Bearer ${username}`,
                "Content-Type": "application/json",
            },
        }
    )
}

// The API_set_new_pfp function calls the API to change the PFP of the user
async function API_set_new_pfp(username, newPfpBlob) {
    await axios.post(`/user/${username}/photo`,
        {
            new_photo: newPfpBlob
        },
        {
            headers: {
                Authorization: `Bearer ${username}`,
                "Content-Type": "application/json",
            },
        }
    )
}

async function API_get_conversations(username) {
    let userChats = null;

    await axios.get(`/user/${username}/chats`,
        {
            headers: {
                Authorization: `Bearer ${username}`,
                "Content-Type": "application/json",
            },
        }
    ).then((result) => {
        userChats = result.data;
    })

    return userChats;
}

async function API_get_all_users(username) {
    let allUsers = null;

    await axios.get(`/users`,
        {
            headers: {
                Authorization: `Bearer ${username}`,
                "Content-Type": "application/json",
            },
        }
    ).then((result) => {
        allUsers = result.data;
    })

    return allUsers;
}

async function API_set_new_username(user_id, new_username) {
    await axios.post(`/user/${user_id}/id`, {
            new_username: new_username
        },
        {
            headers: {
                Authorization: `Bearer ${user_id}`,
                "Content-Type": "application/json",
            },
        })
}

export {
    API_login,
    API_set_new_display_name,
    API_set_new_biography,
    API_set_new_pfp,
    API_get_conversations,
    API_get_all_users,
    API_set_new_username
};