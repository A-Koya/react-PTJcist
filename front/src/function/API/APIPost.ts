import { responseMessage } from "@/types/user"
import axios from "axios"

export const APIPost = async<T>(url: string, data: T): Promise<responseMessage> => {
    const res: responseMessage = await axios.post("https://localhost/" + url, data)
    return res
}