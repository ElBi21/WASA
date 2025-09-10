import axios from "axios";

const __API_URL__ = "http://localhost:3000"

const instance = axios.create({
	baseURL: __API_URL__,
	timeout: 1000 * 5
});

export default instance;
