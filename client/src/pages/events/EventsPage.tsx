import { useEffect, useState } from "react"
import type React from "react"
import type { ChangeEvent, FormEvent, JSX, MouseEvent } from "react"
import { Link } from "react-router-dom"
import RegistrationModal from "../../components/registrationModal/RegistrationModal"
import universityDomains from "../../data/domains.json"

const EventsPage: React.FC = (): JSX.Element => {
    const [hubSuccess, setHubSuccess] = useState<string>("")
    const [isRegistrationOpen, setIsRegistrationOpen] = useState<boolean>(false)
    const [selectedEventTitle, setSelectedEventTitle] = useState<string>("Выбранное событие")

    const [hubForm, setHubForm] = useState({
        fullName: "",
        email: "",
        university: "",
        date: "",
        period: "",
        zone: "",
        seats: "1",
    })
    const [hubIsStudent, setHubIsStudent] = useState<boolean>(false)
    const [hubStudentError, setHubStudentError] = useState<string>("Регистрация на мероприятия доступна только студентам.")

    const getDomainFromEmail = (email: string): string => {
        const atIndex = email.lastIndexOf("@")
        if (atIndex === -1) {
            return ""
        }

        return email.slice(atIndex + 1).trim().toLowerCase()
    }

    useEffect(() => {
        if (!hubIsStudent) {
            setHubStudentError("Регистрация на мероприятия доступна только студентам.")
            if (hubForm.university) {
                setHubForm((previous) => ({ ...previous, university: "" }))
            }
            return
        }

        const domain = getDomainFromEmail(hubForm.email)
        if (!domain) {
            setHubStudentError("Укажите студенческую почту университета.")
            if (hubForm.university) {
                setHubForm((previous) => ({ ...previous, university: "" }))
            }
            return
        }

        const foundUniversity = (universityDomains as Record<string, string>)[domain]
        if (!foundUniversity) {
            setHubStudentError("Домен не найден в списке вузов РФ. Используйте студенческую почту университета.")
            if (hubForm.university) {
                setHubForm((previous) => ({ ...previous, university: "" }))
            }
            return
        }

        setHubStudentError("")
        if (hubForm.university !== foundUniversity) {
            setHubForm((previous) => ({ ...previous, university: foundUniversity }))
        }
    }, [hubForm.email, hubForm.university, hubIsStudent])

    const handleHubChange = (event: ChangeEvent<HTMLInputElement | HTMLSelectElement>): void => {
        const { name, value } = event.target
        if (name === "university") {
            return
        }

        setHubForm((previous) => ({ ...previous, [name]: value }))
    }

    const handleHubStudentToggle = (event: ChangeEvent<HTMLInputElement>): void => {
        setHubIsStudent(event.target.checked)
    }

    const handleHubSubmit = (event: FormEvent<HTMLFormElement>): void => {
        event.preventDefault()

        if (!hubIsStudent) {
            setHubStudentError("Регистрация на мероприятия доступна только студентам.")
            return
        }

        if (!hubForm.university) {
            setHubStudentError("Не удалось определить университет по домену email.")
            return
        }

        setHubSuccess(`Бронь в зоне «${hubForm.zone}» на ${hubForm.date} (${hubForm.period}) для ${hubForm.university} принята.`)
        setHubForm({
            fullName: "",
            email: "",
            university: "",
            date: "",
            period: "",
            zone: "",
            seats: "1",
        })
        setHubIsStudent(false)
        setHubStudentError("")
    }

    const handleRegistrationClick = (event: MouseEvent<HTMLAnchorElement>): void => {
        event.preventDefault()
        const eventCard = event.currentTarget.closest(".event-card")
        const titleElement = eventCard?.querySelector(".event-card__title")
        setSelectedEventTitle(titleElement?.textContent ?? "Выбранное событие")
        setIsRegistrationOpen(true)
    }

    return (
        <main>
            <section className="section section--events">
                <div className="container">
                    <div className="section__header">
                        <div>
                            <h2 className="section__title">События</h2>
                            <p className="section__subtitle">Мастер-классы, встречи, лайв-кодинги и практикумы. Выберите удобные даты.</p>
                        </div>
                        <button className="btn btn--sm btn--outline">Сбросить</button>
                    </div>
                    <div className="events-grid">
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 25.04.2026 19:00–21:00</span>
                            </div>
                            <h3 className="event-card__title">Практикум: Введение в FastAPI</h3>
                            <p className="event-card__description">Быстрый старт с FastAPI: структура проекта, валидация схем и
                                тестирование.</p>
                            <p className="event-card__role">Наставник: Алексей Смирнов</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" onClick={handleRegistrationClick} className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">10.05.2026 18:30–20:00</span>
                            </div>
                            <h3 className="event-card__title">Круглый стол: Архитектура микросервисов</h3>
                            <p className="event-card__description">Плюсы и минусы микросервисов, границы сервисов, коммуникации и
                                наблюдаемость.</p>
                            <p className="event-card__role">Наставники: Ольга Романова, Никита Лебедев</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" onClick={handleRegistrationClick} className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 03.06.2026 20:00–21:30</span>
                            </div>
                            <h3 className="event-card__title">Лайв-кодинг: React + TypeScript</h3>
                            <p className="event-card__description">Пишем приложение с управлением состоянием, формами и оптимизациями
                                рендеринга.</p>
                            <p className="event-card__role">Наставник: Светлана Яковлева</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" onClick={handleRegistrationClick} className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 22.05.2026 19:00–21:00</span>
                            </div>
                            <h3 className="event-card__title">Workshop: Profiling Go-сервисов</h3>
                            <p className="event-card__description">CPU/Memory профили, pprof, flamegraph и поиск узких мест в проде.</p>
                            <p className="event-card__role">Наставник: Екатерина Зайцева</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" onClick={handleRegistrationClick} className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 30.04.2026 18:00–19:00</span>
                            </div>
                            <h3 className="event-card__title">Q&A с Solution Architect</h3>
                            <p className="event-card__description">Ответы на вопросы об архитектуре, интеграциях и безопасности приложений.
                            </p>
                            <p className="event-card__role">Наставник: Татьяна Соколова</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" onClick={handleRegistrationClick} className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 12.07.2026 11:00–15:00</span>
                            </div>
                            <h3 className="event-card__title">Интенсив: Docker для разработчиков</h3>
                            <p className="event-card__description">Сборка образов, оптимизация Dockerfile, многоконтейнерные окружения и
                                best practices.</p>
                            <p className="event-card__role">Наставник: Игорь Кузнецов</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" onClick={handleRegistrationClick} className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">05.05.2026 17:00–18:30</span>
                            </div>
                            <h3 className="event-card__title">AMA: Карьера в Data Science</h3>
                            <p className="event-card__description">Рынок, навыки, собесы и портфолио. Живые ответы на вопросы участников.
                            </p>
                            <p className="event-card__role">Наставник: Мария Орлова</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" onClick={handleRegistrationClick} className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 18.06.2026 19:00–21:00</span>
                            </div>
                            <h3 className="event-card__title">Практикум: CI/CD с GitHub Actions</h3>
                            <p className="event-card__description">Сборка, тесты, релизы и деплой. Секреты, матрицы, кеширование и
                                артефакты.</p>
                            <p className="event-card__role">Наставник: Денис Пак</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" onClick={handleRegistrationClick} className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                        <div className="event-card">
                            <div className="event-card__header">
                                <span className="event-card__time">🕒 15.04.2026 18:00–19:00</span>
                            </div>
                            <h3 className="event-card__title">Разбор резюме и портфолио</h3>
                            <p className="event-card__description">Как структурировать опыт, выбрать проекты и выделиться в откликах.</p>
                            <p className="event-card__role">Ведущий: Команда «Студент и Т»</p>
                            <div className="mentor-card__actions">
                                <Link to="/profile" onClick={handleRegistrationClick} className="btn btn--primary btn--sm">Регистрация</Link>
                            </div>
                        </div>
                    </div>

                    <section className="booking" aria-label="Форма бронирования места в хабе">
                        <div className="section__header booking__header">
                            <div>
                                <h2 className="section__title">Бронирование</h2>
                                <p className="section__subtitle">Забронируйте место для работы в хабе. Бронирование доступно только студентам с университетской почтой.</p>
                            </div>
                        </div>

                        <div className="booking__grid booking__grid--single">
                            <article className="booking-card">
                                <h3 className="booking-card__title">Бронирование места в хабе</h3>
                                <p className="booking-card__subtitle">Выберите дату, зону и количество мест для команды.</p>

                                <form className="booking-form" onSubmit={handleHubSubmit}>
                                    <label className="booking-form__label" htmlFor="hub-fullName">Контактное лицо</label>
                                    <input
                                        id="hub-fullName"
                                        name="fullName"
                                        type="text"
                                        value={hubForm.fullName}
                                        onChange={handleHubChange}
                                        placeholder="Петрова Анна"
                                        required
                                    />

                                    <label className="booking-form__label" htmlFor="hub-email">Email</label>
                                    <input
                                        id="hub-email"
                                        name="email"
                                        type="email"
                                        value={hubForm.email}
                                        onChange={handleHubChange}
                                        placeholder="student@university.ru"
                                        required
                                    />

                                    <label className="booking-form__student-toggle" htmlFor="hub-isStudent">
                                        <input
                                            id="hub-isStudent"
                                            name="isStudent"
                                            type="checkbox"
                                            className="booking-form__student-input"
                                            checked={hubIsStudent}
                                            onChange={handleHubStudentToggle}
                                        />
                                        <span className="booking-form__student-box" aria-hidden="true" />
                                        <span className="booking-form__student-text">Я студент</span>
                                    </label>

                                    <label className="booking-form__label" htmlFor="hub-university">Университет</label>
                                    <input
                                        id="hub-university"
                                        name="university"
                                        type="text"
                                        value={hubForm.university}
                                        onChange={handleHubChange}
                                        placeholder="Определится автоматически по домену email"
                                        readOnly
                                        required
                                    />

                                    {hubStudentError ? <p className="booking-form__error">{hubStudentError}</p> : null}

                                    <div className="booking-form__row">
                                        <div className="booking-form__field">
                                            <label className="booking-form__label" htmlFor="hub-date">Дата</label>
                                            <input
                                                id="hub-date"
                                                name="date"
                                                type="date"
                                                value={hubForm.date}
                                                onChange={handleHubChange}
                                                required
                                            />
                                        </div>
                                        <div className="booking-form__field">
                                            <label className="booking-form__label" htmlFor="hub-period">Время</label>
                                            <select
                                                id="hub-period"
                                                name="period"
                                                value={hubForm.period}
                                                onChange={handleHubChange}
                                                required
                                            >
                                                <option value="">Выберите слот</option>
                                                <option value="09:00–13:00">09:00–13:00</option>
                                                <option value="13:00–17:00">13:00–17:00</option>
                                                <option value="17:00–21:00">17:00–21:00</option>
                                            </select>
                                        </div>
                                    </div>

                                    <div className="booking-form__row">
                                        <div className="booking-form__field">
                                            <label className="booking-form__label" htmlFor="hub-zone">Зона</label>
                                            <select
                                                id="hub-zone"
                                                name="zone"
                                                value={hubForm.zone}
                                                onChange={handleHubChange}
                                                required
                                            >
                                                <option value="">Выберите зону</option>
                                                <option value="Опен-спейс">Опен-спейс</option>
                                                <option value="Переговорная">Переговорная</option>
                                                <option value="Тихая зона">Тихая зона</option>
                                            </select>
                                        </div>
                                        <div className="booking-form__field">
                                            <label className="booking-form__label" htmlFor="hub-seats">Мест</label>
                                            <select
                                                id="hub-seats"
                                                name="seats"
                                                value={hubForm.seats}
                                                onChange={handleHubChange}
                                                required
                                            >
                                                <option value="1">1</option>
                                                <option value="2">2</option>
                                                <option value="3">3</option>
                                                <option value="4">4</option>
                                                <option value="5">5</option>
                                                <option value="6">6</option>
                                            </select>
                                        </div>
                                    </div>

                                    <button type="submit" className="btn btn--primary" disabled={!hubIsStudent || !!hubStudentError}>Забронировать место</button>
                                </form>

                                {hubSuccess ? <p className="booking-card__success">{hubSuccess}</p> : null}
                            </article>
                        </div>
                    </section>
                </div>
            </section>

            <RegistrationModal
                isOpen={isRegistrationOpen}
                onClose={() => setIsRegistrationOpen(false)}
                eventTitle={selectedEventTitle}
            />
        </main>
    )
}

export default EventsPage


