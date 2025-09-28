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

async function base64_to_img(base64, reference) {
    let format = "png"

    if (!base64.startsWith('data:image/')) {
        base64 = `data:image/${format};base64,` + base64;
    }

    reference.value = base64;
}


export {
    img_to_base64,
    base64_to_img
}