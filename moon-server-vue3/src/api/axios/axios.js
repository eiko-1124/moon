import axios from "axios";
import config from "./config"
import { requestSuccess } from "./interceptors";

const _axios = axios.create(config)

_axios.interceptors.request.use(requestSuccess)

export default _axios