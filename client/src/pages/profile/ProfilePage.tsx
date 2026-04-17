import { isAxiosError } from "axios"
import { useEffect, useState } from "react"
import type React from "react"
import type { FormEvent, JSX, MouseEvent } from "react"
import { useDispatch, useSelector } from "react-redux"
import { Link, Navigate, useNavigate } from "react-router-dom"
import { authApi, bookingsApi, eventsApi, hubsApi, mentorsApi, requestsApi, usersApi } from "../../api/api"
import { selectUser, setUser } from "../../redux/features/users/userSlice"
import type { AppDispatch } from "../../redux/store"
import type { IBooking, IRequest } from "../../types"

type AdminTab = "hub" | "admin" | "mentor" | "event"

const ProfilePage: React.FC = (): JSX.Element => {
    const dispatch = useDispatch<AppDispatch>()
    const navigate = useNavigate()
    const user = useSelector(selectUser)

    if (!user) {
        return <Navigate to="/login" replace />
    }

    const roleLabel = user.role ? user.role : "Участник"
    const isAdmin = user.role?.toLowerCase() === "admin"
    const isMentor = user.role?.toLowerCase() === "mentor"
    const isStudent = user.role?.toLowerCase() === "student"
    const telegram = user.telegram?.trim() ? user.telegram : "Не указан"
    const phone = user.phone?.trim() ? user.phone : "Не указан"

    const [activeAdminTab, setActiveAdminTab] = useState<AdminTab>("hub")
    const [adminMessage, setAdminMessage] = useState<string>("")
    const [adminError, setAdminError] = useState<string>("")
    const [isSubmittingAdminAction, setIsSubmittingAdminAction] = useState<boolean>(false)

    const [hubForm, setHubForm] = useState({
        name: "",
        address: "",
        status: "open",
        city: "",
        description: "",
        schedule: "",
        occupancy: "10",
    })

    const [adminForm, setAdminForm] = useState({
        fullname: "",
        email: "",
        telegram: "",
        phone: "",
    })

    const [mentorForm, setMentorForm] = useState({
        fullname: "",
        email: "",
        telegram: "",
        phone: "",
        hubId: "",
    })

    const [eventForm, setEventForm] = useState({
        name: "",
        description: "",
        startTime: "",
        endTime: "",
        hubId: "",
        mentorId: "",
    })

    const [mentorRequests, setMentorRequests] = useState<IRequest[]>([])
    const [studentRequests, setStudentRequests] = useState<IRequest[]>([])
    const [studentBookings, setStudentBookings] = useState<IBooking[]>([])
    const [recordsError, setRecordsError] = useState<string>("")
    const [isLoadingRecords, setIsLoadingRecords] = useState<boolean>(false)

    const clearAdminMessages = (): void => {
        setAdminMessage("")
        setAdminError("")
    }

    const formatDate = (value: string): string => {
        const date = new Date(value)
        if (Number.isNaN(date.getTime())) {
            return value
        }

        return new Intl.DateTimeFormat("ru-RU", {
            day: "2-digit",
            month: "2-digit",
            year: "numeric",
            hour: "2-digit",
            minute: "2-digit",
        }).format(date)
    }

    useEffect(() => {
        const loadRoleRecords = async (): Promise<void> => {
            if (!isMentor && !isStudent) {
                return
            }

            setIsLoadingRecords(true)
            setRecordsError("")

            try {
                if (isMentor) {
                    const mentor = await mentorsApi.getMentorByUserId(user.id)
                    if (!mentor.mentorId) {
                        setMentorRequests([])
                        return
                    }

                    const requests = await requestsApi.getRequestsByMentorId(mentor.mentorId)
                    setMentorRequests(requests)
                    return
                }

                const requests = await requestsApi.getRequestsByUserId(user.id).catch((error: unknown) => {
                    if (isAxiosError(error) && error.response?.status === 404) {
                        return []
                    }

                    throw error
                })

                const bookings = await bookingsApi.getBookingsByUserId(user.id).catch((error: unknown) => {
                    if (isAxiosError(error) && error.response?.status === 404) {
                        return []
                    }

                    throw error
                })

                setStudentRequests(requests)
                setStudentBookings(bookings)
            } catch (error) {
                if (isAxiosError(error) && error.response?.status === 404) {
                    if (isMentor) {
                        setMentorRequests([])
                    }
                    if (isStudent) {
                        setStudentRequests([])
                        setStudentBookings([])
                    }
                    return
                }

                if (isAxiosError(error) && typeof error.response?.data?.message === "string") {
                    setRecordsError(error.response.data.message)
                } else {
                    setRecordsError("Не удалось загрузить записи")
                }
            } finally {
                setIsLoadingRecords(false)
            }
        }

        void loadRoleRecords()
    }, [isMentor, isStudent, user.id])

    const handleCreateHub = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
        event.preventDefault()
        clearAdminMessages()
        setIsSubmittingAdminAction(true)

        try {
            await hubsApi.saveHub({
                hub_name: hubForm.name.trim(),
                address: hubForm.address.trim(),
                status: hubForm.status.trim(),
                city: hubForm.city.trim(),
                description: hubForm.description.trim(),
                schedule: hubForm.schedule.trim(),
                occupancy: Number(hubForm.occupancy),
            })

            setHubForm({
                name: "",
                address: "",
                status: "open",
                city: "",
                description: "",
                schedule: "",
                occupancy: "10",
            })
            setAdminMessage("Хаб успешно добавлен.")
        } catch (error) {
            if (isAxiosError(error) && typeof error.response?.data?.message === "string") {
                setAdminError(error.response.data.message)
            } else {
                setAdminError("Не удалось добавить хаб.")
            }
        } finally {
            setIsSubmittingAdminAction(false)
        }
    }

    const handleCreateAdmin = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
        event.preventDefault()
        clearAdminMessages()
        setIsSubmittingAdminAction(true)

        try {
            await usersApi.saveUser({
                fullname: adminForm.fullname.trim(),
                user_role: "admin",
                email: adminForm.email.trim(),
                telegram: adminForm.telegram.trim(),
                phone: adminForm.phone.trim(),
            })

            setAdminForm({
                fullname: "",
                email: "",
                telegram: "",
                phone: "",
            })
            setAdminMessage("Новый администратор зарегистрирован.")
        } catch (error) {
            if (isAxiosError(error) && typeof error.response?.data?.message === "string") {
                setAdminError(error.response.data.message)
            } else {
                setAdminError("Не удалось зарегистрировать администратора.")
            }
        } finally {
            setIsSubmittingAdminAction(false)
        }
    }

    const handleCreateMentor = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
        event.preventDefault()
        clearAdminMessages()
        setIsSubmittingAdminAction(true)

        try {
            await usersApi.saveUser({
                fullname: mentorForm.fullname.trim(),
                user_role: "mentor",
                email: mentorForm.email.trim(),
                telegram: mentorForm.telegram.trim(),
                phone: mentorForm.phone.trim(),
                hub_id: Number(mentorForm.hubId),
            })

            setMentorForm({
                fullname: "",
                email: "",
                telegram: "",
                phone: "",
                hubId: "",
            })
            setAdminMessage("Новый ментор зарегистрирован.")
        } catch (error) {
            if (isAxiosError(error) && typeof error.response?.data?.message === "string") {
                setAdminError(error.response.data.message)
            } else {
                setAdminError("Не удалось зарегистрировать ментора.")
            }
        } finally {
            setIsSubmittingAdminAction(false)
        }
    }

    const handleCreateEvent = async (event: FormEvent<HTMLFormElement>): Promise<void> => {
        event.preventDefault()
        clearAdminMessages()
        setIsSubmittingAdminAction(true)

        try {
            const startIso = new Date(eventForm.startTime).toISOString()
            const endIso = new Date(eventForm.endTime).toISOString()

            await eventsApi.saveEvent({
                name: eventForm.name.trim(),
                description: eventForm.description.trim(),
                start_time: startIso,
                end_time: endIso,
                hub_id: Number(eventForm.hubId),
                mentor_id: Number(eventForm.mentorId),
            })

            setEventForm({
                name: "",
                description: "",
                startTime: "",
                endTime: "",
                hubId: "",
                mentorId: "",
            })
            setAdminMessage("Событие успешно добавлено.")
        } catch (error) {
            if (isAxiosError(error) && typeof error.response?.data?.message === "string") {
                setAdminError(error.response.data.message)
            } else {
                setAdminError("Не удалось добавить событие.")
            }
        } finally {
            setIsSubmittingAdminAction(false)
        }
    }

    const handleLogout = async (event: MouseEvent<HTMLAnchorElement>): Promise<void> => {
        event.preventDefault()

        try {
            await authApi.logout()
        } finally {
            dispatch(setUser(null))
            navigate("/login", { replace: true })
        }
    }

    return (
        <main className="container">
            <div className="profile__layout">
                <aside className="profile__card">
                    <div className="profile__header">
                        <div className="profile__avatar">👤</div>
                        <div className="profile__meta">
                            <h2>{user.fullname}</h2>
                            <p>{roleLabel}</p>
                            <span className="profile__status-badge">● Авторизован</span>
                        </div>
                    </div>

                    <div className="info-grid">
                        <div className="info-item">
                            <span className="info-item__label">ФИО</span>
                            <span className="info-item__value">{user.fullname}</span>
                        </div>
                        <div className="info-item">
                            <span className="info-item__label">Email</span>
                            <span className="info-item__value">{user.email}</span>
                        </div>
                        <div className="info-item">
                            <span className="info-item__label">Telegram</span>
                            <span className="info-item__value">{telegram}</span>
                        </div>
                        <div className="info-item">
                            <span className="info-item__label">Телефон</span>
                            <span className="info-item__value">{phone}</span>
                        </div>
                        {/* <div className="info-item">
                            <span className="info-item__label">Дата регистрации</span>
                            <span className="info-item__value">12.03.2025</span>
                        </div> */}
                    </div>

                    <h4 className="profile__section-title">О себе</h4>
                    <p className="profile__bio">
                        Профиль загружен из текущей сессии. После обновления страницы данные восстанавливаются по cookie.
                    </p>

                    {isAdmin ? (
                        <div className="profile__admin-panel">
                            <h4 className="profile__section-title">Панель администратора</h4>

                            <div className="profile__admin-tabs">
                                <button type="button" className={`btn btn--sm ${activeAdminTab === "hub" ? "btn--primary" : "btn--outline"}`} onClick={() => setActiveAdminTab("hub")}>Добавление хаба</button>
                                <button type="button" className={`btn btn--sm ${activeAdminTab === "admin" ? "btn--primary" : "btn--outline"}`} onClick={() => setActiveAdminTab("admin")}>Регистрация админов</button>
                                <button type="button" className={`btn btn--sm ${activeAdminTab === "mentor" ? "btn--primary" : "btn--outline"}`} onClick={() => setActiveAdminTab("mentor")}>Регистрация менторов</button>
                                <button type="button" className={`btn btn--sm ${activeAdminTab === "event" ? "btn--primary" : "btn--outline"}`} onClick={() => setActiveAdminTab("event")}>Добавление событий</button>
                            </div>

                            {adminMessage ? <p className="booking-card__success">{adminMessage}</p> : null}
                            {adminError ? <p className="booking-form__error">{adminError}</p> : null}

                            {activeAdminTab === "hub" ? (
                                <form className="booking-form profile__admin-form" onSubmit={handleCreateHub}>
                                    <label className="booking-form__label" htmlFor="hub-name">Название хаба</label>
                                    <input id="hub-name" type="text" value={hubForm.name} onChange={(e) => setHubForm((prev) => ({ ...prev, name: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="hub-address">Адрес</label>
                                    <input id="hub-address" type="text" value={hubForm.address} onChange={(e) => setHubForm((prev) => ({ ...prev, address: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="hub-city">Город</label>
                                    <input id="hub-city" type="text" value={hubForm.city} onChange={(e) => setHubForm((prev) => ({ ...prev, city: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="hub-status">Статус</label>
                                    <input id="hub-status" type="text" value={hubForm.status} onChange={(e) => setHubForm((prev) => ({ ...prev, status: e.target.value }))} placeholder="open" required />

                                    <label className="booking-form__label" htmlFor="hub-description">Описание</label>
                                    <input id="hub-description" type="text" value={hubForm.description} onChange={(e) => setHubForm((prev) => ({ ...prev, description: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="hub-schedule">Расписание</label>
                                    <input id="hub-schedule" type="text" value={hubForm.schedule} onChange={(e) => setHubForm((prev) => ({ ...prev, schedule: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="hub-occupancy">Вместимость</label>
                                    <input id="hub-occupancy" type="number" min={1} value={hubForm.occupancy} onChange={(e) => setHubForm((prev) => ({ ...prev, occupancy: e.target.value }))} required />

                                    <button type="submit" className="btn btn--primary btn--sm" disabled={isSubmittingAdminAction}>{isSubmittingAdminAction ? "Сохраняем..." : "Создать хаб"}</button>
                                </form>
                            ) : null}

                            {activeAdminTab === "admin" ? (
                                <form className="booking-form profile__admin-form" onSubmit={handleCreateAdmin}>
                                    <label className="booking-form__label" htmlFor="admin-fullname">ФИО</label>
                                    <input id="admin-fullname" type="text" value={adminForm.fullname} onChange={(e) => setAdminForm((prev) => ({ ...prev, fullname: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="admin-email">Email</label>
                                    <input id="admin-email" type="email" value={adminForm.email} onChange={(e) => setAdminForm((prev) => ({ ...prev, email: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="admin-telegram">Telegram</label>
                                    <input id="admin-telegram" type="text" value={adminForm.telegram} onChange={(e) => setAdminForm((prev) => ({ ...prev, telegram: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="admin-phone">Телефон</label>
                                    <input id="admin-phone" type="text" value={adminForm.phone} onChange={(e) => setAdminForm((prev) => ({ ...prev, phone: e.target.value }))} required />

                                    <button type="submit" className="btn btn--primary btn--sm" disabled={isSubmittingAdminAction}>{isSubmittingAdminAction ? "Сохраняем..." : "Зарегистрировать админа"}</button>
                                </form>
                            ) : null}

                            {activeAdminTab === "mentor" ? (
                                <form className="booking-form profile__admin-form" onSubmit={handleCreateMentor}>
                                    <label className="booking-form__label" htmlFor="mentor-fullname">ФИО</label>
                                    <input id="mentor-fullname" type="text" value={mentorForm.fullname} onChange={(e) => setMentorForm((prev) => ({ ...prev, fullname: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="mentor-email">Email</label>
                                    <input id="mentor-email" type="email" value={mentorForm.email} onChange={(e) => setMentorForm((prev) => ({ ...prev, email: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="mentor-telegram">Telegram</label>
                                    <input id="mentor-telegram" type="text" value={mentorForm.telegram} onChange={(e) => setMentorForm((prev) => ({ ...prev, telegram: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="mentor-phone">Телефон</label>
                                    <input id="mentor-phone" type="text" value={mentorForm.phone} onChange={(e) => setMentorForm((prev) => ({ ...prev, phone: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="mentor-hub-id">ID хаба</label>
                                    <input id="mentor-hub-id" type="number" min={1} value={mentorForm.hubId} onChange={(e) => setMentorForm((prev) => ({ ...prev, hubId: e.target.value }))} required />

                                    <button type="submit" className="btn btn--primary btn--sm" disabled={isSubmittingAdminAction}>{isSubmittingAdminAction ? "Сохраняем..." : "Зарегистрировать ментора"}</button>
                                </form>
                            ) : null}

                            {activeAdminTab === "event" ? (
                                <form className="booking-form profile__admin-form" onSubmit={handleCreateEvent}>
                                    <label className="booking-form__label" htmlFor="event-name">Название события</label>
                                    <input id="event-name" type="text" value={eventForm.name} onChange={(e) => setEventForm((prev) => ({ ...prev, name: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="event-description">Описание</label>
                                    <input id="event-description" type="text" value={eventForm.description} onChange={(e) => setEventForm((prev) => ({ ...prev, description: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="event-start">Начало</label>
                                    <input id="event-start" type="datetime-local" value={eventForm.startTime} onChange={(e) => setEventForm((prev) => ({ ...prev, startTime: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="event-end">Окончание</label>
                                    <input id="event-end" type="datetime-local" value={eventForm.endTime} onChange={(e) => setEventForm((prev) => ({ ...prev, endTime: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="event-hub">ID хаба</label>
                                    <input id="event-hub" type="number" min={1} value={eventForm.hubId} onChange={(e) => setEventForm((prev) => ({ ...prev, hubId: e.target.value }))} required />

                                    <label className="booking-form__label" htmlFor="event-mentor">ID ментора</label>
                                    <input id="event-mentor" type="number" min={1} value={eventForm.mentorId} onChange={(e) => setEventForm((prev) => ({ ...prev, mentorId: e.target.value }))} required />

                                    <button type="submit" className="btn btn--primary btn--sm" disabled={isSubmittingAdminAction}>{isSubmittingAdminAction ? "Сохраняем..." : "Создать событие"}</button>
                                </form>
                            ) : null}
                        </div>
                    ) : null}

                    {/*<h4 className="profile__section-title">Интересы и стек</h4>
                    <div className="tags-container">
                        <span className="tag tag--stack">JavaScript</span>
                        <span className="tag tag--stack">React</span>
                        <span className="tag tag--stack">TypeScript</span>
                        <span className="tag tag--stack">UI/UX</span>
                        <span className="tag tag--stack">Figma</span>
                    </div>*/}

                    <div className="actions-row profile__actions">
                        <Link to="/profile" className="btn btn--primary btn--sm">Редактировать профиль</Link>
                        <Link to="/profile" className="btn btn--sm btn--outline">Настройки аккаунта</Link>
                    </div>
                    <Link to="/login" className="link--danger" onClick={(event) => void handleLogout(event)}>Выйти из аккаунта</Link>
                </aside>

                <section className="notifications__section">
                    <div className="notifications__header">
                        <h2>{isMentor ? "Записи к ментору" : "Мои записи"}</h2>
                    </div>

                    {isMentor ? (
                        <div className="notifications__list" id="notifList">
                            {isLoadingRecords ? <p className="hubs-search__empty">Загрузка записей...</p> : null}
                            {recordsError ? <p className="booking-form__error">{recordsError}</p> : null}
                            {!isLoadingRecords && !recordsError && mentorRequests.length === 0 ? <p className="hubs-search__empty">Пока нет записей к вам.</p> : null}
                            {!isLoadingRecords && !recordsError && mentorRequests.length > 0 ? (
                                <div className="notifications__list">
                                    {mentorRequests.map((request) => (
                                        <div key={request.id} className="notifications__item notifications__item--unread">
                                            <div className="notifications__icon">🧑‍🏫</div>
                                            <div className="notifications__content">
                                                <div className="notifications__title">Заявка #{request.id}</div>
                                                <div className="notifications__description">{request.message}</div>
                                                <div className="notifications__time">Студент ID: {request.userId}</div>
                                            </div>
                                        </div>
                                    ))}
                                </div>
                            ) : null}
                        </div>
                    ) : null}

                    {isStudent ? (
                        <div className="notifications__list" id="notifList">
                            <h4 className="profile__section-title">Мои записи к менторам</h4>
                            {isLoadingRecords ? <p className="hubs-search__empty">Загрузка записей...</p> : null}
                            {recordsError ? <p className="booking-form__error">{recordsError}</p> : null}
                            {!isLoadingRecords && !recordsError && studentRequests.length === 0 ? <p className="hubs-search__empty">У вас пока нет записей к менторам.</p> : null}
                            {!isLoadingRecords && !recordsError && studentRequests.length > 0 ? (
                                <div className="notifications__list">
                                    {studentRequests.map((request) => (
                                        <div key={request.id} className="notifications__item notifications__item--unread">
                                            <div className="notifications__icon">📝</div>
                                            <div className="notifications__content">
                                                <div className="notifications__title">Заявка к ментору #{request.id}</div>
                                                <div className="notifications__description">{request.message}</div>
                                                <div className="notifications__time">Ментор ID: {request.mentorId ?? "не указан"}</div>
                                            </div>
                                        </div>
                                    ))}
                                </div>
                            ) : null}

                            <h4 className="profile__section-title">Мои регистрации на события</h4>
                            {!isLoadingRecords && !recordsError && studentBookings.length === 0 ? <p className="hubs-search__empty">У вас пока нет регистраций на события.</p> : null}
                            {!isLoadingRecords && !recordsError && studentBookings.length > 0 ? (
                                <div className="notifications__list">
                                    {studentBookings.map((booking) => (
                                        <div key={booking.id} className="notifications__item">
                                            <div className="notifications__icon">📅</div>
                                            <div className="notifications__content">
                                                <div className="notifications__title">Регистрация #{booking.id}</div>
                                                <div className="notifications__description">{booking.bookingZone}</div>
                                                <div className="notifications__time">{formatDate(booking.bookingDate)}</div>
                                            </div>
                                        </div>
                                    ))}
                                </div>
                            ) : null}
                        </div>
                    ) : null}

                    {!isMentor && !isStudent ? (
                        <div className="notifications__list" id="notifList">
                            <p className="hubs-search__empty">Для вашей роли пока нет персональных записей.</p>
                        </div>
                    ) : null}
                </section>
            </div>
        </main>
    )
}

export default ProfilePage