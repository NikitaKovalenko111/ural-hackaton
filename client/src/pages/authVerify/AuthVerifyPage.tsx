import { useEffect, useState } from "react"
import type React from "react"
import type { JSX } from "react/jsx-runtime"
import { isAxiosError } from "axios"
import { Link, useNavigate, useSearchParams } from "react-router-dom"
import { useDispatch } from "react-redux"
import { authApi } from "../../api/api"
import { setUser } from "../../redux/features/users/userSlice"
import type { AppDispatch } from "../../redux/store"

type VerifyState = "loading" | "success" | "error"

const AuthVerifyPage: React.FC = (): JSX.Element => {
    const [searchParams] = useSearchParams()
    const navigate = useNavigate()
    const dispatch = useDispatch<AppDispatch>()

    const [status, setStatus] = useState<VerifyState>("loading")
    const [message, setMessage] = useState<string>("Проверяем ссылку для входа...")

    useEffect(() => {
        const token = searchParams.get("token")

        if (!token) {
            setStatus("error")
            setMessage("Ссылка для входа некорректна: отсутствует токен.")
            return
        }

        const verify = async (): Promise<void> => {
            try {
                const response = await authApi.verifyMagicLink(token)
                dispatch(setUser({
                    id: response.user.id,
                    fullname: response.user.fullname,
                    email: response.user.email,
                    role: response.user.role,
                    telegram: response.user.telegram ?? "",
                    phone: response.user.phone ?? "",
                }))

                setStatus("success")
                setMessage("Вход выполнен успешно. Перенаправляем в профиль...")
                setTimeout(() => navigate("/profile", { replace: true }), 1200)
            } catch (error) {
                if (isAxiosError(error)) {
                    const backendMessage = typeof error.response?.data?.message === "string"
                        ? error.response.data.message
                        : "Не удалось подтвердить ссылку."

                    setMessage(backendMessage)
                } else {
                    setMessage("Не удалось подтвердить ссылку.")
                }
                setStatus("error")
            }
        }

        void verify()
    }, [dispatch, navigate, searchParams])

    return (
        <main>
            <section className="auth-layout">
                <div className="container">
                    <div className="auth-card">
                        <div className="auth-card__header">
                            <span className="auth-card__badge">Авторизация</span>
                            <h1 className="auth-card__title">Подтверждение входа</h1>
                            <p className="auth-card__subtitle">{message}</p>
                        </div>

                        {status === "loading" ? (
                            <p className="auth-card__info">Пожалуйста, подождите несколько секунд.</p>
                        ) : null}

                        {status === "success" ? (
                            <p className="auth-card__success">Готово! Сейчас откроется профиль.</p>
                        ) : null}

                        {status === "error" ? (
                            <p className="auth-card__error">Ссылка недействительна или срок ее действия истек.</p>
                        ) : null}

                        <div className="auth-card__links">
                            <Link to="/login" className="link--text">Отправить ссылку заново</Link>
                            <Link to="/hubs" className="link--small">Вернуться к списку хабов</Link>
                        </div>
                    </div>
                </div>
            </section>
        </main>
    )
}

export default AuthVerifyPage
