import { login, responseMessage } from '@/types/user'
import { useState } from 'react'
import { Alert } from '../atoms/Alert';
import { Button } from '@/components/atoms/Button';
import cist_logo from "@/public/cist_logo.png"
import cist_background from "@/public/cist_background.png"
import { APIPost } from '@/function/API/APIPost';
import { Link } from 'react-router-dom';


export const AddUser = () => {
    const [userInfo, setUserInfo] = useState<login>({ id: "", password: "" })
    const [errorText, setErrorText] = useState<string>('')
    const changeName = (e: any) => {
        setUserInfo({ ...userInfo, id: e.target.value })
    }
    const changePass = (e: any) => {
        setUserInfo({ ...userInfo, password: e.target.value })
    }
    const createUser = () => {
        const func = async () => {
            console.log(userInfo)
            setErrorText("")
            const res: responseMessage = await APIPost<login>("account/create", userInfo)
            if (res.errorMessage) {
                setErrorText(res.errorMessage)
            }
            setUserInfo({ id: "", password: "" })
        }
        func()
    }
    return (
        <div style={{
            backgroundImage: `url(${cist_background})`,
            backgroundSize: '100% 100%',
            backgroundPosition: 'center',
            backgroundRepeat: 'no-repeat',
            height: '100vh',
        }} className=''>
            <div className="bg-green-300 p-8 max-w-lg rounded-xl mx-auto">
                <img src={cist_logo} alt="大学ロゴマーク" width={300} className='mx-auto mb-6' />
                <div className="font-bold text-xl text-center mb-8 border-b-2 border-slate-600 pb-4 mx-auto">新規登録</div>
                <div className='flex flex-col space-y-8'>
                    <label className="max-w-lg">
                        <div className="font-semibold mb-2">学籍番号:</div>
                        <input className='border border-black rounded-md w-full h-10 px-2' type="text" name='userName' value={userInfo.id} onChange={(e) => changeName(e)} />
                    </label>
                    <label className="max-w-lg">
                        <div className="font-semibold mb-2">password:</div>
                        <input className='border border-black rounded-md w-full h-10 px-2' type="text" name='password' value={userInfo.password} onChange={(e) => changePass(e)} />
                    </label>
                    {errorText && (<Alert text={errorText} />)}
                    <div className='flex justify-between mt-1'>
                        <Button text="新規登録" onClick={() => createUser()} />
                        <Link to={"/"}>
                            <Button text="ログインページへ" onClick={() => createUser()} />
                        </Link>
                    </div>
                </div>
            </div>
        </div >
    )
}
