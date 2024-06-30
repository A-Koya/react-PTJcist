type buttonType = {
    text: string
    onClick: () => void
}

export const Button = ({ text, onClick }: buttonType) => {
    return (
        <button className='font-semibold px-6 py-2 border-2 border-black rounded-md bg-slate-200 hover:bg-orange-200 hover:border-orange-200' onClick={onClick}>{text}</button>
    )
}