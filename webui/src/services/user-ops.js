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
    }).catch((error) => {
        // TODO: Call error message popup
        console.log(error)
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
    ).catch((error) => {
        // TODO: Call an error message
        console.log(error)
    })
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
    ).catch((error) => {
        // TODO: Call an error message
        console.log(error)
    })
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
    ).catch((error) => {
        // TODO: Call an error message
        console.log(error)
    })
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
    }).catch((error) => {
        // TODO: Call an error message
        console.log(error)
    })

    return userChats;
}

export {
    API_login,
    API_set_new_display_name,
    API_set_new_biography,
    API_set_new_pfp,
    API_get_conversations
};