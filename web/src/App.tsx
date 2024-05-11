import {FC} from 'react'
import useMount from '@hooks/useMount.ts'

import { token, goLogin } from '@/network/api.ts'

export const App:FC = () => {
    useMount(() => {
        if (import.meta.env.MODE === 'production') {
            if (!token) goLogin()
        }
    })

    return <></>
}
export default App
