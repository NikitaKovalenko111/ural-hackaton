import { useEffect, useState } from "react"
import type React from "react"
import type { ChangeEvent, FormEvent, JSX } from "react"
import { isAxiosError } from "axios"
import { Link } from "react-router-dom"
import universityDomains from "../../data/domains.json"
import { usersApi } from "../../api/api"

type RegistrationFormState = {
    fullName: string
    email: string
    university: string
    telegram: string
}

const initialForm: RegistrationFormState = {
    fullName: "",
    email: "",
    university: "",
    telegram: "",
}

const getDomainFromEmail = (email: string): string => {
    const atIndex = email.lastIndexOf("@")
    if (atIndex === -1) {
        return ""
    }

    return email.slice(atIndex + 1).trim().toLowerCase()
}

const RegisterPage: React.FC = (): JSX.Element => {
    const [form, setForm] = useState<RegistrationFormState>(initialForm)
    const [isStudent, setIsStudent] = useState<boolean>(false)
    const [studentError, setStudentError] = useState<string>("Регистрация доступна только студентам.")
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false)
    const [errorMessage, setErrorMessage] = useState<string>("")
    const [successMessage, setSuccessMessage] = useState<string>("")

    useEffect(() => {
        if (!isStudent) {
            setStudentError("Регистрация доступна только студентам.")
            if (form.university) {
                setForm((previous) => ({ ...previous, university: "" }))
            }
            return
        }

        const domain = getDomainFromEmail(form.email)
        if (!domain) {
            setStudentError("Укажите студенческую почту университета.")
            if (form.university) {
                setForm((previous) => ({ ...previous, university: "" }))
            }
            return
        }

        const foundUniversity = (universityDomains as Record<string, string>)[domain]
        if (!foundUniversity) {
            setStudentError("Домен не найден в списке вузов РФ. Используйте студенческую почту университета.")
            if (form.university) {
                setForm((previous) => ({ ...previous, university: "" }))
            }
            return
        }

        setStudentError("")
        if (form.university !== foundUniversity) {
            setForm((previous) => ({ ...previous, university: foundUniversity }))
        }
    }, [form.email, form.university, isStudent])

    const handleChange = (event: ChangeEvent<HTMLInputElement | HTMLTextAreaElement>): void => {
        const { name, value } = event.target
        if (name === "university") {
            return
        }

        setForm((previous) => ({ ...previous, [name]: value }))
    }

    const handleStudentToggle = (event: ChangeEvent<HTMLInputElement>): void => {
        setIsStudent(event.target.checked)
    }

    const handleSubmit = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
        event.preventDefault()
        setErrorMessage("")
        setSuccessMessage("")

        if (!isStudent) {
            setStudentError("Регистрация доступна только студентам.")
            return
        }

        if (!form.university) {
            setStudentError("Не удалось определить университет по домену email.")
            return
        }

        const trimmedEmail = form.email.trim()
        const trimmedFullName = form.fullName.trim()
        const trimmedTelegram = form.telegram.trim()

        if (!trimmedFullName || !trimmedEmail || !trimmedTelegram) {
            setErrorMessage("Заполните обязательные поля.")
            return
        }

        setIsSubmitting(true)

        try {
            await usersApi.saveUser({
                fullname: trimmedFullName,
                user_role: "student",
                email: trimmedEmail,
                telegram: trimmedTelegram,
                phone: "",
            })

            setSuccessMessage(`Регистрация завершена для ${trimmedFullName}. Университет определён как «${form.university}». Теперь войдите по email.`)
            setForm(initialForm)
            setIsStudent(false)
            setStudentError("Регистрация доступна только студентам.")
        } catch (error) {
            if (isAxiosError(error)) {
                const backendMessage = typeof error.response?.data?.message === "string"
                    ? error.response.data.message
                    : "Не удалось зарегистрировать пользователя. Попробуйте позже."

                setErrorMessage(backendMessage)
            } else {
                setErrorMessage("Не удалось зарегистрировать пользователя. Попробуйте позже.")
            }
        } finally {
            setIsSubmitting(false)
        }
    }

    return (
        <main>
            <section className="booking">
                <div className="container">
                    <div className="section__header booking__header">
                        <div>
                            <h1 className="section__title">Регистрация</h1>
                            <p className="section__subtitle">Заполните профиль и подтвердите, что вы студент. Университет определится автоматически по домену почты.</p>
                        </div>
                    </div>

                    <div className="booking__grid booking__grid--single">
                        <article className="booking-card">
                            <h2 className="booking-card__title">Создание аккаунта</h2>
                            <p className="booking-card__subtitle">Форма почти такая же, как при записи на событие. Нужна студенческая почта и контакты для связи.</p>

                            <form className="booking-form" onSubmit={handleSubmit}>
                                <label className="booking-form__label" htmlFor="register-fullName">ФИО</label>
                                <input
                                    id="register-fullName"
                                    name="fullName"
                                    type="text"
                                    placeholder="Иванов Иван"
                                    value={form.fullName}
                                    onChange={handleChange}
                                    required
                                />

                                <label className="booking-form__label" htmlFor="register-email">Email</label>
                                <input
                                    id="register-email"
                                    name="email"
                                    type="email"
                                    placeholder="student@university.ru"
                                    value={form.email}
                                    onChange={handleChange}
                                    required
                                />

                                <label className="booking-form__student-toggle" htmlFor="register-isStudent">
                                    <input
                                        id="register-isStudent"
                                        name="isStudent"
                                        type="checkbox"
                                        className="booking-form__student-input"
                                        checked={isStudent}
                                        onChange={handleStudentToggle}
                                    />
                                    <span className="booking-form__student-box" aria-hidden="true" />
                                    <span className="booking-form__student-text">Я студент</span>
                                </label>

                                <label className="booking-form__label" htmlFor="register-university">Университет</label>
                                <input
                                    id="register-university"
                                    name="university"
                                    type="text"
                                    value={form.university}
                                    onChange={handleChange}
                                    placeholder="Определится автоматически по домену email"
                                    readOnly
                                    required
                                />

                                {studentError ? <p className="booking-form__error">{studentError}</p> : null}

                                <label className="booking-form__label" htmlFor="register-telegram">Telegram</label>
                                <input
                                    id="register-telegram"
                                    name="telegram"
                                    type="text"
                                    placeholder="@username"
                                    value={form.telegram}
                                    onChange={handleChange}
                                    required
                                />

                                {errorMessage ? <p className="booking-form__error">{errorMessage}</p> : null}

                                <button type="submit" className="btn btn--primary btn--block" disabled={!isStudent || !!studentError || isSubmitting}>
                                    {isSubmitting ? "Регистрируем..." : "Зарегистрироваться"}
                                </button>
                            </form>

                            {successMessage ? <p className="booking-card__success">{successMessage}</p> : null}
                        </article>
                    </div>

                    <div className="auth-card__links" style={{ marginTop: "24px" }}>
                        <Link to="/login" className="link--text">Уже есть аккаунт? Войти</Link>
                        <Link to="/hubs" className="link--small">Вернуться к списку хабов</Link>
                    </div>
                </div>
            </section>
        </main>
    )
}

export default RegisterPage
