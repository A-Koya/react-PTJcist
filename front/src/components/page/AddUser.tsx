import { login } from '@/types/user'
import { useState } from 'react'
import { Alert } from '../atoms/Alert';
import { OrangeButton } from '@/components/atoms/OrangeButton';
import cist_logo from "@/public/cist_logo.png"
import cist_background from "@/public/cist_background.png"


export const AddUser = () => {
    const [loginInfo, setLoginInfo] = useState<login>()
    const [errorText, setErrorText] = useState<string>('')
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
                    <div className="max-w-lg">
                        <div className="font-semibold mb-2">user name:</div>
                        <input className='border border-black rounded-md w-full h-10 px-2' type="text" name='userName' />
                    </div>
                    <div className="max-w-lg">
                        <div className="font-semibold mb-2">password:</div>
                        <input className='border border-black rounded-md w-full h-10 px-2' type="text" name='password' />
                    </div>
                    {errorText && (<Alert text={errorText} />)}
                    <div className='flex justify-between mt-1'>
                        <OrangeButton text="新規登録" />
                        <OrangeButton text="loginページ" />
                    </div>
                </div>
            </div>
        </div >
    )
}
