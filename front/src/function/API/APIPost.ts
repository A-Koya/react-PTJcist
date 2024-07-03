import { responseMessage } from "@/types/user"
import axios from "axios"

export const APIPost = async<T>(url: string, data: T): Promise<responseMessage> => {
    const res: responseMessage = await axios.post("http://localhost:8080/" + url, data, {
        headers: {
            'Content-Type': 'application/json'
        }
    })
    return res
}