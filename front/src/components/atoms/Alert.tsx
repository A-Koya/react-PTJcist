export const Alert = ({ text }: { text: string }) => {
    return (
        <div className="mt-2 bg-red-400 text-sm text-white rounded-lg p-4" role="alert">
            <span className="font-bold">Danger</span>{text}
        </div>
    )
}
