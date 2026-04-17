import { useEffect, useState } from "react"
import type React from "react"
import type { ChangeEvent, FormEvent, JSX } from "react"
import { Link } from "react-router-dom"
import { isAxiosError } from "axios"
import { authApi } from "../../api/api"

const REMEMBERED_EMAIL_KEY = "remembered_login_email"

type LoginFormState = {
    email: string
    remember: boolean
}

const LoginPage: React.FC = (): JSX.Element => {
    const [form, setForm] = useState<LoginFormState>({
        email: "",
        remember: false,
    })
    const [isSubmitted, setIsSubmitted] = useState<boolean>(false)
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false)
    const [errorMessage, setErrorMessage] = useState<string>("")

    useEffect(() => {
        const rememberedEmail = localStorage.getItem(REMEMBERED_EMAIL_KEY)
        if (!rememberedEmail) {
            return
        }

        setForm((previous) => ({
            ...previous,
            email: rememberedEmail,
            remember: true,
        }))
    }, [])

    const handleInputChange = (event: ChangeEvent<HTMLInputElement>): void => {
        const { name, value, type, checked } = event.target
        setForm((previous) => ({
            ...previous,
            [name]: type === "checkbox" ? checked : value,
        }))
    }

    const handleSubmit = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
        event.preventDefault()
        setErrorMessage("")
        setIsSubmitted(false)

        const email = form.email.trim()
        if (!email) {
            setErrorMessage("Укажите email для входа")
            return
        }

        setIsSubmitting(true)

        try {
            await authApi.requestMagicLink(email)

            if (form.remember) {
                localStorage.setItem(REMEMBERED_EMAIL_KEY, email)
            } else {
                localStorage.removeItem(REMEMBERED_EMAIL_KEY)
            }

            setIsSubmitted(true)
        } catch (error) {
            if (isAxiosError(error)) {
                const backendMessage = typeof error.response?.data?.message === "string"
                    ? error.response.data.message
                    : "Не удалось отправить письмо. Попробуйте позже."

                setErrorMessage(backendMessage)
            } else {
                setErrorMessage("Не удалось отправить письмо. Попробуйте позже.")
            }
        } finally {
            setIsSubmitting(false)
        }
    }

    return (
        <main>
            <section className="auth-layout">
                <div className="container">
                    <div className="auth-card">
                        <div className="auth-card__header">
                            <span className="auth-card__badge">Добро пожаловать</span>
                            <h1 className="auth-card__title">Авторизация</h1>
                            <p className="auth-card__subtitle">Введите email и получите письмо со ссылкой для входа.</p>
                        </div>

                        <form className="auth-form" onSubmit={handleSubmit}>
                            <label className="auth-form__label" htmlFor="email">Email</label>
                            <input
                                id="email"
                                name="email"
                                type="email"
                                value={form.email}
                                onChange={handleInputChange}
                                placeholder="you@example.com"
                                required
                            />

                            <label className="auth-form__checkbox" htmlFor="remember">
                                <input
                                    id="remember"
                                    name="remember"
                                    type="checkbox"
                                    checked={form.remember}
                                    onChange={handleInputChange}
                                />
                                <span>Запомнить email на этом устройстве</span>
                            </label>

                            <button type="submit" className="btn btn--primary btn--block" disabled={isSubmitting}>
                                {isSubmitting ? "Отправляем..." : "Отправить письмо для входа"}
                            </button>
                        </form>

                        {errorMessage ? (
                            <p className="auth-card__error">{errorMessage}</p>
                        ) : null}

                        {isSubmitted ? (
                            <p className="auth-card__success">
                                Письмо для входа отправлено на {form.email.trim()}. Проверьте почту и перейдите по ссылке.
                            </p>
                        ) : null}

                        <div className="auth-card__links">
                            <Link to="/register" className="link--text">Нет аккаунта? Зарегистрироваться</Link>
                            <Link to="/hubs" className="link--small">Вернуться к списку хабов</Link>
                        </div>
                    </div>
                </div>
            </section>
        </main>
    )
}

export default LoginPage
