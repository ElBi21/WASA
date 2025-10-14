import * as register_ops from "./registration";
import {ref, onMounted} from 'vue';

async function img_to_base64(image) {
    let reader = new FileReader();
    let base64Pfp = null;
    reader.onloadend = async () => {
        base64Pfp = reader.result
            .replace("data:", "")
            .replace(/^.+,/, "");

        return base64Pfp;
    }

    reader.readAsDataURL(image);
    return base64Pfp;
}


async function retrieveFromStorage() {
    let userDataFromLogin = JSON.parse(localStorage.getItem("user"));
    return userDataFromLogin
}


export {
    img_to_base64,
    retrieveFromStorage
}