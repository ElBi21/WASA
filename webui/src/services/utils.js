import * as register_ops from "./registration";
import {ref, onMounted} from 'vue';

async function img_to_base64(image) {
    return new Promise((resolve, reject) => {
        let reader = new FileReader();
        reader.onloadend = () => {
            let base64Pfp = reader.result
                .replace("data:", "")
                .replace(/^.+,/, "");

            resolve(base64Pfp);
        };

        reader.onerror = reject;
        reader.readAsDataURL(image);
    });
}


async function retrieveFromStorage() {
    let userDataFromLogin = JSON.parse(sessionStorage.getItem("user"));
    return userDataFromLogin
}


async function clearSessionStorage() {
    sessionStorage.removeItem("user");
}


export {
    img_to_base64,
    retrieveFromStorage,
    clearSessionStorage
}