import { useState } from "react"
import type React from "react"
import type { ChangeEvent, FormEvent, JSX } from "react"
import { Link } from "react-router-dom"

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

    const handleInputChange = (event: ChangeEvent<HTMLInputElement>): void => {
        const { name, value, type, checked } = event.target
        setForm((previous) => ({
            ...previous,
            [name]: type === "checkbox" ? checked : value,
        }))
    }

    const handleSubmit = (event: FormEvent<HTMLFormElement>): void => {
        event.preventDefault()
        setIsSubmitted(true)
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

                            <button type="submit" className="btn btn--primary btn--block">Отправить письмо для входа</button>
                        </form>

                        {isSubmitted ? (
                            <p className="auth-card__success">
                                Письмо для входа отправлено на {form.email}. Проверьте почту и перейдите по ссылке.
                            </p>
                        ) : null}

                        <div className="auth-card__links">
                            <Link to="/profile" className="link--text">Нет аккаунта? Заполнить профиль</Link>
                            <Link to="/hubs" className="link--small">Вернуться к списку хабов</Link>
                        </div>
                    </div>
                </div>
            </section>
        </main>
    )
}

export default LoginPage
