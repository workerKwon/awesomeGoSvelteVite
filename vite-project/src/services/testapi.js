import api from "../axios/api.js";

export default {
    ping: () => api.get('/ping'),
    albums: () => api.get("/albums"),
    postAlbum: (params) => api.post("/albums", params)
}