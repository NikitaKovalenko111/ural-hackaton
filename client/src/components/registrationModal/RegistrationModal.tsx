import { useEffect, useState } from "react"
import type React from "react"
import type { ChangeEvent, FormEvent, JSX } from "react"
import universityDomains from "../../data/domains.json"

type RegistrationModalProps = {
    isOpen: boolean
    onClose: () => void
    eventTitle?: string
    registrationType?: "event" | "mentor"
}

type FormState = {
    fullName: string
    email: string
    university: string
    telegram: string
    comment: string
}

const initialForm: FormState = {
    fullName: "",
    email: "",
    university: "",
    telegram: "",
    comment: "",
}

const getDomainFromEmail = (email: string): string => {
    const atIndex = email.lastIndexOf("@")
    if (atIndex === -1) {
        return ""
    }

    return email.slice(atIndex + 1).trim().toLowerCase()
}

const RegistrationModal: React.FC<RegistrationModalProps> = ({ isOpen, onClose, eventTitle, registrationType = "event" }): JSX.Element | null => {
    const [form, setForm] = useState<FormState>(initialForm)
    const [isStudent, setIsStudent] = useState<boolean>(false)
    const [studentError, setStudentError] = useState<string>("")
    const [successMessage, setSuccessMessage] = useState<string>("")
    const isMentorRegistration = registrationType === "mentor"

    useEffect(() => {
        if (!isOpen) {
            return undefined
        }

        const handleEscape = (event: KeyboardEvent): void => {
            if (event.key === "Escape") {
                onClose()
            }
        }

        window.addEventListener("keydown", handleEscape)
        return () => window.removeEventListener("keydown", handleEscape)
    }, [isOpen, onClose])

    useEffect(() => {
        if (!isOpen) {
            setForm(initialForm)
            setIsStudent(false)
            setStudentError("")
            setSuccessMessage("")
        }
    }, [isOpen])

    useEffect(() => {
        if (!isStudent) {
            setStudentError("Регистрация на мероприятия доступна только студентам.")
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

    const handleSubmit = (event: FormEvent<HTMLFormElement>): void => {
        event.preventDefault()

        if (!isStudent) {
            setStudentError("Регистрация на мероприятия доступна только студентам.")
            return
        }

        if (!form.university) {
            setStudentError("Не удалось определить университет по домену email.")
            return
        }

        const subjectText = isMentorRegistration ? "к ментору" : "на мероприятие"
        setSuccessMessage(`Заявка отправлена ${subjectText}${eventTitle ? ` «${eventTitle}»` : ""}. Мы свяжемся с вами по email ${form.email}.`)
        setForm(initialForm)
        setIsStudent(false)
        setStudentError("")
    }

    if (!isOpen) {
        return null
    }

    return (
        <div className="registration-modal" role="dialog" aria-modal="true" aria-label="Форма записи" onClick={onClose}>
            <div className="registration-modal__card" onClick={(event) => event.stopPropagation()}>
                <button
                    type="button"
                    className="registration-modal__close"
                    onClick={onClose}
                    aria-label="Закрыть окно"
                >
                    ×
                </button>

                <h3 className="registration-modal__title">{isMentorRegistration ? "Запись к ментору" : "Запись на мероприятие"}</h3>
                <p className="registration-modal__subtitle">
                    {eventTitle
                        ? isMentorRegistration
                            ? `Ментор: ${eventTitle}`
                            : `Событие: ${eventTitle}`
                        : isMentorRegistration
                            ? "Оставьте контакты, и мы подтвердим запись к ментору."
                            : "Оставьте контакты, и мы подтвердим участие."}
                </p>

                <form className="registration-modal__form" onSubmit={handleSubmit}>
                    <label className="registration-modal__label" htmlFor="registration-fullName">ФИО</label>
                    <input
                        id="registration-fullName"
                        name="fullName"
                        type="text"
                        placeholder="Иванов Иван"
                        value={form.fullName}
                        onChange={handleChange}
                        required
                    />

                    <label className="registration-modal__label" htmlFor="registration-email">Email</label>
                    <input
                        id="registration-email"
                        name="email"
                        type="email"
                        placeholder="student@university.ru"
                        value={form.email}
                        onChange={handleChange}
                        required
                    />

                    <label className="registration-modal__student-toggle" htmlFor="registration-isStudent">
                        <input
                            id="registration-isStudent"
                            name="isStudent"
                            type="checkbox"
                            className="registration-modal__student-input"
                            checked={isStudent}
                            onChange={handleStudentToggle}
                        />
                        <span className="registration-modal__student-box" aria-hidden="true" />
                        <span className="registration-modal__student-text">Я студент</span>
                    </label>

                    <label className="registration-modal__label" htmlFor="registration-university">Университет</label>
                    <input
                        id="registration-university"
                        name="university"
                        type="text"
                        placeholder="Определится автоматически по домену email"
                        value={form.university}
                        onChange={handleChange}
                        readOnly
                        required
                    />

                    {studentError ? <p className="registration-modal__error">{studentError}</p> : null}

                    <label className="registration-modal__label" htmlFor="registration-telegram">Telegram</label>
                    <input
                        id="registration-telegram"
                        name="telegram"
                        type="text"
                        placeholder="@username"
                        value={form.telegram}
                        onChange={handleChange}
                        required
                    />

                    <label className="registration-modal__label" htmlFor="registration-comment">Комментарий</label>
                    <textarea
                        id="registration-comment"
                        name="comment"
                        placeholder="Уточнения по участию"
                        value={form.comment}
                        onChange={handleChange}
                    />

                    <button type="submit" className="btn btn--primary btn--block" disabled={!isStudent || !!studentError}>Отправить заявку</button>
                </form>

                {successMessage ? <p className="registration-modal__success">{successMessage}</p> : null}
            </div>
        </div>
    )
}

export default RegistrationModal
